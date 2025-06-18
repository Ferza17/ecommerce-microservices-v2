import { GrpcOptions } from '@nestjs/microservices';
import { ConsulService } from '../../config/consul.service';
export declare class GrpcClientOptions {
    private readonly consulConfig;
    constructor(consulConfig: ConsulService);
    getGRPCConfig(): Promise<GrpcOptions>;
}
