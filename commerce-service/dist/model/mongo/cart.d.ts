import mongoose, { HydratedDocument } from 'mongoose';
export type CartDocument = HydratedDocument<CartItem>;
export declare class CartItem {
    _id: mongoose.Types.ObjectId;
    productId: string;
    userId: string;
    qty: number;
    price: number;
    created_at: Date;
    updated_at: Date;
}
export declare const CartSchema: mongoose.Schema<CartItem, mongoose.Model<CartItem, any, any, any, mongoose.Document<unknown, any, CartItem, any> & CartItem & Required<{
    _id: mongoose.Types.ObjectId;
}> & {
    __v: number;
}, any>, {}, {}, {}, {}, mongoose.DefaultSchemaOptions, CartItem, mongoose.Document<unknown, {}, mongoose.FlatRecord<CartItem>, {}> & mongoose.FlatRecord<CartItem> & Required<{
    _id: mongoose.Types.ObjectId;
}> & {
    __v: number;
}>;
