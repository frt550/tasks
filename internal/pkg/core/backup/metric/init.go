package metric

import (
	metricPkg "tasks/internal/pkg/core/metric"
)

func Instance() metricPkg.Interface {
	return metrics
}

var metrics metricPkg.Interface

func init() {
	metrics = metricPkg.New()
	metrics.Register(
		"backup_req_in_ok",
		"Total number of inbound requests to backup service that are succeeded",
	)
	metrics.Register(
		"backup_req_in_err",
		"Total number of inbound requests to backup service that are failed",
	)
	metrics.Register(
		"backup_req_out_ok",
		"Total number of outbound requests from backup service that are succeeded",
	)
	metrics.Register(
		"backup_req_out_err",
		"Total number of outbound requests from backup service that are failed",
	)
}
