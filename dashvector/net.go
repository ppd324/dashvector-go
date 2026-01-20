/**
 * @Author: peidong pei
 * @Description: 描述
 * @File: net
 * @Date: 2026/1/16 16:24
 * @Software : GoLand
 */

package dashvector

type commonResponse[T any] struct {
	RequestId string `json:"request_id"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Output    T      `json:"output"`
}
