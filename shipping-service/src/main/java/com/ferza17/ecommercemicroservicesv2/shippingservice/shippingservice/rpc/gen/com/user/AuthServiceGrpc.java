package com.user;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.73.0)",
    comments = "Source: v1/user/service.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class AuthServiceGrpc {

  private AuthServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "user.AuthService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.user.AuthUserRegisterRequest,
      com.google.protobuf.Empty> getAuthUserRegisterMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AuthUserRegister",
      requestType = com.user.AuthUserRegisterRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.user.AuthUserRegisterRequest,
      com.google.protobuf.Empty> getAuthUserRegisterMethod() {
    io.grpc.MethodDescriptor<com.user.AuthUserRegisterRequest, com.google.protobuf.Empty> getAuthUserRegisterMethod;
    if ((getAuthUserRegisterMethod = AuthServiceGrpc.getAuthUserRegisterMethod) == null) {
      synchronized (AuthServiceGrpc.class) {
        if ((getAuthUserRegisterMethod = AuthServiceGrpc.getAuthUserRegisterMethod) == null) {
          AuthServiceGrpc.getAuthUserRegisterMethod = getAuthUserRegisterMethod =
              io.grpc.MethodDescriptor.<com.user.AuthUserRegisterRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AuthUserRegister"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.user.AuthUserRegisterRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new AuthServiceMethodDescriptorSupplier("AuthUserRegister"))
              .build();
        }
      }
    }
    return getAuthUserRegisterMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.user.AuthUserLoginByEmailAndPasswordRequest,
      com.google.protobuf.Empty> getAuthUserLoginByEmailAndPasswordMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AuthUserLoginByEmailAndPassword",
      requestType = com.user.AuthUserLoginByEmailAndPasswordRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.user.AuthUserLoginByEmailAndPasswordRequest,
      com.google.protobuf.Empty> getAuthUserLoginByEmailAndPasswordMethod() {
    io.grpc.MethodDescriptor<com.user.AuthUserLoginByEmailAndPasswordRequest, com.google.protobuf.Empty> getAuthUserLoginByEmailAndPasswordMethod;
    if ((getAuthUserLoginByEmailAndPasswordMethod = AuthServiceGrpc.getAuthUserLoginByEmailAndPasswordMethod) == null) {
      synchronized (AuthServiceGrpc.class) {
        if ((getAuthUserLoginByEmailAndPasswordMethod = AuthServiceGrpc.getAuthUserLoginByEmailAndPasswordMethod) == null) {
          AuthServiceGrpc.getAuthUserLoginByEmailAndPasswordMethod = getAuthUserLoginByEmailAndPasswordMethod =
              io.grpc.MethodDescriptor.<com.user.AuthUserLoginByEmailAndPasswordRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AuthUserLoginByEmailAndPassword"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.user.AuthUserLoginByEmailAndPasswordRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new AuthServiceMethodDescriptorSupplier("AuthUserLoginByEmailAndPassword"))
              .build();
        }
      }
    }
    return getAuthUserLoginByEmailAndPasswordMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.user.AuthUserVerifyOtpRequest,
      com.user.AuthUserVerifyOtpResponse> getAuthUserVerifyOtpMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AuthUserVerifyOtp",
      requestType = com.user.AuthUserVerifyOtpRequest.class,
      responseType = com.user.AuthUserVerifyOtpResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.user.AuthUserVerifyOtpRequest,
      com.user.AuthUserVerifyOtpResponse> getAuthUserVerifyOtpMethod() {
    io.grpc.MethodDescriptor<com.user.AuthUserVerifyOtpRequest, com.user.AuthUserVerifyOtpResponse> getAuthUserVerifyOtpMethod;
    if ((getAuthUserVerifyOtpMethod = AuthServiceGrpc.getAuthUserVerifyOtpMethod) == null) {
      synchronized (AuthServiceGrpc.class) {
        if ((getAuthUserVerifyOtpMethod = AuthServiceGrpc.getAuthUserVerifyOtpMethod) == null) {
          AuthServiceGrpc.getAuthUserVerifyOtpMethod = getAuthUserVerifyOtpMethod =
              io.grpc.MethodDescriptor.<com.user.AuthUserVerifyOtpRequest, com.user.AuthUserVerifyOtpResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AuthUserVerifyOtp"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.user.AuthUserVerifyOtpRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.user.AuthUserVerifyOtpResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AuthServiceMethodDescriptorSupplier("AuthUserVerifyOtp"))
              .build();
        }
      }
    }
    return getAuthUserVerifyOtpMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.user.AuthUserLogoutByTokenRequest,
      com.user.AuthUserLogoutByTokenResponse> getAuthUserLogoutByTokenMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AuthUserLogoutByToken",
      requestType = com.user.AuthUserLogoutByTokenRequest.class,
      responseType = com.user.AuthUserLogoutByTokenResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.user.AuthUserLogoutByTokenRequest,
      com.user.AuthUserLogoutByTokenResponse> getAuthUserLogoutByTokenMethod() {
    io.grpc.MethodDescriptor<com.user.AuthUserLogoutByTokenRequest, com.user.AuthUserLogoutByTokenResponse> getAuthUserLogoutByTokenMethod;
    if ((getAuthUserLogoutByTokenMethod = AuthServiceGrpc.getAuthUserLogoutByTokenMethod) == null) {
      synchronized (AuthServiceGrpc.class) {
        if ((getAuthUserLogoutByTokenMethod = AuthServiceGrpc.getAuthUserLogoutByTokenMethod) == null) {
          AuthServiceGrpc.getAuthUserLogoutByTokenMethod = getAuthUserLogoutByTokenMethod =
              io.grpc.MethodDescriptor.<com.user.AuthUserLogoutByTokenRequest, com.user.AuthUserLogoutByTokenResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AuthUserLogoutByToken"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.user.AuthUserLogoutByTokenRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.user.AuthUserLogoutByTokenResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AuthServiceMethodDescriptorSupplier("AuthUserLogoutByToken"))
              .build();
        }
      }
    }
    return getAuthUserLogoutByTokenMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.user.AuthUserVerifyAccessControlRequest,
      com.user.AuthUserVerifyAccessControlResponse> getAuthUserVerifyAccessControlMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AuthUserVerifyAccessControl",
      requestType = com.user.AuthUserVerifyAccessControlRequest.class,
      responseType = com.user.AuthUserVerifyAccessControlResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.user.AuthUserVerifyAccessControlRequest,
      com.user.AuthUserVerifyAccessControlResponse> getAuthUserVerifyAccessControlMethod() {
    io.grpc.MethodDescriptor<com.user.AuthUserVerifyAccessControlRequest, com.user.AuthUserVerifyAccessControlResponse> getAuthUserVerifyAccessControlMethod;
    if ((getAuthUserVerifyAccessControlMethod = AuthServiceGrpc.getAuthUserVerifyAccessControlMethod) == null) {
      synchronized (AuthServiceGrpc.class) {
        if ((getAuthUserVerifyAccessControlMethod = AuthServiceGrpc.getAuthUserVerifyAccessControlMethod) == null) {
          AuthServiceGrpc.getAuthUserVerifyAccessControlMethod = getAuthUserVerifyAccessControlMethod =
              io.grpc.MethodDescriptor.<com.user.AuthUserVerifyAccessControlRequest, com.user.AuthUserVerifyAccessControlResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AuthUserVerifyAccessControl"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.user.AuthUserVerifyAccessControlRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.user.AuthUserVerifyAccessControlResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AuthServiceMethodDescriptorSupplier("AuthUserVerifyAccessControl"))
              .build();
        }
      }
    }
    return getAuthUserVerifyAccessControlMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.user.AuthServiceVerifyIsExcludedRequest,
      com.user.AuthServiceVerifyIsExcludedResponse> getAuthServiceVerifyIsExcludedMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AuthServiceVerifyIsExcluded",
      requestType = com.user.AuthServiceVerifyIsExcludedRequest.class,
      responseType = com.user.AuthServiceVerifyIsExcludedResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.user.AuthServiceVerifyIsExcludedRequest,
      com.user.AuthServiceVerifyIsExcludedResponse> getAuthServiceVerifyIsExcludedMethod() {
    io.grpc.MethodDescriptor<com.user.AuthServiceVerifyIsExcludedRequest, com.user.AuthServiceVerifyIsExcludedResponse> getAuthServiceVerifyIsExcludedMethod;
    if ((getAuthServiceVerifyIsExcludedMethod = AuthServiceGrpc.getAuthServiceVerifyIsExcludedMethod) == null) {
      synchronized (AuthServiceGrpc.class) {
        if ((getAuthServiceVerifyIsExcludedMethod = AuthServiceGrpc.getAuthServiceVerifyIsExcludedMethod) == null) {
          AuthServiceGrpc.getAuthServiceVerifyIsExcludedMethod = getAuthServiceVerifyIsExcludedMethod =
              io.grpc.MethodDescriptor.<com.user.AuthServiceVerifyIsExcludedRequest, com.user.AuthServiceVerifyIsExcludedResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AuthServiceVerifyIsExcluded"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.user.AuthServiceVerifyIsExcludedRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.user.AuthServiceVerifyIsExcludedResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AuthServiceMethodDescriptorSupplier("AuthServiceVerifyIsExcluded"))
              .build();
        }
      }
    }
    return getAuthServiceVerifyIsExcludedMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.user.AuthUserFindUserByTokenRequest,
      com.user.AuthUserFindUserByTokenResponse> getAuthUserFindUserByTokenMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "AuthUserFindUserByToken",
      requestType = com.user.AuthUserFindUserByTokenRequest.class,
      responseType = com.user.AuthUserFindUserByTokenResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.user.AuthUserFindUserByTokenRequest,
      com.user.AuthUserFindUserByTokenResponse> getAuthUserFindUserByTokenMethod() {
    io.grpc.MethodDescriptor<com.user.AuthUserFindUserByTokenRequest, com.user.AuthUserFindUserByTokenResponse> getAuthUserFindUserByTokenMethod;
    if ((getAuthUserFindUserByTokenMethod = AuthServiceGrpc.getAuthUserFindUserByTokenMethod) == null) {
      synchronized (AuthServiceGrpc.class) {
        if ((getAuthUserFindUserByTokenMethod = AuthServiceGrpc.getAuthUserFindUserByTokenMethod) == null) {
          AuthServiceGrpc.getAuthUserFindUserByTokenMethod = getAuthUserFindUserByTokenMethod =
              io.grpc.MethodDescriptor.<com.user.AuthUserFindUserByTokenRequest, com.user.AuthUserFindUserByTokenResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "AuthUserFindUserByToken"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.user.AuthUserFindUserByTokenRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.user.AuthUserFindUserByTokenResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AuthServiceMethodDescriptorSupplier("AuthUserFindUserByToken"))
              .build();
        }
      }
    }
    return getAuthUserFindUserByTokenMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static AuthServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AuthServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AuthServiceStub>() {
        @java.lang.Override
        public AuthServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AuthServiceStub(channel, callOptions);
        }
      };
    return AuthServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static AuthServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AuthServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AuthServiceBlockingV2Stub>() {
        @java.lang.Override
        public AuthServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AuthServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return AuthServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static AuthServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AuthServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AuthServiceBlockingStub>() {
        @java.lang.Override
        public AuthServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AuthServiceBlockingStub(channel, callOptions);
        }
      };
    return AuthServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static AuthServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AuthServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AuthServiceFutureStub>() {
        @java.lang.Override
        public AuthServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AuthServiceFutureStub(channel, callOptions);
        }
      };
    return AuthServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    default void authUserRegister(com.user.AuthUserRegisterRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAuthUserRegisterMethod(), responseObserver);
    }

    /**
     */
    default void authUserLoginByEmailAndPassword(com.user.AuthUserLoginByEmailAndPasswordRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAuthUserLoginByEmailAndPasswordMethod(), responseObserver);
    }

    /**
     */
    default void authUserVerifyOtp(com.user.AuthUserVerifyOtpRequest request,
        io.grpc.stub.StreamObserver<com.user.AuthUserVerifyOtpResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAuthUserVerifyOtpMethod(), responseObserver);
    }

    /**
     */
    default void authUserLogoutByToken(com.user.AuthUserLogoutByTokenRequest request,
        io.grpc.stub.StreamObserver<com.user.AuthUserLogoutByTokenResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAuthUserLogoutByTokenMethod(), responseObserver);
    }

    /**
     */
    default void authUserVerifyAccessControl(com.user.AuthUserVerifyAccessControlRequest request,
        io.grpc.stub.StreamObserver<com.user.AuthUserVerifyAccessControlResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAuthUserVerifyAccessControlMethod(), responseObserver);
    }

    /**
     */
    default void authServiceVerifyIsExcluded(com.user.AuthServiceVerifyIsExcludedRequest request,
        io.grpc.stub.StreamObserver<com.user.AuthServiceVerifyIsExcludedResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAuthServiceVerifyIsExcludedMethod(), responseObserver);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    default void authUserFindUserByToken(com.user.AuthUserFindUserByTokenRequest request,
        io.grpc.stub.StreamObserver<com.user.AuthUserFindUserByTokenResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getAuthUserFindUserByTokenMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service AuthService.
   */
  public static abstract class AuthServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return AuthServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service AuthService.
   */
  public static final class AuthServiceStub
      extends io.grpc.stub.AbstractAsyncStub<AuthServiceStub> {
    private AuthServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AuthServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AuthServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    public void authUserRegister(com.user.AuthUserRegisterRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAuthUserRegisterMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void authUserLoginByEmailAndPassword(com.user.AuthUserLoginByEmailAndPasswordRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAuthUserLoginByEmailAndPasswordMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void authUserVerifyOtp(com.user.AuthUserVerifyOtpRequest request,
        io.grpc.stub.StreamObserver<com.user.AuthUserVerifyOtpResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAuthUserVerifyOtpMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void authUserLogoutByToken(com.user.AuthUserLogoutByTokenRequest request,
        io.grpc.stub.StreamObserver<com.user.AuthUserLogoutByTokenResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAuthUserLogoutByTokenMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void authUserVerifyAccessControl(com.user.AuthUserVerifyAccessControlRequest request,
        io.grpc.stub.StreamObserver<com.user.AuthUserVerifyAccessControlResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAuthUserVerifyAccessControlMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void authServiceVerifyIsExcluded(com.user.AuthServiceVerifyIsExcludedRequest request,
        io.grpc.stub.StreamObserver<com.user.AuthServiceVerifyIsExcludedResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAuthServiceVerifyIsExcludedMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public void authUserFindUserByToken(com.user.AuthUserFindUserByTokenRequest request,
        io.grpc.stub.StreamObserver<com.user.AuthUserFindUserByTokenResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getAuthUserFindUserByTokenMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service AuthService.
   */
  public static final class AuthServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<AuthServiceBlockingV2Stub> {
    private AuthServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AuthServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AuthServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    public com.google.protobuf.Empty authUserRegister(com.user.AuthUserRegisterRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAuthUserRegisterMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty authUserLoginByEmailAndPassword(com.user.AuthUserLoginByEmailAndPasswordRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAuthUserLoginByEmailAndPasswordMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.user.AuthUserVerifyOtpResponse authUserVerifyOtp(com.user.AuthUserVerifyOtpRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAuthUserVerifyOtpMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.user.AuthUserLogoutByTokenResponse authUserLogoutByToken(com.user.AuthUserLogoutByTokenRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAuthUserLogoutByTokenMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.user.AuthUserVerifyAccessControlResponse authUserVerifyAccessControl(com.user.AuthUserVerifyAccessControlRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAuthUserVerifyAccessControlMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.user.AuthServiceVerifyIsExcludedResponse authServiceVerifyIsExcluded(com.user.AuthServiceVerifyIsExcludedRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAuthServiceVerifyIsExcludedMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public com.user.AuthUserFindUserByTokenResponse authUserFindUserByToken(com.user.AuthUserFindUserByTokenRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAuthUserFindUserByTokenMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service AuthService.
   */
  public static final class AuthServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<AuthServiceBlockingStub> {
    private AuthServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AuthServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AuthServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    public com.google.protobuf.Empty authUserRegister(com.user.AuthUserRegisterRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAuthUserRegisterMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty authUserLoginByEmailAndPassword(com.user.AuthUserLoginByEmailAndPasswordRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAuthUserLoginByEmailAndPasswordMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.user.AuthUserVerifyOtpResponse authUserVerifyOtp(com.user.AuthUserVerifyOtpRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAuthUserVerifyOtpMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.user.AuthUserLogoutByTokenResponse authUserLogoutByToken(com.user.AuthUserLogoutByTokenRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAuthUserLogoutByTokenMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.user.AuthUserVerifyAccessControlResponse authUserVerifyAccessControl(com.user.AuthUserVerifyAccessControlRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAuthUserVerifyAccessControlMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.user.AuthServiceVerifyIsExcludedResponse authServiceVerifyIsExcluded(com.user.AuthServiceVerifyIsExcludedRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAuthServiceVerifyIsExcludedMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public com.user.AuthUserFindUserByTokenResponse authUserFindUserByToken(com.user.AuthUserFindUserByTokenRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getAuthUserFindUserByTokenMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service AuthService.
   */
  public static final class AuthServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<AuthServiceFutureStub> {
    private AuthServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AuthServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AuthServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> authUserRegister(
        com.user.AuthUserRegisterRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAuthUserRegisterMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> authUserLoginByEmailAndPassword(
        com.user.AuthUserLoginByEmailAndPasswordRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAuthUserLoginByEmailAndPasswordMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.user.AuthUserVerifyOtpResponse> authUserVerifyOtp(
        com.user.AuthUserVerifyOtpRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAuthUserVerifyOtpMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.user.AuthUserLogoutByTokenResponse> authUserLogoutByToken(
        com.user.AuthUserLogoutByTokenRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAuthUserLogoutByTokenMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.user.AuthUserVerifyAccessControlResponse> authUserVerifyAccessControl(
        com.user.AuthUserVerifyAccessControlRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAuthUserVerifyAccessControlMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.user.AuthServiceVerifyIsExcludedResponse> authServiceVerifyIsExcluded(
        com.user.AuthServiceVerifyIsExcludedRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAuthServiceVerifyIsExcludedMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.user.AuthUserFindUserByTokenResponse> authUserFindUserByToken(
        com.user.AuthUserFindUserByTokenRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getAuthUserFindUserByTokenMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_AUTH_USER_REGISTER = 0;
  private static final int METHODID_AUTH_USER_LOGIN_BY_EMAIL_AND_PASSWORD = 1;
  private static final int METHODID_AUTH_USER_VERIFY_OTP = 2;
  private static final int METHODID_AUTH_USER_LOGOUT_BY_TOKEN = 3;
  private static final int METHODID_AUTH_USER_VERIFY_ACCESS_CONTROL = 4;
  private static final int METHODID_AUTH_SERVICE_VERIFY_IS_EXCLUDED = 5;
  private static final int METHODID_AUTH_USER_FIND_USER_BY_TOKEN = 6;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_AUTH_USER_REGISTER:
          serviceImpl.authUserRegister((com.user.AuthUserRegisterRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_AUTH_USER_LOGIN_BY_EMAIL_AND_PASSWORD:
          serviceImpl.authUserLoginByEmailAndPassword((com.user.AuthUserLoginByEmailAndPasswordRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_AUTH_USER_VERIFY_OTP:
          serviceImpl.authUserVerifyOtp((com.user.AuthUserVerifyOtpRequest) request,
              (io.grpc.stub.StreamObserver<com.user.AuthUserVerifyOtpResponse>) responseObserver);
          break;
        case METHODID_AUTH_USER_LOGOUT_BY_TOKEN:
          serviceImpl.authUserLogoutByToken((com.user.AuthUserLogoutByTokenRequest) request,
              (io.grpc.stub.StreamObserver<com.user.AuthUserLogoutByTokenResponse>) responseObserver);
          break;
        case METHODID_AUTH_USER_VERIFY_ACCESS_CONTROL:
          serviceImpl.authUserVerifyAccessControl((com.user.AuthUserVerifyAccessControlRequest) request,
              (io.grpc.stub.StreamObserver<com.user.AuthUserVerifyAccessControlResponse>) responseObserver);
          break;
        case METHODID_AUTH_SERVICE_VERIFY_IS_EXCLUDED:
          serviceImpl.authServiceVerifyIsExcluded((com.user.AuthServiceVerifyIsExcludedRequest) request,
              (io.grpc.stub.StreamObserver<com.user.AuthServiceVerifyIsExcludedResponse>) responseObserver);
          break;
        case METHODID_AUTH_USER_FIND_USER_BY_TOKEN:
          serviceImpl.authUserFindUserByToken((com.user.AuthUserFindUserByTokenRequest) request,
              (io.grpc.stub.StreamObserver<com.user.AuthUserFindUserByTokenResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getAuthUserRegisterMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.user.AuthUserRegisterRequest,
              com.google.protobuf.Empty>(
                service, METHODID_AUTH_USER_REGISTER)))
        .addMethod(
          getAuthUserLoginByEmailAndPasswordMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.user.AuthUserLoginByEmailAndPasswordRequest,
              com.google.protobuf.Empty>(
                service, METHODID_AUTH_USER_LOGIN_BY_EMAIL_AND_PASSWORD)))
        .addMethod(
          getAuthUserVerifyOtpMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.user.AuthUserVerifyOtpRequest,
              com.user.AuthUserVerifyOtpResponse>(
                service, METHODID_AUTH_USER_VERIFY_OTP)))
        .addMethod(
          getAuthUserLogoutByTokenMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.user.AuthUserLogoutByTokenRequest,
              com.user.AuthUserLogoutByTokenResponse>(
                service, METHODID_AUTH_USER_LOGOUT_BY_TOKEN)))
        .addMethod(
          getAuthUserVerifyAccessControlMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.user.AuthUserVerifyAccessControlRequest,
              com.user.AuthUserVerifyAccessControlResponse>(
                service, METHODID_AUTH_USER_VERIFY_ACCESS_CONTROL)))
        .addMethod(
          getAuthServiceVerifyIsExcludedMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.user.AuthServiceVerifyIsExcludedRequest,
              com.user.AuthServiceVerifyIsExcludedResponse>(
                service, METHODID_AUTH_SERVICE_VERIFY_IS_EXCLUDED)))
        .addMethod(
          getAuthUserFindUserByTokenMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.user.AuthUserFindUserByTokenRequest,
              com.user.AuthUserFindUserByTokenResponse>(
                service, METHODID_AUTH_USER_FIND_USER_BY_TOKEN)))
        .build();
  }

  private static abstract class AuthServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    AuthServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.user.ServiceProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("AuthService");
    }
  }

  private static final class AuthServiceFileDescriptorSupplier
      extends AuthServiceBaseDescriptorSupplier {
    AuthServiceFileDescriptorSupplier() {}
  }

  private static final class AuthServiceMethodDescriptorSupplier
      extends AuthServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    AuthServiceMethodDescriptorSupplier(java.lang.String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (AuthServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new AuthServiceFileDescriptorSupplier())
              .addMethod(getAuthUserRegisterMethod())
              .addMethod(getAuthUserLoginByEmailAndPasswordMethod())
              .addMethod(getAuthUserVerifyOtpMethod())
              .addMethod(getAuthUserLogoutByTokenMethod())
              .addMethod(getAuthUserVerifyAccessControlMethod())
              .addMethod(getAuthServiceVerifyIsExcludedMethod())
              .addMethod(getAuthUserFindUserByTokenMethod())
              .build();
        }
      }
    }
    return result;
  }
}
