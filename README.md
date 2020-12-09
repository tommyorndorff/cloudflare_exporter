<p align="center"><img src="https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/320/apple/271/sun-behind-large-cloud_1f325-fe0f.png" width="120px"></p>
<h1 align="center">cloudflare_exporter</h1>
<p align="center"><a href="https://prometheus.io/">Prometheus</a> metrics <a href="https://prometheus.io/docs/instrumenting/exporters/">exporter</a> for <a href="https://www.cloudflare.com/">Cloudflare</a> <a href="https://www.cloudflare.com/analytics/">Analytics</a></p>


# Description

[Prometheus](https://prometheus.io/) metrics [exporter](https://prometheus.io/docs/instrumenting/exporters/) for [Cloudflare](https://www.cloudflare.com/) [Analytics](https://www.cloudflare.com/analytics/).

This piece of software has one mission: gather Cloudflare Site analytics from the [Cloudflare API](https://api.cloudflare.com/), and present them in Prometheus' [Exposition](https://prometheus.io/docs/instrumenting/exposition_formats/) format, for Prometheus to scrape.


# Usage

### Considerations

* **Tune your configuration**. All configuration is done through *environment variables*:
  * `CLOUDFLARE_EMAIL`: *(optional)* email used for Cloudflare API email authentication
  * `CLOUDFLARE_KEY`: *(optional)* key used for Cloudflare API email authentication
  * `CLOUDFLARE_TOKEN`: *(optional)* token used for Cloudflare API token authentication
  * `CLOUDFLARE_USER_SERVICE_KEY`: *(optional)* key used for Cloudflare API user service key authentication
  * `CLOUDFLARE_ZONES`: *(required)* comma-separated list of zone names to scrape for metrics (e.g. `example.com,example.org`)
  * `CLOUDFLARE_SCRAPE_ANALYTICS_SINCE`: *(default: `24h`)* `since` parameter of calls to the Cloudflare Analytics API
  * `EXPORTER_LISTEN_ADDR`: *(default: `127.0.0.1:9199`)* address for the exporter to bind to
* **Beware of rate limiting**, Cloudflare's API has a base limit of [1200 requests every 5 minutes](https://support.cloudflare.com/hc/en-us/articles/200171456-How-many-API-calls-can-I-make-). I recommend configuring your [Prometheis](https://prometheus.io/docs/introduction/faq/#what-is-the-plural-of-prometheus) to scrape `cloudflare_exporter` once every 1-5 minutes.

### With the prebuilt container image

Available on [Docker Hub](https://hub.docker.com) as [`docker.io/ricardbejarano/cloudflare_exporter`](https://hub.docker.com/r/ricardbejarano/cloudflare_exporter):

- [`1.0`, `latest` *(Dockerfile)*](exporter/Dockerfile) (about `13.7MB`)

Also available on [Quay](https://quay.io) as [`quay.io/ricardbejarano/cloudflare_exporter`](https://quay.io/repository/ricardbejarano/cloudflare_exporter):

- [`1.0`, `latest` *(Dockerfile)*](exporter/Dockerfile) (about `13.7MB`)

Any of both registries will do, example:

```bash
docker run -it -p 9199:9199 -e CLOUDFLARE_TOKEN="***" -e CLOUDFLARE_ZONES="example.com,example.org" quay.io/ricardbejarano/cloudflare_exporter
```

### Building the container image from source

First clone the repository, and `cd` into it:

```bash
cd exporter/
docker build -t cloudflare_exporter .
```

Now run it:

```bash
docker run -it -p 9199:9199 -e CLOUDFLARE_TOKEN="***" -e CLOUDFLARE_ZONES="example.com,example.org" cloudflare_exporter
```

### Building the binary from source

First clone the repository, and `cd` into it.

```bash
cd exporter/
make
```

Now run it:

```bash
export CLOUDFLARE_TOKEN="***"  # WARNING: remember to erase this from your Bash history
export CLOUDFLARE_ZONES="example.com,example.org"
./exporter
```

***Pro tip:** during development, feel free to use `make fast` to avoid unnecessary `clean`+`go get ...` when testing new code.*


# FAQs

#### What are the differences with [wehkamp/docker-prometheus-cloudflare-exporter](https://github.com/wehkamp/docker-prometheus-cloudflare-exporter)?

* `cloudflare_exporter` scrapes metrics from the Analytics Dashboard API, available to all Cloudflare customers for free
* `cloudflare_exporter` is not strictly tied to Docker (there's a Docker image, though, see [Usage](#usage))
* `cloudflare_exporter` is written in Go, instead of Python

#### What are the differences with [criteo/cloudflare-exporter](https://github.com/criteo/cloudflare-exporter)?

* `cloudflare_exporter` scrapes metrics from the Analytics Dashboard API, available to all Cloudflare customers for free
* `cloudflare_exporter` is written in Go, instead of Python

#### But does it support pulling analytics by colocation like the others?

Not at the moment, as I require an Enterprise account for development and I don't have one. Feel free to contribute that feature!

#### Does it support Cloudflare [Web Analytics](https://blog.cloudflare.com/privacy-first-web-analytics/)?

I finished writing `cloudflare_exporter` exactly 1 day before [Cloudflare announced their Web Analytics service is now available for free](https://blog.cloudflare.com/privacy-first-web-analytics/).

That said, I'd love to include support for that. Once the API and the [Go API client library](https://godoc.org/github.com/prometheus/client_golang/prometheus) support it, I will integrate that into `cloudflare_exporter`. Last I checked (Dec 9th, 2020) neither support Web Analytics yet.

#### What will happen when Cloudflare deprecates their Analytics REST API on March 1st, 2021?

I expect Cloudflare to update their [official Go API client library](https://godoc.org/github.com/prometheus/client_golang/prometheus) to support the new [Analytics GraphQL API](https://developers.cloudflare.com/analytics/graphql-api) before deprecation on March 1st, 2020.
I also expect the Go API client library not to change its interface in the process, but I will fix whatever incompatibilities surge during the transition, if any.

#### What features are coming for `cloudflare_exporter`?

At this stage, development for `cloudflare_exporter` is paused with the following exceptions:
* Bugs in functionality (raise a GitHub Issue)
* Security vulnerabilities (raise a GitHub Issue asking for a GPG public key)
* Third-party contributions (make a pull request)

#### What features would `cloudflare_exporter` be open to merge?

Basically, any usage analytics data that Cloudflare offers over their APIs, including:
* [Cloudflare DNS](https://www.cloudflare.com/dns/) analytics
* [Cloudflare Web Analytics](https://www.cloudflare.com/web-analytics/)
* [Cloudflare Argo](https://www.cloudflare.com/products/argo-smart-routing/) analytics
* Cloudflare Site Analytics [by colocation](https://api.cloudflare.com/#zone-analytics-analytics-by-co-locations)
* [Cloudflare WAF](https://www.cloudflare.com/waf/) analytics
* [Cloudflare Workers KV](https://www.cloudflare.com/products/workers-kv/) analytics
* [Cloudflare Spectrum](https://www.cloudflare.com/products/cloudflare-spectrum/) analytics

#### What features won't `cloudflare_exporter` implement?

For now, these are the features that have been decided not to be implemented, and the reasoning behind those decisions:
* **Pseudo-metrics about site configuration**, to monitor changes to Cloudflare Site configuration. *Why? This feature is out of scope (aggregated, usage-derived analytics) for this project.*
* **Preventive rate limiting**, to avoid Cloudflare-side rate limiting. *Why? There's already one mechanism available to avoid getting rate limited at the Cloudflare level: increasing `scrape_interval` in Prometheus for `cloudflare_exporter`.*

#### How is it licensed?

See [LICENSE](LICENSE) for more details.
