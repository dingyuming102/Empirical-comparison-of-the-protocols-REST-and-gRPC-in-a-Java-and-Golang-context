// go test -v -count 2 -bench . smallTransGRPCClient_test.go -benchtime 1s
package main

import (
	"context"
	"go_code/rest_vs_grpc_in_java_and_golang/proto"
	"go_code/rest_vs_grpc_in_java_and_golang/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	"testing"
)

var routineNum = 1000
var size = routineNum
var data = utils.SmallTransInitData(size)

func BenchmarkSmallTransGRPCClientSequential(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GRPCClientSequential(data)
	}
}

func BenchmarkSmallTransGRPCClientParallel(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GRPCClientParallel(routineNum, data)
		}
	})
}

//func TestMain(m *testing.M) {
//	go server.GoGRPCServer()
//	os.Exit(m.Run())
//}

func GRPCClientSequential(data []int32) {

	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	unaryMode(0, data[0], conn)
	conn.Close()

	return
}

func GRPCClientParallel(routineNum int, data []int32) {

	wg := sync.WaitGroup{}
	wg.Add(routineNum)
	for i := 0; i < routineNum; i++ {
		go func(i int) {
			defer wg.Done()

			conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				panic(err)
			}
			//unaryMode(i, data[i], conn)
			c := proto.NewSquareGRPCClient(conn)
			_, err = c.Square(context.Background(), &proto.SmallTransaction{Id: int32(i), Data: data[i]})
			if err != nil {
				panic(err)
			}
			conn.Close()

		}(i)
	}

	wg.Wait()
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

// chapter8/sources/benchmark-impl/sequential_test.go
//var (
//	m     map[int64]struct{} = make(map[int64]struct{}, 10)
//	mu    sync.Mutex
//	round int64 = 1
//)
//
//func BenchmarkSequential(b *testing.B) {
//	fmt.Printf("\ngoroutine[%d] enter BenchmarkSequential: round[%d], b.N[%d]\n", tls.ID(), atomic.LoadInt64(&round), b.N)
//	defer func() {
//		atomic.AddInt64(&round, 1)
//	}()
//
//	for i := 0; i < b.N; i++ {
//		mu.Lock()
//		_, ok := m[round]
//		if !ok {
//			m[round] = struct{}{}
//			fmt.Printf("goroutine[%d] enter loop in BenchmarkSequential: round[%d], b.N[%d]\n", tls.ID(), atomic.LoadInt64(&round), b.N)
//		}
//		mu.Unlock()
//	}
//	fmt.Printf("goroutine[%d] exit BenchmarkSequential: round[%d], b.N[%d]\n", tls.ID(), atomic.LoadInt64(&round), b.N)
//}

//var large_routineNum = 200 //1440
//var large_size = 65
//var large_data = utils.LargeTransInitData(large_size)
//
//func BenchmarkLargeTransGRPCClientSequential(b *testing.B) {
//
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			GRPCLargeClientSequential(large_data)
//		}
//	})
//}
//
//func GRPCLargeClientSequential(data [][][]float64) {
//
//	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		panic(err)
//	}
//	bidirectSideStreamMode(0, data, conn)
//	conn.Close()
//
//	return
//}
//
//func bidirectSideStreamMode(id int, data [][][]float64, conn *grpc.ClientConn) {
//	c := proto.NewMatrixOpGRPCClient(conn)
//
//	allStr, _ := c.Multiplier(context.Background())
//	defer conn.Close()
//	wg := sync.WaitGroup{}
//	wg.Add(2)
//	go func() {
//		defer wg.Done()
//
//		//死循环所以无法断开连接
//		for i := 0; i < size; i++ {
//			recvData, _ := allStr.Recv()
//			//fmt.Println("Received message: ", id, i, len(utils.Unpaser2DArray(data)))
//			recvData.Id++ //为了使用data通过编译
//		}
//	}()
//
//	go func(id int32) {
//		defer wg.Done()
//		for _, matrix := range data {
//			m := utils.Parse2D(matrix)
//			_ = allStr.Send(&proto.LargeTransaction{Id: id, Matrix: m})
//			//time.Sleep(time.Second)
//		}
//	}(int32(id))
//
//	wg.Wait()
//	return
//}
