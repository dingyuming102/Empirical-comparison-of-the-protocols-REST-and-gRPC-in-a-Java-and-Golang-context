package com.example.java_grpc;

import com.example.proto.FilterGRPCGrpc;
import com.example.proto.MatrixOpGRPCGrpc;
import com.example.proto.SquareGRPCGrpc;
import com.example.proto.Transaction;
import com.google.protobuf.ByteString;

import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.stub.StreamObserver;
import org.opencv.core.Mat;
import org.opencv.core.MatOfByte;
import org.opencv.imgcodecs.Imgcodecs;
import org.opencv.imgproc.Imgproc;

import java.io.IOException;
import java.util.Arrays;
import java.util.List;

public class GRPCServer {
    protected int port = 8080;
    protected Server server;
    public long[] procTime = new long[32];

    public GRPCServer() throws IOException, InterruptedException {
        // build gRPC server
        server = ServerBuilder.forPort(port)
//                .maxInboundMessageSize(50 * 1024 * 1024)
                .addService(new SquareServiceImpl())
                .addService(new FilterServiceImpl())
                .addService(new MatrixOpServiceImpl())
                .build();
    }

    public void start() throws IOException, InterruptedException {
        // start
        server.start();

        System.out.println("Server started, listening on " + port);
        System.out.println("Language, method, transaction, id, time");

        // shutdown hook
        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            System.out.println("gRPC server is shutting down!");
            server.shutdown();
            System.out.println(mean(procTime) / 1000000);
        }));

        server.awaitTermination();
    }

    static class SquareServiceImpl extends SquareGRPCGrpc.SquareGRPCImplBase {
        @Override
        public void square(Transaction.SmallTransaction request, StreamObserver<Transaction.SmallTransaction> responseObserver) {

//            long init_proc = System.nanoTime();

            int data = request.getData();
            int ret = data * data;
            Transaction.SmallTransaction reply = Transaction.SmallTransaction.newBuilder().setData(ret).build();

//            long end_proc = System.nanoTime();

            responseObserver.onNext(reply);

            responseObserver.onCompleted();

//            System.out.printf("Java, gRPC, Small, %d, %d\n", request.getId(), (end_proc-init_proc));

        }
    }

    class FilterServiceImpl extends FilterGRPCGrpc.FilterGRPCImplBase {

        @Override
        public void grayscaleFilter(Transaction.MediumTransaction request, StreamObserver<Transaction.MediumTransaction> responseObserver) {

//            long init_proc = System.nanoTime();

            //nu.pattern.OpenCV.loadShared();
            nu.pattern.OpenCV.loadLocally();

            /* Receive and decode input */
            byte[] img = request.getData(0).toByteArray();
            Mat src = Imgcodecs.imdecode(new MatOfByte(img), Imgcodecs.IMREAD_UNCHANGED), dest = new Mat();

            /* Apply filter */
            //long init_op = System.nanoTime();
            Imgproc.cvtColor(src, dest, Imgproc.COLOR_RGB2GRAY, src.channels());
            //long end_op = System.nanoTime();

            /* Decode and create output */
            MatOfByte dest_img = new MatOfByte();
            Imgcodecs.imencode(".jpg", dest, dest_img);

            Transaction.MediumTransaction reply = Transaction.MediumTransaction.newBuilder().addData(ByteString.copyFrom(dest_img.toArray())).build();
//            Transaction.MediumTransaction reply = Transaction.MediumTransaction.newBuilder().addData(request.getData(0)).build();

//            long end_proc = System.nanoTime();
////            procTime[request.getId()] = end_proc-init_proc;
//            System.out.printf("Java, gRPC, Medium, %d, %d\n", request.getId(), end_proc-init_proc);

            responseObserver.onNext(reply);

            responseObserver.onCompleted();
        }
    }

    static class MatrixOpServiceImpl extends MatrixOpGRPCGrpc.MatrixOpGRPCImplBase {
        @Override
        public StreamObserver<Transaction.LargeTransaction> multiplier(StreamObserver<Transaction.LargeTransaction> responseObserver) {
            return new LargeTransactionObserver(responseObserver);
        }

        class LargeTransactionObserver implements StreamObserver<Transaction.LargeTransaction> {
            private double[][] matrix1 = null;
            private StreamObserver<Transaction.LargeTransaction> responseObserver;

            public LargeTransactionObserver(StreamObserver<Transaction.LargeTransaction> responseObserver) {
                this.responseObserver = responseObserver;
            }

//            @Override
//            public void onNext(Transaction.LargeTransaction largeTransaction) {
//
//
//                Double[][] matrix2 = unpaser2DArray(largeTransaction);
//                if (matrix1 == null) {
//                    matrix1 = genIdentityMatrix(matrix2);
//                }
//
//                try {
//                    matrix1 = matrixMultiplier(matrix1, matrix2);
//                } catch (Exception e) {
//                    e.printStackTrace();
//                }
//
//                Transaction.LargeTransaction reply = parse2D(matrix1, largeTransaction.getId());
//
//                this.responseObserver.onNext(reply);
//            }
            @Override
            public void onNext(Transaction.LargeTransaction largeTransaction) {


                double[][] matrix2 = unpaser2DArray(largeTransaction);
                if (matrix1 == null) {
                    matrix1 = genIdentityMatrix(matrix2);
                }

                try {
                    matrix1 = matrixMultiplier(matrix1, matrix2);
                } catch (Exception e) {
                    e.printStackTrace();
                }

                Transaction.LargeTransaction reply = parse2D(matrix1, largeTransaction.getId());

                this.responseObserver.onNext(reply);
            }

            @Override
            public void onError(Throwable throwable) {

            }

            @Override
            public void onCompleted() {
                this.responseObserver.onCompleted();
            }

//            public static Transaction.LargeTransaction parse2D(Double[][] matrix, int id) {
//                Transaction.LargeTransaction.Builder rowsBuilder = Transaction.LargeTransaction.newBuilder().setId(id);
//                for (int i = 0; i < matrix.length; i++) {
//                    Transaction.LargeTransaction.Row r = Transaction.LargeTransaction.Row.newBuilder().addAllRow(Arrays.asList(matrix[i])).build();
//                    rowsBuilder.addMatrix(r);
//                }
//
//                return rowsBuilder.build();
//            }
            public static Transaction.LargeTransaction parse2D(double[][] matrix, int id) {
//                Double[][] inverse = new Double[matrix.length][matrix[0].length];
//                for (int i = 0; i < matrix.length; i++)
//                    for (int j = 0; j < matrix[0].length; j++)
//                        inverse[i][j] = matrix[i][j];
//
//                Transaction.LargeTransaction.Builder rowsBuilder = Transaction.LargeTransaction.newBuilder().setId(id);
//                for (int i = 0; i < matrix.length; i++) {
//                    Transaction.LargeTransaction.Row r = Transaction.LargeTransaction.Row.newBuilder().addAllRow(Arrays.asList(inverse[i])).build();
//                    rowsBuilder.addMatrix(r);
//                }
                Transaction.LargeTransaction.Builder rowsBuilder = Transaction.LargeTransaction.newBuilder().setId(id);
                for (int i = 0; i < matrix.length; i++) {
                    Transaction.LargeTransaction.Row.Builder r = Transaction.LargeTransaction.Row.newBuilder();
                    for (int j = 0; j < matrix[0].length; j++) {
                        r.addRow(matrix[i][j]);
                    }
                    rowsBuilder.addMatrix(r);
                }

                return rowsBuilder.build();
            }

//            public static Double[][] unpaser2DArray(Transaction.LargeTransaction largeTransaction) {
//                List<Transaction.LargeTransaction.Row> rows = largeTransaction.getMatrixList();
////                double[][] ret = new double[rows.size()][];
////                ret = rows.toArray(ret);
//
//                Double[][] inverse = new Double[rows.size()][rows.size()];
//                for (int i = 0; i < rows.size(); i++)
//                    for (int j = 0; j < rows.size(); j++)
//                        inverse[i][j] = rows.get(i).getRow(j);
//
//                return inverse;
//            }
            public static double[][] unpaser2DArray(Transaction.LargeTransaction largeTransaction) {
                List<Transaction.LargeTransaction.Row> rows = largeTransaction.getMatrixList();
            //                double[][] ret = new double[rows.size()][];
            //                ret = rows.toArray(ret);

                double[][] inverse = new double[rows.size()][rows.size()];
                for (int i = 0; i < rows.size(); i++)
                    for (int j = 0; j < rows.size(); j++)
                        inverse[i][j] = rows.get(i).getRow(j);

                return inverse;
            }

//            public static Double[][] genIdentityMatrix(Double[][] matrix) {
//                int rowsNum = matrix.length;
//
//                Double[][] idMatrix = new Double[rowsNum][rowsNum];
//
//                for(int i = 0; i < rowsNum; i++)
//                    for(int j = 0; j < rowsNum; j++)
//                        if (i == j)
//                            idMatrix[i][j] = 1.0;
//                        else
//                            idMatrix[i][j] = 0.0;
//
//                return idMatrix;
//            }
            public static double[][] genIdentityMatrix(double[][] matrix) {
                int rowsNum = matrix.length;

                double[][] idMatrix = new double[rowsNum][rowsNum];

                for(int i = 0; i < rowsNum; i++)
                    for(int j = 0; j < rowsNum; j++)
                        if (i == j)
                            idMatrix[i][j] = 1.0;
                        else
                            idMatrix[i][j] = 0.0;

                return idMatrix;
            }

            /**
             矩阵乘法
             a为m行k列, b为k行n列
             */
//            public static Double[][] matrixMultiplier(Double[][] a, Double[][] b) throws Exception {
//                int m = a.length;
//                int k = b.length;
//                int n = b[0].length;
//
//                if (a[0].length != b.length) {
//                    throw new Exception("矩阵无法相乘");
//                }
//
//                Double[][] res = new Double[m][n];
//                for (int i = 0; i < m; i++) {
//                    for (int j = 0; j < n; j++) {
//                        Double resI = 0.0;
//                        for (int x = 0; x < k; x++) {
//                            resI += a[i][x] * b[x][j];
//                        }
//                        res[i][j] = resI;
//                    }
//
//                }
//
//                return res;
//            }
            public static double[][] matrixMultiplier(double[][] a, double[][] b) throws Exception {
                int m = a.length;
                int k = b.length;
                int n = b[0].length;

                if (a[0].length != b.length) {
                    throw new Exception("矩阵无法相乘");
                }

                double[][] res = new double[m][n];
                for (int i = 0; i < m; i++) {
                    for (int j = 0; j < n; j++) {
                        double resI = 0.0;
                        for (int x = 0; x < k; x++) {
                            resI += a[i][x] * b[x][j];
                        }
                        res[i][j] = resI;
                    }

                }

                return res;
            }
        }
    }

    public static long mean(long[] arr) {
        long result = 0;
        for (long item: arr)
            result += item;
        return result / arr.length;
    }

    public static void main(String[] args) {

        GRPCServer grpcServer = null;
        try {
            grpcServer = new GRPCServer();
            grpcServer.start();
        } catch (IOException e) {
            e.printStackTrace();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

    }

}
