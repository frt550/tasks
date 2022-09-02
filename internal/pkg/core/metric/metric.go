package metric

import (
	"tasks/internal/pkg/core/logger"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Interface interface {
	Inc(name string)
	register(name string, help string)
}

func Instance() Interface {
	return metrics
}

var metrics Interface

type metricsStruct struct {
	m map[string]prometheus.Counter
}

func (c *metricsStruct) register(name string, help string) {
	c.m[name] = promauto.NewCounter(prometheus.CounterOpts{
		Name: name,
		Help: help,
	})
}

func (c *metricsStruct) Inc(name string) {
	if metric, ok := c.m[name]; ok {
		metric.Inc()
	} else {
		logger.Logger.Sugar().Infof("metric '%s' is not registered", name)
	}
}

func init() {
	metrics = &metricsStruct{
		m: map[string]prometheus.Counter{},
	}
	metrics.register(
		"backup_req_in_ok",
		"Total number of inbound requests to backup service that are succeeded",
	)
	metrics.register(
		"backup_req_in_err",
		"Total number of inbound requests to backup service that are failed",
	)
	metrics.register(
		"backup_req_out_ok",
		"Total number of outbound requests from backup service that are succeeded",
	)
	metrics.register(
		"backup_req_out_err",
		"Total number of outbound requests from backup service that are failed",
	)
}
