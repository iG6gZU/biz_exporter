package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	counter1 = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "biz_request_total",
		Help: "当前业务调用量",
	}, []string{"topic", "biz"})
	counter2 = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "biz_success_total",
		Help: "当前业务成功量",
	}, []string{"topic", "biz"})
	counter3 = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "biz_timeout_total",
		Help: "当前业务超时量",
	}, []string{"topic", "biz"})
	counter4 = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "biz_fail_total",
		Help: "当前业务失败量",
	}, []string{"topic", "biz", "code", "mes"})
	Counters = make([]*prometheus.CounterVec, 4)
)

type kafkaInfo struct {
	brokers  []string
	group    string
	topics   []string
	username string
	password string
	cluster  string
}

func init() {
	prometheus.MustRegister(counter1, counter2, counter3, counter4)
}

func main() {
	fmt.Println("exporter启动...")
	Counters[0] = counter1
	Counters[1] = counter2
	Counters[2] = counter3
	Counters[3] = counter4

	kafka1 := kafkaConn{
		brokers:  []string{},
		group:    "monitor",
		topics:   []string{},
		username: "",
		password: "",
		cluster:  "",
	}

	kafka2 := kafkaConn{
		brokers:  []string{},
		group:    "monitor",
		topics:   []string{},
		username: "",
		password: "",
		cluster:  "",
	}

	go kafka1.consumer()
	go kafka2.consumer()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":", nil))
}
