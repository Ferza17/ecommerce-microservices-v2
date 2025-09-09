import {Module} from '@nestjs/common';
import {RabbitmqInfrastructure} from './rabbitmq/rabbitmq';
import {ConfigModule} from '@nestjs/config';
import {ClientsModule, Transport} from '@nestjs/microservices';
import {glob} from 'fast-glob';
import {join} from 'path';
import {Service} from '../enum/service';
import {ProductRpcService} from './rpc/product.rpc.service';
import {UserRpcService} from './rpc/user.rpc.service';
import {JaegerTelemetryService} from './telemetry/jaeger.telemetry.service';
import {ConsulModule} from '../config/consul.module';
import {ConsulService} from '../config/consul.service';
import {RabbitMQRootAsync} from '../config/configRoot';


const projectRoot = join(__dirname, '../../');
const protoFiles = glob.sync('proto/**/*.proto', {
    cwd: projectRoot,
    absolute: true,
});

if (!protoFiles.length) {
    throw new Error('No .proto files found in proto/**/*.proto');
}

@Module({
    imports: [
        ConfigModule.forRoot(),
        ClientsModule.registerAsync([
            {
                name: Service.ProductService.toString(),
                imports: [ConsulModule],
                inject: [ConsulService],
                useFactory: async (config: ConsulService) => ({
                    transport: Transport.GRPC,
                    options: {
                        url: `${await config.get('/services/product/RPC_HOST')}:${await config.get('/services/product/RPC_PORT')}`,
                        package: `product`,
                        protoPath: protoFiles,
                        loader: {
                            includeDirs: [
                                join(projectRoot, 'proto'),
                            ],
                            oneofs: true,
                            keepCase: true,
                            defaults: true,
                        },
                    },
                }),
            },
        ]),
        ClientsModule.registerAsync([
            {
                name: Service.UserService.toString(),
                imports: [ConsulModule],
                inject: [ConsulService],
                useFactory: async (configService: ConsulService) => ({
                    transport: Transport.GRPC,
                    options: {
                        url: `${ await configService.get('/services/user/RPC_HOST')}:${await configService.get('/services/user/RPC_PORT')}`,
                        package: 'user',
                        protoPath: protoFiles,
                        loader: {
                            includeDirs: [
                                join(projectRoot, 'proto'),
                            ],
                            oneofs: true,
                            keepCase: true,
                            defaults: true,
                        },
                    },
                }),
            },
        ]),
        RabbitMQRootAsync,
        ConsulModule,
    ],
    providers: [RabbitmqInfrastructure, ProductRpcService, UserRpcService, JaegerTelemetryService],
    exports: [RabbitmqInfrastructure, ProductRpcService, UserRpcService, JaegerTelemetryService],
})
export class InfrastructureModule {
}
