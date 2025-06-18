import { ConsulService } from '../../config/consul.service';
export declare class RabbitmqConsumer {
    private readonly consulConfig;
    private readonly logger;
    constructor(consulConfig: ConsulService);
    Serve(): Promise<void>;
}
