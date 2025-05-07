import { Controller, Logger } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { CartService } from './cart.service';
import { Metadata } from '@grpc/grpc-js';
import { Header } from '../../enum/header';
import { CreateCartItemRequest, CreateCartItemResponse } from '../../model/rpc/cartMessage';


@Controller('cart-controller')
export class CartController {
  private readonly logger = new Logger(CartController.name);

  constructor(
    private readonly cartService: CartService,
  ) {
  }

  @GrpcMethod('CartService', 'CreateCartItem')
  async createCartItem(req: CreateCartItemRequest, metadata: Metadata): Promise<CreateCartItemResponse> {
    const requestId: string = metadata.get(Header.X_REQUEST_ID)[0].toString();
    try {
      return await this.cartService.createCartItem(requestId, req);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      throw e;
    }
  }
}
