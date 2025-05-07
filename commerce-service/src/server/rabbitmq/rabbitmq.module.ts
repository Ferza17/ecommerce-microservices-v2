import { Module } from '@nestjs/common';
import { configRoot } from '../../config/configRoot';
import { ModuleModule } from '../../module/module.module';
import { InfrastructureModule } from '../../infrastructure/infrastructure.module';

@Module({
  imports: [
    configRoot,
    ModuleModule,
    InfrastructureModule,
  ],
})
export class RabbitmqModule {
}
