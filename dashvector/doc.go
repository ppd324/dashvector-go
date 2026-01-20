/**
 * @Author: peidong pei
 * @Description: 描述
 * @File: doc
 * @Date: 2026/1/16 17:05
 * @Software : GoLand
 */

package dashvector

import (
	"encoding/json"
	"math"
)

type Document struct {
	Id           *string                  `json:"id,omitempty"`
	Vector       []VectorValue            `json:"vector"`
	Fields       map[string]any           `json:"fields"`
	Vectors      map[string][]VectorValue `json:"vectors"`
	SparseVector any                      `json:"sparse_vector"`
}

type VectorValue struct {
	Iv *int
	Fv *float64
}

func NewVectorValueFromFloat(v float64) VectorValue {
	return VectorValue{
		Fv: &v,
		Iv: nil,
	}
}

func NewVectorValueFromInt(v int) VectorValue {
	return VectorValue{
		Iv: &v,
		Fv: nil,
	}
}

func (v *VectorValue) MarshalJSON() ([]byte, error) {
	if v.Iv != nil {
		return json.Marshal(v.Iv)
	}
	return json.Marshal(v.Fv)
}

func (v *VectorValue) UnmarshalJSON(data []byte) error {
	var f float64
	if err := json.Unmarshal(data, &f); err != nil {
		return err
	}

	if f == math.Trunc(f) {
		// 整数
		i := int(f)
		v.Iv = &i
		v.Fv = nil
	} else {
		v.Fv = &f
		v.Iv = nil
	}
	return nil
}

// 文档化
type Documentation interface {
	MarshalDocument() *Document
	UnmarshalDocument(doc *Document)
}
