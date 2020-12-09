package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func ListenAndServe(addr string) error {
	http.HandleFunc("/", httpHandleFuncRoot)
	http.HandleFunc("/metrics", httpHandleFuncMetrics)

	return http.ListenAndServe(addr, nil)
}

func httpHandleFuncRoot(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Cloudflare metrics exporter for Prometheus.\n"))
	w.Write([]byte("Metrics available at /metrics.\n"))
	w.Write([]byte("\n"))
	w.Write([]byte("Copyright (c) 2020 Ricard Bejarano\n"))
}

func httpHandleFuncMetrics(w http.ResponseWriter, r *http.Request) {
	cloudflare_metrics.update()
	promhttp.Handler().ServeHTTP(w, r)
}
