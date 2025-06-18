import { OnModuleInit } from '@nestjs/common';
import { Context, Span } from '@opentelemetry/api';
import { ConsulService } from '../../config/consul.service';
export declare class JaegerTelemetryService implements OnModuleInit {
    private configService;
    private readonly logger;
    private tt;
    constructor(configService: ConsulService);
    onModuleInit(): Promise<void>;
    tracer(operationName: string, ctx?: Context): Span;
}
