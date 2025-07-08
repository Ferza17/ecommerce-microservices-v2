package com.payment;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.73.0)",
    comments = "Source: v1/payment/service.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class PaymentProviderServiceGrpc {

  private PaymentProviderServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "payment.PaymentProviderService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.payment.FindPaymentProvidersRequest,
      com.payment.FindPaymentProvidersResponse> getFindPaymentProvidersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "FindPaymentProviders",
      requestType = com.payment.FindPaymentProvidersRequest.class,
      responseType = com.payment.FindPaymentProvidersResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.payment.FindPaymentProvidersRequest,
      com.payment.FindPaymentProvidersResponse> getFindPaymentProvidersMethod() {
    io.grpc.MethodDescriptor<com.payment.FindPaymentProvidersRequest, com.payment.FindPaymentProvidersResponse> getFindPaymentProvidersMethod;
    if ((getFindPaymentProvidersMethod = PaymentProviderServiceGrpc.getFindPaymentProvidersMethod) == null) {
      synchronized (PaymentProviderServiceGrpc.class) {
        if ((getFindPaymentProvidersMethod = PaymentProviderServiceGrpc.getFindPaymentProvidersMethod) == null) {
          PaymentProviderServiceGrpc.getFindPaymentProvidersMethod = getFindPaymentProvidersMethod =
              io.grpc.MethodDescriptor.<com.payment.FindPaymentProvidersRequest, com.payment.FindPaymentProvidersResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "FindPaymentProviders"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.payment.FindPaymentProvidersRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.payment.FindPaymentProvidersResponse.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentProviderServiceMethodDescriptorSupplier("FindPaymentProviders"))
              .build();
        }
      }
    }
    return getFindPaymentProvidersMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.payment.FindPaymentProviderByIdRequest,
      com.payment.Provider> getFindPaymentProviderByIdMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "FindPaymentProviderById",
      requestType = com.payment.FindPaymentProviderByIdRequest.class,
      responseType = com.payment.Provider.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.payment.FindPaymentProviderByIdRequest,
      com.payment.Provider> getFindPaymentProviderByIdMethod() {
    io.grpc.MethodDescriptor<com.payment.FindPaymentProviderByIdRequest, com.payment.Provider> getFindPaymentProviderByIdMethod;
    if ((getFindPaymentProviderByIdMethod = PaymentProviderServiceGrpc.getFindPaymentProviderByIdMethod) == null) {
      synchronized (PaymentProviderServiceGrpc.class) {
        if ((getFindPaymentProviderByIdMethod = PaymentProviderServiceGrpc.getFindPaymentProviderByIdMethod) == null) {
          PaymentProviderServiceGrpc.getFindPaymentProviderByIdMethod = getFindPaymentProviderByIdMethod =
              io.grpc.MethodDescriptor.<com.payment.FindPaymentProviderByIdRequest, com.payment.Provider>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "FindPaymentProviderById"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.payment.FindPaymentProviderByIdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.payment.Provider.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentProviderServiceMethodDescriptorSupplier("FindPaymentProviderById"))
              .build();
        }
      }
    }
    return getFindPaymentProviderByIdMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static PaymentProviderServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PaymentProviderServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PaymentProviderServiceStub>() {
        @java.lang.Override
        public PaymentProviderServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PaymentProviderServiceStub(channel, callOptions);
        }
      };
    return PaymentProviderServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static PaymentProviderServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PaymentProviderServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PaymentProviderServiceBlockingV2Stub>() {
        @java.lang.Override
        public PaymentProviderServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PaymentProviderServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return PaymentProviderServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static PaymentProviderServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PaymentProviderServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PaymentProviderServiceBlockingStub>() {
        @java.lang.Override
        public PaymentProviderServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PaymentProviderServiceBlockingStub(channel, callOptions);
        }
      };
    return PaymentProviderServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static PaymentProviderServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PaymentProviderServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PaymentProviderServiceFutureStub>() {
        @java.lang.Override
        public PaymentProviderServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PaymentProviderServiceFutureStub(channel, callOptions);
        }
      };
    return PaymentProviderServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    default void findPaymentProviders(com.payment.FindPaymentProvidersRequest request,
        io.grpc.stub.StreamObserver<com.payment.FindPaymentProvidersResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getFindPaymentProvidersMethod(), responseObserver);
    }

    /**
     */
    default void findPaymentProviderById(com.payment.FindPaymentProviderByIdRequest request,
        io.grpc.stub.StreamObserver<com.payment.Provider> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getFindPaymentProviderByIdMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service PaymentProviderService.
   */
  public static abstract class PaymentProviderServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return PaymentProviderServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service PaymentProviderService.
   */
  public static final class PaymentProviderServiceStub
      extends io.grpc.stub.AbstractAsyncStub<PaymentProviderServiceStub> {
    private PaymentProviderServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PaymentProviderServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PaymentProviderServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public void findPaymentProviders(com.payment.FindPaymentProvidersRequest request,
        io.grpc.stub.StreamObserver<com.payment.FindPaymentProvidersResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getFindPaymentProvidersMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void findPaymentProviderById(com.payment.FindPaymentProviderByIdRequest request,
        io.grpc.stub.StreamObserver<com.payment.Provider> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getFindPaymentProviderByIdMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service PaymentProviderService.
   */
  public static final class PaymentProviderServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<PaymentProviderServiceBlockingV2Stub> {
    private PaymentProviderServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PaymentProviderServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PaymentProviderServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public com.payment.FindPaymentProvidersResponse findPaymentProviders(com.payment.FindPaymentProvidersRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindPaymentProvidersMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.payment.Provider findPaymentProviderById(com.payment.FindPaymentProviderByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindPaymentProviderByIdMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service PaymentProviderService.
   */
  public static final class PaymentProviderServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<PaymentProviderServiceBlockingStub> {
    private PaymentProviderServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PaymentProviderServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PaymentProviderServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public com.payment.FindPaymentProvidersResponse findPaymentProviders(com.payment.FindPaymentProvidersRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindPaymentProvidersMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.payment.Provider findPaymentProviderById(com.payment.FindPaymentProviderByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindPaymentProviderByIdMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service PaymentProviderService.
   */
  public static final class PaymentProviderServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<PaymentProviderServiceFutureStub> {
    private PaymentProviderServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PaymentProviderServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PaymentProviderServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.payment.FindPaymentProvidersResponse> findPaymentProviders(
        com.payment.FindPaymentProvidersRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getFindPaymentProvidersMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.payment.Provider> findPaymentProviderById(
        com.payment.FindPaymentProviderByIdRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getFindPaymentProviderByIdMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_FIND_PAYMENT_PROVIDERS = 0;
  private static final int METHODID_FIND_PAYMENT_PROVIDER_BY_ID = 1;

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
        case METHODID_FIND_PAYMENT_PROVIDERS:
          serviceImpl.findPaymentProviders((com.payment.FindPaymentProvidersRequest) request,
              (io.grpc.stub.StreamObserver<com.payment.FindPaymentProvidersResponse>) responseObserver);
          break;
        case METHODID_FIND_PAYMENT_PROVIDER_BY_ID:
          serviceImpl.findPaymentProviderById((com.payment.FindPaymentProviderByIdRequest) request,
              (io.grpc.stub.StreamObserver<com.payment.Provider>) responseObserver);
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
          getFindPaymentProvidersMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.payment.FindPaymentProvidersRequest,
              com.payment.FindPaymentProvidersResponse>(
                service, METHODID_FIND_PAYMENT_PROVIDERS)))
        .addMethod(
          getFindPaymentProviderByIdMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.payment.FindPaymentProviderByIdRequest,
              com.payment.Provider>(
                service, METHODID_FIND_PAYMENT_PROVIDER_BY_ID)))
        .build();
  }

  private static abstract class PaymentProviderServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    PaymentProviderServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.payment.ServiceProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("PaymentProviderService");
    }
  }

  private static final class PaymentProviderServiceFileDescriptorSupplier
      extends PaymentProviderServiceBaseDescriptorSupplier {
    PaymentProviderServiceFileDescriptorSupplier() {}
  }

  private static final class PaymentProviderServiceMethodDescriptorSupplier
      extends PaymentProviderServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    PaymentProviderServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (PaymentProviderServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new PaymentProviderServiceFileDescriptorSupplier())
              .addMethod(getFindPaymentProvidersMethod())
              .addMethod(getFindPaymentProviderByIdMethod())
              .build();
        }
      }
    }
    return result;
  }
}
