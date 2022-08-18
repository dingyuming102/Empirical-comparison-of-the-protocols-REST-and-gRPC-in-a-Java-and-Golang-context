package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go_code/rest_vs_grpc_in_java_and_golang/proto"
	"go_code/rest_vs_grpc_in_java_and_golang/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"sync"
	"time"
)

var routineNum = 2000 //1440

var size = 65
var timer []time.Duration = make([]time.Duration, routineNum)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")
var blockprofile = flag.String("blockprofile", "", "write block profile to this file")
var traceprofile = flag.String("traceprofile", "", "write trace profile to file")

func main() {
	flag.Parse()

	if *blockprofile != "" {
		f, err := os.Create(*blockprofile)
		if err != nil {
			log.Fatal(err)
		}
		runtime.SetBlockProfileRate(1)
		p := pprof.Lookup("block")

		go func() {
			http.Handle("/metrics", promhttp.Handler())
			http.ListenAndServe(":2112", nil)
		}()
		largeTransGRPCClient()

		p.WriteTo(f, 0)
		f.Close()
	}

	if *traceprofile != "" {
		f, err := os.Create(*traceprofile)
		if err != nil {
			log.Fatal(err)
		}
		trace.Start(f)

		largeTransGRPCClient2()

		defer f.Close()
		defer trace.Stop()
	}

	//for i := 0; i < 30; i++ {
	//	go largeTransGRPCClient()
	//	time.Sleep(time.Second)
	//}
	largeTransGRPCClient2()
	//time.Sleep(time.Second * 10)
	fmt.Println(utils.Mean(timer[:]) / int64(time.Millisecond))
	fmt.Println(utils.StdDev(timer) / float64(time.Millisecond))
	fmt.Println(utils.Min(timer[:]) / int64(time.Millisecond))
	fmt.Println(utils.Max(timer) / int64(time.Millisecond))
	fmt.Println(timer)
}

func largeTransGRPCClient() {
	fmt.Println("Language, method, transaction, id, time")
	data := utils.LargeTransInitData(size)

	wg := sync.WaitGroup{}
	wg.Add(routineNum)
	for i := 0; i < routineNum; i++ {
		go func(i int) {
			defer wg.Done()

			start := time.Now()

			conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				panic(err)
			}
			bidirectSideStreamMode(i, data, conn)
			conn.Close()

			elapsed := time.Since(start)
			fmt.Printf("Go, gRPC, Large, %d, %d\n", i, elapsed)
			timer[i] = elapsed
		}(i)

		if i%(routineNum/30) == 0 && i != 0 {
			time.Sleep(time.Second)
		}
	}

	wg.Wait()

	return
}

func largeTransGRPCClient2() {
	fmt.Println("Language, method, transaction, id, time")
	data := utils.LargeTransInitData(size)

	for i := 0; i < routineNum; i++ {

		start := time.Now()

		conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		bidirectSideStreamMode(i, data, conn)
		conn.Close()

		elapsed := time.Since(start)
		//fmt.Printf("Go, gRPC, Large, %d, %d\n", i, elapsed)
		timer[i] = elapsed

	}

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
		for i := 0; i < size; i++ {
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

//func serverSideStreamMode(conn *grpc.ClientConn) {
//	c := proto.NewGreeterClient(conn)
//	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "慕课网"})
//	for {
//		a, err := res.Recv()
//		if err != nil {
//			fmt.Println(err)
//			break
//		}
//		fmt.Println(a)
//	}
//}
//
//func clientSideStreamMode(conn *grpc.ClientConn) {
//	c := proto.NewGreeterClient(conn)
//	putS, _ := c.PutStream(context.Background())
//	i := 0
//	for true {
//		i++
//		putS.Send(&proto.StreamReqData{
//			Data: fmt.Sprintf("慕课网%d", i),
//		})
//		time.Sleep(time.Second)
//		if i > 10 {
//			break
//		}
//	}
//}
