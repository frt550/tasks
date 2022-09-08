package metric

import (
	"tasks/internal/pkg/core/logger"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus/client_golang/prometheus/promauto"
)

type Interface interface {
	Inc(name string)
	Register(name string, help string)
}

type MetricsStruct struct {
	m map[string]prometheus.Counter
}

func New() *MetricsStruct {
	return &MetricsStruct{
		m: map[string]prometheus.Counter{},
	}
}

func (c *MetricsStruct) Register(name string, help string) {
	c.m[name] = promauto.NewCounter(prometheus.CounterOpts{
		Name: name,
		Help: help,
	})
}

func (c *MetricsStruct) Inc(name string) {
	if metric, ok := c.m[name]; ok {
		metric.Inc()
	} else {
		logger.Logger.Sugar().Infof("metric '%s' is not registered", name)
	}
}
