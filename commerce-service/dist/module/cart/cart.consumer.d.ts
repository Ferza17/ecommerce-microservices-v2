import { RmqContext } from '@nestjs/microservices';
import { CreateCartItemRequest, CreateCartItemResponse, DeleteCartItemByIdRequest, DeleteCartItemByIdResponse, UpdateCartItemByIdRequest, UpdateCartItemByIdResponse } from '../../model/rpc/gen/commerce/v1/cartMessage';
import { CartService } from './cart.service';
import { JaegerTelemetryService } from '../../infrastructure/telemetry/jaeger.telemetry.service';
export declare class CartConsumer {
    private readonly cartService;
    private readonly otel;
    private readonly logger;
    constructor(cartService: CartService, otel: JaegerTelemetryService);
    consumeCreateCartItem(data: CreateCartItemRequest, ctx: RmqContext): Promise<CreateCartItemResponse>;
    updateCartItemByIdRequest(data: UpdateCartItemByIdRequest, ctx: RmqContext): Promise<UpdateCartItemByIdResponse>;
    deleteCartItemByIdRequest(data: DeleteCartItemByIdRequest, ctx: RmqContext): Promise<DeleteCartItemByIdResponse>;
}
