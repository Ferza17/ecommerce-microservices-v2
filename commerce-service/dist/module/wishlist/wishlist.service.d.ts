import { ProductRpcService } from '../../infrastructure/rpc/product.rpc.service';
import { UserRpcService } from '../../infrastructure/rpc/user.rpc.service';
import { RabbitmqInfrastructure } from '../../infrastructure/rabbitmq/rabbitmq';
export declare class WishlistService {
    private readonly rabbitMQInfrastructure;
    private readonly productRpcService;
    private readonly userRpcService;
    private readonly logger;
    constructor(rabbitMQInfrastructure: RabbitmqInfrastructure, productRpcService: ProductRpcService, userRpcService: UserRpcService);
}
