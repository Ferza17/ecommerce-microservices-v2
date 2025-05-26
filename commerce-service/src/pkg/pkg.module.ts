import { Module } from '@nestjs/common';
import { CircuitBreakerService } from './circuit-breaker/circuit-breaker.service';

@Module({
  providers: [CircuitBreakerService],
})
export class PkgModule {
}
