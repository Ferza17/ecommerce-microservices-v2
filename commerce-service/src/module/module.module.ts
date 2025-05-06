import { Module } from '@nestjs/common';
import { CartModule } from './cart/cart.module';
import { OrderService } from './order/order.service';
import { OrderModule } from './order/order.module';
import { WishlistModule } from './wishlist/wishlist.module';

@Module({
  exports: [],
  imports: [CartModule, OrderModule, WishlistModule],
  providers: [OrderService],
})
export class ModuleModule {}
