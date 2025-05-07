import { Controller, Logger } from '@nestjs/common';
import { Ctx, EventPattern, MessagePattern, Payload, RmqContext } from '@nestjs/microservices';
import { CartService } from './cart.service';
import { Header } from '../../enum/header';
import {
  CreateCartItemRequest,
  CreateCartItemResponse,
  UpdateCartItemByIdRequest,
  UpdateCartItemByIdResponse,
} from '../../model/rpc/cartMessage';
import { RoutingKey } from '../../enum/routingKey';


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
    const pattern = ctx.getPattern();
    this.logger.log(`requestId: ${requestId} , data: ${JSON.stringify(data)}`);
    this.logger.log(`Received pattern: ${pattern}`);
  }


  @MessagePattern(RoutingKey.CART_CREATED)
  async createCartItem(@Payload() data: CreateCartItemRequest, @Ctx() context: RmqContext): Promise<CreateCartItemResponse> {
    const { properties: { headers } } = context.getMessage();
    const requestId: string = headers[Header.X_REQUEST_ID];
    this.logger.log(`requestId: ${requestId} , data: ${JSON.stringify(data)}`);
    this.logger.log(`createCartItem`);
    try {
      return await this.cartService.createCartItem(requestId, data);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      throw e;
    }
  }

  @MessagePattern(RoutingKey.CART_UPDATED)
  async updateCartItemByIdRequest(@Payload() data: UpdateCartItemByIdRequest, @Ctx() context: RmqContext): Promise<UpdateCartItemByIdResponse> {
    const { properties: { headers } } = context.getMessage();
    const requestId: string = headers[Header.X_REQUEST_ID];
    this.logger.log(`requestId: ${requestId} , data: ${JSON.stringify(data)}`);
    this.logger.log(`updateCartItemByIdRequest`);
    try {
      return await this.cartService.updateCartItemByIdRequest(requestId, data);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      throw e;
    }
  }
}
