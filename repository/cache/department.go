package cache

import (
	"context"
	"encoding/json"
	"gin-web/models"
	"gin-web/pkg/redis"
)

type DepartmentCache struct {
	redis *redis.CommonRedisClient
}

func NewDepartmentCache(r *redis.CommonRedisClient) *DepartmentCache {
	return &DepartmentCache{
		redis: r,
	}
}

func (d *DepartmentCache) CacheDepartment(ctx context.Context, key string, depts []*models.Department) error {
	result, err := json.Marshal(depts)
	if err != nil {
		return err
	}
	return d.redis.Client.Set(ctx, key, string(result), 0).Err()
}

func (d *DepartmentCache) GetDepartmentCache(ctx context.Context, key string) ([]*models.Department, error) {
	result, err := d.redis.Client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var depts []*models.Department
	if err = json.Unmarshal([]byte(result), &depts); err != nil {
		return nil, err
	}
	return depts, nil
}
