package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go_code/rest_vs_grpc_in_java_and_golang/utils"
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"strconv"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")
var blockprofile = flag.String("blockprofile", "", "write block profile to this file")
var traceprofile = flag.String("traceprofile", "", "write trace profile to file")

type Matrix struct {
	Id   int
	Data [][]float64
}

func main() {

	//runtime.GOMAXPROCS(1)

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
		go GoRestServer()
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
		go GoRestServer()
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
		go GoRestServer()
		time.Sleep(time.Second * 50)

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

		go GoRestServer()
		time.Sleep(time.Second * 5)
	}
	GoRestServer()
}

func GoRestServer() {
	//r := gin.Default()
	runtime.GOMAXPROCS(1)
	r := gin.New()
	r.UseH2C = true

	r.GET("/greeting", func(c *gin.Context) {
		barr, _ := json.Marshal(&Matrix{Id: 1, Data: [][]float64{{1, 2, 3}, {4, 5, 6}}})
		jsData := Matrix{}
		_ = json.Unmarshal(barr, &jsData)

		c.JSON(200, gin.H{
			"message": "pong",
			"id":      1,
			"data":    [][]float64{{1, 2, 3}, {4, 5, 6}},
		})
	})

	r.GET("/rest/unary/square/:number", func(c *gin.Context) {

		start := time.Now()

		body := c.Param("number")
		data, err := strconv.Atoi(body)
		if err != nil {
			panic(err)
		}
		resData := utils.Square(int32(data))
		fmt.Println(resData)
		res := strconv.Itoa(int(resData))

		//time.Sleep(time.Second * 1000)
		elapsed := time.Since(start)
		fmt.Printf("Go, REST, Small, %d, %d\n", resData, elapsed)

		//c.String(200, res)
		c.JSON(200, res)

	})

	r.POST("/rest/unary/square/", func(c *gin.Context) {

		//start := time.Now()

		body, err := c.GetRawData()
		var data int32
		json.Unmarshal(body, &data)
		if err != nil {
			panic(err)
		}
		resData := utils.Square(data)

		//time.Sleep(time.Second * 1000)
		//elapsed := time.Since(start)
		//fmt.Printf("Go, REST, Small, %d, %d\n", resData, elapsed)

		//c.String(200, res)
		c.JSON(200, resData)

	})

	//限制上传最大尺寸
	//r.MaxMultipartMemory = 8 << 20
	//r.POST("/testpost", func(c *gin.Context) {
	//	file, err := c.FormFile("file")
	//	if err != nil {
	//		c.String(500, "上传图片出错")
	//	}
	//	// c.JSON(200, gin.H{"message": file.Header.Context})
	//	c.SaveUploadedFile(file, file.Filename)
	//	c.String(http.StatusOK, file.Filename)
	//})

	r.POST("/rest/unary/GrayscaleFilter", func(c *gin.Context) {

		//start := time.Now()

		barr, _ := c.GetRawData()

		var barr2 []byte
		json.Unmarshal(barr, &barr2)
		dst_en := utils.GrayscaleFilter(barr2)

		//elapsed := time.Since(start)
		//fmt.Printf("Go, REST, Small, %d, %d\n", intData, elapsed)
		//fmt.Println(elapsed.Milliseconds())

		//c.Data(200, "imageByteData", dst_en)
		c.JSON(200, dst_en)
		//c.JSON(200, barr2)
	})

	r.POST("/rest/unary/Multiplier", func(c *gin.Context) {

		//start := time.Now()

		barr, _ := c.GetRawData()

		var data [][][]float64
		err := json.Unmarshal(barr, &data)
		if err != nil {
			panic(err)
		}

		replyData := generateEmptyMatrixList(data)
		replyData = matrixMultiplication(data, replyData)

		jsonData, err := json.Marshal(replyData)
		if err != nil {
			panic(err)
		}

		c.Data(200, "", jsonData)

		//elapsed := time.Since(start)
		//fmt.Printf("Go, REST, Large, %d\n", elapsed)
	})

	r.Run() // listen and serve on 0.0.0.0:8080

	//if err := r.RunTLS(":8080", "crt/server.crt", "crt/server.key"); err != nil {
	//	fmt.Println(err)
	//}
}

func generateEmptyMatrixList(data [][][]float64) [][][]float64 {

	replyData := make([][][]float64, len(data))
	for i := range data {
		replyData[i] = make([][]float64, len(data))
		for j := range data[i] {
			replyData[i][j] = make([]float64, len(data))
		}
	}

	return replyData
}

func matrixMultiplication(data [][][]float64, replyData [][][]float64) [][][]float64 {

	var matrix1 [][]float64
	for i := 0; i < len(data); i++ {

		matrix2 := data[i]
		if matrix1 == nil {
			matrix1 = utils.GenIdentityMatrix(matrix2)
		}
		matrix1 = utils.MatrixMultiplier(matrix1, matrix2)
		replyData[i] = matrix1
	}

	return replyData
}
