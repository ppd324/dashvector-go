/**
 * @Author: peidong pei
 * @Description: 描述
 * @File: type
 * @Date: 2026/1/16 15:40
 * @Software : GoLand
 */

package schema

type FieldType string

const (
	FieldTypeString      FieldType = "STRING"
	FieldTypeInt         FieldType = "INT"
	FieldTypeFloat       FieldType = "FLOAT"
	FieldTypeLong        FieldType = "LONG"
	FieldTypeStringArray FieldType = "ARRAY_STRING"
	FieldTypeIntArray    FieldType = "ARRAY_INT"
	FieldTypeFloatArray  FieldType = "ARRAY_FLOAT"
	FieldTypeLongArray   FieldType = "ARRAY_LONG"
)

type VectorType string

const (
	VectorTypeFloat VectorType = "FLOAT"
	VectorTypeInt   VectorType = "INT"
)

type MetricType string

const (
	MetricTypeEuclidean  MetricType = "euclidean"
	MetricTypeCosine     MetricType = "cosine"
	MetricTypeDotproduct MetricType = "dotproduct"
)
