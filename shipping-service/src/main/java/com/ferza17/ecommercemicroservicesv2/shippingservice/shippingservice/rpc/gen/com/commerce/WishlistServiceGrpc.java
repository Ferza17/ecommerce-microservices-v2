package com.commerce;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.73.0)",
    comments = "Source: v1/commerce/service.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class WishlistServiceGrpc {

  private WishlistServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "commerce.WishlistService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.commerce.FindWishlistItemWithPaginationRequest,
      com.commerce.FindWishlistItemWithPaginationResponse> getFindWishlistItemWithPaginationMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "FindWishlistItemWithPagination",
      requestType = com.commerce.FindWishlistItemWithPaginationRequest.class,
      responseType = com.commerce.FindWishlistItemWithPaginationResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.commerce.FindWishlistItemWithPaginationRequest,
      com.commerce.FindWishlistItemWithPaginationResponse> getFindWishlistItemWithPaginationMethod() {
    io.grpc.MethodDescriptor<com.commerce.FindWishlistItemWithPaginationRequest, com.commerce.FindWishlistItemWithPaginationResponse> getFindWishlistItemWithPaginationMethod;
    if ((getFindWishlistItemWithPaginationMethod = WishlistServiceGrpc.getFindWishlistItemWithPaginationMethod) == null) {
      synchronized (WishlistServiceGrpc.class) {
        if ((getFindWishlistItemWithPaginationMethod = WishlistServiceGrpc.getFindWishlistItemWithPaginationMethod) == null) {
          WishlistServiceGrpc.getFindWishlistItemWithPaginationMethod = getFindWishlistItemWithPaginationMethod =
              io.grpc.MethodDescriptor.<com.commerce.FindWishlistItemWithPaginationRequest, com.commerce.FindWishlistItemWithPaginationResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "FindWishlistItemWithPagination"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.commerce.FindWishlistItemWithPaginationRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.commerce.FindWishlistItemWithPaginationResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WishlistServiceMethodDescriptorSupplier("FindWishlistItemWithPagination"))
              .build();
        }
      }
    }
    return getFindWishlistItemWithPaginationMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.commerce.CreateWishlistItemRequest,
      com.commerce.CreateWishlistItemResponse> getCreateWishlistItemMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateWishlistItem",
      requestType = com.commerce.CreateWishlistItemRequest.class,
      responseType = com.commerce.CreateWishlistItemResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.commerce.CreateWishlistItemRequest,
      com.commerce.CreateWishlistItemResponse> getCreateWishlistItemMethod() {
    io.grpc.MethodDescriptor<com.commerce.CreateWishlistItemRequest, com.commerce.CreateWishlistItemResponse> getCreateWishlistItemMethod;
    if ((getCreateWishlistItemMethod = WishlistServiceGrpc.getCreateWishlistItemMethod) == null) {
      synchronized (WishlistServiceGrpc.class) {
        if ((getCreateWishlistItemMethod = WishlistServiceGrpc.getCreateWishlistItemMethod) == null) {
          WishlistServiceGrpc.getCreateWishlistItemMethod = getCreateWishlistItemMethod =
              io.grpc.MethodDescriptor.<com.commerce.CreateWishlistItemRequest, com.commerce.CreateWishlistItemResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateWishlistItem"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.commerce.CreateWishlistItemRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.commerce.CreateWishlistItemResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WishlistServiceMethodDescriptorSupplier("CreateWishlistItem"))
              .build();
        }
      }
    }
    return getCreateWishlistItemMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.commerce.DeleteWishlistItemByIdRequest,
      com.commerce.DeleteWishlistItemByIdResponse> getDeleteWishlistItemByIdMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DeleteWishlistItemById",
      requestType = com.commerce.DeleteWishlistItemByIdRequest.class,
      responseType = com.commerce.DeleteWishlistItemByIdResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.commerce.DeleteWishlistItemByIdRequest,
      com.commerce.DeleteWishlistItemByIdResponse> getDeleteWishlistItemByIdMethod() {
    io.grpc.MethodDescriptor<com.commerce.DeleteWishlistItemByIdRequest, com.commerce.DeleteWishlistItemByIdResponse> getDeleteWishlistItemByIdMethod;
    if ((getDeleteWishlistItemByIdMethod = WishlistServiceGrpc.getDeleteWishlistItemByIdMethod) == null) {
      synchronized (WishlistServiceGrpc.class) {
        if ((getDeleteWishlistItemByIdMethod = WishlistServiceGrpc.getDeleteWishlistItemByIdMethod) == null) {
          WishlistServiceGrpc.getDeleteWishlistItemByIdMethod = getDeleteWishlistItemByIdMethod =
              io.grpc.MethodDescriptor.<com.commerce.DeleteWishlistItemByIdRequest, com.commerce.DeleteWishlistItemByIdResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DeleteWishlistItemById"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.commerce.DeleteWishlistItemByIdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.commerce.DeleteWishlistItemByIdResponse.getDefaultInstance()))
              .setSchemaDescriptor(new WishlistServiceMethodDescriptorSupplier("DeleteWishlistItemById"))
              .build();
        }
      }
    }
    return getDeleteWishlistItemByIdMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static WishlistServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<WishlistServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<WishlistServiceStub>() {
        @java.lang.Override
        public WishlistServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new WishlistServiceStub(channel, callOptions);
        }
      };
    return WishlistServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static WishlistServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<WishlistServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<WishlistServiceBlockingV2Stub>() {
        @java.lang.Override
        public WishlistServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new WishlistServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return WishlistServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static WishlistServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<WishlistServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<WishlistServiceBlockingStub>() {
        @java.lang.Override
        public WishlistServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new WishlistServiceBlockingStub(channel, callOptions);
        }
      };
    return WishlistServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static WishlistServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<WishlistServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<WishlistServiceFutureStub>() {
        @java.lang.Override
        public WishlistServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new WishlistServiceFutureStub(channel, callOptions);
        }
      };
    return WishlistServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    default void findWishlistItemWithPagination(com.commerce.FindWishlistItemWithPaginationRequest request,
        io.grpc.stub.StreamObserver<com.commerce.FindWishlistItemWithPaginationResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getFindWishlistItemWithPaginationMethod(), responseObserver);
    }

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    default void createWishlistItem(com.commerce.CreateWishlistItemRequest request,
        io.grpc.stub.StreamObserver<com.commerce.CreateWishlistItemResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateWishlistItemMethod(), responseObserver);
    }

    /**
     */
    default void deleteWishlistItemById(com.commerce.DeleteWishlistItemByIdRequest request,
        io.grpc.stub.StreamObserver<com.commerce.DeleteWishlistItemByIdResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeleteWishlistItemByIdMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service WishlistService.
   */
  public static abstract class WishlistServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return WishlistServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service WishlistService.
   */
  public static final class WishlistServiceStub
      extends io.grpc.stub.AbstractAsyncStub<WishlistServiceStub> {
    private WishlistServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected WishlistServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new WishlistServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public void findWishlistItemWithPagination(com.commerce.FindWishlistItemWithPaginationRequest request,
        io.grpc.stub.StreamObserver<com.commerce.FindWishlistItemWithPaginationResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getFindWishlistItemWithPaginationMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    public void createWishlistItem(com.commerce.CreateWishlistItemRequest request,
        io.grpc.stub.StreamObserver<com.commerce.CreateWishlistItemResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateWishlistItemMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void deleteWishlistItemById(com.commerce.DeleteWishlistItemByIdRequest request,
        io.grpc.stub.StreamObserver<com.commerce.DeleteWishlistItemByIdResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeleteWishlistItemByIdMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service WishlistService.
   */
  public static final class WishlistServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<WishlistServiceBlockingV2Stub> {
    private WishlistServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected WishlistServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new WishlistServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public com.commerce.FindWishlistItemWithPaginationResponse findWishlistItemWithPagination(com.commerce.FindWishlistItemWithPaginationRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindWishlistItemWithPaginationMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    public com.commerce.CreateWishlistItemResponse createWishlistItem(com.commerce.CreateWishlistItemRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateWishlistItemMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.commerce.DeleteWishlistItemByIdResponse deleteWishlistItemById(com.commerce.DeleteWishlistItemByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteWishlistItemByIdMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service WishlistService.
   */
  public static final class WishlistServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<WishlistServiceBlockingStub> {
    private WishlistServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected WishlistServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new WishlistServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public com.commerce.FindWishlistItemWithPaginationResponse findWishlistItemWithPagination(com.commerce.FindWishlistItemWithPaginationRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindWishlistItemWithPaginationMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    public com.commerce.CreateWishlistItemResponse createWishlistItem(com.commerce.CreateWishlistItemRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateWishlistItemMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.commerce.DeleteWishlistItemByIdResponse deleteWishlistItemById(com.commerce.DeleteWishlistItemByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteWishlistItemByIdMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service WishlistService.
   */
  public static final class WishlistServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<WishlistServiceFutureStub> {
    private WishlistServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected WishlistServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new WishlistServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.commerce.FindWishlistItemWithPaginationResponse> findWishlistItemWithPagination(
        com.commerce.FindWishlistItemWithPaginationRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getFindWishlistItemWithPaginationMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.commerce.CreateWishlistItemResponse> createWishlistItem(
        com.commerce.CreateWishlistItemRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateWishlistItemMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.commerce.DeleteWishlistItemByIdResponse> deleteWishlistItemById(
        com.commerce.DeleteWishlistItemByIdRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeleteWishlistItemByIdMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_FIND_WISHLIST_ITEM_WITH_PAGINATION = 0;
  private static final int METHODID_CREATE_WISHLIST_ITEM = 1;
  private static final int METHODID_DELETE_WISHLIST_ITEM_BY_ID = 2;

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
        case METHODID_FIND_WISHLIST_ITEM_WITH_PAGINATION:
          serviceImpl.findWishlistItemWithPagination((com.commerce.FindWishlistItemWithPaginationRequest) request,
              (io.grpc.stub.StreamObserver<com.commerce.FindWishlistItemWithPaginationResponse>) responseObserver);
          break;
        case METHODID_CREATE_WISHLIST_ITEM:
          serviceImpl.createWishlistItem((com.commerce.CreateWishlistItemRequest) request,
              (io.grpc.stub.StreamObserver<com.commerce.CreateWishlistItemResponse>) responseObserver);
          break;
        case METHODID_DELETE_WISHLIST_ITEM_BY_ID:
          serviceImpl.deleteWishlistItemById((com.commerce.DeleteWishlistItemByIdRequest) request,
              (io.grpc.stub.StreamObserver<com.commerce.DeleteWishlistItemByIdResponse>) responseObserver);
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
          getFindWishlistItemWithPaginationMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.commerce.FindWishlistItemWithPaginationRequest,
              com.commerce.FindWishlistItemWithPaginationResponse>(
                service, METHODID_FIND_WISHLIST_ITEM_WITH_PAGINATION)))
        .addMethod(
          getCreateWishlistItemMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.commerce.CreateWishlistItemRequest,
              com.commerce.CreateWishlistItemResponse>(
                service, METHODID_CREATE_WISHLIST_ITEM)))
        .addMethod(
          getDeleteWishlistItemByIdMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.commerce.DeleteWishlistItemByIdRequest,
              com.commerce.DeleteWishlistItemByIdResponse>(
                service, METHODID_DELETE_WISHLIST_ITEM_BY_ID)))
        .build();
  }

  private static abstract class WishlistServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    WishlistServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.commerce.ServiceProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("WishlistService");
    }
  }

  private static final class WishlistServiceFileDescriptorSupplier
      extends WishlistServiceBaseDescriptorSupplier {
    WishlistServiceFileDescriptorSupplier() {}
  }

  private static final class WishlistServiceMethodDescriptorSupplier
      extends WishlistServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    WishlistServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (WishlistServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new WishlistServiceFileDescriptorSupplier())
              .addMethod(getFindWishlistItemWithPaginationMethod())
              .addMethod(getCreateWishlistItemMethod())
              .addMethod(getDeleteWishlistItemByIdMethod())
              .build();
        }
      }
    }
    return result;
  }
}
