# Monitoring `amppackager` in production

Once you've run `amppackager` server in production, you may want to monitor its
health and performance, as well as the performance of the underlying requests to
the AMP document server. 

`amppackager` provides a few [Prometheus](https://prometheus.io/) metrics for such monitoring. They are available via the `/metrics` endpoint.

Prometheus is a powerful monitoring framework. We encourage you to fully utilize
it to make your monitoring convenient, scalable and automated. A few things we
recommend you do:
1.  Explore the metrics `amppackager` provides by following or skimming through
    the tutorial below. Pick the essential metrics you'd like to monitor.
1.  [Set up a Prometheus server](https://prometheus.io/docs/prometheus/latest/getting_started/).
1.  Set up all `amppackager` replicas as targets for the Prometheus server. Use
    the [multi-target exporter
    pattern](https://prometheus.io/docs/guides/multi-target-exporter/#understanding-and-using-the-multi-target-exporter-pattern).
1.  Try
    [querying](https://prometheus.io/docs/prometheus/latest/querying/basics/)
    the essential metrics you chose.
1.  Visualize the metrics via
    [grafana](https://prometheus.io/docs/visualization/grafana/).
1.  Setup [alerts](https://prometheus.io/docs/alerting/latest/overview/) that
    will notify you of abnormal behavior (e.g. latencies growing beyond 60
    seconds - see more [examples](#example-alerts) below).

The sections below walk you through the available metrics, explain how to
manually check them via command line, and how to interpret the results. All the
command line examples are for Linux OS.

## Monitoring health

To check `amppackager`'s health, `curl` its `/healthz` endpoint:

```console
$ curl https://localhost:8080/healthz
```

If the server is up and has a fresh, valid certificate, it will respond with
`ok`. If not, it will provide an error message.

## Monitoring performance

You can take a step further and check a few performance metrics, both for
requests handled by `amppackager`, and for the underlying gateway requests that
`amppackager` sends to the AMP document server. Get all the metrics by `curl`ing
the `/metrics` endpoint. 

There is also [Go pprof](https://pkg.go.dev/net/http/pprof) available on `/debug/pprof`.

## Example: monitoring total requests count

The example command below fetches all the available metrics. It then greps the report for `amppackager_http_duration_seconds_count` metric. This metric counts the HTTP requests the `amppackager` server has processed since it's been up.

```console
$ curl -s https://127.0.0.1:8080/metrics | grep amppackager_http_duration_seconds_count

amppackager_http_duration_seconds_count{code="200",handler="healthz"} 3
amppackager_http_duration_seconds_count{code="200",handler="validityMap"} 5
amppackager_http_duration_seconds_count{code="200",handler="signer"} 6
amppackager_http_duration_seconds_count{code="502",handler="signer"} 4
amppackager_http_duration_seconds_count{code="404",handler="handler_not_assigned"} 1
```
 
The example stats above are broken down by response HTTP code, and by the
internal amppackager's module (handler) that has handled the request. The stats
report 3 requests to the `healthz` handler that got a 200 response (OK), 4
requests to the `signer` handler that got a 502 response (Bad Gateway) etc.

## `amppackager`'s handlers 

The table below lists `amppackager`'s handlers accounted for by the metrics:

| Handler | Responsibility |
|-|-|
| signer | Handles `/priv/doc` requests. Fetches the AMP document, signs it and returns it. |
| certCache | Handles `/amppkg/cert` requests. Returns your Signed Exchange certificate. |
| validityMap | Handles `/amppkg/validity` requests. Returns the [validity data](https://tools.ietf.org/html/draft-yasskin-httpbis-origin-signed-exchanges-impl-00#section-3.6) referred by `amppackager`'s signatures. |
| healthz | Handles `/healthz` requests. Checks if `amppackager` is running and has a valid, fresh certificate. |
| metrics | Handles `/metrics` requests. Reports performance metrics for `amppackager` and for the underlying gateway requests to the AMP document server. |

## Metrics labels: breakdown by handler and response code

For some metrics like `amppackager_http_duration_seconds_count` the stats in the
`/metrics` response are grouped into buckets by two dimensions: the handler
name, and the HTTP response code. E.g. note the buckets in the the example
above:

```console
amppackager_http_duration_seconds_count{code="200",handler="healthz"} 3
amppackager_http_duration_seconds_count{code="200",handler="signer"} 6
amppackager_http_duration_seconds_count{code="502",handler="signer"} 4
```

Invalid requests that were not routed by `amppackager` to any handler are
assigned a special label `handler_not_assigned`:

```console
amppackager_http_duration_seconds_count{code="404",handler="handler_not_assigned"} 1
```

Some metrics only make sense for a particular handler. E.g.
`amppackager_signer_gateway_duration_seconds` and other metrics related to gateway
requests are only related to `signer` handler's operation. Such metrics are only
broken down into buckets by the response code, not by the handler. 

__Labels__ are key-value properties of buckets that indicate the specific
values of the breakdown dimensions, e.g. `code="200"` or `handler="healthz"`.

## Metric types

The two types of metrics are *counters* and *histograms*.

Counter metrics like `amppackager_signer_documents_total` are monotonic increasing.
They track the accumlated increase since the start of the process.

[Histogram metrics](https://prometheus.io/docs/practices/histograms/) like `amppackager_request_duration_seconds` track observed values
in buckets. They also track the number of observations.

For more information, see the [Official Prometheus documentation](https://prometheus.io/docs/concepts/metric_types/).

## Available metrics

The table below lists the key available metrics, along with their types and
labels.

| Metric | Metric type | Explanation | Broken down by HTTP response code? | Broken down by [handler](https://github.com/ampproject/amppackager/blob/doc/monitoring.md#amppackagers-handlers)? |
|--|--|--|--|--|
| amppackager_http_duration_seconds | [Histogram](#metric-types) | `amppackager`'s handlers' latencies in seconds, measured from the moment the handler starts processing the request, to the moment the response is returned. | Yes | Yes |
| amppackager_signer_gateway_requests_total | Counter | Total number of underlying requests sent by `signer` handler to the AMP document server. | Yes | No, specific to [`signer` handler](#amppackagers-handlers). |
| amppackager_signer_gateway_duration_seconds | [Histogram](#metric-types) | Latencies (in seconds) of gateway requests to the AMP document server. | Yes | No, specific to [`signer` handler](#amppackagers-handlers). |
| amppackager_signer_signed_amp_documents_size_bytes | [Histogram](#metric-types) | Actual size (in bytes) of gateway response body from AMP document server. Reported only if signer decided to sign, not return an error or proxy unsigned. | No, specific to 200 (OK) responses. | No, specific to [`signer` handler](#amppackagers-handlers). |
| amppackager_signer_documents_total | Counter | Total number of successful underlying requests to AMP document server, broken down by status based on the action signer has taken: sign or proxy unsigned. Does not account for requests to `amppackager` that resulted in an HTTP error. | No, specific to 200 (OK) responses. | No, specific to [`signer` handler](#amppackagers-handlers). |

## More examples

Check all the stats (note - this command produces numerous metrics that are self-documented; the key ones are also explained in the table above):

```console
$ curl https://127.0.0.1:8080/metrics
```

Check stats for the `total_gateway_requests_by_code` metric: 

```console
$ curl https://127.0.0.1:8080/metrics | grep total_gateway_requests_by_code
```

Check stats for the `amppackager_signer_gateway_duration_seconds_count` metric, for requests
that got an OK response (200) : 

```console
$ curl https://127.0.0.1:8080/metrics | grep amppackager_signer_gateway_duration_seconds | grep code=\"200\"
```

To check the 90th percentile latency, you can use [PromQL](https://prometheus.io/docs/prometheus/latest/querying/basics/) to calculate this from the histogram data.

For example, requests that got an OK response (200) from the `signer` handler:

```
histogram_quantile(
  0.9,
  sum by (le) (
    rate(amppackager_request_duration_seconds{handler="signer",status="200"}[5m])
  )
)
```

## Performance metrics lifetime

The metrics provided by the `/metrics` endpoint are reset upon the execution of
the `amppackager` server binary. Every request to `/metrics` is served with the
stats accumulated since the server's been up, up to the time of the request, but
not including the request itself. 

## Example alerts

Below are a few examples of indicators of possibly abnormal behavior of
`amppackager` and/or the underlying AMP document server. Feel free to adjust the
numbers and check these manually, or setup automatic alerts in Prometheus:

* Non-200 responses count going beyond 1% of all requests.
* Latencies 90 percentile going beyond 60 seconds (of either server).
* Document size 90 percentile going beyond 3.5MB.
* Unsigned documents count going beyond 1% of all documents.

When designing the alerts for your setup, pay special attention to
[requirements](README.md#limitations) that `amppackager` imposes on the AMP
documents you serve.
