/**
 * @Author: peidong pei
 * @Description: 描述
 * @File: schema
 * @Date: 2026/1/16 15:58
 * @Software : GoLand
 */

package schema

type Schema struct {
	Fields  []Field
	Vectors []Vector
	DType   VectorType
	Metric  MetricType
}

func New(fields []Field, vectors []Vector, options ...Option) (*Schema, error) {
	s := &Schema{
		Fields:  fields,
		Vectors: vectors,
		DType:   VectorTypeFloat,
		Metric:  MetricTypeCosine,
	}
	for _, option := range options {
		option(s)
	}
	if err := s.validate(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Schema) validate() error {
	for _, field := range s.Fields {
		if err := field.Validate(); err != nil {
			return err
		}
	}
	for _, vector := range s.Vectors {
		if err := vector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SchemaBuilder struct {
	fields  []Field
	vectors []Vector
	dType   VectorType
	metric  MetricType
}

func NewBuilder() *SchemaBuilder {
	return &SchemaBuilder{}
}

func (b *SchemaBuilder) AddField(fields ...Field) *SchemaBuilder {
	b.fields = append(b.fields, fields...)
	return b
}

func (b *SchemaBuilder) AddVector(vectors ...Vector) *SchemaBuilder {
	b.vectors = append(b.vectors, vectors...)
	return b
}

func (b *SchemaBuilder) SetVectorType(dType VectorType) *SchemaBuilder {
	b.dType = dType
	return b
}

func (b *SchemaBuilder) SetMetricType(metricType MetricType) *SchemaBuilder {
	b.metric = metricType
	return b
}

func (b *SchemaBuilder) Build() (*Schema, error) {
	s := &Schema{
		Fields:  b.fields,
		Vectors: b.vectors,
		DType:   b.dType,
		Metric:  b.metric,
	}
	if err := s.validate(); err != nil {
		return nil, err
	}
	return s, nil
}
