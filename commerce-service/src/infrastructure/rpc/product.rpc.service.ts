import { Inject, Injectable, Logger, OnModuleInit } from '@nestjs/common';
import { ProductServiceService } from '../../model/rpc/productServices';
import { ClientGrpc } from '@nestjs/microservices';
import { Service } from '../../enum/service';
import { FindProductByIdRequest, Product } from '../../model/rpc/productMessage';
import { lastValueFrom, Observable } from 'rxjs';
import { Metadata } from '@grpc/grpc-js';
import { Header } from '../../enum/header';
import { JaegerTelemetryService } from '../telemetry/jaeger.telemetry.service';
import { Context } from '@opentelemetry/api';

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
      return await lastValueFrom(this.productService.findProductById(req, metadata));
    } catch (e) {
      this.logger.error(`requestId: ${requestId} , error: ${e.message}`);
      throw e;
    }
  }
}
