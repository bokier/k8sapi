package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func main() {
	// 跳过证书验证
	var tlsConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	// 设置超时时间
	var transport http.RoundTripper = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
		DisableCompression:    true,
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		serverUrl,_ := url.Parse("https://10.100.0.2:6443")
		log.Println(request.URL.Path)
		p := httputil.NewSingleHostReverseProxy(serverUrl)
		p.Transport = transport
		p.ServeHTTP(writer,request)
	})

	log.Println("[info] 开始代理 k8sapi, :11111 --> :6443")
	err := http.ListenAndServe("0.0.0.0:11111",nil)
	if err != nil {
		log.Fatal(err)
	}
}