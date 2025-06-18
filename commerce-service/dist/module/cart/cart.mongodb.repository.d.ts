import { CartDocument, CartItem } from '../../model/mongo/cart';
import { Model } from 'mongoose';
import { CreateCartItemRequest, FindCartItemsWithPaginationRequest, FindCartItemsWithPaginationResponse } from '../../model/rpc/gen/commerce/v1/cartMessage';
import { JaegerTelemetryService } from '../../infrastructure/telemetry/jaeger.telemetry.service';
import { Context } from '@opentelemetry/api';
export declare class CartMongodbRepository {
    private cartModel;
    private readonly otel;
    private readonly logger;
    constructor(cartModel: Model<CartDocument>, otel: JaegerTelemetryService);
    CreateCartItem(requestId: string, request: CreateCartItemRequest, context?: Context): Promise<string | null>;
    FindCartItemById(requestId: string, id: string, context?: Context): Promise<CartItem | null>;
    FindCartItemByProductId(requestId: string, productId: string, context?: Context): Promise<CartItem | null>;
    FindCartItemsWithPagination(requestId: string, request: FindCartItemsWithPaginationRequest, context?: Context): Promise<FindCartItemsWithPaginationResponse>;
    UpdateCartItemById(requestId: string, id: string, request: CreateCartItemRequest, context?: Context): Promise<CartItem | null>;
    DeleteCartItemById(requestId: string, id: string, context?: Context): Promise<void>;
}
