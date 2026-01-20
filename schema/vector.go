/**
 * @Author: peidong pei
 * @Description: 描述
 * @File: vector
 * @Date: 2026/1/16 15:48
 * @Software : GoLand
 */

package schema

import "errors"

type Vector struct {
	Name      string     `json:"name"`
	Dimension int        `json:"dimension"`
	DType     VectorType `json:"dtype"`
	Metric    MetricType `json:"metric"`
}

func (v *Vector) Validate() error {
	if v.Name == "" {
		return ErrInvalidFieldNameEmpty
	}
	if v.Metric == MetricTypeCosine {
		if v.DType != VectorTypeFloat {
			return errors.Join(ErrInvalidFieldType, errors.New("metric cosine only support float"))
		}
	}
	if v.Dimension < 1 || v.Dimension > 20000 {
		return ErrInvalidDimension
	}
	return nil
}
