package com.example.java_grpc;


import com.example.proto.Transaction;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.google.protobuf.ByteString;
import com.google.protobuf.InvalidProtocolBufferException;
import org.opencv.core.Mat;
import org.opencv.core.MatOfByte;
import org.opencv.imgcodecs.Imgcodecs;

import java.io.IOException;

public class Test {
    public static void main(String[] args) throws IOException {

        nu.pattern.OpenCV.loadLocally();

        Imgcodecs imageCodecs = new Imgcodecs();
        Mat color_img = imageCodecs.imread("asserts/pic/img5.jpg");

        MatOfByte req_img = new MatOfByte();
        Imgcodecs.imencode(".jpg", color_img, req_img);
        byte[] reqByte = req_img.toArray();

        Transaction.MediumTransaction request = Transaction.MediumTransaction.newBuilder().addData(ByteString.copyFrom(reqByte)).build();
        byte[] requestData = request.toByteArray();
        long init_proc = System.nanoTime();
        for (int i = 0; i < 1000; i++) {
//            Transaction.MediumTransaction request = Transaction.MediumTransaction.newBuilder().addData(ByteString.copyFrom(reqByte)).build();
//            byte[] requestData = request.toByteArray();
            Transaction.MediumTransaction reply = Transaction.MediumTransaction.parseFrom(requestData);
        }
        long end_proc = System.nanoTime();
        System.out.println((end_proc-init_proc) / 1000);

        ObjectMapper mapper = new ObjectMapper();
        requestData = mapper.writeValueAsBytes(reqByte);
        init_proc = System.nanoTime();
        for (int i = 0; i < 1000; i++) {
//            ObjectMapper mapper = new ObjectMapper();
//            byte[] requestData = mapper.writeValueAsBytes(reqByte);
            byte[] reply = mapper.readValue(requestData, requestData.getClass());
        }
        end_proc = System.nanoTime();
        System.out.println((end_proc-init_proc) / 1000);


    }
}
