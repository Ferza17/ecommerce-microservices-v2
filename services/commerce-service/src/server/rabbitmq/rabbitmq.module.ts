import { Module } from '@nestjs/common';
import { configRoot, MongooseRootAsync } from '../../config/configRoot';
import { ModuleModule } from '../../module/module.module';

@Module({
  imports: [
    configRoot,
    MongooseRootAsync,
    ModuleModule,
  ],
})
export class RabbitmqModule {}
