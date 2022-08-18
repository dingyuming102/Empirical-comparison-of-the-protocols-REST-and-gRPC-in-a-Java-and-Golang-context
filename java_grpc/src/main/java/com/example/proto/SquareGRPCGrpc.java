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
public final class SquareGRPCGrpc {

  private SquareGRPCGrpc() {}

  public static final String SERVICE_NAME = "SquareGRPC";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.example.proto.Transaction.SmallTransaction,
      com.example.proto.Transaction.SmallTransaction> getSquareMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "Square",
      requestType = com.example.proto.Transaction.SmallTransaction.class,
      responseType = com.example.proto.Transaction.SmallTransaction.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.example.proto.Transaction.SmallTransaction,
      com.example.proto.Transaction.SmallTransaction> getSquareMethod() {
    io.grpc.MethodDescriptor<com.example.proto.Transaction.SmallTransaction, com.example.proto.Transaction.SmallTransaction> getSquareMethod;
    if ((getSquareMethod = SquareGRPCGrpc.getSquareMethod) == null) {
      synchronized (SquareGRPCGrpc.class) {
        if ((getSquareMethod = SquareGRPCGrpc.getSquareMethod) == null) {
          SquareGRPCGrpc.getSquareMethod = getSquareMethod =
              io.grpc.MethodDescriptor.<com.example.proto.Transaction.SmallTransaction, com.example.proto.Transaction.SmallTransaction>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "Square"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.example.proto.Transaction.SmallTransaction.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.example.proto.Transaction.SmallTransaction.getDefaultInstance()))
              .setSchemaDescriptor(new SquareGRPCMethodDescriptorSupplier("Square"))
              .build();
        }
      }
    }
    return getSquareMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static SquareGRPCStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SquareGRPCStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SquareGRPCStub>() {
        @java.lang.Override
        public SquareGRPCStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SquareGRPCStub(channel, callOptions);
        }
      };
    return SquareGRPCStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static SquareGRPCBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SquareGRPCBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SquareGRPCBlockingStub>() {
        @java.lang.Override
        public SquareGRPCBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SquareGRPCBlockingStub(channel, callOptions);
        }
      };
    return SquareGRPCBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static SquareGRPCFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<SquareGRPCFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<SquareGRPCFutureStub>() {
        @java.lang.Override
        public SquareGRPCFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new SquareGRPCFutureStub(channel, callOptions);
        }
      };
    return SquareGRPCFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class SquareGRPCImplBase implements io.grpc.BindableService {

    /**
     */
    public void square(com.example.proto.Transaction.SmallTransaction request,
        io.grpc.stub.StreamObserver<com.example.proto.Transaction.SmallTransaction> responseObserver) {
      asyncUnimplementedUnaryCall(getSquareMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getSquareMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.example.proto.Transaction.SmallTransaction,
                com.example.proto.Transaction.SmallTransaction>(
                  this, METHODID_SQUARE)))
          .build();
    }
  }

  /**
   */
  public static final class SquareGRPCStub extends io.grpc.stub.AbstractAsyncStub<SquareGRPCStub> {
    private SquareGRPCStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SquareGRPCStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SquareGRPCStub(channel, callOptions);
    }

    /**
     */
    public void square(com.example.proto.Transaction.SmallTransaction request,
        io.grpc.stub.StreamObserver<com.example.proto.Transaction.SmallTransaction> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getSquareMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class SquareGRPCBlockingStub extends io.grpc.stub.AbstractBlockingStub<SquareGRPCBlockingStub> {
    private SquareGRPCBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SquareGRPCBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SquareGRPCBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.example.proto.Transaction.SmallTransaction square(com.example.proto.Transaction.SmallTransaction request) {
      return blockingUnaryCall(
          getChannel(), getSquareMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class SquareGRPCFutureStub extends io.grpc.stub.AbstractFutureStub<SquareGRPCFutureStub> {
    private SquareGRPCFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected SquareGRPCFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new SquareGRPCFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.example.proto.Transaction.SmallTransaction> square(
        com.example.proto.Transaction.SmallTransaction request) {
      return futureUnaryCall(
          getChannel().newCall(getSquareMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_SQUARE = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final SquareGRPCImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(SquareGRPCImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_SQUARE:
          serviceImpl.square((com.example.proto.Transaction.SmallTransaction) request,
              (io.grpc.stub.StreamObserver<com.example.proto.Transaction.SmallTransaction>) responseObserver);
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

  private static abstract class SquareGRPCBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    SquareGRPCBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.example.proto.Transaction.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("SquareGRPC");
    }
  }

  private static final class SquareGRPCFileDescriptorSupplier
      extends SquareGRPCBaseDescriptorSupplier {
    SquareGRPCFileDescriptorSupplier() {}
  }

  private static final class SquareGRPCMethodDescriptorSupplier
      extends SquareGRPCBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    SquareGRPCMethodDescriptorSupplier(String methodName) {
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
      synchronized (SquareGRPCGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new SquareGRPCFileDescriptorSupplier())
              .addMethod(getSquareMethod())
              .build();
        }
      }
    }
    return result;
  }
}
