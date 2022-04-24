package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// metrics
func Register() {
	err := prometheus.Register(functionLatency)
	if err != nil {
		fmt.Println(err)
	}
}

const (
	MetricsNamespace = "cloudnative"
)

// NewExecutionTimer provides a timer for Updater's RunOnce execution
func NewTimer() *ExecutionTimer {
	return NewExecutionTimer(functionLatency)
}

var (
	functionLatency = CreateExecutionTimeMetric(MetricsNamespace,
		"Time spent.")
)

// NewExecutionTimer provides a timer for admission latency; call ObserveXXX() on it to measure
func NewExecutionTimer(histo *prometheus.HistogramVec) *ExecutionTimer {
	now := time.Now()
	return &ExecutionTimer{
		histo: histo,
		start: now,
		last:  now,
	}
}

// ObserveTotal measures the execution time from the creation of the ExecutionTimer
func (t *ExecutionTimer) ObserveTotal() {
	(*t.histo).WithLabelValues("total").Observe(time.Now().Sub(t.start).Seconds())
}

// CreateExecutionTimeMetric prepares a new histogram labeled with execution step
func CreateExecutionTimeMetric(namespace string, help string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "execution_latency_seconds",
			Help:      help,
			Buckets:   prometheus.ExponentialBuckets(0.001, 2, 15),
		}, []string{"step"},
	)
}

// ExecutionTimer measures execution time of a computation, split into major steps
// usual usage pattern is: timer := NewExecutionTimer(...) ; compute ; timer.ObserveStep() ; ... ; timer.ObserveTotal()
type ExecutionTimer struct {
	histo *prometheus.HistogramVec
	start time.Time
	last  time.Time
}

/*

   1. 接收客户端 request，并将 request 中带的 header 写入 response header
   2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
   3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
   4. 当访问 localhost/healthz 时，应返回 200
*/

func main() {
	Register()
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/images", images)
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/healthz", healthz)
	// Use the default  DefaultServerMux
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func images(w http.ResponseWriter, r *http.Request) {
	timer := NewTimer()
	defer timer.ObserveTotal()
	randInt := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(randInt))
	w.Write([]byte(fmt.Sprintf("<h1>%d<h1>", randInt)))
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
