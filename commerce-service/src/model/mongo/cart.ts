import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import mongoose, { HydratedDocument } from 'mongoose';

export type CartDocument = HydratedDocument<CartItem>;

@Schema()
export class CartItem {
  _id: mongoose.Types.ObjectId;

  @Prop()
  productId: string;

  @Prop()
  userId: string;

  @Prop()
  qty: number;

  @Prop()
  price: number;

  @Prop({ default: Date.now })
  created_at: Date;

  @Prop({ default: Date.now })
  updated_at: Date;
}

export const CartSchema = SchemaFactory.createForClass(CartItem);
