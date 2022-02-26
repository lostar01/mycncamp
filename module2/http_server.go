package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

/*

   1. 接收客户端 request，并将 request 中带的 header 写入 response header
   2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
   3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
   4. 当访问 localhost/healthz 时，应返回 200
*/
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	// Use the default  DefaultServerMux
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	os.Setenv("VERSION", "v1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)

	//接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range r.Header {
		for _, vv := range v {
			w.Header().Set(k, vv)
		}
	}

	//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出

	io.WriteString(w, "<h1>云原生培训训练营</h1>")
	clientIp := getCurrentIp(r)
	log.Printf("Success!  client ip %s  response code 200", clientIp)
}

func getCurrentIp(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

//当访问 localhost/healthz 时，应返回 200
func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
	clientIp := getCurrentIp(r)
	log.Printf("Success!  client ip %s  response code 200", clientIp)

}
