import { Module } from '@nestjs/common';
import { CartModule } from './cart/cart.module';
import { WishlistModule } from './wishlist/wishlist.module';
import { InfrastructureModule } from '../infrastructure/infrastructure.module';

@Module({
  exports: [],
  imports: [
    CartModule,
    WishlistModule,
    InfrastructureModule,
  ],
})
export class ModuleModule {
}
