import { Injectable, Logger } from '@nestjs/common';
import { CartMongodbRepository } from './cart.mongodb.repository';
import {
  CreateCartItemResponse,
  CreateCartItemRequest,
  UpdateCartItemByIdRequest,
  UpdateCartItemByIdResponse,
} from '../../model/rpc/cartMessage';
import { ICartItem } from '../../model/mongo/cart';

@Injectable()
export class CartService {
  private readonly logger = new Logger(CartService.name);

  constructor(
    private readonly cartItemRepository: CartMongodbRepository,
  ) {
  }

  async createCartItem(requestId: string, req: CreateCartItemRequest): Promise<CreateCartItemResponse> {

    let id: string = '';
    try {
      const cartItem: ICartItem = {
        productId: req.productId,
        userId: req.userId,
        qty: req.qty,
        price: req.price,
        created_at: new Date(),
        updated_at: new Date(),
      } as ICartItem;
      id = await this.cartItemRepository.CreateCartItem(requestId, cartItem);
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
    }

    return { id };
  }

  async updateCartItemByIdRequest(requestId: string, req: UpdateCartItemByIdRequest): Promise<UpdateCartItemByIdResponse> {
    return { id: req.id };
  }
}
