package utils

import (
	"go_code/rest_vs_grpc_in_java_and_golang/proto"
	"math/rand"
)

//func Parse2D(matrix [][]float64) proto.LargeTransaction {
//	var rows []*proto.LargeTransaction_Row
//	for _, arr := range matrix {
//		r := proto.LargeTransaction_Row{Row: arr}
//		rows = append(rows, &r)
//	}
//	return proto.LargeTransaction{Matrix: rows}
//}

func LargeTransInitData(size int) [][][]float64 {
	data := make([][][]float64, size)
	for i := range data {
		data[i] = make([][]float64, size)
		for j := range data[i] {
			data[i][j] = make([]float64, size)
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				data[i][j][k] = rand.Float64()
			}
		}
	}

	return data
}

func Parse2D(matrix [][]float64) []*proto.LargeTransaction_Row {
	var rows []*proto.LargeTransaction_Row
	for _, arr := range matrix {
		r := proto.LargeTransaction_Row{Row: arr}
		rows = append(rows, &r)
	}
	return rows
}

func Unpaser2DArray(data *proto.LargeTransaction) [][]float64 {
	var matrix [][]float64
	for _, row := range data.Matrix {
		matrix = append(matrix, row.Row)
	}
	return matrix
}

func GenIdentityMatrix(matrix [][]float64) [][]float64 {
	rowsNum := len(matrix)

	idMatrix := make([][]float64, rowsNum)
	for i := range idMatrix {
		idMatrix[i] = make([]float64, rowsNum)
	}

	for i := 0; i < rowsNum; i++ {
		for j := 0; j < rowsNum; j++ {
			if i == j {
				idMatrix[i][j] = 1
			} else {
				idMatrix[i][j] = 0
			}
		}
	}
	return idMatrix
}

/**
  矩阵乘法
  a为m行k列, b为k行n列
*/
func MatrixMultiplier(a [][]float64, b [][]float64) [][]float64 {
	m := len(a)
	k := len(b)
	n := len(b[0])

	if len(a[0]) != len(b) {
		panic("矩阵无法相乘")
	}

	res := make([][]float64, m)
	for i := 0; i < m; i++ {
		resI := make([]float64, k)
		for j := 0; j < n; j++ {
			for kk := 0; kk < k; kk++ {
				resI[kk] += a[i][j] * b[j][kk]
			}
		}
		res[i] = resI
	}
	return res
}
