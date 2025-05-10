import { Injectable, Logger } from '@nestjs/common';
import { ProductRpcService } from '../../infrastructure/rpc/product.rpc.service';
import { UserRpcService } from '../../infrastructure/rpc/user.rpc.service';
import { RabbitmqInfrastructure } from '../../infrastructure/rabbitmq/rabbitmq';

@Injectable()
export class WishlistService {
  private readonly logger = new Logger(WishlistService.name);

  constructor(
    private readonly rabbitMQInfrastructure: RabbitmqInfrastructure,
    private readonly productRpcService: ProductRpcService,
    private readonly userRpcService: UserRpcService,
  ) {}
}
