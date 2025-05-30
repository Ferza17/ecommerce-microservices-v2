import { Controller, Logger } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { CartService } from './cart.service';
import { Metadata } from '@grpc/grpc-js';
import { Header } from '../../enum/header';
import {
  CartItem,
  CreateCartItemRequest,
  CreateCartItemResponse, DeleteCartItemByIdRequest, DeleteCartItemByIdResponse,
  FindCartItemByIdRequest,
  FindCartItemsWithPaginationRequest,
  FindCartItemsWithPaginationResponse,
} from '../../model/rpc/gen/commerce/v1/cartMessage';
import { JaegerTelemetryService } from '../../infrastructure/telemetry/jaeger.telemetry.service';
import { context, propagation } from '@opentelemetry/api';


@Controller('cart-controller')
export class CartController {
  private readonly logger = new Logger(CartController.name);

  constructor(
    private readonly cartService: CartService,
    private readonly otel: JaegerTelemetryService,
  ) {
  }

  @GrpcMethod('CartService', 'CreateCartItem')
  async createCartItem(req: CreateCartItemRequest, metadata: Metadata): Promise<CreateCartItemResponse> {
    const requestId: string = metadata.get(Header.X_REQUEST_ID)[0].toString();
    const spanCtx = propagation.extract(context.active(), metadata.getMap());
    const span = this.otel.tracer('Controller.ConsumeCreateCartItem', spanCtx);
    try {
      return await this.cartService.CreateCartItem(requestId, req);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }

  @GrpcMethod('CartService', 'FindCartItemById')
  async findCartItemById(req: FindCartItemByIdRequest, metadata: Metadata): Promise<CartItem> {
    const requestId: string = metadata.get(Header.X_REQUEST_ID)[0].toString();
    const spanCtx = propagation.extract(context.active(), metadata.getMap());
    const span = this.otel.tracer('Controller.FindCartItemById', spanCtx);
    try {
      return await this.cartService.FindCartItemById(requestId, req);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }

  @GrpcMethod('CartService', 'FindCartItemsWithPagination')
  async FindCartItemsWithPagination(req: FindCartItemsWithPaginationRequest, metadata: Metadata): Promise<FindCartItemsWithPaginationResponse> {
    const requestId: string = metadata.get(Header.X_REQUEST_ID)[0].toString();
    const spanCtx = propagation.extract(context.active(), metadata.getMap());
    const span = this.otel.tracer('Controller.FindCartItemsWithPagination', spanCtx);
    try {
      return await this.cartService.FindCartItemsWithPagination(requestId, req);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }

  @GrpcMethod('CartService', 'DeleteCartItemById')
  async DeleteCartItemById(req: DeleteCartItemByIdRequest, metadata: Metadata): Promise<DeleteCartItemByIdResponse> {
    const requestId: string = metadata.get(Header.X_REQUEST_ID)[0].toString();
    const spanCtx = propagation.extract(context.active(), metadata.getMap());
    const span = this.otel.tracer('Controller.ConsumeCreateCartItem', spanCtx);
    try {
      return await this.cartService.DeleteCartItemById(requestId, req);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      span.recordException(e);
      throw e;
    } finally {
      span.end();
    }
  }

}
