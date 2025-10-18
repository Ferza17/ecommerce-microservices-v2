package com.ferza17.ecommercemicroservicesv2.commerceservice.module.event;

import com.ferza17.ecommercemicroservicesv2.proto.v1.event.Model;
import io.opentelemetry.api.trace.Span;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.context.Scope;
import org.springframework.beans.factory.annotation.Autowired;

@org.springframework.stereotype.Service
public class EventUseCase {
    @Autowired
    private Tracer tracer;

    public void AppendEvent(
            Model.Event event){
        Span span = this.tracer.spanBuilder("EventUseCase.AppendEvent").startSpan();
        try (Scope scope = span.makeCurrent()) {
        } catch (Exception e) {
            span.recordException(e);
            throw new RuntimeException(e);
        } finally {
            span.end();
        }
    }
}
