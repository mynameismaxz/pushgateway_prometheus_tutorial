package main

import (
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {
	// declare output
	gatewayUrl := "http://localhost:9091"

	// create gauge
	gauge := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "gaugeExample",
		Help: "Example of a gauge",
	}, []string{"label1", "label2", "label3"})

	// create registry for prometheus
	reg := prometheus.NewRegistry()
	reg.MustRegister(gauge)

	// random value to gauge
	go func() {
		for {
			gauge.With(prometheus.Labels{
				"label1": "test",
				"label2": "222",
				"label3": "333",
			}).Set(float64(time.Now().Unix()))
			time.Sleep(3 * time.Second)
		}
	}()

	// push to gateway
	for {
		log.Printf("Push to gateway: %s\n", gatewayUrl)
		err := push.New(gatewayUrl, "gaugeExample").Collector(reg).Push()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)
	}

}
