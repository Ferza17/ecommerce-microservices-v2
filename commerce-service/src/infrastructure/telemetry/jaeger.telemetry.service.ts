import { Injectable, Logger, OnModuleInit } from '@nestjs/common';
import { Service } from '../../enum/service';
import { Context, Span, SpanOptions, Tracer } from '@opentelemetry/api';
import { JaegerExporter } from '@opentelemetry/exporter-jaeger';
import opentelemetry from '@opentelemetry/api';
import { BasicTracerProvider, SimpleSpanProcessor } from '@opentelemetry/sdk-trace-base';
import { ATTR_SERVICE_NAME } from '@opentelemetry/semantic-conventions';
import { AsyncLocalStorageContextManager } from '@opentelemetry/context-async-hooks';
import { CompositePropagator, W3CBaggagePropagator, W3CTraceContextPropagator } from '@opentelemetry/core';
import { Resource } from '@opentelemetry/resources';
import { ConsulService } from '../../config/consul.service';


@Injectable()
export class JaegerTelemetryService implements OnModuleInit {
  private readonly logger = new Logger(JaegerTelemetryService.name);
  private tt: Tracer;

  constructor(private configService: ConsulService,) {
  }

  async onModuleInit() {
    const exporter = new JaegerExporter({
      endpoint: `${await this.configService.get(`/telemetry/jaeger/JAEGER_TELEMETRY_HOST`)}:${await this.configService.get(`/telemetry/jaeger/JAEGER_TELEMETRY_HOST`)}/api/traces`,
    });

    const provider = new BasicTracerProvider({
      resource: new Resource({
        [ATTR_SERVICE_NAME]: Service.CommerceService.toString(),
      }),
    });
    provider.addSpanProcessor(new SimpleSpanProcessor(exporter));
    // provider.addSpanProcessor(new SimpleSpanProcessor(new ConsoleSpanExporter()));
    opentelemetry.context.setGlobalContextManager(new AsyncLocalStorageContextManager());
    opentelemetry.propagation.setGlobalPropagator(new CompositePropagator({
      propagators: [
        new W3CTraceContextPropagator(),
        new W3CBaggagePropagator()],
    }));
    opentelemetry.trace.setGlobalTracerProvider(provider);
    this.tt = opentelemetry.trace.getTracer(Service.CommerceService.toString());
  }

  tracer(operationName: string, ctx?: Context): Span {
    const spanOptions: SpanOptions = {};
    let span: Span;
    if (ctx) {
      span = this.tt.startSpan(operationName, spanOptions, ctx);
    } else {
      span = this.tt.startSpan(operationName, spanOptions, opentelemetry.context.active());
    }
    return span;
  }
}