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

var routineNum = 1000
var size = routineNum
var timer = make([]time.Duration, routineNum)

var httpVersion = "http/2"
var url = "http://127.0.0.1:8080/rest/unary/square/"

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

		smallTransRESTClient()

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

		smallTransRESTClient()
	}

	for i := 0; i < 4; i++ {
		go smallTransRESTClient()
		time.Sleep(time.Second)
	}
	//smallTransRESTClient2()
	fmt.Println(utils.Mean(timer[:]))
	fmt.Println(utils.StdDev(timer))
	fmt.Println(utils.Min(timer[:]))
	fmt.Println(utils.Max(timer))
	fmt.Println(timer)
}

func smallTransRESTClient() {
	fmt.Println("Language, method, transaction, id, time")
	data := utils.SmallTransInitData(size)

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
				//resp, err = client.Get(url + strconv.Itoa(int(data[i])))
				reqData, _ := json.Marshal(data[i])
				resp, err = client.Post(url, "application/stream+json", bytes.NewReader(reqData))
			} else {
				//resp, err = http.Get(url + strconv.Itoa(int(data[i])))
				reqData, _ := json.Marshal(data[i])
				resp, err = http.Post(url, "application/stream+json", bytes.NewReader(reqData))

			}
			if err != nil {
				panic(err)
			}
			resData, _ := io.ReadAll(resp.Body)
			//strData := string(resData)
			var intData int32
			json.Unmarshal(resData, &intData)

			//err = resp.Body.Close()
			client.CloseIdleConnections()

			elapsed := time.Since(start)
			//fmt.Printf("Go, REST, Small, %d, %d\n", intData, elapsed)
			timer[i] = elapsed
		}(i)
	}

	wg.Wait()

	return
}

func smallTransRESTClient2() {
	fmt.Println("Language, method, transaction, id, time")
	data := utils.SmallTransInitData(size)

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
			//resp, err = client.Get(url + strconv.Itoa(int(data[i])))
			reqData, _ := json.Marshal(data[i])
			resp, err = client.Post(url, "application/stream+json", bytes.NewReader(reqData))
		} else {
			client = http.Client{}
			reqData, _ := json.Marshal(data[i])
			resp, err = client.Post(url, "application/stream+json", bytes.NewReader(reqData))

		}
		if err != nil {
			panic(err)
		}
		resData, _ := io.ReadAll(resp.Body)
		//strData := string(resData)
		var intData int32
		json.Unmarshal(resData, &intData)

		//err = resp.Body.Close()
		client.CloseIdleConnections()

		elapsed := time.Since(start)
		//fmt.Printf("Go, REST, Small, %d, %d\n", intData, elapsed)
		timer[i] = elapsed

		//time.Sleep(time.Millisecond)
	}

	return
}
