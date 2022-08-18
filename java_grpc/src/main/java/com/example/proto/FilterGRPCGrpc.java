package com.example.proto;

import io.grpc.MethodDescriptor;
import io.grpc.stub.annotations.RpcMethod;

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
public final class FilterGRPCGrpc {

  private FilterGRPCGrpc() {}

  public static final String SERVICE_NAME = "FilterGRPC";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.example.proto.Transaction.MediumTransaction,
      com.example.proto.Transaction.MediumTransaction> getGrayscaleFilterMethod;

  @RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GrayscaleFilter",
      requestType = Transaction.MediumTransaction.class,
      responseType = Transaction.MediumTransaction.class,
      methodType = MethodDescriptor.MethodType.UNARY)
  public static MethodDescriptor<Transaction.MediumTransaction,
      Transaction.MediumTransaction> getGrayscaleFilterMethod() {
    MethodDescriptor<Transaction.MediumTransaction, Transaction.MediumTransaction> getGrayscaleFilterMethod;
    if ((getGrayscaleFilterMethod = FilterGRPCGrpc.getGrayscaleFilterMethod) == null) {
      synchronized (FilterGRPCGrpc.class) {
        if ((getGrayscaleFilterMethod = FilterGRPCGrpc.getGrayscaleFilterMethod) == null) {
          FilterGRPCGrpc.getGrayscaleFilterMethod = getGrayscaleFilterMethod =
              MethodDescriptor.<Transaction.MediumTransaction, Transaction.MediumTransaction>newBuilder()
              .setType(MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GrayscaleFilter"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  Transaction.MediumTransaction.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  Transaction.MediumTransaction.getDefaultInstance()))
              .setSchemaDescriptor(new FilterGRPCMethodDescriptorSupplier("GrayscaleFilter"))
              .build();
        }
      }
    }
    return getGrayscaleFilterMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static FilterGRPCStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<FilterGRPCStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<FilterGRPCStub>() {
        @java.lang.Override
        public FilterGRPCStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new FilterGRPCStub(channel, callOptions);
        }
      };
    return FilterGRPCStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static FilterGRPCBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<FilterGRPCBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<FilterGRPCBlockingStub>() {
        @java.lang.Override
        public FilterGRPCBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new FilterGRPCBlockingStub(channel, callOptions);
        }
      };
    return FilterGRPCBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static FilterGRPCFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<FilterGRPCFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<FilterGRPCFutureStub>() {
        @java.lang.Override
        public FilterGRPCFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new FilterGRPCFutureStub(channel, callOptions);
        }
      };
    return FilterGRPCFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class FilterGRPCImplBase implements io.grpc.BindableService {

    /**
     */
    public void grayscaleFilter(com.example.proto.Transaction.MediumTransaction request,
        io.grpc.stub.StreamObserver<com.example.proto.Transaction.MediumTransaction> responseObserver) {
      asyncUnimplementedUnaryCall(getGrayscaleFilterMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGrayscaleFilterMethod(),
            asyncUnaryCall(
              new MethodHandlers<
                com.example.proto.Transaction.MediumTransaction,
                com.example.proto.Transaction.MediumTransaction>(
                  this, METHODID_GRAYSCALE_FILTER)))
          .build();
    }
  }

  /**
   */
  public static final class FilterGRPCStub extends io.grpc.stub.AbstractAsyncStub<FilterGRPCStub> {
    private FilterGRPCStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FilterGRPCStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new FilterGRPCStub(channel, callOptions);
    }

    /**
     */
    public void grayscaleFilter(com.example.proto.Transaction.MediumTransaction request,
        io.grpc.stub.StreamObserver<com.example.proto.Transaction.MediumTransaction> responseObserver) {
      asyncUnaryCall(
          getChannel().newCall(getGrayscaleFilterMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class FilterGRPCBlockingStub extends io.grpc.stub.AbstractBlockingStub<FilterGRPCBlockingStub> {
    private FilterGRPCBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FilterGRPCBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new FilterGRPCBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.example.proto.Transaction.MediumTransaction grayscaleFilter(com.example.proto.Transaction.MediumTransaction request) {
      return blockingUnaryCall(
          getChannel(), getGrayscaleFilterMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class FilterGRPCFutureStub extends io.grpc.stub.AbstractFutureStub<FilterGRPCFutureStub> {
    private FilterGRPCFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected FilterGRPCFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new FilterGRPCFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.example.proto.Transaction.MediumTransaction> grayscaleFilter(
        com.example.proto.Transaction.MediumTransaction request) {
      return futureUnaryCall(
          getChannel().newCall(getGrayscaleFilterMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GRAYSCALE_FILTER = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final FilterGRPCImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(FilterGRPCImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GRAYSCALE_FILTER:
          serviceImpl.grayscaleFilter((com.example.proto.Transaction.MediumTransaction) request,
              (io.grpc.stub.StreamObserver<com.example.proto.Transaction.MediumTransaction>) responseObserver);
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

  private static abstract class FilterGRPCBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    FilterGRPCBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.example.proto.Transaction.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("FilterGRPC");
    }
  }

  private static final class FilterGRPCFileDescriptorSupplier
      extends FilterGRPCBaseDescriptorSupplier {
    FilterGRPCFileDescriptorSupplier() {}
  }

  private static final class FilterGRPCMethodDescriptorSupplier
      extends FilterGRPCBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    FilterGRPCMethodDescriptorSupplier(String methodName) {
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
      synchronized (FilterGRPCGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new FilterGRPCFileDescriptorSupplier())
              .addMethod(getGrayscaleFilterMethod())
              .build();
        }
      }
    }
    return result;
  }
}
