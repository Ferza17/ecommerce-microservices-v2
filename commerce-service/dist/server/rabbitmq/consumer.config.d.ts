import { RmqOptions } from '@nestjs/microservices';
import { Queue } from '../../enum/queue';
import { ConsulService } from '../../config/consul.service';
export declare class RabbitmqOptions {
    private readonly consulConfig;
    private queue;
    constructor(consulConfig: ConsulService, queue: Queue);
    getRabbitmqOptions(): Promise<RmqOptions>;
}
