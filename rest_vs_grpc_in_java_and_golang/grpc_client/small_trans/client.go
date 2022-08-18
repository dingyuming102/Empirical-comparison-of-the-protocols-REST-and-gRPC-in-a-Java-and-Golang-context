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

var routineNum = 100
var size = routineNum
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
		smallTransGRPCClient()

		p.WriteTo(f, 0)
		f.Close()
	}

	if *traceprofile != "" {
		f, err := os.Create(*traceprofile)
		if err != nil {
			log.Fatal(err)
		}
		trace.Start(f)

		smallTransGRPCClient()

		defer f.Close()
		defer trace.Stop()
	}

	//for i := 0; i < 10; i++ {
	//	go smallTransGRPCClient()
	//	time.Sleep(time.Second)
	//}
	smallTransGRPCClient2()
	fmt.Println(utils.Mean(timer[:]))
	fmt.Println(utils.StdDev(timer[:]))
	fmt.Println(utils.Min(timer[:]))
	fmt.Println(utils.Max(timer[:]))

}

func smallTransGRPCClient() {
	fmt.Println("Language, method, transaction, id, time")
	data := utils.SmallTransInitData(size)

	wg := sync.WaitGroup{}
	wg.Add(routineNum)
	for i := 0; i < routineNum; i++ {
		go func(i int) {
			defer wg.Done()

			start := time.Now()

			conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				panic(err)
			}
			unaryMode(i, data[i], conn)
			conn.Close()

			elapsed := time.Since(start)
			//fmt.Printf("Go, gRPC, Small, %d, %d\n", i, elapsed)
			timer[i] = elapsed
		}(i)
	}

	wg.Wait()
	return
}

func smallTransGRPCClient2() {
	fmt.Println("Language, method, transaction, id, time")
	data := utils.SmallTransInitData(size)

	for i := 0; i < routineNum; i++ {

		start := time.Now()

		conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		unaryMode(i, data[i], conn)
		conn.Close()

		elapsed := time.Since(start)
		//fmt.Printf("Go, gRPC, Small, %d, %d\n", i, elapsed)
		timer[i] = elapsed
	}

	return
}

func unaryMode(id int, data int32, conn *grpc.ClientConn) {
	c := proto.NewSquareGRPCClient(conn)
	_, err := c.Square(context.Background(), &proto.SmallTransaction{Id: int32(id), Data: data})
	if err != nil {
		panic(err)
	}
	//fmt.Println(r.Data)
}
