package com.event;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 * <pre>
 * Service definitions
 * </pre>
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.73.0)",
    comments = "Source: v1/event/service.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class EventStoreServiceGrpc {

  private EventStoreServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "event.EventStoreService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.event.Event,
      com.event.StoreEventResponse> getStoreEventMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "StoreEvent",
      requestType = com.event.Event.class,
      responseType = com.event.StoreEventResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.event.Event,
      com.event.StoreEventResponse> getStoreEventMethod() {
    io.grpc.MethodDescriptor<com.event.Event, com.event.StoreEventResponse> getStoreEventMethod;
    if ((getStoreEventMethod = EventStoreServiceGrpc.getStoreEventMethod) == null) {
      synchronized (EventStoreServiceGrpc.class) {
        if ((getStoreEventMethod = EventStoreServiceGrpc.getStoreEventMethod) == null) {
          EventStoreServiceGrpc.getStoreEventMethod = getStoreEventMethod =
              io.grpc.MethodDescriptor.<com.event.Event, com.event.StoreEventResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "StoreEvent"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.event.Event.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.event.StoreEventResponse.getDefaultInstance()))
              .setSchemaDescriptor(new EventStoreServiceMethodDescriptorSupplier("StoreEvent"))
              .build();
        }
      }
    }
    return getStoreEventMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.event.StoreEventsRequest,
      com.event.StoreEventResponse> getStoreEventsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "StoreEvents",
      requestType = com.event.StoreEventsRequest.class,
      responseType = com.event.StoreEventResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.event.StoreEventsRequest,
      com.event.StoreEventResponse> getStoreEventsMethod() {
    io.grpc.MethodDescriptor<com.event.StoreEventsRequest, com.event.StoreEventResponse> getStoreEventsMethod;
    if ((getStoreEventsMethod = EventStoreServiceGrpc.getStoreEventsMethod) == null) {
      synchronized (EventStoreServiceGrpc.class) {
        if ((getStoreEventsMethod = EventStoreServiceGrpc.getStoreEventsMethod) == null) {
          EventStoreServiceGrpc.getStoreEventsMethod = getStoreEventsMethod =
              io.grpc.MethodDescriptor.<com.event.StoreEventsRequest, com.event.StoreEventResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "StoreEvents"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.event.StoreEventsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.event.StoreEventResponse.getDefaultInstance()))
              .setSchemaDescriptor(new EventStoreServiceMethodDescriptorSupplier("StoreEvents"))
              .build();
        }
      }
    }
    return getStoreEventsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.event.EventStreamRequest,
      com.event.EventStreamResponse> getGetEventStreamMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetEventStream",
      requestType = com.event.EventStreamRequest.class,
      responseType = com.event.EventStreamResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.event.EventStreamRequest,
      com.event.EventStreamResponse> getGetEventStreamMethod() {
    io.grpc.MethodDescriptor<com.event.EventStreamRequest, com.event.EventStreamResponse> getGetEventStreamMethod;
    if ((getGetEventStreamMethod = EventStoreServiceGrpc.getGetEventStreamMethod) == null) {
      synchronized (EventStoreServiceGrpc.class) {
        if ((getGetEventStreamMethod = EventStoreServiceGrpc.getGetEventStreamMethod) == null) {
          EventStoreServiceGrpc.getGetEventStreamMethod = getGetEventStreamMethod =
              io.grpc.MethodDescriptor.<com.event.EventStreamRequest, com.event.EventStreamResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetEventStream"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.event.EventStreamRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.event.EventStreamResponse.getDefaultInstance()))
              .setSchemaDescriptor(new EventStoreServiceMethodDescriptorSupplier("GetEventStream"))
              .build();
        }
      }
    }
    return getGetEventStreamMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.event.EventStreamRequest,
      com.event.Event> getStreamEventsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "StreamEvents",
      requestType = com.event.EventStreamRequest.class,
      responseType = com.event.Event.class,
      methodType = io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
  public static io.grpc.MethodDescriptor<com.event.EventStreamRequest,
      com.event.Event> getStreamEventsMethod() {
    io.grpc.MethodDescriptor<com.event.EventStreamRequest, com.event.Event> getStreamEventsMethod;
    if ((getStreamEventsMethod = EventStoreServiceGrpc.getStreamEventsMethod) == null) {
      synchronized (EventStoreServiceGrpc.class) {
        if ((getStreamEventsMethod = EventStoreServiceGrpc.getStreamEventsMethod) == null) {
          EventStoreServiceGrpc.getStreamEventsMethod = getStreamEventsMethod =
              io.grpc.MethodDescriptor.<com.event.EventStreamRequest, com.event.Event>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.SERVER_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "StreamEvents"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.event.EventStreamRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.event.Event.getDefaultInstance()))
              .setSchemaDescriptor(new EventStoreServiceMethodDescriptorSupplier("StreamEvents"))
              .build();
        }
      }
    }
    return getStreamEventsMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static EventStoreServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<EventStoreServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<EventStoreServiceStub>() {
        @java.lang.Override
        public EventStoreServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new EventStoreServiceStub(channel, callOptions);
        }
      };
    return EventStoreServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static EventStoreServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<EventStoreServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<EventStoreServiceBlockingV2Stub>() {
        @java.lang.Override
        public EventStoreServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new EventStoreServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return EventStoreServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static EventStoreServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<EventStoreServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<EventStoreServiceBlockingStub>() {
        @java.lang.Override
        public EventStoreServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new EventStoreServiceBlockingStub(channel, callOptions);
        }
      };
    return EventStoreServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static EventStoreServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<EventStoreServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<EventStoreServiceFutureStub>() {
        @java.lang.Override
        public EventStoreServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new EventStoreServiceFutureStub(channel, callOptions);
        }
      };
    return EventStoreServiceFutureStub.newStub(factory, channel);
  }

  /**
   * <pre>
   * Service definitions
   * </pre>
   */
  public interface AsyncService {

    /**
     * <pre>
     * Store a single event
     * </pre>
     */
    default void storeEvent(com.event.Event request,
        io.grpc.stub.StreamObserver<com.event.StoreEventResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getStoreEventMethod(), responseObserver);
    }

    /**
     * <pre>
     * Store multiple events atomically
     * </pre>
     */
    default void storeEvents(com.event.StoreEventsRequest request,
        io.grpc.stub.StreamObserver<com.event.StoreEventResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getStoreEventsMethod(), responseObserver);
    }

    /**
     * <pre>
     * Get events by aggregate
     * </pre>
     */
    default void getEventStream(com.event.EventStreamRequest request,
        io.grpc.stub.StreamObserver<com.event.EventStreamResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetEventStreamMethod(), responseObserver);
    }

    /**
     * <pre>
     * Stream events in real-time
     * </pre>
     */
    default void streamEvents(com.event.EventStreamRequest request,
        io.grpc.stub.StreamObserver<com.event.Event> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getStreamEventsMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service EventStoreService.
   * <pre>
   * Service definitions
   * </pre>
   */
  public static abstract class EventStoreServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return EventStoreServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service EventStoreService.
   * <pre>
   * Service definitions
   * </pre>
   */
  public static final class EventStoreServiceStub
      extends io.grpc.stub.AbstractAsyncStub<EventStoreServiceStub> {
    private EventStoreServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected EventStoreServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new EventStoreServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * Store a single event
     * </pre>
     */
    public void storeEvent(com.event.Event request,
        io.grpc.stub.StreamObserver<com.event.StoreEventResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getStoreEventMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Store multiple events atomically
     * </pre>
     */
    public void storeEvents(com.event.StoreEventsRequest request,
        io.grpc.stub.StreamObserver<com.event.StoreEventResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getStoreEventsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Get events by aggregate
     * </pre>
     */
    public void getEventStream(com.event.EventStreamRequest request,
        io.grpc.stub.StreamObserver<com.event.EventStreamResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetEventStreamMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * Stream events in real-time
     * </pre>
     */
    public void streamEvents(com.event.EventStreamRequest request,
        io.grpc.stub.StreamObserver<com.event.Event> responseObserver) {
      io.grpc.stub.ClientCalls.asyncServerStreamingCall(
          getChannel().newCall(getStreamEventsMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service EventStoreService.
   * <pre>
   * Service definitions
   * </pre>
   */
  public static final class EventStoreServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<EventStoreServiceBlockingV2Stub> {
    private EventStoreServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected EventStoreServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new EventStoreServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * Store a single event
     * </pre>
     */
    public com.event.StoreEventResponse storeEvent(com.event.Event request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getStoreEventMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Store multiple events atomically
     * </pre>
     */
    public com.event.StoreEventResponse storeEvents(com.event.StoreEventsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getStoreEventsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get events by aggregate
     * </pre>
     */
    public com.event.EventStreamResponse getEventStream(com.event.EventStreamRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetEventStreamMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Stream events in real-time
     * </pre>
     */
    @io.grpc.ExperimentalApi("https://github.com/grpc/grpc-java/issues/10918")
    public io.grpc.stub.BlockingClientCall<?, com.event.Event>
        streamEvents(com.event.EventStreamRequest request) {
      return io.grpc.stub.ClientCalls.blockingV2ServerStreamingCall(
          getChannel(), getStreamEventsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service EventStoreService.
   * <pre>
   * Service definitions
   * </pre>
   */
  public static final class EventStoreServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<EventStoreServiceBlockingStub> {
    private EventStoreServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected EventStoreServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new EventStoreServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * Store a single event
     * </pre>
     */
    public com.event.StoreEventResponse storeEvent(com.event.Event request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getStoreEventMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Store multiple events atomically
     * </pre>
     */
    public com.event.StoreEventResponse storeEvents(com.event.StoreEventsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getStoreEventsMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Get events by aggregate
     * </pre>
     */
    public com.event.EventStreamResponse getEventStream(com.event.EventStreamRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetEventStreamMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * Stream events in real-time
     * </pre>
     */
    public java.util.Iterator<com.event.Event> streamEvents(
        com.event.EventStreamRequest request) {
      return io.grpc.stub.ClientCalls.blockingServerStreamingCall(
          getChannel(), getStreamEventsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service EventStoreService.
   * <pre>
   * Service definitions
   * </pre>
   */
  public static final class EventStoreServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<EventStoreServiceFutureStub> {
    private EventStoreServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected EventStoreServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new EventStoreServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * Store a single event
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.event.StoreEventResponse> storeEvent(
        com.event.Event request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getStoreEventMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Store multiple events atomically
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.event.StoreEventResponse> storeEvents(
        com.event.StoreEventsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getStoreEventsMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * Get events by aggregate
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.event.EventStreamResponse> getEventStream(
        com.event.EventStreamRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetEventStreamMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_STORE_EVENT = 0;
  private static final int METHODID_STORE_EVENTS = 1;
  private static final int METHODID_GET_EVENT_STREAM = 2;
  private static final int METHODID_STREAM_EVENTS = 3;

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
        case METHODID_STORE_EVENT:
          serviceImpl.storeEvent((com.event.Event) request,
              (io.grpc.stub.StreamObserver<com.event.StoreEventResponse>) responseObserver);
          break;
        case METHODID_STORE_EVENTS:
          serviceImpl.storeEvents((com.event.StoreEventsRequest) request,
              (io.grpc.stub.StreamObserver<com.event.StoreEventResponse>) responseObserver);
          break;
        case METHODID_GET_EVENT_STREAM:
          serviceImpl.getEventStream((com.event.EventStreamRequest) request,
              (io.grpc.stub.StreamObserver<com.event.EventStreamResponse>) responseObserver);
          break;
        case METHODID_STREAM_EVENTS:
          serviceImpl.streamEvents((com.event.EventStreamRequest) request,
              (io.grpc.stub.StreamObserver<com.event.Event>) responseObserver);
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
          getStoreEventMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.event.Event,
              com.event.StoreEventResponse>(
                service, METHODID_STORE_EVENT)))
        .addMethod(
          getStoreEventsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.event.StoreEventsRequest,
              com.event.StoreEventResponse>(
                service, METHODID_STORE_EVENTS)))
        .addMethod(
          getGetEventStreamMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.event.EventStreamRequest,
              com.event.EventStreamResponse>(
                service, METHODID_GET_EVENT_STREAM)))
        .addMethod(
          getStreamEventsMethod(),
          io.grpc.stub.ServerCalls.asyncServerStreamingCall(
            new MethodHandlers<
              com.event.EventStreamRequest,
              com.event.Event>(
                service, METHODID_STREAM_EVENTS)))
        .build();
  }

  private static abstract class EventStoreServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    EventStoreServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.event.ServiceProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("EventStoreService");
    }
  }

  private static final class EventStoreServiceFileDescriptorSupplier
      extends EventStoreServiceBaseDescriptorSupplier {
    EventStoreServiceFileDescriptorSupplier() {}
  }

  private static final class EventStoreServiceMethodDescriptorSupplier
      extends EventStoreServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    EventStoreServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (EventStoreServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new EventStoreServiceFileDescriptorSupplier())
              .addMethod(getStoreEventMethod())
              .addMethod(getStoreEventsMethod())
              .addMethod(getGetEventStreamMethod())
              .addMethod(getStreamEventsMethod())
              .build();
        }
      }
    }
    return result;
  }
}
