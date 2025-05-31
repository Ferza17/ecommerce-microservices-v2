import { Injectable, Logger } from '@nestjs/common';
import { CartDocument, CartItem } from '../../model/mongo/cart';
import { InjectModel } from '@nestjs/mongoose';
import mongoose, { Model } from 'mongoose';
import {
  CreateCartItemRequest,
  FindCartItemsWithPaginationRequest,
  FindCartItemsWithPaginationResponse,
} from '../../model/rpc/gen/commerce/v1/cartMessage';
import { JaegerTelemetryService } from '../../infrastructure/telemetry/jaeger.telemetry.service';
import { Context } from '@opentelemetry/api';


@Injectable()
export class CartMongodbRepository {
  private readonly logger = new Logger(CartMongodbRepository.name);

  constructor(
    @InjectModel(CartItem.name)
    private cartModel: Model<CartDocument>,
    private readonly otel: JaegerTelemetryService,
  ) {
  }

  async CreateCartItem(requestId: string, request: CreateCartItemRequest, context?: Context): Promise<string | null> {
    const span = this.otel.tracer('Repository.CreateCartItem', context);
    try {
      const resp = new this.cartModel({
        ...request,
        created_at: new Date(),
        updated_at: new Date(),
      });
      await resp.save();
      return resp._id.toString();
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      return null;
    } finally {
      span.end();
    }
  }

  async FindCartItemById(requestId: string, id: string, context?: Context): Promise<CartItem | null> {
    const span = this.otel.tracer('Repository.FindCartItemById', context);
    try {
      return await this.cartModel.findOne({ _id: id });
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      return null;
    } finally {
      span.end();
    }
  }

  async FindCartItemByProductId(requestId: string, productId: string, context?: Context): Promise<CartItem | null> {
    const span = this.otel.tracer('Repository.FindCartItemByProductId', context);
    try {
      return await this.cartModel.findOne({ productId: productId });
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      return null;
    } finally {
      span.end();
    }
  }

  async FindCartItemsWithPagination(requestId: string, request: FindCartItemsWithPaginationRequest, context?: Context): Promise<FindCartItemsWithPaginationResponse> {
    const span = this.otel.tracer('Repository.FindCartItemsWithPagination', context);
    try {
      const skip = (request.page - 1) * request.limit;
      const query: mongoose.FilterQuery<CartDocument> = {};
      if (request.userId && request.userId.trim() !== '') {
        query.userId = request.userId;
      }

      if (request.productIds && request.productIds.length > 0) {
        query.productId = {
          '$in': request.productIds,
        };
      }

      const cartItems = await this.cartModel
        .find(query)
        .skip(skip)
        .limit(request.limit)
        .exec();

      const total = await this.cartModel.countDocuments().exec();

      return FindCartItemsWithPaginationResponse.create({
        items: cartItems,
        total: total,
        page: request.page,
        limit: request.limit,
      });

    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }

  async UpdateCartItemById(requestId: string, id: string, request: CreateCartItemRequest, context?: Context): Promise<CartItem | null> {
    const span = this.otel.tracer('Repository.UpdateCartItemById', context);
    try {
      return await this.cartModel.findByIdAndUpdate(id, request);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }

  async DeleteCartItemById(requestId: string, id: string, context?: Context) {
    const span = this.otel.tracer('Repository.DeleteCartItemById', context);
    try {
      await this.cartModel.findByIdAndDelete(id);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }
}