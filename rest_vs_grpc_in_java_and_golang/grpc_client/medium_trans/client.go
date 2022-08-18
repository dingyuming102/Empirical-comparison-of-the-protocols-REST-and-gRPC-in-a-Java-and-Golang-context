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

var routineNum = 64
var size = routineNum
var srcFileName = "asserts/pic/img5.jpg"
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
		mediumTransGRPCClient()

		p.WriteTo(f, 0)
		f.Close()
	}

	if *traceprofile != "" {
		f, err := os.Create(*traceprofile)
		if err != nil {
			log.Fatal(err)
		}
		trace.Start(f)

		mediumTransGRPCClient()

		defer f.Close()
		defer trace.Stop()
	}

	for i := 0; i < 15; i++ {
		go mediumTransGRPCClient()
		time.Sleep(time.Second)
	}
	//mediumTransGRPCClient2()
	fmt.Println(utils.Mean(timer[:]) / int64(time.Millisecond))
	fmt.Println(utils.StdDev(timer) / float64(time.Millisecond))
	fmt.Println(utils.Min(timer[:]) / int64(time.Millisecond))
	fmt.Println(utils.Max(timer) / int64(time.Millisecond))
	fmt.Println(timer)
}

func mediumTransGRPCClient() {
	//fmt.Println("Language, method, transaction, id, time")
	data := utils.MediumTransInitData(srcFileName, size)

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
			reply := unaryMode(i, data, conn)
			conn.Close()
			utils.DecodeAndSave(i, reply[0])

			elapsed := time.Since(start)
			timer[i] = elapsed
			//fmt.Printf("Go, gRPC, Small, %d, %d\n", i, elapsed)

		}(i)
	}

	wg.Wait()
	return
}

func mediumTransGRPCClient2() {
	fmt.Println("Language, method, transaction, id, time")
	data := utils.MediumTransInitData(srcFileName, size)

	for i := 0; i < routineNum; i++ {

		start := time.Now()

		conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		reply := unaryMode(i, data, conn)
		conn.Close()
		utils.DecodeAndSave(i, reply[0])

		elapsed := time.Since(start)
		//fmt.Printf("Go, gRPC, Small, %d, %d\n", i, elapsed)
		timer[i] = elapsed

	}

	return
}

func unaryMode(id int, data [][]byte, conn *grpc.ClientConn) [][]byte {
	c := proto.NewFilterGRPCClient(conn)
	r, err := c.GrayscaleFilter(context.Background(), &proto.MediumTransaction{Id: int32(id), Data: data})
	if err != nil {
		panic(err)
	}
	return r.Data
}
