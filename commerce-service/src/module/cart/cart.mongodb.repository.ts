import { Injectable, Logger } from '@nestjs/common';
import { CartDocument, CartItem } from '../../model/mongo/cart';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { CreateCartItemRequest } from '../../model/rpc/cartMessage';


@Injectable()
export class CartMongodbRepository {
  private readonly logger = new Logger(CartMongodbRepository.name);

  constructor(
    @InjectModel(CartItem.name)
    private cartModel: Model<CartDocument>) {
  }


  async CreateCartItem(requestId: string, request: CreateCartItemRequest): Promise<string | null> {
    try {
      const resp = new this.cartModel({
        ...request,
        created_at: new Date(),
        updated_at: new Date(),
      });
      await resp.save();
      return resp._id.toString();
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      return null;
    }
  }

  async FindCartItemById(requestId: string, id: string): Promise<CartItem | null> {
    try {
      return await this.cartModel.findOne({ _id: id });
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      return null;
    }
  }
}