// go build serverprofile.go

// ./serverprofile -cpuprofile=cpu.pprof
// ./serverprofile -memprofile=mem.pprof
// ./serverprofile -blockprofile=block.pprof
// ./serverprofile -traceprofile=trace.pprof

// go tool pprof cpu cpu.pprof
// go tool pprof mem mem.pprof
// go tool pprof block block.pprof
// go tool trace -http="127.0.0.1:8081" trace

// go tool pprof -http=":8081" cpu cpu.pprof
// go tool pprof -http=:8080 dumps/pprof.pprof.goroutine.001.pb.gz

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go_code/rest_vs_grpc_in_java_and_golang/proto"
	"go_code/rest_vs_grpc_in_java_and_golang/utils"
	"google.golang.org/grpc"
	"log"
	"math"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"sync"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")
var blockprofile = flag.String("blockprofile", "", "write block profile to this file")
var traceprofile = flag.String("traceprofile", "", "write trace profile to file")

const PORT = ":8080"

type server struct {
}

func main() {

	runtime.GOMAXPROCS(1)

	go func() {
		//panic(http.ListenAndServe(":8000", nil))
		log.Println(http.ListenAndServe("localhost:8000", nil))
	}()

	flag.Parse()
	if *cpuprofile != "" {
		fmt.Println(*cpuprofile)
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)

		// The server code be test
		go func() {
			http.Handle("/metrics", promhttp.Handler())
			http.ListenAndServe(":2112", nil)
		}()
		go GoGRPCServer()
		time.Sleep(time.Second * 30)

		pprof.StopCPUProfile()
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		//pprof.WriteHeapProfile(f)
		p := pprof.Lookup("allocs")

		go func() {
			http.Handle("/metrics", promhttp.Handler())
			http.ListenAndServe(":2112", nil)
		}()
		go GoGRPCServer()
		time.Sleep(time.Second * 30)

		p.WriteTo(f, 0)
		f.Close()
	}

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
		go GoGRPCServer()
		time.Sleep(time.Second * 30)

		p.WriteTo(f, 0)
		f.Close()
	}

	if *traceprofile != "" {
		f, err := os.Create(*traceprofile)
		if err != nil {
			log.Fatal(err)
		}
		trace.Start(f)
		defer f.Close()
		defer trace.Stop()

		go GoGRPCServer()
		time.Sleep(time.Second * 5)
	}
	GoGRPCServer()
}

func GoGRPCServer() {
	fmt.Println("Server started, listening on " + PORT)
	fmt.Println("Language, method, transaction, id, time")
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterSquareGRPCServer(s, &server{})
	proto.RegisterFilterGRPCServer(s, &server{})
	proto.RegisterMatrixOpGRPCServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}

func (s *server) Square(ctx context.Context, transaction *proto.SmallTransaction) (*proto.SmallTransaction, error) {

	//start := time.Now()

	ret := utils.Square(transaction.Data)

	//elapsed := time.Since(start)
	//
	//fmt.Printf("Go, gRPC, Small, %d, %d\n", i, elapsed)

	//time.Sleep(time.Second * 5)
	return &proto.SmallTransaction{Data: ret}, nil
}

func (s *server) GrayscaleFilter(ctx context.Context, in *proto.MediumTransaction) (*proto.MediumTransaction, error) {

	//init_proc := time.Now().UnixNano() / int64(time.Millisecond)

	dst_en := utils.GrayscaleFilter(in.GetData()[0])
	out_img := proto.MediumTransaction{Data: [][]byte{dst_en}}
	//out_img := proto.MediumTransaction{Data: in.GetData()}

	//end_proc := time.Now().UnixNano() / int64(time.Millisecond)
	//
	//fmt.Printf("Go, gRPC, Medium, %d, %d\n", in.Id, end_proc-init_proc)

	return &out_img, nil
}

//func (s *server) Multiplier(allStr proto.MatrixOpGRPC_MultiplierServer) error {
//
//	start := time.Now()
//
//	var id int32
//
//	link := make(chan [][]float64)
//	done := make(chan bool)
//
//	wg := sync.WaitGroup{}
//	wg.Add(2)
//	go func() {
//		defer wg.Done()
//		defer close(link)
//
//		var matrix1 [][]float64
//		loopNum := math.MaxInt64
//		for i := 0; i < loopNum; i++ {
//			select {
//			case _ = <-done:
//				break
//			default:
//				data, _ := allStr.Recv()
//				if data == nil {
//					continue
//				}
//				id = data.Id
//				matrix2 := utils.Unpaser2DArray(data)
//				if matrix1 == nil {
//					matrix1 = utils.GenIdentityMatrix(matrix2)
//				}
//				matrix1 = utils.MatrixMultiplier(matrix1, matrix2)
//				loopNum = len(matrix1)
//				//time.Sleep(time.Second)
//				link <- matrix1
//			}
//		}
//	}()
//
//	go func() {
//		defer wg.Done()
//		for matrix := range link {
//			m := utils.Parse2D(matrix)
//			_ = allStr.Send(&proto.LargeTransaction{Id: id, Matrix: m})
//			//time.Sleep(time.Second)
//		}
//	}()
//
//	wg.Wait()
//
//	elapsed := time.Since(start)
//	fmt.Printf("Go, gRPC, Large, %d, %d\n", 0, elapsed)
//	//procTime[id] = elapsed
//
//	return nil
//}

func (s *server) Multiplier(allStr proto.MatrixOpGRPC_MultiplierServer) error {

	//start := time.Now()

	var id int32

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		var matrix1 [][]float64
		loopNum := math.MaxInt64
		for i := 0; i < loopNum; i++ {

			data, _ := allStr.Recv()
			if data == nil {
				continue
			}
			id = data.Id
			matrix2 := utils.Unpaser2DArray(data)
			if matrix1 == nil {
				matrix1 = utils.GenIdentityMatrix(matrix2)
			}
			matrix1 = utils.MatrixMultiplier(matrix1, matrix2)
			loopNum = len(matrix1)

			m := utils.Parse2D(matrix1)
			_ = allStr.Send(&proto.LargeTransaction{Id: id, Matrix: m})

		}
	}()

	wg.Wait()

	//elapsed := time.Since(start)
	//fmt.Printf("Go, gRPC, Large, %d, %d\n", 0, elapsed)
	//procTime[id] = elapsed

	return nil
}
