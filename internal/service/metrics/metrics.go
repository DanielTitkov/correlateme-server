package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	UnprocessedUpdateCorrelationsRequests = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "correlateme_unprocessed_update_correlations_requests",
		Help: "Number of request waiting in channel",
	})
	UnprocessedUpdateAggregationsRequests = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "correlateme_unprocessed_update_aggregations_requests",
		Help: "Number of request waiting in channel",
	})
)
