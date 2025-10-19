package com.ferza17.ecommercemicroservicesv2.commerceservice.config;

import com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.inbound.InboundAuthorization;
import com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.inbound.InboundOpenTelemetry;
import com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.inbound.InboundRequestID;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.CorsRegistry;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

@Configuration
public class RestConfig implements WebMvcConfigurer {
    @Autowired
    private InboundAuthorization authorizationInbound;
    @Autowired
    private InboundOpenTelemetry openTelemetryInbound;
    @Autowired
    private InboundRequestID requestIDInbound;

    @Override
    public void addCorsMappings(CorsRegistry registry) {
        registry.addMapping("/api/**")
                .allowedOrigins("*")
                .allowedMethods("GET", "POST", "PUT", "DELETE")
                .allowedHeaders("*")
                .maxAge(3600);
    }

    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        registry.addInterceptor(authorizationInbound).addPathPatterns("/api/**");
        registry.addInterceptor(openTelemetryInbound).addPathPatterns("/api/**");
        registry.addInterceptor(requestIDInbound).addPathPatterns("/api/**");
    }

}
