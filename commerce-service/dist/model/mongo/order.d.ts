import mongoose, { HydratedDocument } from 'mongoose';
export type OrderDocument = HydratedDocument<OrderItem>;
export declare class OrderItem {
    _id: mongoose.Types.ObjectId;
    productId: string;
    userId: string;
    qty: number;
    price: number;
    crated_at?: Date;
    updated_at?: Date;
}
export declare const OrderSchema: mongoose.Schema<OrderItem, mongoose.Model<OrderItem, any, any, any, mongoose.Document<unknown, any, OrderItem, any> & OrderItem & Required<{
    _id: mongoose.Types.ObjectId;
}> & {
    __v: number;
}, any>, {}, {}, {}, {}, mongoose.DefaultSchemaOptions, OrderItem, mongoose.Document<unknown, {}, mongoose.FlatRecord<OrderItem>, {}> & mongoose.FlatRecord<OrderItem> & Required<{
    _id: mongoose.Types.ObjectId;
}> & {
    __v: number;
}>;
