package main

import (
	"flag"

	"github.com/logston/go-check-issues/issues"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

var (
	addr = flag.String(
		"push-address",
		"prometheus-pushgateway.default.svc.cluster.local:9091",
		"The address to push metrics to.",
	)

	pusher *push.Pusher

	counter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "open_github_issue_count",
	})
)

func main() {
	flag.Parse()

	pusher = push.New(*addr, "check_open_issues")
	pusher.Collector(counter)

	issues.FetchIssues()
}
