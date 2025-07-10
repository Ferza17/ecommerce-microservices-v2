package com.ferza17.ecommercemicroservicesv2.shippingservice.shippingservice.config;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Component;

@Component
@ConfigurationProperties(prefix = "")
public class ConsulProperties {

    private String env;
    private String serviceName;
    private String rpcHost;
    private String rpcPort;
    private String httpHost;
    private String httpPort;
    private String metricHttpPort;


}
