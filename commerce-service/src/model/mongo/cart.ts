// cartItem.model.ts

import { Schema, model, Document } from 'mongoose';
import { MongoDBCollection } from '../../enum/mongodbCollection';

export interface ICartItem extends Document {
  _id: string;
  productId: string;
  userId: string;
  qty: number;
  price: number;
  created_at: Date;
  updated_at: Date;
}

const CartItemSchema = new Schema<ICartItem>(
  {
    productId: { type: String, required: true },
    userId: { type: String, required: true },
    qty: { type: Number, required: true },
    price: { type: Number, required: true },
    created_at: { type: Date, required: true },
    updated_at: { type: Date, required: true },
  },
  {
    timestamps: true,
    versionKey: false,
    collection: MongoDBCollection.CartItemCollection,
  },
);

export const CartItem = model<ICartItem>(MongoDBCollection.CartItemCollection, CartItemSchema);