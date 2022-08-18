package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go_code/rest_vs_grpc_in_java_and_golang/utils"
	"golang.org/x/net/http2"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"sync"
	"time"
)

var routineNum = 48
var size = routineNum
var srcFileName = "asserts/pic/img5.jpg"
var timer = make([]time.Duration, routineNum)

var httpVersion = "http/2"
var url = "http://127.0.0.1:8080/rest/unary/GrayscaleFilter"

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

		mediumTransRESTClient()

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

		mediumTransRESTClient()
	}

	for i := 0; i < 35; i++ {
		go mediumTransRESTClient()
		time.Sleep(time.Second)
	}
	//mediumTransRESTClient()
	fmt.Println(utils.Mean(timer[:]) / int64(time.Millisecond))
	fmt.Println(utils.StdDev(timer) / float64(time.Millisecond))
	fmt.Println(utils.Min(timer[:]) / int64(time.Millisecond))
	fmt.Println(utils.Max(timer) / int64(time.Millisecond))
	fmt.Println(timer)
}

func mediumTransRESTClient() {
	fmt.Println("Language, method, transaction, id, time")
	data := utils.MediumTransInitData(srcFileName, size)

	wg := sync.WaitGroup{}
	wg.Add(routineNum)
	for i := 0; i < routineNum; i++ {
		go func(i int) {
			defer wg.Done()

			start := time.Now()

			var client http.Client
			var resp *http.Response
			var err error
			if httpVersion == "http/2" {
				client = http.Client{
					Transport: &http2.Transport{
						// So http2.Transport doesn't complain the URL scheme isn't 'https'
						AllowHTTP: true,
						// Pretend we are dialing a TLS endpoint. (Note, we ignore the passed tls.Config)
						DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
							return net.Dial(network, addr)
						},
					},
				}
				//resp, err = client.Post(url, "image/jpeg", bytes.NewReader(data[0]))
				jsonData, _ := json.Marshal(data[0])
				resp, err = client.Post(url, "image/jpeg", bytes.NewReader(jsonData))
			} else {
				//resp, err = http.Post(url, "image/jpeg", bytes.NewReader(data[0]))
				jsonData, _ := json.Marshal(data[0])
				client = http.Client{}
				resp, err = client.Post(url, "image/jpeg", bytes.NewReader(jsonData))

			}
			if err != nil {
				panic(err)
			}

			reply, err := io.ReadAll(resp.Body)
			err = resp.Body.Close()

			utils.DecodeAndSave(i, reply)

			elapsed := time.Since(start)
			//fmt.Printf("Go, REST, Small, %d, %d\n", intData, elapsed)
			timer[i] = elapsed
		}(i)
	}

	wg.Wait()
	return
}

func mediumTransRESTClient2() {
	fmt.Println("Language, method, transaction, id, time")
	data := utils.MediumTransInitData(srcFileName, size)

	for i := 0; i < routineNum; i++ {

		start := time.Now()

		var client http.Client
		var resp *http.Response
		var err error
		if httpVersion == "http/2" {
			client = http.Client{
				Transport: &http2.Transport{
					// So http2.Transport doesn't complain the URL scheme isn't 'https'
					AllowHTTP: true,
					// Pretend we are dialing a TLS endpoint. (Note, we ignore the passed tls.Config)
					DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
						return net.Dial(network, addr)
					},
				},
			}
			//resp, err = client.Post(url, "image/jpeg", bytes.NewReader(data[0]))
			jsonData, _ := json.Marshal(data[0])
			resp, err = client.Post(url, "image/jpeg", bytes.NewReader(jsonData))
		} else {
			client = http.Client{}
			//resp, err = client.Post(url, "image/jpeg", bytes.NewReader(data[0]))
			jsonData, _ := json.Marshal(data[0])
			resp, err = client.Post(url, "image/jpeg", bytes.NewReader(jsonData))
		}
		if err != nil {
			panic(err)
		}

		reply, err := io.ReadAll(resp.Body)
		err = resp.Body.Close()
		var jsonreply []byte
		json.Unmarshal(reply, &jsonreply)
		utils.DecodeAndSave(i, jsonreply)

		client.CloseIdleConnections()

		elapsed := time.Since(start)
		//fmt.Printf("Go, REST, Small, %d, %d\n", intData, elapsed)
		timer[i] = elapsed

	}

	return
}
