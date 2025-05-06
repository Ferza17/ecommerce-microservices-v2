// cartItem.model.ts

import { Schema, model, Document } from 'mongoose';

export interface ICartItem extends Document {
  id: string;
  productId: string;
  userId: string;
  qty: string;
  price: string;
  created_at: Date;
  updated_at: Date;
}

const CartItemSchema = new Schema<ICartItem>(
  {
    id: { type: String, required: true },
    productId: { type: String, required: true },
    userId: { type: String, required: true },
    qty: { type: String, required: true },
    price: { type: String, required: true },
    created_at: { type: Date, required: true },
    updated_at: { type: Date, required: true },
  },
  {
    timestamps: false, // karena kamu manual handle `created_at` dan `updated_at`
    versionKey: false,
  }
);

export const CartItem = model<ICartItem>('CartItem', CartItemSchema);
