import { Module } from '@nestjs/common';
import { OrderService } from './order.service';
import { OrderController } from './order.controller';
import { InfrastructureModule } from '../../infrastructure/infrastructure.module';

@Module({
  imports: [InfrastructureModule],
  providers: [OrderService],
  controllers: [OrderController],
})
export class OrderModule {}
