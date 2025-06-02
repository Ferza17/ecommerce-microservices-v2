import { Module } from '@nestjs/common';
import { CartController } from './cart.controller';
import { CartService } from './cart.service';
import { CartMongodbRepository } from './cart.mongodb.repository';
import { CartConsumer } from './cart.consumer';
import { MongooseModule } from '@nestjs/mongoose';
import { CartItem, CartSchema } from '../../model/mongo/cart';
import { MongoDBCollection } from '../../enum/mongodbCollection';
import { ClientsModule } from '@nestjs/microservices';
import { ClientModuleAsyncConfig } from '../../infrastructure/rabbitmq/publisher.config';
import { Exchange } from '../../enum/exchange';
import { Queue } from '../../enum/queue';
import { InfrastructureModule } from '../../infrastructure/infrastructure.module';

@Module({
  imports: [
    InfrastructureModule,
    MongooseModule.forFeature([
      { name: CartItem.name, schema: CartSchema, collection: MongoDBCollection.CartItemCollection },
    ]),
  ],
  controllers: [CartController, CartConsumer],
  providers: [CartService, CartMongodbRepository],
})
export class CartModule {
}
