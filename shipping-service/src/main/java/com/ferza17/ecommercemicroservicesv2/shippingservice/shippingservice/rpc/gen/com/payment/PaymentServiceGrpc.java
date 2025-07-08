package com.payment;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.73.0)",
    comments = "Source: v1/payment/service.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class PaymentServiceGrpc {

  private PaymentServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "payment.PaymentService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.payment.FindPaymentByIdRequest,
      com.payment.Payment> getFindPaymentByIdMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "FindPaymentById",
      requestType = com.payment.FindPaymentByIdRequest.class,
      responseType = com.payment.Payment.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.payment.FindPaymentByIdRequest,
      com.payment.Payment> getFindPaymentByIdMethod() {
    io.grpc.MethodDescriptor<com.payment.FindPaymentByIdRequest, com.payment.Payment> getFindPaymentByIdMethod;
    if ((getFindPaymentByIdMethod = PaymentServiceGrpc.getFindPaymentByIdMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getFindPaymentByIdMethod = PaymentServiceGrpc.getFindPaymentByIdMethod) == null) {
          PaymentServiceGrpc.getFindPaymentByIdMethod = getFindPaymentByIdMethod =
              io.grpc.MethodDescriptor.<com.payment.FindPaymentByIdRequest, com.payment.Payment>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "FindPaymentById"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.payment.FindPaymentByIdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.payment.Payment.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("FindPaymentById"))
              .build();
        }
      }
    }
    return getFindPaymentByIdMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.payment.FindPaymentByUserIdAndStatusRequest,
      com.payment.Payment> getFindPaymentByUserIdAndStatusMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "FindPaymentByUserIdAndStatus",
      requestType = com.payment.FindPaymentByUserIdAndStatusRequest.class,
      responseType = com.payment.Payment.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.payment.FindPaymentByUserIdAndStatusRequest,
      com.payment.Payment> getFindPaymentByUserIdAndStatusMethod() {
    io.grpc.MethodDescriptor<com.payment.FindPaymentByUserIdAndStatusRequest, com.payment.Payment> getFindPaymentByUserIdAndStatusMethod;
    if ((getFindPaymentByUserIdAndStatusMethod = PaymentServiceGrpc.getFindPaymentByUserIdAndStatusMethod) == null) {
      synchronized (PaymentServiceGrpc.class) {
        if ((getFindPaymentByUserIdAndStatusMethod = PaymentServiceGrpc.getFindPaymentByUserIdAndStatusMethod) == null) {
          PaymentServiceGrpc.getFindPaymentByUserIdAndStatusMethod = getFindPaymentByUserIdAndStatusMethod =
              io.grpc.MethodDescriptor.<com.payment.FindPaymentByUserIdAndStatusRequest, com.payment.Payment>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "FindPaymentByUserIdAndStatus"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.payment.FindPaymentByUserIdAndStatusRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.payment.Payment.getDefaultInstance()))
              .setSchemaDescriptor(new PaymentServiceMethodDescriptorSupplier("FindPaymentByUserIdAndStatus"))
              .build();
        }
      }
    }
    return getFindPaymentByUserIdAndStatusMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static PaymentServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PaymentServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PaymentServiceStub>() {
        @java.lang.Override
        public PaymentServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PaymentServiceStub(channel, callOptions);
        }
      };
    return PaymentServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static PaymentServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PaymentServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PaymentServiceBlockingV2Stub>() {
        @java.lang.Override
        public PaymentServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PaymentServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return PaymentServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static PaymentServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PaymentServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PaymentServiceBlockingStub>() {
        @java.lang.Override
        public PaymentServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PaymentServiceBlockingStub(channel, callOptions);
        }
      };
    return PaymentServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static PaymentServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<PaymentServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<PaymentServiceFutureStub>() {
        @java.lang.Override
        public PaymentServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new PaymentServiceFutureStub(channel, callOptions);
        }
      };
    return PaymentServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    default void findPaymentById(com.payment.FindPaymentByIdRequest request,
        io.grpc.stub.StreamObserver<com.payment.Payment> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getFindPaymentByIdMethod(), responseObserver);
    }

    /**
     */
    default void findPaymentByUserIdAndStatus(com.payment.FindPaymentByUserIdAndStatusRequest request,
        io.grpc.stub.StreamObserver<com.payment.Payment> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getFindPaymentByUserIdAndStatusMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service PaymentService.
   */
  public static abstract class PaymentServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return PaymentServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service PaymentService.
   */
  public static final class PaymentServiceStub
      extends io.grpc.stub.AbstractAsyncStub<PaymentServiceStub> {
    private PaymentServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PaymentServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PaymentServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public void findPaymentById(com.payment.FindPaymentByIdRequest request,
        io.grpc.stub.StreamObserver<com.payment.Payment> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getFindPaymentByIdMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void findPaymentByUserIdAndStatus(com.payment.FindPaymentByUserIdAndStatusRequest request,
        io.grpc.stub.StreamObserver<com.payment.Payment> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getFindPaymentByUserIdAndStatusMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service PaymentService.
   */
  public static final class PaymentServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<PaymentServiceBlockingV2Stub> {
    private PaymentServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PaymentServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PaymentServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public com.payment.Payment findPaymentById(com.payment.FindPaymentByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindPaymentByIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.payment.Payment findPaymentByUserIdAndStatus(com.payment.FindPaymentByUserIdAndStatusRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindPaymentByUserIdAndStatusMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service PaymentService.
   */
  public static final class PaymentServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<PaymentServiceBlockingStub> {
    private PaymentServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PaymentServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PaymentServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public com.payment.Payment findPaymentById(com.payment.FindPaymentByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindPaymentByIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.payment.Payment findPaymentByUserIdAndStatus(com.payment.FindPaymentByUserIdAndStatusRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindPaymentByUserIdAndStatusMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service PaymentService.
   */
  public static final class PaymentServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<PaymentServiceFutureStub> {
    private PaymentServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected PaymentServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new PaymentServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.payment.Payment> findPaymentById(
        com.payment.FindPaymentByIdRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getFindPaymentByIdMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.payment.Payment> findPaymentByUserIdAndStatus(
        com.payment.FindPaymentByUserIdAndStatusRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getFindPaymentByUserIdAndStatusMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_FIND_PAYMENT_BY_ID = 0;
  private static final int METHODID_FIND_PAYMENT_BY_USER_ID_AND_STATUS = 1;

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
        case METHODID_FIND_PAYMENT_BY_ID:
          serviceImpl.findPaymentById((com.payment.FindPaymentByIdRequest) request,
              (io.grpc.stub.StreamObserver<com.payment.Payment>) responseObserver);
          break;
        case METHODID_FIND_PAYMENT_BY_USER_ID_AND_STATUS:
          serviceImpl.findPaymentByUserIdAndStatus((com.payment.FindPaymentByUserIdAndStatusRequest) request,
              (io.grpc.stub.StreamObserver<com.payment.Payment>) responseObserver);
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
          getFindPaymentByIdMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.payment.FindPaymentByIdRequest,
              com.payment.Payment>(
                service, METHODID_FIND_PAYMENT_BY_ID)))
        .addMethod(
          getFindPaymentByUserIdAndStatusMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.payment.FindPaymentByUserIdAndStatusRequest,
              com.payment.Payment>(
                service, METHODID_FIND_PAYMENT_BY_USER_ID_AND_STATUS)))
        .build();
  }

  private static abstract class PaymentServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    PaymentServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.payment.ServiceProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("PaymentService");
    }
  }

  private static final class PaymentServiceFileDescriptorSupplier
      extends PaymentServiceBaseDescriptorSupplier {
    PaymentServiceFileDescriptorSupplier() {}
  }

  private static final class PaymentServiceMethodDescriptorSupplier
      extends PaymentServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    PaymentServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (PaymentServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new PaymentServiceFileDescriptorSupplier())
              .addMethod(getFindPaymentByIdMethod())
              .addMethod(getFindPaymentByUserIdAndStatusMethod())
              .build();
        }
      }
    }
    return result;
  }
}
