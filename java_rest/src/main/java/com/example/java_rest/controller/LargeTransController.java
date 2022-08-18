package com.example.java_rest.controller;

import com.example.java_rest.services.LargeTransService;
import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
//import io.grpc.stub.ServerCalls;
import org.springframework.web.bind.annotation.*;

@RestController
public class LargeTransController {


    @PostMapping(path = "/rest/unary/Multiplier")
    public @ResponseBody
    double[][][] matrixMultiplier(@RequestBody double[][][] data) throws Exception {

//        long init_proc = System.nanoTime() / 1000000;

        double[][] matrix1 = null;
        double[][][] replyData = new double[data.length][][];
        for (int i = 0; i < data.length; i++) {

            double[][] matrix2 = data[i];
            if (matrix1 == null) {
                matrix1 = LargeTransService.GenIdentityMatrix(matrix2);
            }
            matrix1 = LargeTransService.MatrixMultiplier(matrix1, matrix2);
            replyData[i] = matrix1;
        }

//        long end_proc = System.nanoTime() / 1000000;
//
//        System.out.printf("Java, REST, Large, %d, %d\n", data.length, end_proc-init_proc);

        return replyData;
    }

}
