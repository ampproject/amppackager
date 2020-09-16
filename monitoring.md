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

## Example: monitoring total requests count

The example command below fetches all the available metrics. It then greps the report for `total_requests_by_code_and_url` metric. This metric counts the HTTP requests the `amppackager` server has processed since it's been up. 

```console
$ curl  https://127.0.0.1:8080/metrics | grep total_requests_by_code_and_url

# HELP total_requests_by_code_and_url Total number of requests by HTTP code and URL.
# TYPE total_requests_by_code_and_url counter
total_requests_by_code_and_url{code="200",handler="healthz"} 3
total_requests_by_code_and_url{code="200",handler="validityMap"} 5
total_requests_by_code_and_url{code="200",handler="signer"} 6
total_requests_by_code_and_url{code="502",handler="signer"} 4
total_requests_by_code_and_url{code="404",handler="handler_not_assigned"} 1
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

For some metrics like `total_requests_by_code_and_url` the stats in the
`/metrics` response are grouped into buckets by two dimensions: the handler
name, and the HTTP response code. E.g. note the buckets in the the example
above:

```console
total_requests_by_code_and_url{code="200",handler="healthz"} 3
total_requests_by_code_and_url{code="200",handler="signer"} 6
total_requests_by_code_and_url{code="502",handler="signer"} 4
```

Invalid requests that were not routed by `amppackager` to any handler are
assigned a special label `handler_not_assigned`:

```console
total_requests_by_code_and_url{code="404",handler="handler_not_assigned"} 1
```

Some metrics only make sense for a particular handler. E.g.
`gateway_request_latencies_in_seconds` and other metrics related to gateway
requests are only related to `signer` handler's operation. Such metrics are only
broken down into buckets by the response code, not by the handler. 

__Labels__ are key-value properties of buckets that indicate the specific
values of the breakdown dimensions, e.g. `code="200"` or `handler="healthz"`.

## Metrics types: counters and summaries

The two types of metrics are *counters* and *summaries*. 

For some metrics like `total_requests_by_code_and_url`, each request increases
the total by 1, so the metric is a *counter* that has no historical data, just
the accumulated total. 

For other metrics like `request_latencies_in_seconds`, a distribution of
requests latencies is stored, and a few historical percentiles are reported -
0.5 percentile, 0.9 percentile 0.99 percentile. Such metrics are *summaries*.

## Available metrics

The table below lists the key available metrics, along with their types and
labels.

| Metric | Metric type | Explanation | Broken down by HTTP response code? | Broken down by [handler](https://github.com/MichaelRybak/amppackager/blob/doc/monitoring.md#understanding-stats-broken-down-by-amppackagers-handlers)? | 
|--|--|--|--|--|
| total_requests_by_code_and_url | Counter | Total number of requests handled by `amppackager` since it's been up. | Yes | Yes |
| request_latencies_in_seconds | [Summary](#metrics-types-counters-and-summaries) | `amppackager`'s handlers latencies in seconds, measured from the moment the handler starts processing the request, to the moment the response is returned. | Yes | Yes |
| total_gateway_requests_by_code | Counter | Total number of underlying requests sent by `signer` handler to the AMP document server. | Yes | No, specific to [`signer` handler](#amppackagers-handlers). |
| gateway_request_latencies_in_seconds | [Summary](#metrics-types-counters-and-summaries) | Latencies (in seconds) of gateway requests to the AMP document server. | Yes | No, specific to [`signer` handler](#amppackagers-handlers). |
| signed_amp_documents_size_in_bytes | [Summary](#metrics-types-counters-and-summaries) | Actual size (in bytes) of gateway response body from AMP document server. Reported only if signer decided to sign, not return an error or proxy unsigned. | No, specific to 200 (OK) responses. | No, specific to [`signer` handler](#amppackagers-handlers). |
| documents_signed_vs_unsigned | Counter | Total number of successful underlying requests to AMP document server, broken down by status based on the action signer has taken: sign or proxy unsigned. Does not account for requests to `amppackager` that resulted in an HTTP error. | No, specific to 200 (OK) responses. | No, specific to [`signer` handler](#amppackagers-handlers). |

## Understanding percentiles reported by Summaries

A [percentile](https://en.wikipedia.org/wiki/Percentile) indicates the value
below which a given percentage of observations in a group of observations falls.
E.g. if in a room of 100 people, 80% are shorter than you, then you are the 80%
percentile. The 50% percentile is also known as
[median](https://en.wikipedia.org/wiki/Median).

For summary metrics like `request_latencies_in_seconds`, the `/metrics` endpoint provides three percentiles: 0.5, 0.9, 0.99. 

To get an idea of how long it __usually__ takes `amppackager` to handle a
request, look at the respective 0.5 percentile. To check the rare, __worst case
scenarios__, look at 0.9 and 0.99 percentiles.

Consider the following example. Let's say you're interested in the stats for the `request_latencies_in_seconds` metric, specifically for requests that got an OK response (200) from the `signer` handler:

```console
$ curl https://127.0.0.1:8080/metrics | grep request_latencies_in_seconds | grep signer | grep code=\"200\"

request_latencies_in_seconds{code="200",handler="signer",quantile="0.5"} 0.023
request_latencies_in_seconds{code="200",handler="signer",quantile="0.9"} 0.237
request_latencies_in_seconds{code="200",handler="signer",quantile="0.99"} 0.238
request_latencies_in_seconds_sum{code="200",handler="signer"} 661.00
request_latencies_in_seconds_count{code="200",handler="signer"} 10000
```

According to the example stats above, `signer` has handled 10000 requests.
Consider the respective 10000 latencies ranked from smallest to largest. The
latency of the request handled the fastest gets ranked 1, and of the one handled
the slowest - 10000. According to the stats above, the latency ranked 5001 is
0.023s, the latency ranked 9001 is 0.237s, and the latency ranked 9901 is
0.238s. 

__The conclusion for the example above__: successful signer requests are usually handled within 0.023s, but occasionally may take up to 0.238s.

Note that the results provided for latencies and other summaries may be off by a
few rank positions in the ranking. This is due to metrics engine optimizations
that allow to not store all the historical data, therefore saving RAM
significantly. The results are still accurate enough for performance monitoring.


Also note that every stat (e.g. 0.9 percentile latency) provided by the metrics
is an actual historical value that has been seen by the server, not an
approximation.

### Mean vs percentiles

In the example above all the 10000 requests were handled in 661 seconds, which
means the mean (average) latency was 0.0661s. This value is ~3 times larger than
the median. So which one more accurately represents the "typical" scenario? Why
not look at mean instead of looking at percentiles?

Median (0.5 percentile) is [more stable against
outliers](https://en.wikipedia.org/wiki/Median) than mean, and therefore gives a
better understanding of the typical response time. At the same time the 0.9 and
0.99 percentiles give you a good idea about the large outliers, i.e. abnormally
slow response times.

## More examples

Check all the stats (note - this command produces numerous metrics that are self-documented; the key ones are also explained in the table above):

```console
$ curl https://127.0.0.1:8080/metrics
```

Check stats for the `total_gateway_requests_by_code` metric: 

```console
$ curl https://127.0.0.1:8080/metrics | grep total_gateway_requests_by_code
```

Check stats for the `gateway_request_latencies_in_seconds` metric, for requests
that got an OK response (200) : 

```console
$ curl https://127.0.0.1:8080/metrics | grep gateway_request_latencies_in_seconds | grep code=\"200\"
```

Check the 0.9 percentile latency for the `request_latencies_in_seconds` metric,
for requests that got an OK response (200) from the `signer` handler:

```console
$ curl https://127.0.0.1:8080/metrics | grep request_latencies_in_seconds | grep signer | grep code=\"200\" | grep quantile=\"0.9\"
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