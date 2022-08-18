package main

import (
	protobuf "github.com/golang/protobuf/proto"
	"go_code/rest_vs_grpc_in_java_and_golang/proto"
	"gocv.io/x/gocv"
	"testing"
)

func BenchmarkMediumTransProtobuf(b *testing.B) {

	for n := 0; n < b.N; n++ {
		color_img := gocv.IMRead("asserts/pic/img5.jpg", gocv.IMReadUnchanged)
		req_img, err := gocv.IMEncode(".jpg", color_img)
		if err != nil {
			panic(err)
		}
		mediumData := make([][]byte, 1)
		mediumData[0] = req_img
		_, _ = protobuf.Marshal(&proto.MediumTransaction{Id: 0, Data: mediumData})
	}
}
