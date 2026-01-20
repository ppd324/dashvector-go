/**
 * @Author: peidong pei
 * @Description: 描述
 * @File: example_test
 * @Date: 2026/1/16 16:39
 * @Software : GoLand
 */

package dashvector_test

import (
	"context"
	"fmt"
	"time"

	"github.com/ppd324/dashvector-go/dashvector"
)

func ExampleNew() {
	client := dashvector.New("endpoint", "apikey")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	coll, err := client.ListCollections(ctx)
	if err != nil {
		panic(err)
	}
	for _, c := range coll {
		fmt.Println(c)
	}
	client.Close()
}
