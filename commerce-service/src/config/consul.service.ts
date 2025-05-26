import { Inject, Injectable, Logger, OnModuleInit } from '@nestjs/common';
import Consul from 'consul';
import { ConfigService } from '@nestjs/config';

@Injectable()
export class ConsulService implements OnModuleInit {
  private readonly logger = new Logger(ConsulService.name);
  private consul: Consul;

  constructor(@Inject() private readonly configService: ConfigService) {
    this.consul = new Consul({
      host: configService.get<string>('CONSUL_HOST'),
      port: configService.get<number>('CONSUL_PORT'),
    });
  }

  async onModuleInit() {
    await this.consul.agent.service.register({
      name: await this.get('/services/commerce/SERVICE_NAME'),
      port: parseInt(await this.get('/services/commerce/RPC_PORT')),
      address: await this.get('/services/commerce/RPC_HOST'),
      tags: ['v1'],
    });
  }

  async get(key: string): Promise<string> {
    const k = `${this.configService.get<string>('ENV')}${key}`;
    const pair = await this.consul.kv.get(k);
    const value = pair?.Value;
    if (!value) {
      this.logger.error(`Key ${key} not found`);
      throw new Error(`Key ${key} not found in Consul`);
    }
    return value.toString();
  }
}
