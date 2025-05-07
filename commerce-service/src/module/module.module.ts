import { Module } from '@nestjs/common';
import { CartModule } from './cart/cart.module';
import { OrderModule } from './order/order.module';
import { WishlistModule } from './wishlist/wishlist.module';

@Module({
  exports: [],
  imports: [CartModule, OrderModule, WishlistModule],
})
export class ModuleModule {}
