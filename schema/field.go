/**
 * @Author: peidong pei
 * @Description: 描述
 * @File: field
 * @Date: 2026/1/16 15:47
 * @Software : GoLand
 */

package schema

type Field struct {
	Name string    `json:"name"`
	Type FieldType `json:"type"`
}

func (f *Field) Validate() error {
	if f.Name == "" {
		return ErrInvalidFieldNameEmpty
	}
	switch f.Type {
	case FieldTypeString, FieldTypeInt, FieldTypeFloat, FieldTypeLong, FieldTypeStringArray, FieldTypeIntArray, FieldTypeFloatArray, FieldTypeLongArray:
		return nil
	default:
		return ErrInvalidFieldType
	}
}
