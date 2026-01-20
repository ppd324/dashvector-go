/**
 * @Author: peidong pei
 * @Description: 描述
 * @File: option
 * @Date: 2026/1/16 15:13
 * @Software : GoLand
 */

package dashvector

import "time"

type Option func(*Client)

func WithEndpoint(endpoint string) Option {
	return func(c *Client) {
		c.endpoint = endpoint
	}
}

func WithApiKey(apiKey string) Option {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.readTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.writeTimeout = timeout
	}
}

func WithDialTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.dialTimeout = timeout
	}
}

func WithConnPoolSize(size int) Option {
	return func(c *Client) {
		c.connPoolSize = size
	}
}

func WithHttpTrace(trace bool) Option {
	return func(client *Client) {
		client.httpTrace = trace
	}
}

type InsertOptions struct {
	Partition string
	AutoId    bool
}

type InsertOption func(*InsertOptions)

func WithPartition(partition string) InsertOption {
	return func(option *InsertOptions) {
		option.Partition = partition
	}
}

func WithAutoId(autoId bool) InsertOption {
	return func(option *InsertOptions) {
		option.AutoId = autoId
	}
}

type SearchQuery struct {
	Id            *string                  `json:"id,omitempty"`
	Vector        []VectorValue            `json:"vector,omitempty"`
	TopK          *int                     `json:"topk,omitempty"`
	Filter        *string                  `json:"filter,omitempty"`
	IncludeVector *bool                    `json:"include_vector,omitempty"`
	IncludeFields []string                 `json:"include_fields,omitempty"`
	Partition     *string                  `json:"partition,omitempty"`
	Vectors       map[string][]VectorValue `json:"vectors,omitempty"`
}

type SearchOption func(*SearchQuery)

func WithVector(vector []VectorValue) SearchOption {
	return func(query *SearchQuery) {
		query.Vector = vector
	}
}

func WithTopK(topK int) SearchOption {
	return func(query *SearchQuery) {
		query.TopK = &topK
	}
}
func WithFilter(filter string) SearchOption {
	return func(query *SearchQuery) {
		query.Filter = &filter
	}
}
func WithIncludeVector(includeVector bool) SearchOption {
	return func(query *SearchQuery) {
		query.IncludeVector = &includeVector
	}
}
func WithIncludeFields(includeFields []string) SearchOption {
	return func(query *SearchQuery) {
		query.IncludeFields = includeFields
	}
}

func WithId(id string) SearchOption {
	return func(query *SearchQuery) {
		query.Id = &id
	}
}
