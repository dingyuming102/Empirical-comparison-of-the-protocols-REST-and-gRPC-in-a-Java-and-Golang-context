package com.example.java_rest.services;

import java.util.Random;

public class LargeTransService {

    private double[][][] data;

    public double[][][] getData() {
        return this.data;
    }

    public LargeTransService(int size) {
        this.data = LargeTransInitData(size);
    }

    public LargeTransService(double[][][] data) {
        this.data = data;
    }

    public double[][][] LargeTransInitData(int size) {
        double[][][] data = new double[size][size][size];

        Random r = new Random();

        for(int i = 0; i < size; i++)
            for(int j = 0; j < size; j++)
                for(int k = 0; k < size; k++)
                    data[i][j][k] = r.nextDouble();

        return data;
    }

    public static double[][] GenIdentityMatrix(double[][] matrix) {
        int rowsNum = matrix.length;

        double[][] idMatrix = new double[rowsNum][rowsNum];

        for(int i = 0; i < rowsNum; i++)
            for(int j = 0; j < rowsNum; j++)
                if (i == j)
                    idMatrix[i][j] = 1;
                else
                    idMatrix[i][j] = 0;

        return idMatrix;
    }

    /**
     矩阵乘法
     a为m行k列, b为k行n列
     */
    public static double[][] MatrixMultiplier(double[][] a, double[][] b) throws Exception {
        int m = a.length;
        int k = b.length;
        int n = b[0].length;

        if (a[0].length != b.length) {
            throw new Exception("矩阵无法相乘");
        }

        double[][] res = new double[m][n];
        for (int i = 0; i < m; i++) {
            double[] resI = new double[k];
            for (int j = 0; j < n; j++) {
                for (int kk = 0; kk < k; kk++) {
                    resI[kk] += a[i][j] * b[j][kk];
                }
            }
            res[i] = resI;
        }

        return res;
    }

}
