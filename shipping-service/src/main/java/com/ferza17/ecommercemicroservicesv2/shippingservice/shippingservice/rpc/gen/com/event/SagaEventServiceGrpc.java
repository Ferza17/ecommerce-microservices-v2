package com.event;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.73.0)",
    comments = "Source: v1/event/service.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class SagaEventServiceGrpc {

  private SagaEventServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "event.SagaEventService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.event.SagaEvent,
      com.event.StoreSagaEventResponse> getStoreSagaEventMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "StoreSagaEvent",
      requestType = com.event.SagaEvent.class,
      responseType = com.event.StoreSagaEventResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.event.SagaEvent,
      com.event.StoreSagaEventResponse> getStoreSagaEventMethod() {
    io.grpc.MethodDescriptor<com.event.SagaEvent, com.event.StoreSagaEventResponse> getStoreSagaEventMethod;
    if ((getStoreSagaEventMethod = SagaEventServiceGrpc.getStoreSagaEventMethod) == null) {
      synchronized (SagaEventServiceGrpc.class) {
        if ((getStoreSagaEventMethod = SagaEventServiceGrpc.getStoreSagaEventMethod) == null) {
          SagaEventServiceGrpc.getStoreSagaEventMethod = getStoreSagaEventMethod =
              io.grpc.MethodDescriptor.<com.event.SagaEvent, com.event.StoreSagaEventResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "StoreSagaEvent"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.event.SagaEvent.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.event.StoreSagaEventResponse.getDefaultInstance()))
              .setSchemaDescriptor(new SagaEventServiceMethodDescriptorSupplier("StoreSagaEvent"))
              .build();
        }
      }
    }
    return getStoreSagaEventMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.event.CompensationEvent,
      com.event.StoreCompensationEventResponse> getStoreCompensationEventMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "StoreCompensationEvent",
      requestType = com.event.CompensationEvent.class,
      responseType = com.event.StoreCompensationEventResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.event.CompensationEvent,
      com.event.StoreCompensationEventResponse> getStoreCompensationEventMethod() {
    io.grpc.MethodDescriptor<com.event.CompensationEvent, com.event.StoreCompensationEventResponse> getStoreCompensationEventMethod;
    if ((getStoreCompensationEventMethod = SagaEventServiceGrpc.getStoreCompensationEventMethod) == null) {
      synchronized (SagaEventServiceGrpc.class) {
        if ((getStoreCompensationEventMethod = SagaEventServiceGrpc.getStoreCompensationEventMethod) == null) {
          SagaEventServiceGrpc.getStoreCompensationEventMethod = getStoreCompensationEventMethod =
              io.grpc.MethodDescriptor.<com.event.CompensationEvent, com.event.StoreCompensationEventResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "StoreCompensationEvent"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.event.CompensationEvent.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.event.StoreCompensationEventResponse.getDefaultInstance()))
              .setSchemaDescriptor(new SagaEventServiceMethodDescriptorSupplier("StoreCompensationEvent"))
              .build();
        }
      }
    }
    return getStoreCompensationEventMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.event.SagaStreamRequest,
      com.event.SagaStreamResponse> getGetSagaEventsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetSagaEvents",
      requestType = com.event.SagaStreamRequest.class,
      responseType = com.event.SagaStreamResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.event.SagaStreamRequest,
      com.event.SagaStreamResponse> getGetSagaEventsMethod() {
    io.grpc.MethodDescriptor<com.event.SagaStreamRequest, com.event.SagaStreamResponse> getGetSagaEventsMethod;
    if ((getGetSagaEventsMethod = SagaEventServiceGrpc.getGetSagaEventsMethod) == null) {
      synchronized (SagaEventServiceGrpc.class) {
        if ((getGetSagaEventsMethod = SagaEventServiceGrpc.getGetSagaEventsMethod) == null) {
          SagaEventServiceGrpc.getGetSagaEventsMethod = getGetSagaEventsMethod =
              io.grpc.MethodDescriptor.<com.event.SagaStreamRequest, com.event.SagaStreamResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetSagaEvents"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.event.SagaStreamRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.event.SagaStreamResponse.getDefaultInstance()))
              .setSchemaDescriptor(new SagaEventServiceMethodDescriptorSupplier("GetSagaEvents"))
              .build();
        }
      }
    }
    return getGetSagaEventsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.event.SagaStreamRequest,
      com.event.SagaEvent> getStreamSagaEventsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "StreamSagaEvents",
      requestType = com.event.SagaStreamRequest.class,
      responseType = com.event.SagaEvent.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<com.event.SagaStreamRequest,
      com.event.SagaEvent> getStreamSagaEventsMethod() {
    io.grpc.MethodDescriptor<com.event.SagaStreamRequest, com.event.SagaEvent> getStreamSagaEventsMethod;
    if ((getStreamSagaEventsMethod = SagaEventServiceGrpc.getStreamSagaEventsMethod) == null) {
      synchronized (SagaEventServiceGrpc.class) {
        if ((getStreamSagaEventsMethod = SagaEventServiceGrpc.getStreamSagaEventsMethod) == null) {
          SagaEventServiceGrpc.getStreamSagaEventsMethod = getStreamSagaEventsMethod =
              io.grpc.MethodDescriptor.<com.event.SagaStreamRequest, com.event.SagaEvent>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "StreamSagaEvents"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.event.SagaStreamRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.event.SagaEvent.getDefaultInstance()))
              .setSchemaDescriptor(new SagaEventServiceMethodDescriptorSupplier("StreamSagaEvents"))
              .build();
        }
      }
    }
    return getStreamSagaEventsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.event.GetCompensationEventsRequest,
      com.event.GetCompensationEventsResponse> getGetCompensationEventsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetCompensationEvents",
      requestType = com.event.GetCompensationEventsRequest.class,
      responseType = com.event.GetCompensationEventsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.event.GetCompensationEventsRequest,
      com.event.GetCompensationEventsResponse> getGetCompensationEventsMethod() {
    io.grpc.MethodDescriptor<com.event.GetCompensationEventsRequest, com.event.GetCompensationEventsResponse> getGetCompensationEventsMethod;
    if ((getGetCompensationEventsMethod = SagaEventServiceGrpc.getGetCompensationEventsMethod) == null) {
      synchronized (SagaEventServiceGrpc.class) {
        if ((getGetCompensationEventsMethod = SagaEventServiceGrpc.getGetCompensationEventsMethod) == null) {
          SagaEventServiceGrpc.getGetCompensationEventsMethod = getGetCompensationEventsMethod =
              io.grpc.MethodDescriptor.<com.event.GetCompensationEventsRequest, com.event.GetCompensationEventsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetCompensationEvents"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.event.GetCompensationEventsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.event.GetCompensationEventsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new SagaEventServiceMethodDescriptorSupplier("GetCompensationEvents"))
              .build();
        }
      }
    }
    return getGetCompensationEventsMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static SagaEventServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SagaEventServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SagaEventServiceStub>() {
        @java.lang.Override
        public SagaEventServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SagaEventServiceStub(channel, callOptions);
        }
      };
    return SagaEventServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static SagaEventServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SagaEventServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SagaEventServiceBlockingV2Stub>() {
        @java.lang.Override
        public SagaEventServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SagaEventServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return SagaEventServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static SagaEventServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SagaEventServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SagaEventServiceBlockingStub>() {
        @java.lang.Override
        public SagaEventServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SagaEventServiceBlockingStub(channel, callOptions);
        }
      };
    return SagaEventServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static SagaEventServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SagaEventServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SagaEventServiceFutureStub>() {
        @java.lang.Override
        public SagaEventServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SagaEventServiceFutureStub(channel, callOptions);
        }
      };
    return SagaEventServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     * <pre>
     * Store saga event
     * </pre>
     */
    default void storeSagaEvent(com.event.SagaEvent request,
        io.grpc.stub.StreamObserver<com.event.StoreSagaEventResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getStoreSagaEventMethod(), responseObserver);
    }

    /**
     * <pre>
     * Store compensation event
     * </pre>
     */
    default void storeCompensationEvent(com.event.CompensationEvent request,
        io.grpc.stub.StreamObserver<com.event.StoreCompensationEventResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getStoreCompensationEventMethod(), responseObserver);
    }

    /**
     * <pre>
     * Get saga events
     * </pre>
     */
    default void getSagaEvents(com.event.SagaStreamRequest request,
        io.grpc.stub.StreamObserver<com.event.SagaStreamResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetSagaEventsMethod(), responseObserver);
    }

    /**
     * <pre>
     * Stream saga events
     * </pre>
     */
    default void streamSagaEvents(com.event.SagaStreamRequest request,
        io.grpc.stub.StreamObserver<com.event.SagaEvent> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getStreamSagaEventsMethod(), responseObserver);
    }

    /**
     * <pre>
     * Get compensation events for a saga
     * </pre>
     */
    default void getCompensationEvents(com.event.GetCompensationEventsRequest request,
        io.grpc.stub.StreamObserver<com.event.GetCompensationEventsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetCompensationEventsMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service SagaEventService.
   */
  public static abstract class SagaEventServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return SagaEventServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service SagaEventService.
   */
  public static final class SagaEventServiceStub
      extends io.grpc.stub.AbstractAsyncStub<SagaEventServiceStub> {
    private SagaEventServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SagaEventServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SagaEventServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * Store saga event
     * </pre>
     */
    public void storeSagaEvent(com.event.SagaEvent request,
        io.grpc.stub.StreamObserver<com.event.StoreSagaEventResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getStoreSagaEventMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Store compensation event
     * </pre>
     */
    public void storeCompensationEvent(com.event.CompensationEvent request,
        io.grpc.stub.StreamObserver<com.event.StoreCompensationEventResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getStoreCompensationEventMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Get saga events
     * </pre>
     */
    public void getSagaEvents(com.event.SagaStreamRequest request,
        io.grpc.stub.StreamObserver<com.event.SagaStreamResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetSagaEventsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Stream saga events
     * </pre>
     */
    public void streamSagaEvents(com.event.SagaStreamRequest request,
        io.grpc.stub.StreamObserver<com.event.SagaEvent> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getStreamSagaEventsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Get compensation events for a saga
     * </pre>
     */
    public void getCompensationEvents(com.event.GetCompensationEventsRequest request,
        io.grpc.stub.StreamObserver<com.event.GetCompensationEventsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetCompensationEventsMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service SagaEventService.
   */
  public static final class SagaEventServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<SagaEventServiceBlockingV2Stub> {
    private SagaEventServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SagaEventServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SagaEventServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * Store saga event
     * </pre>
     */
    public com.event.StoreSagaEventResponse storeSagaEvent(com.event.SagaEvent request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getStoreSagaEventMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Store compensation event
     * </pre>
     */
    public com.event.StoreCompensationEventResponse storeCompensationEvent(com.event.CompensationEvent request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getStoreCompensationEventMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get saga events
     * </pre>
     */
    public com.event.SagaStreamResponse getSagaEvents(com.event.SagaStreamRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetSagaEventsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Stream saga events
     * </pre>
     */
    @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/10918")
    public io.grpc.stub.BlockingClientCall<?, com.event.SagaEvent>
        streamSagaEvents(com.event.SagaStreamRequest request) {
      return io.grpc.stub.ClientCalls.blockingV2ServerStreamingCall(
          getChannel(), getStreamSagaEventsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get compensation events for a saga
     * </pre>
     */
    public com.event.GetCompensationEventsResponse getCompensationEvents(com.event.GetCompensationEventsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetCompensationEventsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service SagaEventService.
   */
  public static final class SagaEventServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<SagaEventServiceBlockingStub> {
    private SagaEventServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SagaEventServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SagaEventServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * Store saga event
     * </pre>
     */
    public com.event.StoreSagaEventResponse storeSagaEvent(com.event.SagaEvent request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getStoreSagaEventMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Store compensation event
     * </pre>
     */
    public com.event.StoreCompensationEventResponse storeCompensationEvent(com.event.CompensationEvent request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getStoreCompensationEventMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get saga events
     * </pre>
     */
    public com.event.SagaStreamResponse getSagaEvents(com.event.SagaStreamRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetSagaEventsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Stream saga events
     * </pre>
     */
    public java.util.Iterator<com.event.SagaEvent> streamSagaEvents(
        com.event.SagaStreamRequest request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getStreamSagaEventsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get compensation events for a saga
     * </pre>
     */
    public com.event.GetCompensationEventsResponse getCompensationEvents(com.event.GetCompensationEventsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetCompensationEventsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service SagaEventService.
   */
  public static final class SagaEventServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<SagaEventServiceFutureStub> {
    private SagaEventServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SagaEventServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SagaEventServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * Store saga event
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.event.StoreSagaEventResponse> storeSagaEvent(
        com.event.SagaEvent request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getStoreSagaEventMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Store compensation event
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.event.StoreCompensationEventResponse> storeCompensationEvent(
        com.event.CompensationEvent request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getStoreCompensationEventMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Get saga events
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.event.SagaStreamResponse> getSagaEvents(
        com.event.SagaStreamRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetSagaEventsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Get compensation events for a saga
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.event.GetCompensationEventsResponse> getCompensationEvents(
        com.event.GetCompensationEventsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetCompensationEventsMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_STORE_SAGA_EVENT = 0;
  private static final int METHODID_STORE_COMPENSATION_EVENT = 1;
  private static final int METHODID_GET_SAGA_EVENTS = 2;
  private static final int METHODID_STREAM_SAGA_EVENTS = 3;
  private static final int METHODID_GET_COMPENSATION_EVENTS = 4;

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
        case METHODID_STORE_SAGA_EVENT:
          serviceImpl.storeSagaEvent((com.event.SagaEvent) request,
              (io.grpc.stub.StreamObserver<com.event.StoreSagaEventResponse>) responseObserver);
          break;
        case METHODID_STORE_COMPENSATION_EVENT:
          serviceImpl.storeCompensationEvent((com.event.CompensationEvent) request,
              (io.grpc.stub.StreamObserver<com.event.StoreCompensationEventResponse>) responseObserver);
          break;
        case METHODID_GET_SAGA_EVENTS:
          serviceImpl.getSagaEvents((com.event.SagaStreamRequest) request,
              (io.grpc.stub.StreamObserver<com.event.SagaStreamResponse>) responseObserver);
          break;
        case METHODID_STREAM_SAGA_EVENTS:
          serviceImpl.streamSagaEvents((com.event.SagaStreamRequest) request,
              (io.grpc.stub.StreamObserver<com.event.SagaEvent>) responseObserver);
          break;
        case METHODID_GET_COMPENSATION_EVENTS:
          serviceImpl.getCompensationEvents((com.event.GetCompensationEventsRequest) request,
              (io.grpc.stub.StreamObserver<com.event.GetCompensationEventsResponse>) responseObserver);
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
          getStoreSagaEventMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.event.SagaEvent,
              com.event.StoreSagaEventResponse>(
                service, METHODID_STORE_SAGA_EVENT)))
        .addMethod(
          getStoreCompensationEventMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.event.CompensationEvent,
              com.event.StoreCompensationEventResponse>(
                service, METHODID_STORE_COMPENSATION_EVENT)))
        .addMethod(
          getGetSagaEventsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.event.SagaStreamRequest,
              com.event.SagaStreamResponse>(
                service, METHODID_GET_SAGA_EVENTS)))
        .addMethod(
          getStreamSagaEventsMethod(),
          io.grpc.stub.ServerCalls.asyncServerStreamingCall(
            new MethodHandlers<
              com.event.SagaStreamRequest,
              com.event.SagaEvent>(
                service, METHODID_STREAM_SAGA_EVENTS)))
        .addMethod(
          getGetCompensationEventsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.event.GetCompensationEventsRequest,
              com.event.GetCompensationEventsResponse>(
                service, METHODID_GET_COMPENSATION_EVENTS)))
        .build();
  }

  private static abstract class SagaEventServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    SagaEventServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.event.ServiceProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("SagaEventService");
    }
  }

  private static final class SagaEventServiceFileDescriptorSupplier
      extends SagaEventServiceBaseDescriptorSupplier {
    SagaEventServiceFileDescriptorSupplier() {}
  }

  private static final class SagaEventServiceMethodDescriptorSupplier
      extends SagaEventServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    SagaEventServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (SagaEventServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new SagaEventServiceFileDescriptorSupplier())
              .addMethod(getStoreSagaEventMethod())
              .addMethod(getStoreCompensationEventMethod())
              .addMethod(getGetSagaEventsMethod())
              .addMethod(getStreamSagaEventsMethod())
              .addMethod(getGetCompensationEventsMethod())
              .build();
        }
      }
    }
    return result;
  }
}
