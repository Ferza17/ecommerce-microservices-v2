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
} from '../../model/rpc/cartMessage';
import { RabbitmqInfrastructure } from '../../infrastructure/rabbitmq/rabbitmq';
import { EventStore } from '../../model/rpc/eventStoreMessage';
import { SagaStatus } from '../../enum/sagaStatus';
import { Service } from '../../enum/service';
import { Queue } from '../../enum/queue';
import { ProductRpcService } from '../../infrastructure/rpc/product.rpc.service';
import { FindProductByIdRequest } from '../../model/rpc/productMessage';
import { UserRpcService } from '../../infrastructure/rpc/user.rpc.service';

@Injectable()
export class CartService {
  private readonly logger = new Logger(CartService.name);

  constructor(
    private readonly cartItemRepository: CartMongodbRepository,
    private readonly rabbitMQInfrastructure: RabbitmqInfrastructure,
    private readonly productRpcService: ProductRpcService,
    private readonly userRpcService: UserRpcService,
  ) {
  }

  async CreateCartItem(requestId: string, req: CreateCartItemRequest): Promise<CreateCartItemResponse> {
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
      const user = await this.userRpcService.findUserById(
        requestId,
        FindProductByIdRequest.create({ id: req.userId }),
      );

      if (user === null) {
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

      // Find If Product in Cart Item Already Created
      const cartItem = await this.cartItemRepository.FindCartItemByProductId(requestId, req.productId);
      if (cartItem !== null) {
        await this.cartItemRepository.UpdateCartItemById(requestId, cartItem._id.toString(), {
          price: product.price * (req.qty + cartItem.qty),
          qty: req.qty + cartItem.qty,
          userId: req.userId,
          productId: req.productId,
        });
        id = cartItem._id.toString() || '';
      } else {
        req.price = product.price * req.qty;
        const result = await this.cartItemRepository.CreateCartItem(requestId, req);
        if (result !== null) {
        }
      }

      event.status = SagaStatus.SUCCESS.toString();
      event.previousState = cartItem || undefined;
      await this.rabbitMQInfrastructure.publishEventCreated(requestId, event);
      return { id: id };
    } catch (e) {
      event.status = SagaStatus.FAILED.toString();
      await this.rabbitMQInfrastructure.publishEventCreated(requestId, event);
      throw e;
    }
  }

  async FindCartItemById(requestId: string, req: FindCartItemByIdRequest): Promise<CartItem> {
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
      throw e;
    }
  }

  async FindCartItemsWithPagination(requestId: string, req: FindCartItemsWithPaginationRequest): Promise<FindCartItemsWithPaginationResponse> {
    try {
      return await this.cartItemRepository.FindCartItemsWithPagination(requestId, req);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      throw e;
    }
  }

  async UpdateCartItemByIdRequest(requestId: string, req: UpdateCartItemByIdRequest): Promise<UpdateCartItemByIdResponse> {
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
      if (user === null) {
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
      throw e;
    }
  }

  async DeleteCartItemById(requestId: string, req: DeleteCartItemByIdRequest): Promise<DeleteCartItemByIdResponse> {
    try {
      await this.cartItemRepository.DeleteCartItemById(requestId, req.id);
      return {
        message: 'Cart item deleted successfully',
      };
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      throw e;
    }
  }
}
