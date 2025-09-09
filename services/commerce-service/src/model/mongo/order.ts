import mongoose, { HydratedDocument } from 'mongoose';
import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { CartItem } from './cart';

export type OrderDocument = HydratedDocument<OrderItem>;

@Schema()
export class OrderItem {
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
  crated_at?: Date;
  @Prop({ default: Date.now })
  updated_at?: Date;
}

export const OrderSchema = SchemaFactory.createForClass(OrderItem);
