/**
 * @Author: peidong pei
 * @Description: 描述
 * @File: error
 * @Date: 2026/1/16 15:53
 * @Software : GoLand
 */

package schema

import "errors"

var (
	ErrInvalidFieldType      = errors.New("invalid field type")
	ErrInvalidDimension      = errors.New("invalid dimension,only support (1,2000]")
	ErrInvalidFieldNameEmpty = errors.New("field name is empty")
)
