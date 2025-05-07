import { Injectable, Logger } from '@nestjs/common';
import { CartItem, ICartItem } from '../../model/mongo/cart';


@Injectable()
export class CartMongodbRepository {
  private readonly logger = new Logger(CartMongodbRepository.name);

  constructor() {
  }

  async CreateCartItem(requestId: string, request: ICartItem): Promise<string> {
    try {
      const cartItem = await new CartItem(request).save();
      return await cartItem.id;
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      throw e;
    }
  }
}