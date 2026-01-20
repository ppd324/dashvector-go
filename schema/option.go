/**
 * @Author: peidong pei
 * @Description: 描述
 * @File: option
 * @Date: 2026/1/16 16:05
 * @Software : GoLand
 */

package schema

type Option func(schema *Schema)

func WithDType(dtype VectorType) Option {
	return func(schema *Schema) {
		schema.DType = dtype
	}
}

func WithMetricType(metricType MetricType) Option {
	return func(schema *Schema) {
		schema.Metric = metricType
	}
}
