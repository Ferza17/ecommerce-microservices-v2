import { Controller } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { CartService } from './cart.service';
import { Metadata } from '@grpc/grpc-js';
import { Header } from '../../enum/header';
import { CartItem } from '../../model/rpc/cartMessage';


@Controller('cart')
export class CartController {
  constructor(private readonly cartService: CartService) {
  }

  @GrpcMethod('CartService', 'CreateCartItem')
  createCartItem(req: CartItem, metadata: Metadata): any {
    return {
      id: '12',
    };
  }
}
