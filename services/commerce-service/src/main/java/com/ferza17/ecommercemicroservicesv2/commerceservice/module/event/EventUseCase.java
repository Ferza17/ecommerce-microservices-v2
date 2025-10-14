package com.ferza17.ecommercemicroservicesv2.commerceservice.module.event;

import com.ferza17.ecommercemicroservicesv2.proto.v1.event.Model;

@org.springframework.stereotype.Service
public class EventUseCase {
    public void AppendEvent(Model.Event event){
        //TODO:
        // 1. Find Event By aggregate type and aggregate id
        // 2. If Exists then next version else version 1
        // 3. Insert via Sink Connector Commerce Event Stores
    }
}
