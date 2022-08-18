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

var routineNum = 200
var size = 65
var timer = make([]time.Duration, routineNum)

var url = "http://127.0.0.1:8080/rest/unary/Multiplier"
var httpVersion = "http/2"
var encoding = "json"

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

		largeTransRESTClient()

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

		largeTransRESTClient2()
	}

	//for i := 0; i < 10; i++ {
	//	largeTransRESTClient()
	//	time.Sleep(time.Second)
	//}
	largeTransRESTClient2()
	fmt.Println(utils.Mean(timer[:]) / int64(time.Millisecond))
	fmt.Println(utils.StdDev(timer) / float64(time.Millisecond))
	fmt.Println(utils.Min(timer[:]) / int64(time.Millisecond))
	fmt.Println(utils.Max(timer) / int64(time.Millisecond))
	fmt.Println(timer)
}

func largeTransRESTClient() {
	fmt.Println("Language, method, transaction, id, time")
	data := utils.LargeTransInitData(size)

	wg := sync.WaitGroup{}
	wg.Add(routineNum)
	for i := 0; i < routineNum; i++ {
		go func(i int) {
			defer wg.Done()

			start := time.Now()

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
				jsonData, err := json.Marshal(&data)
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

			elapsed := time.Since(start)
			//fmt.Println(len(reply))
			fmt.Printf("Go, REST, Large, %d, %d\n", i, elapsed/1000000)
			timer[i] = elapsed
		}(i)

		if i%(routineNum/30) == 0 && i != 0 {
			time.Sleep(time.Second)
		}
	}

	wg.Wait()

	return
}

func largeTransRESTClient2() {
	fmt.Println("Language, method, transaction, id, time")
	data := utils.LargeTransInitData(size)

	for i := 0; i < routineNum; i++ {

		start := time.Now()

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
			jsonData, err := json.Marshal(&data)

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

		elapsed := time.Since(start)
		//fmt.Println(len(reply))
		//fmt.Printf("Go, REST, Large, %d, %d\n", i, elapsed)
		timer[i] = elapsed
	}

	return
}
