import { Module } from '@nestjs/common';
import { CartController } from './cart.controller';
import { CartService } from './cart.service';
import { CartMongodbRepository } from './cart.mongodb.repository';
import { CartConsumer } from './cart.consumer';

@Module({
  controllers: [CartController, CartConsumer],
  providers: [CartService, CartMongodbRepository],
})
export class CartModule {
}
