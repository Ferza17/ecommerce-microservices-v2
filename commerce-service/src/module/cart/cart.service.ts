import { Injectable, Logger } from '@nestjs/common';
import { CartMongodbRepository } from './cart.mongodb.repository';
import {
  CartItem,
  CreateCartItemRequest,
  CreateCartItemResponse, DeleteCartItemByIdRequest, DeleteCartItemByIdResponse,
  FindCartItemByIdRequest,
  FindCartItemsWithPaginationRequest,
  FindCartItemsWithPaginationResponse,
  UpdateCartItemByIdRequest,
  UpdateCartItemByIdResponse,
} from '../../model/rpc/gen/commerce/v1/cartMessage';
import { RabbitmqInfrastructure } from '../../infrastructure/rabbitmq/rabbitmq';
import { EventStore } from '../../model/rpc/gen/event/v1/eventStoreMessage';
import { SagaStatus } from '../../enum/sagaStatus';
import { Service } from '../../enum/service';
import { Queue } from '../../enum/queue';
import { ProductRpcService } from '../../infrastructure/rpc/product.rpc.service';
import { FindProductByIdRequest, Product } from '../../model/rpc/gen/product/v1/productMessage';
import { UserRpcService } from '../../infrastructure/rpc/user.rpc.service';
import { JaegerTelemetryService } from '../../infrastructure/telemetry/jaeger.telemetry.service';
import { Context } from '@opentelemetry/api';
import { FindUserByIdRequest, User } from '../../model/rpc/gen/user/v1/userMessage';

@Injectable()
export class CartService {
  private readonly logger = new Logger(CartService.name);

  constructor(
    private readonly cartItemRepository: CartMongodbRepository,
    private readonly rabbitMQInfrastructure: RabbitmqInfrastructure,
    private readonly productRpcService: ProductRpcService,
    private readonly userRpcService: UserRpcService,
    private readonly otel: JaegerTelemetryService,
  ) {
  }

  async CreateCartItem(requestId: string, req: CreateCartItemRequest, context?: Context): Promise<CreateCartItemResponse> {
    const span = this.otel.tracer('Service.CreateCartItem', context);
    let id: string = '';
    let event: EventStore = {
      service: Service.CommerceService.toString(),
      eventType: Queue.CART_CREATED.toString(),
      requestId: requestId,
      payload: req,
      updatedAt: new Date(),
      createdAt: new Date(),
      id: '',
      status: '',
    };

    try {
      // Validate User
      const user: User = await this.userRpcService.findUserById(
        requestId,
        FindUserByIdRequest.create({ id: req.userId }),
        context,
      );

      if (user === null) {
        throw new Error('User not found');
      }

      // Validate Product
      const product: Product = await this.productRpcService.findProductById(
        requestId,
        FindProductByIdRequest.create({ id: req.productId }),
        context,
      );

      if (product === null) {
        throw new Error('Product not found');
      }

      if (product.stock < req.qty) {
        throw new Error('Product stock is not enough');
      }

      // Find If Product in Cart Item Already Created
      const cartItem = await this.cartItemRepository.FindCartItemByProductId(requestId, req.productId);
      if (cartItem !== null) {
        await this.cartItemRepository.UpdateCartItemById(requestId, cartItem._id.toString(), {
          price: product.price * (req.qty + cartItem.qty),
          qty: req.qty + cartItem.qty,
          userId: user.id,
          productId: req.productId,
        }, context);
        id = cartItem._id.toString() || '';
      } else {
        req.price = product.price * req.qty;
        const result = await this.cartItemRepository.CreateCartItem(requestId, req, context);
        if (result !== null) {
        }
      }

      event.status = SagaStatus.SUCCESS.toString();
      event.previousState = cartItem || undefined;
      await this.rabbitMQInfrastructure.publishEventCreated(requestId, event, context);
      return { id: id };
    } catch (e) {
      event.status = SagaStatus.FAILED.toString();
      await this.rabbitMQInfrastructure.publishEventCreated(requestId, event, context);
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }

  async FindCartItemById(requestId: string, req: FindCartItemByIdRequest, context?: Context): Promise<CartItem> {
    const span = this.otel.tracer('Service.FindCartItemById', context);
    try {
      const cartItem = await this.cartItemRepository.FindCartItemById(requestId, req.id);
      if (!cartItem) {
        throw new Error('Cart item not found');
      }
      return {
        id: cartItem._id.toString(),
        productId: cartItem.productId,
        userId: cartItem.userId,
        qty: cartItem.qty,
        price: cartItem.price,
        cratedAt: cartItem.created_at,
        updatedAt: cartItem.updated_at,
      };
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }

  async FindCartItemsWithPagination(requestId: string, req: FindCartItemsWithPaginationRequest, context?: Context): Promise<FindCartItemsWithPaginationResponse> {
    const span = this.otel.tracer('Service.FindCartItemsWithPagination', context);
    try {
      return await this.cartItemRepository.FindCartItemsWithPagination(requestId, req);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }

  async UpdateCartItemByIdRequest(requestId: string, req: UpdateCartItemByIdRequest, context?: Context): Promise<UpdateCartItemByIdResponse> {
    const span = this.otel.tracer('Service.UpdateCartItemByIdRequest', context);
    let event: EventStore = {
      service: Service.CommerceService.toString(),
      eventType: Queue.CART_UPDATED.toString(),
      requestId: requestId,
      payload: req,
      updatedAt: new Date(),
      createdAt: new Date(),
      id: '',
      status: '',
    };
    try {
      // Validate User
      const user = await this.userRpcService.findUserById(
        requestId,
        FindProductByIdRequest.create({ id: req.userId }),
      );
      if (user === null || user === undefined) {
        throw new Error('User not found');
      }

      // Validate Product
      const product = await this.productRpcService.findProductById(
        requestId,
        FindProductByIdRequest.create({ id: req.productId }),
      );
      if (product === null) {
        throw new Error('Product not found');
      }
      if (product.stock < req.qty) {
        throw new Error('Product stock is not enough');
      }
      req.price = product.price * req.qty;

      // Validate Cart Item ID
      const cartItem = await this.cartItemRepository.FindCartItemById(requestId, req.id);
      if (cartItem === null) {
        throw new Error('Cart item not found');
      }

      await this.cartItemRepository.UpdateCartItemById(requestId, req.id, req);
      event.status = SagaStatus.SUCCESS.toString();
      await this.rabbitMQInfrastructure.publishEventCreated(requestId, event);
      return { id: req.id };
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      event.status = SagaStatus.FAILED.toString();
      await this.rabbitMQInfrastructure.publishEventCreated(requestId, event);
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }

  async DeleteCartItemById(requestId: string, req: DeleteCartItemByIdRequest, context?: Context): Promise<DeleteCartItemByIdResponse> {
    const span = this.otel.tracer('Service.DeleteCartItemById', context);
    try {
      await this.cartItemRepository.DeleteCartItemById(requestId, req.id);
      return {
        message: 'Cart item deleted successfully',
      };
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }
}
