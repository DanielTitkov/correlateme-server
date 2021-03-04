package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	UnprocessedUpdateCorrelationsRequests = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "correlateme_unprocessed_update_correlation_request",
		Help: "Number of request waiting in channel",
	})
)
