package handler

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Deployment struct {
	Name string `redis:"name"`
}

type CacheWrapper struct {
	rdb *redis.Client
}

func NewCacheWrapper(r *redis.Client) *CacheWrapper {
	return &CacheWrapper{rdb: r}
}

func (c *CacheWrapper) GetDeploymentByName(ctx *context.Context, name string) []Deployment {
	
	d := Deployment{}

	c.rdb.Get(ctx, name).Scan(&d)
	return []Deployment{d}
}

func (c *CacheWrapper) GetAllDeployments(ctx *context.Context) []string {

	allkeys := c.rdb.HGetAll(ctx, "deployment-keys").Val()
	return allkeys
}

func (c *CacheWrapper) SetDeployment(ctx *context.Context) error {
	return nil
}
