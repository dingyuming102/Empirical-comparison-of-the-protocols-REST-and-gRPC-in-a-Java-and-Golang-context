package com.example.proto;

import static io.grpc.MethodDescriptor.generateFullMethodName;
import static io.grpc.stub.ClientCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ClientCalls.asyncClientStreamingCall;
import static io.grpc.stub.ClientCalls.asyncServerStreamingCall;
import static io.grpc.stub.ClientCalls.asyncUnaryCall;
import static io.grpc.stub.ClientCalls.blockingServerStreamingCall;
import static io.grpc.stub.ClientCalls.blockingUnaryCall;
import static io.grpc.stub.ClientCalls.futureUnaryCall;
import static io.grpc.stub.ServerCalls.asyncBidiStreamingCall;
import static io.grpc.stub.ServerCalls.asyncClientStreamingCall;
import static io.grpc.stub.ServerCalls.asyncServerStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnaryCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedStreamingCall;
import static io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.26.0)",
    comments = "Source: transaction.proto")
public final class MatrixOpGRPCGrpc {

  private MatrixOpGRPCGrpc() {}

  public static final String SERVICE_NAME = "MatrixOpGRPC";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.example.proto.Transaction.LargeTransaction,
      com.example.proto.Transaction.LargeTransaction> getMultiplierMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Multiplier",
      requestType = com.example.proto.Transaction.LargeTransaction.class,
      responseType = com.example.proto.Transaction.LargeTransaction.class,
      methodType = io.grpc.MethodDescriptor.MethodType.BIDI_STREAMING)
  public static io.grpc.MethodDescriptor<com.example.proto.Transaction.LargeTransaction,
      com.example.proto.Transaction.LargeTransaction> getMultiplierMethod() {
    io.grpc.MethodDescriptor<com.example.proto.Transaction.LargeTransaction, com.example.proto.Transaction.LargeTransaction> getMultiplierMethod;
    if ((getMultiplierMethod = MatrixOpGRPCGrpc.getMultiplierMethod) == null) {
      synchronized (MatrixOpGRPCGrpc.class) {
        if ((getMultiplierMethod = MatrixOpGRPCGrpc.getMultiplierMethod) == null) {
          MatrixOpGRPCGrpc.getMultiplierMethod = getMultiplierMethod =
              io.grpc.MethodDescriptor.<com.example.proto.Transaction.LargeTransaction, com.example.proto.Transaction.LargeTransaction>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.BIDI_STREAMING)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Multiplier"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.example.proto.Transaction.LargeTransaction.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.example.proto.Transaction.LargeTransaction.getDefaultInstance()))
              .setSchemaDescriptor(new MatrixOpGRPCMethodDescriptorSupplier("Multiplier"))
              .build();
        }
      }
    }
    return getMultiplierMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static MatrixOpGRPCStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MatrixOpGRPCStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MatrixOpGRPCStub>() {
        @java.lang.Override
        public MatrixOpGRPCStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MatrixOpGRPCStub(channel, callOptions);
        }
      };
    return MatrixOpGRPCStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static MatrixOpGRPCBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MatrixOpGRPCBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MatrixOpGRPCBlockingStub>() {
        @java.lang.Override
        public MatrixOpGRPCBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MatrixOpGRPCBlockingStub(channel, callOptions);
        }
      };
    return MatrixOpGRPCBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static MatrixOpGRPCFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MatrixOpGRPCFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MatrixOpGRPCFutureStub>() {
        @java.lang.Override
        public MatrixOpGRPCFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MatrixOpGRPCFutureStub(channel, callOptions);
        }
      };
    return MatrixOpGRPCFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class MatrixOpGRPCImplBase implements io.grpc.BindableService {

    /**
     * <pre>
     *双向流模式
     * </pre>
     */
    public io.grpc.stub.StreamObserver<com.example.proto.Transaction.LargeTransaction> multiplier(
        io.grpc.stub.StreamObserver<com.example.proto.Transaction.LargeTransaction> responseObserver) {
      return asyncUnimplementedStreamingCall(getMultiplierMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getMultiplierMethod(),
            asyncBidiStreamingCall(
              new MethodHandlers<
                com.example.proto.Transaction.LargeTransaction,
                com.example.proto.Transaction.LargeTransaction>(
                  this, METHODID_MULTIPLIER)))
          .build();
    }
  }

  /**
   */
  public static final class MatrixOpGRPCStub extends io.grpc.stub.AbstractAsyncStub<MatrixOpGRPCStub> {
    private MatrixOpGRPCStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MatrixOpGRPCStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MatrixOpGRPCStub(channel, callOptions);
    }

    /**
     * <pre>
     *双向流模式
     * </pre>
     */
    public io.grpc.stub.StreamObserver<com.example.proto.Transaction.LargeTransaction> multiplier(
        io.grpc.stub.StreamObserver<com.example.proto.Transaction.LargeTransaction> responseObserver) {
      return asyncBidiStreamingCall(
          getChannel().newCall(getMultiplierMethod(), getCallOptions()), responseObserver);
    }
  }

  /**
   */
  public static final class MatrixOpGRPCBlockingStub extends io.grpc.stub.AbstractBlockingStub<MatrixOpGRPCBlockingStub> {
    private MatrixOpGRPCBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MatrixOpGRPCBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MatrixOpGRPCBlockingStub(channel, callOptions);
    }
  }

  /**
   */
  public static final class MatrixOpGRPCFutureStub extends io.grpc.stub.AbstractFutureStub<MatrixOpGRPCFutureStub> {
    private MatrixOpGRPCFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MatrixOpGRPCFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MatrixOpGRPCFutureStub(channel, callOptions);
    }
  }

  private static final int METHODID_MULTIPLIER = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final MatrixOpGRPCImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(MatrixOpGRPCImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_MULTIPLIER:
          return (io.grpc.stub.StreamObserver<Req>) serviceImpl.multiplier(
              (io.grpc.stub.StreamObserver<com.example.proto.Transaction.LargeTransaction>) responseObserver);
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class MatrixOpGRPCBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    MatrixOpGRPCBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.example.proto.Transaction.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("MatrixOpGRPC");
    }
  }

  private static final class MatrixOpGRPCFileDescriptorSupplier
      extends MatrixOpGRPCBaseDescriptorSupplier {
    MatrixOpGRPCFileDescriptorSupplier() {}
  }

  private static final class MatrixOpGRPCMethodDescriptorSupplier
      extends MatrixOpGRPCBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    MatrixOpGRPCMethodDescriptorSupplier(String methodName) {
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
      synchronized (MatrixOpGRPCGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new MatrixOpGRPCFileDescriptorSupplier())
              .addMethod(getMultiplierMethod())
              .build();
        }
      }
    }
    return result;
  }
}
