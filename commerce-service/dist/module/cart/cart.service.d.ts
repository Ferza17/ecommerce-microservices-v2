import { CartMongodbRepository } from './cart.mongodb.repository';
import { CartItem, CreateCartItemRequest, CreateCartItemResponse, DeleteCartItemByIdRequest, DeleteCartItemByIdResponse, FindCartItemByIdRequest, FindCartItemsWithPaginationRequest, FindCartItemsWithPaginationResponse, UpdateCartItemByIdRequest, UpdateCartItemByIdResponse } from '../../model/rpc/gen/commerce/v1/cartMessage';
import { RabbitmqInfrastructure } from '../../infrastructure/rabbitmq/rabbitmq';
import { ProductRpcService } from '../../infrastructure/rpc/product.rpc.service';
import { UserRpcService } from '../../infrastructure/rpc/user.rpc.service';
import { JaegerTelemetryService } from '../../infrastructure/telemetry/jaeger.telemetry.service';
import { Context } from '@opentelemetry/api';
export declare class CartService {
    private readonly cartItemRepository;
    private readonly rabbitMQInfrastructure;
    private readonly productRpcService;
    private readonly userRpcService;
    private readonly otel;
    private readonly logger;
    constructor(cartItemRepository: CartMongodbRepository, rabbitMQInfrastructure: RabbitmqInfrastructure, productRpcService: ProductRpcService, userRpcService: UserRpcService, otel: JaegerTelemetryService);
    CreateCartItem(requestId: string, req: CreateCartItemRequest, context?: Context): Promise<CreateCartItemResponse>;
    FindCartItemById(requestId: string, req: FindCartItemByIdRequest, context?: Context): Promise<CartItem>;
    FindCartItemsWithPagination(requestId: string, req: FindCartItemsWithPaginationRequest, context?: Context): Promise<FindCartItemsWithPaginationResponse>;
    UpdateCartItemByIdRequest(requestId: string, req: UpdateCartItemByIdRequest, context?: Context): Promise<UpdateCartItemByIdResponse>;
    DeleteCartItemById(requestId: string, req: DeleteCartItemByIdRequest, context?: Context): Promise<DeleteCartItemByIdResponse>;
}
