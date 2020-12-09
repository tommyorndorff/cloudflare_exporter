package main

import (
	"log"
)

var (
	CLOUDFLARE_EMAIL                  = EnvString("CLOUDFLARE_EMAIL", "")                     // (optional) email used for Cloudflare API email authentication
	CLOUDFLARE_KEY                    = EnvString("CLOUDFLARE_KEY", "")                       // (optional) key used for Cloudflare API email authentication
	CLOUDFLARE_TOKEN                  = EnvString("CLOUDFLARE_TOKEN", "")                     // (optional) token used for Cloudflare API token authentication
	CLOUDFLARE_USER_SERVICE_KEY       = EnvString("CLOUDFLARE_USER_SERVICE_KEY", "")          // (optional) key used for Cloudflare API user service key authentication
	CLOUDFLARE_ZONES                  = EnvString("CLOUDFLARE_ZONES", "")                     // (required) comma-separated list of zone names to scrape for metrics (e.g. "example.com,example.org")
	CLOUDFLARE_SCRAPE_ANALYTICS_SINCE = EnvString("CLOUDFLARE_SCRAPE_ANALYTICS_SINCE", "24h") // (optional) `since` parameter of calls to the Cloudflare Analytics API ("Free" tenants have a minimum of 24h)
	EXPORTER_LISTEN_ADDR              = EnvString("EXPORTER_LISTEN_ADDR", "127.0.0.1:9199")   // (optional) address for the exporter to bind to

	cloudflare_metrics *CloudflareMetrics
)

func main() {
	var err error
	cloudflare_metrics, err = New(CLOUDFLARE_EMAIL, CLOUDFLARE_KEY, CLOUDFLARE_TOKEN, CLOUDFLARE_USER_SERVICE_KEY, CLOUDFLARE_ZONES, CLOUDFLARE_SCRAPE_ANALYTICS_SINCE)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("serving metrics at http://%v/metrics\n", EXPORTER_LISTEN_ADDR)
	log.Fatal(ListenAndServe(EXPORTER_LISTEN_ADDR))
}
