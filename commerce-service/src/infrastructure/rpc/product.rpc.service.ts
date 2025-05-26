import { Inject, Injectable, Logger, OnModuleInit } from '@nestjs/common';
import { ProductServiceService } from '../../model/rpc/productServices';
import { ClientGrpc } from '@nestjs/microservices';
import { Service } from '../../enum/service';
import { FindProductByIdRequest, Product } from '../../model/rpc/productMessage';
import { Metadata } from '@grpc/grpc-js';
import { Header } from '../../enum/header';
import { JaegerTelemetryService } from '../telemetry/jaeger.telemetry.service';
import { Context } from '@opentelemetry/api';
import CircuitBreaker from 'opossum';

@Injectable()
export class ProductRpcService implements OnModuleInit {
  private readonly logger = new Logger(ProductRpcService.name);
  private productService: any;

  constructor(
    @Inject(Service.ProductService.toString()) private client: ClientGrpc,
    private readonly otel: JaegerTelemetryService,
  ) {
  }

  onModuleInit() {
    this.productService = this.client.getService<typeof ProductServiceService>('ProductService');
  }

  async findProductById(requestId: string, req: FindProductByIdRequest, context?: Context): Promise<Product> {
    const span = this.otel.tracer('RpcService.findProductById', context);
    try {
      const metadata = new Metadata();
      metadata.set(Header.X_REQUEST_ID, requestId);
      const breaker = new CircuitBreaker(this.productService.findProductById(req, metadata), {
        timeout: 1000,
        errorThresholdPercentage: 50,
        resetTimeout: 10000,
      });

      breaker.on('success', (result, latencyMs) => {
        this.logger.log(`findUserById requestId: ${requestId} , latencyMs: ${latencyMs}`);
      })
      breaker.on('failure', (result, latencyMs) => {
        this.logger.log(`findUserById requestId: ${requestId} , latencyMs: ${latencyMs}`);
      })
      breaker.on('open', () => {
        this.logger.log(`findUserById requestId: ${requestId} , circuit breaker is open`);
      })
      breaker.on('close', () => {
        this.logger.log(`findUserById requestId: ${requestId} , circuit breaker is closed`);
      })
      breaker.on('halfOpen', () => {
        this.logger.log(`findUserById requestId: ${requestId} , circuit breaker is half open`);
      })
      return await breaker.fire() as Product;
    } catch (e) {
      span.recordException(e);
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      throw e;
    }finally {
      span.end();
    }
  }
}
