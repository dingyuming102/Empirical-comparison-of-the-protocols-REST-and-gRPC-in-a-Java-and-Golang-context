package main

/*
import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go_code/rest_vs_grpc_in_java_and_golang/proto"
	"go_code/rest_vs_grpc_in_java_and_golang/utils"
	"google.golang.org/grpc"
	"math"
	"net"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

const PORT = ":8080"

type server struct {
}

var procTime [1]time.Duration

func main() {
	//go func() {
	//	log.Println(http.ListenAndServe("localhost:8000", nil))
	//}()
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()
	go GoGRPCServer()
	time.Sleep(time.Second * 60)
	fmt.Println(utils.Mean(procTime[:]) / int64(time.Millisecond))
	fmt.Println(utils.StdDev(procTime[:]) / float64(time.Millisecond))
	fmt.Println(utils.Min(procTime[:]) / int64(time.Millisecond))
	fmt.Println(utils.Max(procTime[:]) / int64(time.Millisecond))
	fmt.Println(procTime[:])
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
	//fmt.Printf("Go, gRPC, Small, %d, %d\n", transaction.Id, elapsed)
	//time.Sleep(time.Second * 5)
	return &proto.SmallTransaction{Data: ret}, nil
}

func (s *server) GrayscaleFilter(ctx context.Context, in *proto.MediumTransaction) (*proto.MediumTransaction, error) {

	start := time.Now()

	dst_en := utils.GrayscaleFilter(in.GetData()[0])

	out_img := proto.MediumTransaction{Data: [][]byte{dst_en}}

	elapsed := time.Since(start)
	//fmt.Printf("Go, gRPC, Medium, %d, %d\n", in.Id, elapsed)
	procTime[in.Id] = elapsed

	return &out_img, nil
}

func (s *server) Multiplier(allStr proto.MatrixOpGRPC_MultiplierServer) error {

	start := time.Now()

	var id int32

	link := make(chan [][]float64)
	done := make(chan bool)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer close(link)

		var matrix1 [][]float64
		loopNum := math.MaxInt64
		for i := 0; i < loopNum; i++ {
			select {
			case _ = <-done:
				break
			default:
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
				//time.Sleep(time.Second)
				link <- matrix1
			}
		}
	}()

	go func() {
		defer wg.Done()
		for matrix := range link {
			m := utils.Parse2D(matrix)
			_ = allStr.Send(&proto.LargeTransaction{Id: id, Matrix: m})
			//time.Sleep(time.Second)
		}
	}()

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Go, gRPC, Large, %d, %d\n", 0, elapsed)
	//procTime[id] = elapsed

	return nil
}
*/
