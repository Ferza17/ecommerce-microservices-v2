import { OnModuleInit } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { FindUserByIdRequest, User } from '../../model/rpc/gen/user/v1/userMessage';
import { JaegerTelemetryService } from '../telemetry/jaeger.telemetry.service';
import { Context } from '@opentelemetry/api';
export declare class UserRpcService implements OnModuleInit {
    private client;
    private readonly otel;
    private readonly logger;
    private userService;
    constructor(client: ClientGrpc, otel: JaegerTelemetryService);
    onModuleInit(): void;
    findUserById(requestId: string, req: FindUserByIdRequest, context?: Context): Promise<User>;
}
