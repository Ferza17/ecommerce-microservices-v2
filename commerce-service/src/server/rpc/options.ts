import { addReflectionToGrpcConfig } from 'nestjs-grpc-reflection';
import { GrpcOptions, Transport } from '@nestjs/microservices';
import { glob } from 'fast-glob';
import { join } from 'path';
import { Injectable } from '@nestjs/common';
import { ConsulService } from '../../config/consul.service';


@Injectable()
export class GrpcClientOptions {
  constructor(private readonly consulConfig: ConsulService) {
  }

  async getGRPCConfig(): Promise<GrpcOptions> {
    const rpcHost = await this.consulConfig.get('/services/commerce/RPC_HOST');
    const rpcPort = await this.consulConfig.get('/services/commerce/RPC_PORT') || '5000';

    return addReflectionToGrpcConfig({
      transport: Transport.GRPC,
      options: {
        url: `${rpcHost}:${rpcPort}`,
        package: 'pb',
        protoPath: glob.sync(['proto/**/*.proto'], {
          cwd: join(__dirname, '../../../'),
          absolute: true,
        }),
        loader: {
          oneofs: true,
        },
      },
    });
  }
}

