import { addReflectionToGrpcConfig } from 'nestjs-grpc-reflection';
import { GrpcOptions, Transport } from '@nestjs/microservices';
import { glob } from 'fast-glob';
import { join } from 'path';
import { Injectable } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';

@Injectable()
export class GrpcClientOptions {
  constructor(private readonly configService: ConfigService) {}

  get getGRPCConfig(): GrpcOptions {
    const rpcHost = this.configService.get<number>('RPC_HOST') || 50054;
    const rpcPort = this.configService.get<number>('RPC_PORT') || 5000;


    return addReflectionToGrpcConfig({
      transport: Transport.GRPC,
      options: {
        url: `${rpcHost}:${rpcPort}`,
        package: 'proto',
        protoPath: glob.sync(['proto/*.proto'], {
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

