package com.ferza17.ecommercemicroservicesv2.commerceservice.interceptor.inbound;

import io.grpc.*;
import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.jboss.logging.MDC;
import org.springframework.core.Ordered;
import org.springframework.core.annotation.Order;
import org.springframework.grpc.server.GlobalServerInterceptor;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;
import javax.security.sasl.AuthenticationException;
import java.io.IOException;

import static com.ferza17.ecommercemicroservicesv2.commerceservice.pkg.context.BaseContext.*;

@GlobalServerInterceptor
@Order(Ordered.HIGHEST_PRECEDENCE + 2)
@Component
public class AuthorizationInbound extends OncePerRequestFilter implements ServerInterceptor {
    // GRPC
    @Override
    public <ReqT, RespT> ServerCall.Listener<ReqT> interceptCall(ServerCall<ReqT, RespT> call, Metadata metadata, ServerCallHandler<ReqT, RespT> next) {
        try {
            String tokenFromHeader = metadata.get(AUTHORIZATION_METADATA);

            if (tokenFromHeader == null || tokenFromHeader.isBlank()) {
                throw new AuthenticationException("Authorization header missing");
            }

            String sanitizedToken = tokenFromHeader.replaceAll("(?i)^Bearer\\s+", "");
            if (sanitizedToken.isBlank()) {
                throw new AuthenticationException("Invalid Token");
            }
            MDC.put(AUTHORIZATION_CONTEXT_KEY, tokenFromHeader);
            ServerCall<ReqT, RespT> wrappedCall = new ForwardingServerCall.SimpleForwardingServerCall<>(call) {
                @Override
                public void sendHeaders(Metadata responseHeaders) {
                    responseHeaders.put(AUTHORIZATION_METADATA, MDC.get(AUTHORIZATION_CONTEXT_KEY).toString());
                    super.sendHeaders(responseHeaders);
                }
            };

            return Contexts.interceptCall(Context.current().withValue(Context.key(AUTHORIZATION_CONTEXT_KEY), tokenFromHeader),
                    wrappedCall, metadata, next);

        } catch (AuthenticationException e) {
            call.close(Status.UNAUTHENTICATED.withDescription(e.getMessage()), new Metadata());
            return new ServerCall.Listener<>() {}; // return empty listener to stop call
        } catch (Exception e) {
            call.close(Status.INTERNAL.withDescription(e.getMessage()), new Metadata());
            return new ServerCall.Listener<>() {};
        }
    }

    // REST
    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain) throws ServletException, IOException {

    }
}
