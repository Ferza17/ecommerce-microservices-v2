import { GrpcClientOptions } from './options';
export declare class GrpcServer {
    private readonly grpcClientOptions;
    private readonly logger;
    constructor(grpcClientOptions: GrpcClientOptions);
    Serve(): Promise<void>;
}
