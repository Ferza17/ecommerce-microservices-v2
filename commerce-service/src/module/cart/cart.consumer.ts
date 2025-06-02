import { Controller, Logger } from '@nestjs/common';
import { Ctx, MessagePattern, Payload, RmqContext } from '@nestjs/microservices';
import { Header } from '../../enum/header';
import {
  CreateCartItemRequest,
  CreateCartItemResponse, DeleteCartItemByIdRequest, DeleteCartItemByIdResponse,
  UpdateCartItemByIdRequest,
  UpdateCartItemByIdResponse,
} from '../../model/rpc/gen/commerce/v1/cartMessage';
import { CartService } from './cart.service';
import { Queue } from '../../enum/queue';
import { JaegerTelemetryService } from '../../infrastructure/telemetry/jaeger.telemetry.service';
import { propagation, context } from '@opentelemetry/api';


@Controller()
export class CartConsumer {
  private readonly logger = new Logger(CartConsumer.name);

  constructor(
    private readonly cartService: CartService,
    private readonly otel: JaegerTelemetryService,
  ) {
  }

  @MessagePattern(`${Queue.CART_CREATED}`)
  async consumeCreateCartItem(@Payload() data: CreateCartItemRequest, @Ctx() ctx: RmqContext): Promise<CreateCartItemResponse> {
    const { properties: { headers } } = ctx.getMessage();
    const requestId: string = headers[Header.X_REQUEST_ID];
    const spanCtx = propagation.extract(context.active(), headers);
    const span = this.otel.tracer('Consumer.ConsumeCreateCartItem', spanCtx);
    try {
      return await this.cartService.CreateCartItem(requestId, data, spanCtx);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }

  @MessagePattern(`${Queue.CART_UPDATED}`)
  async updateCartItemByIdRequest(@Payload() data: UpdateCartItemByIdRequest, @Ctx() ctx: RmqContext): Promise<UpdateCartItemByIdResponse> {
    const { properties: { headers } } = ctx.getMessage();
    const requestId: string = headers[Header.X_REQUEST_ID];
    const spanCtx = propagation.extract(context.active(), headers);
    const span = this.otel.tracer('Consumer.UpdateCartItemByIdRequest', spanCtx);
    try {
      return await this.cartService.UpdateCartItemByIdRequest(requestId, <UpdateCartItemByIdRequest>data);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }

  @MessagePattern(`${Queue.CART_DELETED}`)
  async deleteCartItemByIdRequest(@Payload() data: DeleteCartItemByIdRequest, @Ctx() ctx: RmqContext): Promise<DeleteCartItemByIdResponse> {
    const { properties: { headers } } = ctx.getMessage();
    const requestId: string = headers[Header.X_REQUEST_ID];
    const spanCtx = propagation.extract(context.active(), headers);
    const span = this.otel.tracer('Consumer.DeleteCartItemByIdRequest', spanCtx);
    try {
      return await this.cartService.DeleteCartItemById(requestId, data);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }
}