/**
 * @Author: peidong pei
 * @Description: 描述
 * @File: dvcollection
 * @Date: 2026/1/16 18:02
 * @Software : GoLand
 */

package dashvector

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

type dvCollection struct {
	client *Client
	name   string
}

func newDvCollection(client *Client, name string) Collection {
	return &dvCollection{
		client: client,
		name:   name,
	}
}

type dvResult struct {
	DocOp   string `json:"doc_op"`
	Id      string `json:"id"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type insertInput struct {
	Docs []*Document `json:"docs"`
}

func (c *dvCollection) InsertOne(ctx context.Context, doc Documentation, options ...InsertOption) (*Result, error) {
	res, err := c.InsertMany(ctx, []Documentation{doc}, options...)
	if err != nil {
		return nil, err
	}
	return &res[0], nil
}

func (c *dvCollection) InsertMany(ctx context.Context, docs []Documentation, options ...InsertOption) ([]Result, error) {
	path := fmt.Sprintf("collections/%s/docs", c.name)
	insertDocs := make([]*Document, len(docs))
	for i := range docs {
		insertDocs[i] = docs[i].MarshalDocument()
	}
	var insertOptions InsertOptions
	for _, option := range options {
		option(&insertOptions)
	}
	if insertOptions.AutoId {
		for i := range insertDocs {
			insertDocs[i].Id = nil
		}
	}
	res, err := call[insertInput, []dvResult](c.client, ctx, "POST", path, &insertInput{
		Docs: insertDocs,
	})
	if err != nil {
		return nil, errors.Join(ErrInsertFailed, err)
	}
	if res == nil || len(*res) == 0 {
		return nil, ErrInsertFailed
	}
	r := make([]Result, len(*res))
	for i := range *res {
		r[i] = Result{
			Id:  (*res)[i].Id,
			Msg: (*res)[i].Message,
			Op:  (*res)[i].DocOp,
		}
	}
	return r, nil
}

type SearchResult struct {
	docs []Document
	err  error
}

func (s *SearchResult) One(res any) error {
	if len(s.docs) == 0 {
		return ErrEmptyResult
	}
	//断言，查看是否实现Documentation接口
	if doc, ok := res.(Documentation); ok {
		if len(s.docs) > 0 {
			doc.UnmarshalDocument(&s.docs[0])
			return s.err
		}
	}
	return ErrInsertFailed
}

func (s *SearchResult) All(results any) error {
	resultsVal := reflect.ValueOf(results)
	if resultsVal.Kind() != reflect.Ptr {
		return fmt.Errorf("results argument must be a pointer to a slice, but was a %s", resultsVal.Kind())
	}

	sliceVal := resultsVal.Elem()
	if sliceVal.Kind() == reflect.Interface {
		sliceVal = sliceVal.Elem()
	}

	if sliceVal.Kind() != reflect.Slice {
		return fmt.Errorf("results argument must be a pointer to a slice, but was a pointer to %s", sliceVal.Kind())
	}

	elementType := sliceVal.Type().Elem()
	docType := reflect.TypeFor[Documentation]()
	if elementType.Kind() == reflect.Ptr && !elementType.Implements(docType) {
		return fmt.Errorf(
			"slice element type %s must be a pointer type implementing Documentation",
			elementType,
		)
	} else if elementType.Kind() == reflect.Struct {
		ptrType := reflect.PointerTo(elementType)
		if !ptrType.Implements(docType) {
			return fmt.Errorf(
				"slice element type %s must implement Documentation",
				elementType,
			)
		}
	} else {
		return fmt.Errorf(
			"slice element type %s must be a pointer type implementing Documentation",
			elementType,
		)
	}
	sliceVal.Set(reflect.MakeSlice(sliceVal.Type(), 0, len(s.docs)))
	for i := range s.docs {
		// new(T) —— 这里是反射唯一合理的使用点
		var elem reflect.Value
		if elementType.Kind() == reflect.Ptr {
			// 元素是指针类型
			elem = reflect.New(elementType.Elem())
			docImpl := elem.Interface().(Documentation)
			docImpl.UnmarshalDocument(&s.docs[i])
		} else {
			// 元素是值类型，生成值类型，再取地址调用接口
			elem = reflect.New(elementType).Elem()
			docPtr := elem.Addr().Interface()
			docImpl := docPtr.(Documentation)
			docImpl.UnmarshalDocument(&s.docs[i])
		}
		sliceVal.Set(reflect.Append(sliceVal, elem))

	}
	return nil

}

func (c *dvCollection) Search(ctx context.Context, query SearchQuery, options ...SearchOption) *SearchResult {
	path := fmt.Sprintf("collections/%s/query", c.name)
	for _, opt := range options {
		opt(&query)
	}
	res, err := call[SearchQuery, []Document](c.client, ctx, "POST", path, &query)
	if err != nil {
		return &SearchResult{
			err: err,
		}
	}
	return &SearchResult{
		docs: *res,
		err:  nil,
	}

}

func (c *dvCollection) DeleteById(ctx context.Context, Id string) error {
	path := fmt.Sprintf("collections/%s/docs", c.name)
	_, err := call[map[string]any, any](c.client, ctx, "DELETE", path, &map[string]any{
		"ids": []string{Id},
	})
	if err != nil {
		return err
	}
	return nil

}

func (c *dvCollection) DeleteByIds(ctx context.Context, Ids []string) error {
	path := fmt.Sprintf("collections/%s/docs", c.name)
	_, err := call[map[string]any, any](c.client, ctx, "DELETE", path, &map[string]any{
		"ids": Ids,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *dvCollection) Upsert(ctx context.Context, doc Documentation) (*Result, error) {
	res, err := c.UpsertMany(ctx, []Documentation{doc})
	if err != nil {
		return nil, err
	}
	return &res[0], nil
}

func (c *dvCollection) UpsertMany(ctx context.Context, docs []Documentation) ([]Result, error) {
	path := fmt.Sprintf("collections/%s/docs", c.name)
	insertDocs := make([]*Document, len(docs))
	for i := range docs {
		insertDocs[i] = docs[i].MarshalDocument()
	}
	res, err := call[insertInput, []dvResult](c.client, ctx, "PUT", path, &insertInput{Docs: insertDocs})
	if err != nil {
		return nil, err
	}
	return toResults(*res), nil
}

func (c *dvCollection) Name() string {
	return c.name
}

func toResults(res []dvResult) []Result {
	results := make([]Result, len(res))
	for i := range res {
		results[i] = Result{
			Op:  res[i].DocOp,
			Id:  res[i].Id,
			Msg: res[i].Message,
		}
	}
	return results
}
