package com.product;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.73.0)",
    comments = "Source: v1/product/service.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class ProductServiceGrpc {

  private ProductServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "product.ProductService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.product.CreateProductRequest,
      com.google.protobuf.Empty> getCreateProductMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateProduct",
      requestType = com.product.CreateProductRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.product.CreateProductRequest,
      com.google.protobuf.Empty> getCreateProductMethod() {
    io.grpc.MethodDescriptor<com.product.CreateProductRequest, com.google.protobuf.Empty> getCreateProductMethod;
    if ((getCreateProductMethod = ProductServiceGrpc.getCreateProductMethod) == null) {
      synchronized (ProductServiceGrpc.class) {
        if ((getCreateProductMethod = ProductServiceGrpc.getCreateProductMethod) == null) {
          ProductServiceGrpc.getCreateProductMethod = getCreateProductMethod =
              io.grpc.MethodDescriptor.<com.product.CreateProductRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateProduct"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.product.CreateProductRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new ProductServiceMethodDescriptorSupplier("CreateProduct"))
              .build();
        }
      }
    }
    return getCreateProductMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.product.UpdateProductByIdRequest,
      com.google.protobuf.Empty> getUpdateProductByIdMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateProductById",
      requestType = com.product.UpdateProductByIdRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.product.UpdateProductByIdRequest,
      com.google.protobuf.Empty> getUpdateProductByIdMethod() {
    io.grpc.MethodDescriptor<com.product.UpdateProductByIdRequest, com.google.protobuf.Empty> getUpdateProductByIdMethod;
    if ((getUpdateProductByIdMethod = ProductServiceGrpc.getUpdateProductByIdMethod) == null) {
      synchronized (ProductServiceGrpc.class) {
        if ((getUpdateProductByIdMethod = ProductServiceGrpc.getUpdateProductByIdMethod) == null) {
          ProductServiceGrpc.getUpdateProductByIdMethod = getUpdateProductByIdMethod =
              io.grpc.MethodDescriptor.<com.product.UpdateProductByIdRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateProductById"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.product.UpdateProductByIdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new ProductServiceMethodDescriptorSupplier("UpdateProductById"))
              .build();
        }
      }
    }
    return getUpdateProductByIdMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.product.DeleteProductByIdRequest,
      com.google.protobuf.Empty> getDeleteProductByIdMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DeleteProductById",
      requestType = com.product.DeleteProductByIdRequest.class,
      responseType = com.google.protobuf.Empty.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.product.DeleteProductByIdRequest,
      com.google.protobuf.Empty> getDeleteProductByIdMethod() {
    io.grpc.MethodDescriptor<com.product.DeleteProductByIdRequest, com.google.protobuf.Empty> getDeleteProductByIdMethod;
    if ((getDeleteProductByIdMethod = ProductServiceGrpc.getDeleteProductByIdMethod) == null) {
      synchronized (ProductServiceGrpc.class) {
        if ((getDeleteProductByIdMethod = ProductServiceGrpc.getDeleteProductByIdMethod) == null) {
          ProductServiceGrpc.getDeleteProductByIdMethod = getDeleteProductByIdMethod =
              io.grpc.MethodDescriptor.<com.product.DeleteProductByIdRequest, com.google.protobuf.Empty>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DeleteProductById"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.product.DeleteProductByIdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.google.protobuf.Empty.getDefaultInstance()))
              .setSchemaDescriptor(new ProductServiceMethodDescriptorSupplier("DeleteProductById"))
              .build();
        }
      }
    }
    return getDeleteProductByIdMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.product.FindProductsWithPaginationRequest,
      com.product.FindProductsWithPaginationResponse> getFindProductsWithPaginationMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "FindProductsWithPagination",
      requestType = com.product.FindProductsWithPaginationRequest.class,
      responseType = com.product.FindProductsWithPaginationResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.product.FindProductsWithPaginationRequest,
      com.product.FindProductsWithPaginationResponse> getFindProductsWithPaginationMethod() {
    io.grpc.MethodDescriptor<com.product.FindProductsWithPaginationRequest, com.product.FindProductsWithPaginationResponse> getFindProductsWithPaginationMethod;
    if ((getFindProductsWithPaginationMethod = ProductServiceGrpc.getFindProductsWithPaginationMethod) == null) {
      synchronized (ProductServiceGrpc.class) {
        if ((getFindProductsWithPaginationMethod = ProductServiceGrpc.getFindProductsWithPaginationMethod) == null) {
          ProductServiceGrpc.getFindProductsWithPaginationMethod = getFindProductsWithPaginationMethod =
              io.grpc.MethodDescriptor.<com.product.FindProductsWithPaginationRequest, com.product.FindProductsWithPaginationResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "FindProductsWithPagination"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.product.FindProductsWithPaginationRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.product.FindProductsWithPaginationResponse.getDefaultInstance()))
              .setSchemaDescriptor(new ProductServiceMethodDescriptorSupplier("FindProductsWithPagination"))
              .build();
        }
      }
    }
    return getFindProductsWithPaginationMethod;
  }

  private static volatile io.grpc.MethodDescriptor<com.product.FindProductByIdRequest,
      com.product.Product> getFindProductByIdMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "FindProductById",
      requestType = com.product.FindProductByIdRequest.class,
      responseType = com.product.Product.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.product.FindProductByIdRequest,
      com.product.Product> getFindProductByIdMethod() {
    io.grpc.MethodDescriptor<com.product.FindProductByIdRequest, com.product.Product> getFindProductByIdMethod;
    if ((getFindProductByIdMethod = ProductServiceGrpc.getFindProductByIdMethod) == null) {
      synchronized (ProductServiceGrpc.class) {
        if ((getFindProductByIdMethod = ProductServiceGrpc.getFindProductByIdMethod) == null) {
          ProductServiceGrpc.getFindProductByIdMethod = getFindProductByIdMethod =
              io.grpc.MethodDescriptor.<com.product.FindProductByIdRequest, com.product.Product>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "FindProductById"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.product.FindProductByIdRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.product.Product.getDefaultInstance()))
              .setSchemaDescriptor(new ProductServiceMethodDescriptorSupplier("FindProductById"))
              .build();
        }
      }
    }
    return getFindProductByIdMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static ProductServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ProductServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ProductServiceStub>() {
        @java.lang.Override
        public ProductServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ProductServiceStub(channel, callOptions);
        }
      };
    return ProductServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports all types of calls on the service
   */
  public static ProductServiceBlockingV2Stub newBlockingV2Stub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ProductServiceBlockingV2Stub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ProductServiceBlockingV2Stub>() {
        @java.lang.Override
        public ProductServiceBlockingV2Stub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ProductServiceBlockingV2Stub(channel, callOptions);
        }
      };
    return ProductServiceBlockingV2Stub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static ProductServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ProductServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ProductServiceBlockingStub>() {
        @java.lang.Override
        public ProductServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ProductServiceBlockingStub(channel, callOptions);
        }
      };
    return ProductServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static ProductServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<ProductServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<ProductServiceFutureStub>() {
        @java.lang.Override
        public ProductServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new ProductServiceFutureStub(channel, callOptions);
        }
      };
    return ProductServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    default void createProduct(com.product.CreateProductRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateProductMethod(), responseObserver);
    }

    /**
     */
    default void updateProductById(com.product.UpdateProductByIdRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateProductByIdMethod(), responseObserver);
    }

    /**
     */
    default void deleteProductById(com.product.DeleteProductByIdRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeleteProductByIdMethod(), responseObserver);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    default void findProductsWithPagination(com.product.FindProductsWithPaginationRequest request,
        io.grpc.stub.StreamObserver<com.product.FindProductsWithPaginationResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getFindProductsWithPaginationMethod(), responseObserver);
    }

    /**
     */
    default void findProductById(com.product.FindProductByIdRequest request,
        io.grpc.stub.StreamObserver<com.product.Product> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getFindProductByIdMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service ProductService.
   */
  public static abstract class ProductServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return ProductServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service ProductService.
   */
  public static final class ProductServiceStub
      extends io.grpc.stub.AbstractAsyncStub<ProductServiceStub> {
    private ProductServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ProductServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ProductServiceStub(channel, callOptions);
    }

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    public void createProduct(com.product.CreateProductRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateProductMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateProductById(com.product.UpdateProductByIdRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateProductByIdMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void deleteProductById(com.product.DeleteProductByIdRequest request,
        io.grpc.stub.StreamObserver<com.google.protobuf.Empty> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeleteProductByIdMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public void findProductsWithPagination(com.product.FindProductsWithPaginationRequest request,
        io.grpc.stub.StreamObserver<com.product.FindProductsWithPaginationResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getFindProductsWithPaginationMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void findProductById(com.product.FindProductByIdRequest request,
        io.grpc.stub.StreamObserver<com.product.Product> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getFindProductByIdMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service ProductService.
   */
  public static final class ProductServiceBlockingV2Stub
      extends io.grpc.stub.AbstractBlockingStub<ProductServiceBlockingV2Stub> {
    private ProductServiceBlockingV2Stub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ProductServiceBlockingV2Stub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ProductServiceBlockingV2Stub(channel, callOptions);
    }

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    public com.google.protobuf.Empty createProduct(com.product.CreateProductRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateProductMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty updateProductById(com.product.UpdateProductByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateProductByIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty deleteProductById(com.product.DeleteProductByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteProductByIdMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public com.product.FindProductsWithPaginationResponse findProductsWithPagination(com.product.FindProductsWithPaginationRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindProductsWithPaginationMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.product.Product findProductById(com.product.FindProductByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindProductByIdMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do limited synchronous rpc calls to service ProductService.
   */
  public static final class ProductServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<ProductServiceBlockingStub> {
    private ProductServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ProductServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ProductServiceBlockingStub(channel, callOptions);
    }

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    public com.google.protobuf.Empty createProduct(com.product.CreateProductRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateProductMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty updateProductById(com.product.UpdateProductByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateProductByIdMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.google.protobuf.Empty deleteProductById(com.product.DeleteProductByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteProductByIdMethod(), getCallOptions(), request);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public com.product.FindProductsWithPaginationResponse findProductsWithPagination(com.product.FindProductsWithPaginationRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindProductsWithPaginationMethod(), getCallOptions(), request);
    }

    /**
     */
    public com.product.Product findProductById(com.product.FindProductByIdRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getFindProductByIdMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service ProductService.
   */
  public static final class ProductServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<ProductServiceFutureStub> {
    private ProductServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected ProductServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new ProductServiceFutureStub(channel, callOptions);
    }

    /**
     * <pre>
     * COMMAND
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> createProduct(
        com.product.CreateProductRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateProductMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> updateProductById(
        com.product.UpdateProductByIdRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateProductByIdMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.google.protobuf.Empty> deleteProductById(
        com.product.DeleteProductByIdRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeleteProductByIdMethod(), getCallOptions()), request);
    }

    /**
     * <pre>
     * QUERY
     * </pre>
     */
    public com.google.common.util.concurrent.ListenableFuture<com.product.FindProductsWithPaginationResponse> findProductsWithPagination(
        com.product.FindProductsWithPaginationRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getFindProductsWithPaginationMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.product.Product> findProductById(
        com.product.FindProductByIdRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getFindProductByIdMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_PRODUCT = 0;
  private static final int METHODID_UPDATE_PRODUCT_BY_ID = 1;
  private static final int METHODID_DELETE_PRODUCT_BY_ID = 2;
  private static final int METHODID_FIND_PRODUCTS_WITH_PAGINATION = 3;
  private static final int METHODID_FIND_PRODUCT_BY_ID = 4;

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
        case METHODID_CREATE_PRODUCT:
          serviceImpl.createProduct((com.product.CreateProductRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_UPDATE_PRODUCT_BY_ID:
          serviceImpl.updateProductById((com.product.UpdateProductByIdRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_DELETE_PRODUCT_BY_ID:
          serviceImpl.deleteProductById((com.product.DeleteProductByIdRequest) request,
              (io.grpc.stub.StreamObserver<com.google.protobuf.Empty>) responseObserver);
          break;
        case METHODID_FIND_PRODUCTS_WITH_PAGINATION:
          serviceImpl.findProductsWithPagination((com.product.FindProductsWithPaginationRequest) request,
              (io.grpc.stub.StreamObserver<com.product.FindProductsWithPaginationResponse>) responseObserver);
          break;
        case METHODID_FIND_PRODUCT_BY_ID:
          serviceImpl.findProductById((com.product.FindProductByIdRequest) request,
              (io.grpc.stub.StreamObserver<com.product.Product>) responseObserver);
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
          getCreateProductMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.product.CreateProductRequest,
              com.google.protobuf.Empty>(
                service, METHODID_CREATE_PRODUCT)))
        .addMethod(
          getUpdateProductByIdMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.product.UpdateProductByIdRequest,
              com.google.protobuf.Empty>(
                service, METHODID_UPDATE_PRODUCT_BY_ID)))
        .addMethod(
          getDeleteProductByIdMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.product.DeleteProductByIdRequest,
              com.google.protobuf.Empty>(
                service, METHODID_DELETE_PRODUCT_BY_ID)))
        .addMethod(
          getFindProductsWithPaginationMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.product.FindProductsWithPaginationRequest,
              com.product.FindProductsWithPaginationResponse>(
                service, METHODID_FIND_PRODUCTS_WITH_PAGINATION)))
        .addMethod(
          getFindProductByIdMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              com.product.FindProductByIdRequest,
              com.product.Product>(
                service, METHODID_FIND_PRODUCT_BY_ID)))
        .build();
  }

  private static abstract class ProductServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    ProductServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.product.ServiceProto.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("ProductService");
    }
  }

  private static final class ProductServiceFileDescriptorSupplier
      extends ProductServiceBaseDescriptorSupplier {
    ProductServiceFileDescriptorSupplier() {}
  }

  private static final class ProductServiceMethodDescriptorSupplier
      extends ProductServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    ProductServiceMethodDescriptorSupplier(java.lang.String methodName) {
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
      synchronized (ProductServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new ProductServiceFileDescriptorSupplier())
              .addMethod(getCreateProductMethod())
              .addMethod(getUpdateProductByIdMethod())
              .addMethod(getDeleteProductByIdMethod())
              .addMethod(getFindProductsWithPaginationMethod())
              .addMethod(getFindProductByIdMethod())
              .build();
        }
      }
    }
    return result;
  }
}
