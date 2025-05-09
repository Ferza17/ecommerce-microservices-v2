import { Injectable, Logger } from '@nestjs/common';
import { CartMongodbRepository } from './cart.mongodb.repository';
import {
  CreateCartItemResponse,
  CreateCartItemRequest,
  UpdateCartItemByIdRequest,
  UpdateCartItemByIdResponse,
} from '../../model/rpc/cartMessage';
import { RabbitmqInfrastructure } from '../../infrastructure/rabbitmq/rabbitmq';
import { EventStore } from '../../model/rpc/eventStoreMessage';
import { SagaStatus } from '../../enum/sagaStatus';
import { Service } from '../../enum/service';
import { Queue } from '../../enum/queue';

@Injectable()
export class CartService {
  private readonly logger = new Logger(CartService.name);

  constructor(
    private readonly cartItemRepository: CartMongodbRepository,
    private readonly rabbitMQInfrastructure: RabbitmqInfrastructure,
  ) {
  }

  async createCartItem(requestId: string, req: CreateCartItemRequest): Promise<CreateCartItemResponse> {
    let id = '';
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
      const result = await this.cartItemRepository.CreateCartItem(requestId, req);
      if (result !== null) {
        id = result;
      }
      event.status = SagaStatus.SUCCESS.toString();
      await this.rabbitMQInfrastructure.publishEventCreated(requestId,event);
    } catch (e) {
      event.status = SagaStatus.FAILED.toString();
      await this.rabbitMQInfrastructure.publishEventCreated(requestId,event);
      throw e;
    }
    return { id };
  }

  async updateCartItemByIdRequest(requestId: string, req: UpdateCartItemByIdRequest): Promise<UpdateCartItemByIdResponse> {
    try {
      // Validate Id

    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      throw e;
    }

    return { id: req.id };
  }
}
