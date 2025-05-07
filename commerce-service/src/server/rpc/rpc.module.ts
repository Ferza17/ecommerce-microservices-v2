import { Module } from '@nestjs/common';
import { ModuleModule } from '../../module/module.module';
import { InfrastructureModule } from '../../infrastructure/infrastructure.module';
import { configRoot } from '../../config/configRoot';


@Module({
  imports: [
    configRoot,
    ModuleModule,
    InfrastructureModule,
  ],
})
export class RpcServerModule {
}
