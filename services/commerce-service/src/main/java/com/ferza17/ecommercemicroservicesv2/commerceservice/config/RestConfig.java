package com.ferza17.ecommercemicroservicesv2.commerceservice.config;

import com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.inbound.AuthorizationInbound;
import com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.inbound.OpenTelemetryInbound;
import com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.inbound.RequestIDInbound;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.CorsRegistry;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

@Configuration
public class RestConfig implements WebMvcConfigurer {
    @Autowired
    private AuthorizationInbound authorizationInbound;
    @Autowired
    private OpenTelemetryInbound openTelemetryInbound;
    @Autowired
    private RequestIDInbound requestIDInbound;

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
