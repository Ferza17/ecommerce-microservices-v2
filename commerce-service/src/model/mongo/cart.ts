// cartItem.model.ts

import { Schema, model, Document } from 'mongoose';
import { MongoDBCollection } from '../../enum/mongodbCollection';

export interface ICartItem extends Document {
  _id: string;
  productId: string;
  userId: string;
  qty: string;
  price: string;
  created_at: Date;
  updated_at: Date;
}

const CartItemSchema = new Schema<ICartItem>(
  {
    productId: { type: String, required: true },
    userId: { type: String, required: true },
    qty: { type: String, required: true },
    price: { type: String, required: true },
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