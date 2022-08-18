package main

import (
	"encoding/json"
	"fmt"
	protobuf "github.com/golang/protobuf/proto"
	"go_code/rest_vs_grpc_in_java_and_golang/proto"
	"go_code/rest_vs_grpc_in_java_and_golang/utils"
	"gocv.io/x/gocv"
	"io/ioutil"
	"runtime"
	"time"
)

type Matrix struct {
	Id   int
	Data [][]float64
}

func main() {
	//barr, _ := json.Marshal(&Matrix{Id: 1, Data: [][]float64{{1, 2, 3}, {4, 5, 6}}})
	//jsData := Matrix{}
	//_ = json.Unmarshal(barr, &jsData)
	runtime.GOMAXPROCS(1)
	test4()
}

func test1() {
	color_img := gocv.IMRead("asserts/pic/img5.jpg", gocv.IMReadUnchanged)
	req_img, err := gocv.IMEncode(".jpg", color_img)
	if err != nil {
		panic(err)
	}
	data := make([][]byte, 1)
	data[0] = req_img
	fmt.Println(len(data[0]))

	protoData, _ := protobuf.Marshal(&proto.MediumTransaction{Id: 0, Data: data})
	fmt.Println(len(protoData))

	barr, _ := json.Marshal(data)
	fmt.Println(len(barr))

	fileContent, _ := ioutil.ReadFile("asserts/pic/img5.jpg")
	fmt.Println(len(fileContent))
}

func test2() {
	data := utils.LargeTransInitData(65)

	protoData, _ := protobuf.Marshal(&proto.LargeTransaction{
		Id:     0,
		Matrix: utils.Parse2D(data[0]),
	})
	fmt.Println(protoData)
	fmt.Println(len(protoData))

	barr, _ := json.Marshal(data[0])
	fmt.Println(barr)
	fmt.Println(len(barr))
	fmt.Println(data[0])
}

func test3() {
	color_img := gocv.IMRead("asserts/pic/img5.jpg", gocv.IMReadUnchanged)
	req_img, err := gocv.IMEncode(".jpg", color_img)
	if err != nil {
		panic(err)
	}
	data := make([][]byte, 1)
	data[0] = req_img
	fmt.Println(len(data[0]))

	protoData, _ := protobuf.Marshal(&proto.MediumTransaction{Id: 0, Data: data})
	start := time.Now()
	for i := 0; i < 50000; i++ {
		//_, _ = protobuf.Marshal(&proto.MediumTransaction{Id: 0, Data: data})
		protobuf.Unmarshal(protoData, &proto.MediumTransaction{})
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed / 50000)
	//fmt.Println(len(protoData))

	barr, _ := json.Marshal(data)
	start = time.Now()
	for i := 0; i < 5000; i++ {
		//_, _ = json.Marshal(data)
		json.Unmarshal(barr, &[][]byte{})
	}
	elapsed = time.Since(start)
	fmt.Println(elapsed / 5000)
}

//122us 8.29ms
//473us 19.66ms
func test4() {
	data := utils.LargeTransInitData(65)

	//protoData, _ := protobuf.Marshal(&proto.LargeTransaction{
	//	Id:     0,
	//	Matrix: utils.Parse2D(data[0]),
	//})
	start := time.Now()
	for j := 0; j < 10000; j++ {

		for i := 0; i < 65; i++ {
			_, _ = protobuf.Marshal(&proto.LargeTransaction{
				Id:     0,
				Matrix: utils.Parse2D(data[i]),
			})
			//protobuf.Unmarshal(protoData, &proto.LargeTransaction{})
		}

	}
	elapsed := time.Since(start)
	fmt.Println(elapsed / 10000)

	//jsonData, _ := json.Marshal(data)
	start = time.Now()
	for i := 0; i < 1000; i++ {
		_, _ = json.Marshal(data)
		//json.Unmarshal(jsonData, &[][][]float64{})
	}
	elapsed = time.Since(start)
	fmt.Println(elapsed / 1000)

}

func test5() {
	data := utils.LargeTransInitData(75)

	protoData, _ := protobuf.Marshal(&proto.LargeTransaction{
		Id:     0,
		Matrix: utils.Parse2D(data[0]),
	})
	fmt.Println(len(protoData) * len(data))

	jsonData, _ := json.Marshal(data)
	fmt.Println(len(jsonData))
}
