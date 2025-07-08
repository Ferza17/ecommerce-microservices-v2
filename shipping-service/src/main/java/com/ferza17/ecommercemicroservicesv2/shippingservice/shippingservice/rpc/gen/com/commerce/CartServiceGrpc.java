package com.commerce;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.73.0)",
    comments = "Source: v1/commerce/service.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class CartServiceGrpc {

  private CartServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "commerce.CartService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.commerce.CreateCartItemRequest,
      com.commerce.CreateCartItemResponse> getCreateCartItemMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateCartItem",
      requestType = com.commerce.CreateCartItemRequest.class,
      responseType = com.commerce.CreateCartItemResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.commerce.CreateCartItemRequest,
      com.commerce.CreateCartItemResponse> getCreateCartItemMethod() {
    io.grpc.MethodDescriptor<com.commerce.CreateCartItemRequest, com.commerce.CreateCartItemResponse> getCreateCartItemMethod;
    if ((getCreateCartItemMethod = CartServiceGrpc.getCreateCartItemMethod) == null) {
      synchronized (CartServiceGrpc.class) {
        if ((getCreateCartItemMethod = CartServiceGrpc.getCreateCartItemMethod) == null) {
          CartServiceGrpc.getCreateCartItemMethod = getCreateCartItemMethod =
              io.grpc.MethodDescriptor.<com.commerce.CreateCartItemRequest, com.commerce.CreateCartItemResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateCartItem"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.commerce.CreateCartItemRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.commerce.CreateCartItemResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CartServiceMethodDescriptorSupplier("CreateCartItem"))
              .build();
        }
      }
    }
    return getCreateCartItemMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.commerce.FindCartItemByIdRequest,
      com.commerce.CartItem> getFindCartItemByIdMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "FindCartItemById",
      requestType = com.commerce.FindCartItemByIdRequest.class,
      responseType = com.commerce.CartItem.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.commerce.FindCartItemByIdRequest,
      com.commerce.CartItem> getFindCartItemByIdMethod() {
    io.grpc.MethodDescriptor<com.commerce.FindCartItemByIdRequest, com.commerce.CartItem> getFindCartItemByIdMethod;
    if ((getFindCartItemByIdMethod = CartServiceGrpc.getFindCartItemByIdMethod) == null) {
      synchronized (CartServiceGrpc.class) {
        if ((getFindCartItemByIdMethod = CartServiceGrpc.getFindCartItemByIdMethod) == null) {
          CartServiceGrpc.getFindCartItemByIdMethod = getFindCartItemByIdMethod =
              io.grpc.MethodDescriptor.<com.commerce.FindCartItemByIdRequest, com.commerce.CartItem>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "FindCartItemById"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.commerce.FindCartItemByIdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.commerce.CartItem.getDefaultInstance()))
              .setSchemaDescriptor(new CartServiceMethodDescriptorSupplier("FindCartItemById"))
              .build();
        }
      }
    }
    return getFindCartItemByIdMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.commerce.FindCartItemsWithPaginationRequest,
      com.commerce.FindCartItemsWithPaginationResponse> getFindCartItemsWithPaginationMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "FindCartItemsWithPagination",
      requestType = com.commerce.FindCartItemsWithPaginationRequest.class,
      responseType = com.commerce.FindCartItemsWithPaginationResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.commerce.FindCartItemsWithPaginationRequest,
      com.commerce.FindCartItemsWithPaginationResponse> getFindCartItemsWithPaginationMethod() {
    io.grpc.MethodDescriptor<com.commerce.FindCartItemsWithPaginationRequest, com.commerce.FindCartItemsWithPaginationResponse> getFindCartItemsWithPaginationMethod;
    if ((getFindCartItemsWithPaginationMethod = CartServiceGrpc.getFindCartItemsWithPaginationMethod) == null) {
      synchronized (CartServiceGrpc.class) {
        if ((getFindCartItemsWithPaginationMethod = CartServiceGrpc.getFindCartItemsWithPaginationMethod) == null) {
          CartServiceGrpc.getFindCartItemsWithPaginationMethod = getFindCartItemsWithPaginationMethod =
              io.grpc.MethodDescriptor.<com.commerce.FindCartItemsWithPaginationRequest, com.commerce.FindCartItemsWithPaginationResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "FindCartItemsWithPagination"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.commerce.FindCartItemsWithPaginationRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.commerce.FindCartItemsWithPaginationResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CartServiceMethodDescriptorSupplier("FindCartItemsWithPagination"))
              .build();
        }
      }
    }
    return getFindCartItemsWithPaginationMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.commerce.UpdateCartItemByIdRequest,
      com.commerce.UpdateCartItemByIdResponse> getUpdateCartItemByIdMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateCartItemById",
      requestType = com.commerce.UpdateCartItemByIdRequest.class,
      responseType = com.commerce.UpdateCartItemByIdResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.commerce.UpdateCartItemByIdRequest,
      com.commerce.UpdateCartItemByIdResponse> getUpdateCartItemByIdMethod() {
    io.grpc.MethodDescriptor<com.commerce.UpdateCartItemByIdRequest, com.commerce.UpdateCartItemByIdResponse> getUpdateCartItemByIdMethod;
    if ((getUpdateCartItemByIdMethod = CartServiceGrpc.getUpdateCartItemByIdMethod) == null) {
      synchronized (CartServiceGrpc.class) {
        if ((getUpdateCartItemByIdMethod = CartServiceGrpc.getUpdateCartItemByIdMethod) == null) {
          CartServiceGrpc.getUpdateCartItemByIdMethod = getUpdateCartItemByIdMethod =
              io.grpc.MethodDescriptor.<com.commerce.UpdateCartItemByIdRequest, com.commerce.UpdateCartItemByIdResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateCartItemById"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.commerce.UpdateCartItemByIdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.commerce.UpdateCartItemByIdResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CartServiceMethodDescriptorSupplier("UpdateCartItemById"))
              .build();
        }
      }
    }
    return getUpdateCartItemByIdMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.commerce.DeleteCartItemByIdRequest,
      com.commerce.DeleteCartItemByIdResponse> getDeleteCartItemByIdMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DeleteCartItemById",
      requestType = com.commerce.DeleteCartItemByIdRequest.class,
      responseType = com.commerce.DeleteCartItemByIdResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.commerce.DeleteCartItemByIdRequest,
      com.commerce.DeleteCartItemByIdResponse> getDeleteCartItemByIdMethod() {
    io.grpc.MethodDescriptor<com.commerce.DeleteCartItemByIdRequest, com.commerce.DeleteCartItemByIdResponse> getDeleteCartItemByIdMethod;
    if ((getDeleteCartItemByIdMethod = CartServiceGrpc.getDeleteCartItemByIdMethod) == null) {
      synchronized (CartServiceGrpc.class) {
        if ((getDeleteCartItemByIdMethod = CartServiceGrpc.getDeleteCartItemByIdMethod) == null) {
          CartServiceGrpc.getDeleteCartItemByIdMethod = getDeleteCartItemByIdMethod =
              io.grpc.MethodDescriptor.<com.commerce.DeleteCartItemByIdRequest, com.commerce.DeleteCartItemByIdResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DeleteCartItemById"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.commerce.DeleteCartItemByIdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.commerce.DeleteCartItemByIdResponse.getDefaultInstance()))
              .setSchemaDescriptor(new CartServiceMethodDescriptorSupplier("DeleteCartItemById"))
              .build();
        }
      }
    }
    return getDeleteCartItemByIdMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static CartServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<CartServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<CartServiceStub>() {
        @java.lang.Override
        public CartServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new CartServiceStub(channel, callOptions);
        }
      };
    return CartServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static CartServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<CartServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<CartServiceBlockingV2Stub>() {
        @java.lang.Override
        public CartServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new CartServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return CartServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static CartServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<CartServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<CartServiceBlockingStub>() {
        @java.lang.Override
        public CartServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new CartServiceBlockingStub(channel, callOptions);
        }
      };
    return CartServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static CartServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<CartServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<CartServiceFutureStub>() {
        @java.lang.Override
        public CartServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new CartServiceFutureStub(channel, callOptions);
        }
      };
    return CartServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void createCartItem(com.commerce.CreateCartItemRequest request,
        io.grpc.stub.StreamObserver<com.commerce.CreateCartItemResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateCartItemMethod(), responseObserver);
    }

    /**
     */
    default void findCartItemById(com.commerce.FindCartItemByIdRequest request,
        io.grpc.stub.StreamObserver<com.commerce.CartItem> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getFindCartItemByIdMethod(), responseObserver);
    }

    /**
     */
    default void findCartItemsWithPagination(com.commerce.FindCartItemsWithPaginationRequest request,
        io.grpc.stub.StreamObserver<com.commerce.FindCartItemsWithPaginationResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getFindCartItemsWithPaginationMethod(), responseObserver);
    }

    /**
     */
    default void updateCartItemById(com.commerce.UpdateCartItemByIdRequest request,
        io.grpc.stub.StreamObserver<com.commerce.UpdateCartItemByIdResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateCartItemByIdMethod(), responseObserver);
    }

    /**
     */
    default void deleteCartItemById(com.commerce.DeleteCartItemByIdRequest request,
        io.grpc.stub.StreamObserver<com.commerce.DeleteCartItemByIdResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeleteCartItemByIdMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service CartService.
   */
  public static abstract class CartServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return CartServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service CartService.
   */
  public static final class CartServiceStub
      extends io.grpc.stub.AbstractAsyncStub<CartServiceStub> {
    private CartServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CartServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new CartServiceStub(channel, callOptions);
    }

    /**
     */
    public void createCartItem(com.commerce.CreateCartItemRequest request,
        io.grpc.stub.StreamObserver<com.commerce.CreateCartItemResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateCartItemMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void findCartItemById(com.commerce.FindCartItemByIdRequest request,
        io.grpc.stub.StreamObserver<com.commerce.CartItem> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getFindCartItemByIdMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void findCartItemsWithPagination(com.commerce.FindCartItemsWithPaginationRequest request,
        io.grpc.stub.StreamObserver<com.commerce.FindCartItemsWithPaginationResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getFindCartItemsWithPaginationMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateCartItemById(com.commerce.UpdateCartItemByIdRequest request,
        io.grpc.stub.StreamObserver<com.commerce.UpdateCartItemByIdResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateCartItemByIdMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void deleteCartItemById(com.commerce.DeleteCartItemByIdRequest request,
        io.grpc.stub.StreamObserver<com.commerce.DeleteCartItemByIdResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeleteCartItemByIdMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service CartService.
   */
  public static final class CartServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<CartServiceBlockingV2Stub> {
    private CartServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CartServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new CartServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     */
    public com.commerce.CreateCartItemResponse createCartItem(com.commerce.CreateCartItemRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateCartItemMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.commerce.CartItem findCartItemById(com.commerce.FindCartItemByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindCartItemByIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.commerce.FindCartItemsWithPaginationResponse findCartItemsWithPagination(com.commerce.FindCartItemsWithPaginationRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindCartItemsWithPaginationMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.commerce.UpdateCartItemByIdResponse updateCartItemById(com.commerce.UpdateCartItemByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateCartItemByIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.commerce.DeleteCartItemByIdResponse deleteCartItemById(com.commerce.DeleteCartItemByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteCartItemByIdMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service CartService.
   */
  public static final class CartServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<CartServiceBlockingStub> {
    private CartServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CartServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new CartServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.commerce.CreateCartItemResponse createCartItem(com.commerce.CreateCartItemRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateCartItemMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.commerce.CartItem findCartItemById(com.commerce.FindCartItemByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindCartItemByIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.commerce.FindCartItemsWithPaginationResponse findCartItemsWithPagination(com.commerce.FindCartItemsWithPaginationRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindCartItemsWithPaginationMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.commerce.UpdateCartItemByIdResponse updateCartItemById(com.commerce.UpdateCartItemByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateCartItemByIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.commerce.DeleteCartItemByIdResponse deleteCartItemById(com.commerce.DeleteCartItemByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteCartItemByIdMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service CartService.
   */
  public static final class CartServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<CartServiceFutureStub> {
    private CartServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected CartServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new CartServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.commerce.CreateCartItemResponse> createCartItem(
        com.commerce.CreateCartItemRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateCartItemMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.commerce.CartItem> findCartItemById(
        com.commerce.FindCartItemByIdRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getFindCartItemByIdMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.commerce.FindCartItemsWithPaginationResponse> findCartItemsWithPagination(
        com.commerce.FindCartItemsWithPaginationRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getFindCartItemsWithPaginationMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.commerce.UpdateCartItemByIdResponse> updateCartItemById(
        com.commerce.UpdateCartItemByIdRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateCartItemByIdMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.commerce.DeleteCartItemByIdResponse> deleteCartItemById(
        com.commerce.DeleteCartItemByIdRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeleteCartItemByIdMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_CART_ITEM = 0;
  private static final int METHODID_FIND_CART_ITEM_BY_ID = 1;
  private static final int METHODID_FIND_CART_ITEMS_WITH_PAGINATION = 2;
  private static final int METHODID_UPDATE_CART_ITEM_BY_ID = 3;
  private static final int METHODID_DELETE_CART_ITEM_BY_ID = 4;

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
        case METHODID_CREATE_CART_ITEM:
          serviceImpl.createCartItem((com.commerce.CreateCartItemRequest) request,
              (io.grpc.stub.StreamObserver<com.commerce.CreateCartItemResponse>) responseObserver);
          break;
        case METHODID_FIND_CART_ITEM_BY_ID:
          serviceImpl.findCartItemById((com.commerce.FindCartItemByIdRequest) request,
              (io.grpc.stub.StreamObserver<com.commerce.CartItem>) responseObserver);
          break;
        case METHODID_FIND_CART_ITEMS_WITH_PAGINATION:
          serviceImpl.findCartItemsWithPagination((com.commerce.FindCartItemsWithPaginationRequest) request,
              (io.grpc.stub.StreamObserver<com.commerce.FindCartItemsWithPaginationResponse>) responseObserver);
          break;
        case METHODID_UPDATE_CART_ITEM_BY_ID:
          serviceImpl.updateCartItemById((com.commerce.UpdateCartItemByIdRequest) request,
              (io.grpc.stub.StreamObserver<com.commerce.UpdateCartItemByIdResponse>) responseObserver);
          break;
        case METHODID_DELETE_CART_ITEM_BY_ID:
          serviceImpl.deleteCartItemById((com.commerce.DeleteCartItemByIdRequest) request,
              (io.grpc.stub.StreamObserver<com.commerce.DeleteCartItemByIdResponse>) responseObserver);
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
          getCreateCartItemMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.commerce.CreateCartItemRequest,
              com.commerce.CreateCartItemResponse>(
                service, METHODID_CREATE_CART_ITEM)))
        .addMethod(
          getFindCartItemByIdMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.commerce.FindCartItemByIdRequest,
              com.commerce.CartItem>(
                service, METHODID_FIND_CART_ITEM_BY_ID)))
        .addMethod(
          getFindCartItemsWithPaginationMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.commerce.FindCartItemsWithPaginationRequest,
              com.commerce.FindCartItemsWithPaginationResponse>(
                service, METHODID_FIND_CART_ITEMS_WITH_PAGINATION)))
        .addMethod(
          getUpdateCartItemByIdMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.commerce.UpdateCartItemByIdRequest,
              com.commerce.UpdateCartItemByIdResponse>(
                service, METHODID_UPDATE_CART_ITEM_BY_ID)))
        .addMethod(
          getDeleteCartItemByIdMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.commerce.DeleteCartItemByIdRequest,
              com.commerce.DeleteCartItemByIdResponse>(
                service, METHODID_DELETE_CART_ITEM_BY_ID)))
        .build();
  }

  private static abstract class CartServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    CartServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.commerce.ServiceProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("CartService");
    }
  }

  private static final class CartServiceFileDescriptorSupplier
      extends CartServiceBaseDescriptorSupplier {
    CartServiceFileDescriptorSupplier() {}
  }

  private static final class CartServiceMethodDescriptorSupplier
      extends CartServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    CartServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (CartServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new CartServiceFileDescriptorSupplier())
              .addMethod(getCreateCartItemMethod())
              .addMethod(getFindCartItemByIdMethod())
              .addMethod(getFindCartItemsWithPaginationMethod())
              .addMethod(getUpdateCartItemByIdMethod())
              .addMethod(getDeleteCartItemByIdMethod())
              .build();
        }
      }
    }
    return result;
  }
}
