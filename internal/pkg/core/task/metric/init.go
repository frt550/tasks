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
		"task_cache_repository_task_find_all_hit",
		"Total number of hits in cache for task.FindAll repository method in task service",
	)
	metrics.Register(
		"task_cache_repository_task_find_all_miss",
		"Total number of misses in cache for task.FindAll repository method in task service",
	)
	metrics.Register(
		"task_cache_repository_task_find_one_by_id_hit",
		"Total number of hits in cache for task.FindOneById repository method in task service",
	)
	metrics.Register(
		"task_cache_repository_task_find_one_by_id_miss",
		"Total number of misses in cache for task.FindOneById repository method in task service",
	)
}
