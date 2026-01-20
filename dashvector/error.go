/**
 * @Author: peidong pei
 * @Description: 描述
 * @File: error
 * @Date: 2026/1/19 14:57
 * @Software : GoLand
 */

package dashvector

import "errors"

var (
	ErrInsertFailed = errors.New("insert failed")
	ErrEmptyResult  = errors.New("empty result")
	ErrInvalidType  = errors.New("invalid type,unimpl Documentation")
)
