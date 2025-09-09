import { Module } from '@nestjs/common';
import { ModuleModule } from '../../module/module.module';
import { configRoot, MongooseRootAsync } from '../../config/configRoot';

@Module({
  imports: [
    configRoot,
    MongooseRootAsync,
    ModuleModule,
  ],
})
export class RpcServerModule {
}
