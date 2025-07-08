package com.user;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.73.0)",
    comments = "Source: v1/user/service.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class UserServiceGrpc {

  private UserServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "user.UserService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.user.UpdateUserByIdRequest,
      com.google.protobuf.Empty> getUpdateUserByIdMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateUserById",
      requestType = com.user.UpdateUserByIdRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.user.UpdateUserByIdRequest,
      com.google.protobuf.Empty> getUpdateUserByIdMethod() {
    io.grpc.MethodDescriptor<com.user.UpdateUserByIdRequest, com.google.protobuf.Empty> getUpdateUserByIdMethod;
    if ((getUpdateUserByIdMethod = UserServiceGrpc.getUpdateUserByIdMethod) == null) {
      synchronized (UserServiceGrpc.class) {
        if ((getUpdateUserByIdMethod = UserServiceGrpc.getUpdateUserByIdMethod) == null) {
          UserServiceGrpc.getUpdateUserByIdMethod = getUpdateUserByIdMethod =
              io.grpc.MethodDescriptor.<com.user.UpdateUserByIdRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateUserById"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.user.UpdateUserByIdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new UserServiceMethodDescriptorSupplier("UpdateUserById"))
              .build();
        }
      }
    }
    return getUpdateUserByIdMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.user.FindUserByIdRequest,
      com.user.FindUserByIdResponse> getFindUserByIdMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "FindUserById",
      requestType = com.user.FindUserByIdRequest.class,
      responseType = com.user.FindUserByIdResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.user.FindUserByIdRequest,
      com.user.FindUserByIdResponse> getFindUserByIdMethod() {
    io.grpc.MethodDescriptor<com.user.FindUserByIdRequest, com.user.FindUserByIdResponse> getFindUserByIdMethod;
    if ((getFindUserByIdMethod = UserServiceGrpc.getFindUserByIdMethod) == null) {
      synchronized (UserServiceGrpc.class) {
        if ((getFindUserByIdMethod = UserServiceGrpc.getFindUserByIdMethod) == null) {
          UserServiceGrpc.getFindUserByIdMethod = getFindUserByIdMethod =
              io.grpc.MethodDescriptor.<com.user.FindUserByIdRequest, com.user.FindUserByIdResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "FindUserById"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.user.FindUserByIdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.user.FindUserByIdResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserServiceMethodDescriptorSupplier("FindUserById"))
              .build();
        }
      }
    }
    return getFindUserByIdMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.user.FindUserByEmailAndPasswordRequest,
      com.user.FindUserByEmailAndPasswordResponse> getFindUserByEmailAndPasswordMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "FindUserByEmailAndPassword",
      requestType = com.user.FindUserByEmailAndPasswordRequest.class,
      responseType = com.user.FindUserByEmailAndPasswordResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.user.FindUserByEmailAndPasswordRequest,
      com.user.FindUserByEmailAndPasswordResponse> getFindUserByEmailAndPasswordMethod() {
    io.grpc.MethodDescriptor<com.user.FindUserByEmailAndPasswordRequest, com.user.FindUserByEmailAndPasswordResponse> getFindUserByEmailAndPasswordMethod;
    if ((getFindUserByEmailAndPasswordMethod = UserServiceGrpc.getFindUserByEmailAndPasswordMethod) == null) {
      synchronized (UserServiceGrpc.class) {
        if ((getFindUserByEmailAndPasswordMethod = UserServiceGrpc.getFindUserByEmailAndPasswordMethod) == null) {
          UserServiceGrpc.getFindUserByEmailAndPasswordMethod = getFindUserByEmailAndPasswordMethod =
              io.grpc.MethodDescriptor.<com.user.FindUserByEmailAndPasswordRequest, com.user.FindUserByEmailAndPasswordResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "FindUserByEmailAndPassword"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.user.FindUserByEmailAndPasswordRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.user.FindUserByEmailAndPasswordResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserServiceMethodDescriptorSupplier("FindUserByEmailAndPassword"))
              .build();
        }
      }
    }
    return getFindUserByEmailAndPasswordMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static UserServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserServiceStub>() {
        @java.lang.Override
        public UserServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserServiceStub(channel, callOptions);
        }
      };
    return UserServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static UserServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserServiceBlockingV2Stub>() {
        @java.lang.Override
        public UserServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return UserServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static UserServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserServiceBlockingStub>() {
        @java.lang.Override
        public UserServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserServiceBlockingStub(channel, callOptions);
        }
      };
    return UserServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static UserServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserServiceFutureStub>() {
        @java.lang.Override
        public UserServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserServiceFutureStub(channel, callOptions);
        }
      };
    return UserServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    default void updateUserById(com.user.UpdateUserByIdRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateUserByIdMethod(), responseObserver);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    default void findUserById(com.user.FindUserByIdRequest request,
        io.grpc.stub.StreamObserver<com.user.FindUserByIdResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getFindUserByIdMethod(), responseObserver);
    }

    /**
     */
    default void findUserByEmailAndPassword(com.user.FindUserByEmailAndPasswordRequest request,
        io.grpc.stub.StreamObserver<com.user.FindUserByEmailAndPasswordResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getFindUserByEmailAndPasswordMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service UserService.
   */
  public static abstract class UserServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return UserServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service UserService.
   */
  public static final class UserServiceStub
      extends io.grpc.stub.AbstractAsyncStub<UserServiceStub> {
    private UserServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    public void updateUserById(com.user.UpdateUserByIdRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateUserByIdMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public void findUserById(com.user.FindUserByIdRequest request,
        io.grpc.stub.StreamObserver<com.user.FindUserByIdResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getFindUserByIdMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void findUserByEmailAndPassword(com.user.FindUserByEmailAndPasswordRequest request,
        io.grpc.stub.StreamObserver<com.user.FindUserByEmailAndPasswordResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getFindUserByEmailAndPasswordMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service UserService.
   */
  public static final class UserServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<UserServiceBlockingV2Stub> {
    private UserServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    public com.google.protobuf.Empty updateUserById(com.user.UpdateUserByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateUserByIdMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public com.user.FindUserByIdResponse findUserById(com.user.FindUserByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindUserByIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.user.FindUserByEmailAndPasswordResponse findUserByEmailAndPassword(com.user.FindUserByEmailAndPasswordRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindUserByEmailAndPasswordMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service UserService.
   */
  public static final class UserServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<UserServiceBlockingStub> {
    private UserServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    public com.google.protobuf.Empty updateUserById(com.user.UpdateUserByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateUserByIdMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public com.user.FindUserByIdResponse findUserById(com.user.FindUserByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindUserByIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.user.FindUserByEmailAndPasswordResponse findUserByEmailAndPassword(com.user.FindUserByEmailAndPasswordRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindUserByEmailAndPasswordMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service UserService.
   */
  public static final class UserServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<UserServiceFutureStub> {
    private UserServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> updateUserById(
        com.user.UpdateUserByIdRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateUserByIdMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.user.FindUserByIdResponse> findUserById(
        com.user.FindUserByIdRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getFindUserByIdMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.user.FindUserByEmailAndPasswordResponse> findUserByEmailAndPassword(
        com.user.FindUserByEmailAndPasswordRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getFindUserByEmailAndPasswordMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_UPDATE_USER_BY_ID = 0;
  private static final int METHODID_FIND_USER_BY_ID = 1;
  private static final int METHODID_FIND_USER_BY_EMAIL_AND_PASSWORD = 2;

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
        case METHODID_UPDATE_USER_BY_ID:
          serviceImpl.updateUserById((com.user.UpdateUserByIdRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_FIND_USER_BY_ID:
          serviceImpl.findUserById((com.user.FindUserByIdRequest) request,
              (io.grpc.stub.StreamObserver<com.user.FindUserByIdResponse>) responseObserver);
          break;
        case METHODID_FIND_USER_BY_EMAIL_AND_PASSWORD:
          serviceImpl.findUserByEmailAndPassword((com.user.FindUserByEmailAndPasswordRequest) request,
              (io.grpc.stub.StreamObserver<com.user.FindUserByEmailAndPasswordResponse>) responseObserver);
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
          getUpdateUserByIdMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.user.UpdateUserByIdRequest,
              com.google.protobuf.Empty>(
                service, METHODID_UPDATE_USER_BY_ID)))
        .addMethod(
          getFindUserByIdMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.user.FindUserByIdRequest,
              com.user.FindUserByIdResponse>(
                service, METHODID_FIND_USER_BY_ID)))
        .addMethod(
          getFindUserByEmailAndPasswordMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.user.FindUserByEmailAndPasswordRequest,
              com.user.FindUserByEmailAndPasswordResponse>(
                service, METHODID_FIND_USER_BY_EMAIL_AND_PASSWORD)))
        .build();
  }

  private static abstract class UserServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    UserServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.user.ServiceProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("UserService");
    }
  }

  private static final class UserServiceFileDescriptorSupplier
      extends UserServiceBaseDescriptorSupplier {
    UserServiceFileDescriptorSupplier() {}
  }

  private static final class UserServiceMethodDescriptorSupplier
      extends UserServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    UserServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (UserServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new UserServiceFileDescriptorSupplier())
              .addMethod(getUpdateUserByIdMethod())
              .addMethod(getFindUserByIdMethod())
              .addMethod(getFindUserByEmailAndPasswordMethod())
              .build();
        }
      }
    }
    return result;
  }
}
