import { Injectable, Logger } from '@nestjs/common';
import { CartDocument, CartItem } from '../../model/mongo/cart';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import {
  CreateCartItemRequest,
  FindCartItemsWithPaginationRequest,
  FindCartItemsWithPaginationResponse,
} from '../../model/rpc/cartMessage';


@Injectable()
export class CartMongodbRepository {
  private readonly logger = new Logger(CartMongodbRepository.name);

  constructor(
    @InjectModel(CartItem.name)
    private cartModel: Model<CartDocument>) {
  }

  async CreateCartItem(requestId: string, request: CreateCartItemRequest): Promise<string | null> {
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
      return null;
    }
  }

  async FindCartItemById(requestId: string, id: string): Promise<CartItem | null> {
    try {
      return await this.cartModel.findOne({ _id: id });
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      return null;
    }
  }

  async FindCartItemByProductId(requestId: string, productId: string): Promise<CartItem | null> {
    try {
      return await this.cartModel.findOne({ productId: productId });
    }catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      return null;
    }
  }

  async FindCartItemsWithPagination(requestId: string, request: FindCartItemsWithPaginationRequest): Promise<FindCartItemsWithPaginationResponse> {
    try {
      const skip = (request.page - 1) * request.limit;

      const cartItems = await this.cartModel
        .find()
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
      throw e;
    }
  }

  async UpdateCartItemById(requestId: string, id: string, request: CreateCartItemRequest): Promise<CartItem | null> {
    try {
      return await this.cartModel.findByIdAndUpdate(id, request);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      throw e;
    }
  }

  async DeleteCartItemById(requestId: string, id: string) {
    try {
      await this.cartModel.findByIdAndDelete(id);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      throw e;
    }
  }
}