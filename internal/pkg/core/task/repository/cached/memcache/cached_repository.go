package memcache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"tasks/internal/pkg/core/logger"
	"tasks/internal/pkg/core/task/metric"
	"tasks/internal/pkg/core/task/models"
	repositoryPkg "tasks/internal/pkg/core/task/repository"

	"github.com/bradfitz/gomemcache/memcache"
)

type cachedRepository struct {
	repository repositoryPkg.Interface
	cache      *memcache.Client
}

func New(repository repositoryPkg.Interface, cache *memcache.Client) repositoryPkg.Interface {
	return &cachedRepository{
		repository: repository,
		cache:      cache,
	}
}

func (c *cachedRepository) FindAll(ctx context.Context, limit, offset uint64) ([]*models.Task, error) {
	// cache only requests without limit and offset
	if limit == 0 && offset == 0 {
		key := keyFindAll()
		item, errGet := c.cache.Get(key)

		// cache hit
		if errGet == nil {
			var result []*models.Task
			if err := json.Unmarshal(item.Value, &result); err != nil {
				logger.Logger.Sugar().Error(err)
				return c.repository.FindAll(ctx, limit, offset)
			}
			metric.Instance().Inc("task_cache_repository_task_find_all_hit")
			return result, nil
		} else if errors.Is(errGet, memcache.ErrCacheMiss) {
			metric.Instance().Inc("task_cache_repository_task_find_all_miss")
			result, err := c.repository.FindAll(ctx, limit, offset)
			if err != nil {
				logger.Logger.Sugar().Error(err)
				return result, err
			}
			j, err := json.Marshal(result)
			if err != nil {
				logger.Logger.Sugar().Error(err)
				return result, nil
			}
			if err := c.cache.Set(&memcache.Item{
				Key:   key,
				Value: j,
			}); err != nil {
				logger.Logger.Sugar().Error(err)
				return result, nil
			}
			return result, nil
		} else {
			logger.Logger.Sugar().Error(errGet)
			result, err := c.repository.FindAll(ctx, limit, offset)
			if err != nil {
				logger.Logger.Sugar().Error(err)
				return result, err
			}
			return result, nil
		}
	}
	return c.repository.FindAll(ctx, limit, offset)
}

func (c *cachedRepository) Insert(ctx context.Context, task *models.Task) error {
	err := c.repository.Insert(ctx, task)
	if err != nil {
		return err
	}
	// invalidate cache of FindAll
	err = c.cache.Delete(keyFindAll())
	if err != nil && !errors.Is(err, memcache.ErrCacheMiss) {
		logger.Logger.Sugar().Error(err)
		return err
	}
	return nil
}

func (c *cachedRepository) Update(ctx context.Context, task *models.Task) error {
	err := c.repository.Update(ctx, task)
	if err != nil {
		return err
	}
	// invalidate cache of FindAll
	err = c.cache.Delete(keyFindAll())
	if err != nil && !errors.Is(err, memcache.ErrCacheMiss) {
		logger.Logger.Sugar().Error(err)
		return err
	}

	// invalidate cache of FindOneById
	err = c.cache.Delete(keyFindOneById(task.Id))
	if err != nil && !errors.Is(err, memcache.ErrCacheMiss) {
		logger.Logger.Sugar().Error(err)
		return err
	}
	return nil
}

func (c *cachedRepository) DeleteById(ctx context.Context, id uint64) error {
	err := c.repository.DeleteById(ctx, id)
	if err != nil {
		return err
	}
	// invalidate cache of FindAll
	err = c.cache.Delete(keyFindAll())
	if err != nil && !errors.Is(err, memcache.ErrCacheMiss) {
		logger.Logger.Sugar().Error(err)
		return err
	}

	// invalidate cache of FindOneById
	err = c.cache.Delete(keyFindOneById(id))
	if err != nil && !errors.Is(err, memcache.ErrCacheMiss) {
		logger.Logger.Sugar().Error(err)
		return err
	}
	return nil
}

func (c *cachedRepository) FindOneById(ctx context.Context, id uint64) (*models.Task, error) {
	key := keyFindOneById(id)
	item, errGet := c.cache.Get(key)
	// cache hit
	if errGet == nil {
		var result *models.Task
		if err := json.Unmarshal(item.Value, &result); err != nil {
			logger.Logger.Sugar().Error(err)
			return c.repository.FindOneById(ctx, id)
		}
		metric.Instance().Inc("task_cache_repository_task_find_one_by_id_hit")
		return result, nil
	} else if errors.Is(errGet, memcache.ErrCacheMiss) {
		metric.Instance().Inc("task_cache_repository_task_find_one_by_id_miss")
		result, err := c.repository.FindOneById(ctx, id)
		if err != nil {
			logger.Logger.Sugar().Error(err)
			return result, err
		}
		j, err := json.Marshal(result)
		if err != nil {
			logger.Logger.Sugar().Error(err)
			return result, nil
		}
		if err := c.cache.Set(&memcache.Item{
			Key:   key,
			Value: j,
		}); err != nil {
			logger.Logger.Sugar().Error(err)
			return result, nil
		}
		return result, nil
	} else {
		logger.Logger.Sugar().Error(errGet)
		result, err := c.repository.FindOneById(ctx, id)
		if err != nil {
			logger.Logger.Sugar().Error(err)
			return result, err
		}
		return result, nil
	}
}

func keyFindAll() string {
	return "repository.task.FindOneByAll"
}

func keyFindOneById(id uint64) string {
	return fmt.Sprintf("repository.task.FindOneById.%d", id)
}
