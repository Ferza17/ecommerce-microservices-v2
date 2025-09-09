import {addReflectionToGrpcConfig} from 'nestjs-grpc-reflection';
import {GrpcOptions, Transport} from '@nestjs/microservices';
import {glob} from 'fast-glob';
import {join} from 'path';
import {Injectable} from '@nestjs/common';
import {ConsulService} from '../../config/consul.service';


@Injectable()
export class GrpcClientOptions {

    constructor(private readonly consulConfig: ConsulService) {
        this.consulConfig = consulConfig;
    }


    async getGRPCConfig(): Promise<GrpcOptions> {
        await this.consulConfig.onModuleInit()

        const projectRoot = join(__dirname, '../../../');
        const protoFiles = glob.sync('proto/**/*.proto', {
            cwd: projectRoot,
            absolute: true,
        });

        if (!protoFiles.length) {
            throw new Error('No .proto files found in proto/**/*.proto');
        }

        return addReflectionToGrpcConfig({
            transport: Transport.GRPC,
            options: {
                url: `${this.consulConfig.CommerceServiceRpcHost}:${this.consulConfig.CommerceServiceRpcPort}`,
                package: "commerce",
                protoPath: protoFiles,
                loader: {
                    includeDirs: [
                        join(projectRoot, 'proto'),
                    ],
                    oneofs: true,
                },
            },
        });
    }
}

