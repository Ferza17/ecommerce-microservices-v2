import { OnModuleInit } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
export declare class ConsulService implements OnModuleInit {
    private readonly configService;
    private readonly logger;
    private consul;
    constructor(configService: ConfigService);
    onModuleInit(): Promise<void>;
    get(key: string): Promise<string>;
}
