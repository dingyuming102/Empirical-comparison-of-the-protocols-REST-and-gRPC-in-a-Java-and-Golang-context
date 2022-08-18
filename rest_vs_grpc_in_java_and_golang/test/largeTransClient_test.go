package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"go_code/rest_vs_grpc_in_java_and_golang/proto"
	"go_code/rest_vs_grpc_in_java_and_golang/utils"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"net"
	"net/http"
	"sync"
	"testing"
)

var large_routineNum = 200 //1440
var large_size = 65
var large_data = utils.LargeTransInitData(large_size)

func BenchmarkLargeTransGRPCClientSequential(b *testing.B) {

	for n := 0; n < b.N; n++ {
		GRPCLargeClientSequential(large_data)
	}

}

func GRPCLargeClientSequential(data [][][]float64) {

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	bidirectSideStreamMode(0, data, conn)
	conn.Close()

	return
}

func bidirectSideStreamMode(id int, data [][][]float64, conn *grpc.ClientConn) {
	c := proto.NewMatrixOpGRPCClient(conn)

	allStr, _ := c.Multiplier(context.Background())
	defer conn.Close()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()

		//死循环所以无法断开连接
		for i := 0; i < large_size; i++ {
			recvData, _ := allStr.Recv()
			//fmt.Println("Received message: ", id, i, len(utils.Unpaser2DArray(data)))
			recvData.Id++ //为了使用data通过编译
		}
	}()

	go func(id int32) {
		defer wg.Done()
		for _, matrix := range data {
			m := utils.Parse2D(matrix)
			_ = allStr.Send(&proto.LargeTransaction{Id: id, Matrix: m})
			//time.Sleep(time.Second)
		}
	}(int32(id))

	wg.Wait()
	return
}

var url = "http://127.0.0.1:8080/rest/unary/Multiplier"
var httpVersion = "http/2"
var encoding = "json"

func BenchmarkLargeTransRESTClientSequential(b *testing.B) {

	for n := 0; n < b.N; n++ {
		RESTLargeClientSequential(large_data)
	}

}

func RESTLargeClientSequential(large_data [][][]float64) {

	var resp *http.Response
	var client *http.Client
	var err error

	if httpVersion == "http/2" {
		client = &http.Client{
			Transport: &http2.Transport{
				// So http2.Transport doesn't complain the URL scheme isn't 'https'
				AllowHTTP: true,
				// Pretend we are dialing a TLS endpoint. (Note, we ignore the passed tls.Config)
				DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
					return net.Dial(network, addr)
				},
			},
		}
	} else {
		client = http.DefaultClient
	}

	if encoding == "json" {
		jsonData, err := json.Marshal(&large_data)

		if err != nil {
			panic(err)
		}
		resp, err = client.Post(url, "application/stream+json", bytes.NewReader(jsonData))
		if err != nil {
			panic(err)
		}
	} else {

	}

	reply, err := io.ReadAll(resp.Body)
	var replyData [][][]float64

	err = resp.Body.Close()
	client.CloseIdleConnections()
	if err != nil {
		panic(err)
	}

	if encoding == "json" {
		err = json.Unmarshal(reply, &replyData)
		if err != nil {
			panic(err)
		}
	} else {

	}

	return
}
