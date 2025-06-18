import { OnModuleInit } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { FindProductByIdRequest, Product } from '../../model/rpc/gen/product/v1/productMessage';
import { JaegerTelemetryService } from '../telemetry/jaeger.telemetry.service';
import { Context } from '@opentelemetry/api';
export declare class ProductRpcService implements OnModuleInit {
    private client;
    private readonly otel;
    private readonly logger;
    private productService;
    constructor(client: ClientGrpc, otel: JaegerTelemetryService);
    onModuleInit(): void;
    findProductById(requestId: string, req: FindProductByIdRequest, context?: Context): Promise<Product>;
}
