import { Controller, Logger, UseInterceptors } from '@nestjs/common';
import { Ctx, EventPattern, MessagePattern, Payload, RmqContext, Transport } from '@nestjs/microservices';
import { Header } from '../../enum/header';
import {
  CreateCartItemRequest,
  CreateCartItemResponse,
  UpdateCartItemByIdRequest,
  UpdateCartItemByIdResponse,
} from '../../model/rpc/cartMessage';
import { CartService } from './cart.service';
import { Queue } from '../../enum/queue';


@Controller()
export class CartConsumer {
  private readonly logger = new Logger(CartConsumer.name);

  constructor(
    private readonly cartService: CartService,
  ) {
  }

  @MessagePattern()
  handleAnyPattern(@Payload() data: any, @Ctx() ctx: RmqContext) {
    const { properties: { headers } } = ctx.getMessage();
    const requestId: string = headers[Header.X_REQUEST_ID];
    const pattern: string = ctx.getPattern();

    this.logger.log(`requestId: ${requestId} , pattern: ${pattern} , data: ${JSON.stringify(data)}`);
    this.logger.log(`pattern : `, ctx.getPattern());
  }

  @MessagePattern(`${Queue.CART_CREATED}`)
  async consumeCreateCartItem(@Payload() data: CreateCartItemRequest, @Ctx() context: RmqContext): Promise<CreateCartItemResponse> {
    const { properties: { headers } } = context.getMessage();
    const requestId: string = headers[Header.X_REQUEST_ID];
    try {
      return await this.cartService.createCartItem(requestId, data);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      throw e;
    }
  }

  @MessagePattern(`${Queue.CART_UPDATED}`)
  async updateCartItemByIdRequest(@Payload() data: UpdateCartItemByIdRequest, @Ctx() context: RmqContext): Promise<UpdateCartItemByIdResponse> {
    const { properties: { headers } } = context.getMessage();
    const requestId: string = headers[Header.X_REQUEST_ID];
    this.logger.log(`updateCartItemByIdRequest`);

    try {
      return await this.cartService.updateCartItemByIdRequest(requestId, <UpdateCartItemByIdRequest>data);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      throw e;
    }
  }


}