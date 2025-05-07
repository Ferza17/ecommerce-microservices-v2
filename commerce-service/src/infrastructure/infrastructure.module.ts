import { Module } from '@nestjs/common';
import { MongodbInfrastructure } from './mongodb/mongodb.infrastructure';
import { RabbitmqInfrastructure } from './rabbitmq/rabbitmq.infrastructure';
import { RpcInfrastructure } from './rpc/rpc.infrastructure';
@Module({
  providers: [MongodbInfrastructure, RabbitmqInfrastructure, RpcInfrastructure]
})
export class InfrastructureModule {}
