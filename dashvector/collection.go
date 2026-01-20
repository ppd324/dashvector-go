/**
 * @Author: peidong pei
 * @Description: 描述
 * @File: collection
 * @Date: 2026/1/16 15:33
 * @Software : GoLand
 */

package dashvector

import (
	"context"
)

type Result struct {
	Op  string
	Id  string
	Msg string
}

type Collection interface {
	Name() string
	InsertOne(ctx context.Context, doc Documentation, options ...InsertOption) (*Result, error)
	InsertMany(ctx context.Context, docs []Documentation, options ...InsertOption) ([]Result, error)
	Search(ctx context.Context, query SearchQuery, options ...SearchOption) *SearchResult
	DeleteById(ctx context.Context, Id string) error
	DeleteByIds(ctx context.Context, Ids []string) error
	Upsert(ctx context.Context, doc Documentation) (*Result, error)
	UpsertMany(ctx context.Context, docs []Documentation) ([]Result, error)
}
