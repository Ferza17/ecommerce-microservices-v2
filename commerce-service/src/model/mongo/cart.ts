import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { HydratedDocument } from 'mongoose';

export type CartDocument = HydratedDocument<CartItem>;

@Schema()
export class CartItem {
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
