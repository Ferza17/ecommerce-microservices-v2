import { Inject, Injectable, Logger, OnModuleInit } from '@nestjs/common';
import { ProductServiceService } from '../../model/rpc/gen/product/v1/productServices';
import { ClientGrpc } from '@nestjs/microservices';
import { Service } from '../../enum/service';
import { FindProductByIdRequest, Product } from '../../model/rpc/gen/product/v1/productMessage';
import { Metadata } from '@grpc/grpc-js';
import { Header } from '../../enum/header';
import { JaegerTelemetryService } from '../telemetry/jaeger.telemetry.service';
import { Context, propagation } from '@opentelemetry/api';
import CircuitBreaker from 'opossum';
import { lastValueFrom } from 'rxjs';
import { ProductServiceCircuitOptions } from '../../config/circuitOptions';

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
      const breakerFn: (requestId: string, req: FindProductByIdRequest, context?: Context) => Promise<Product> =
        async (requestId, req, context) => {
          const metadata = new Metadata();
          metadata.set(Header.X_REQUEST_ID, requestId);
          if (context) {
            propagation.inject(context, metadata, {
              set: (metadata, key, value) => metadata.set(key, value as string),
            });
          }
          const observableResult = this.productService.findProductById(req, metadata);
          return lastValueFrom(observableResult);
        };

      const cb = new CircuitBreaker<[string, FindProductByIdRequest, Context], Product>(breakerFn, ProductServiceCircuitOptions);
      return await cb.fire(requestId, req, <Context>context);
    } catch (e) {
      span.recordException(e);
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      throw e;
    } finally {
      span.end();
    }
  }
}
