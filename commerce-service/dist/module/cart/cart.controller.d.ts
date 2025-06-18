import { CartService } from './cart.service';
import { Metadata } from '@grpc/grpc-js';
import { CartItem, CreateCartItemRequest, CreateCartItemResponse, DeleteCartItemByIdRequest, DeleteCartItemByIdResponse, FindCartItemByIdRequest, FindCartItemsWithPaginationRequest, FindCartItemsWithPaginationResponse } from '../../model/rpc/gen/commerce/v1/cartMessage';
import { JaegerTelemetryService } from '../../infrastructure/telemetry/jaeger.telemetry.service';
export declare class CartController {
    private readonly cartService;
    private readonly otel;
    private readonly logger;
    constructor(cartService: CartService, otel: JaegerTelemetryService);
    createCartItem(req: CreateCartItemRequest, metadata: Metadata): Promise<CreateCartItemResponse>;
    findCartItemById(req: FindCartItemByIdRequest, metadata: Metadata): Promise<CartItem>;
    FindCartItemsWithPagination(req: FindCartItemsWithPaginationRequest, metadata: Metadata): Promise<FindCartItemsWithPaginationResponse>;
    DeleteCartItemById(req: DeleteCartItemByIdRequest, metadata: Metadata): Promise<DeleteCartItemByIdResponse>;
}
