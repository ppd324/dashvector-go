# dashvector-go

[![Go Reference](https://pkg.go.dev/badge/github.com/ppd324/dashvector-go.svg)](https://pkg.go.dev/github.com/ppd324/dashvector-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/ppd324/dashvector-go)](https://goreportcard.com/report/github.com/ppd324/dashvector-go)
[![License](https://img.shields.io/github/license/ppd324/dashvector-go)](LICENSE)

**dashvector-go** 是阿里云 [DashVector](https://help.aliyun.com/document_detail/2510225.html) 向量检索服务的非官方 Go 语言 SDK。

本项目旨在为 Go 开发者提供简洁、高效的方式来接入 DashVector，轻松实现向量的存储、管理和检索功能，助力构建大模型 RAG、语义搜索和推荐系统等应用。

## ✨ 特性

- **向量操作**：支持向量（Doc）的插入、更新、删除和批量操作。
- **高效检索**：支持向量检索及带过滤条件的混合检索。
- **类型安全**：利用 Go 的强类型特性，减少运行时错误。

## 📦 安装

使用 `go get` 获取最新版本：

```bash
go get [github.com/ppd324/dashvector-go](https://github.com/ppd324/dashvector-go)
```

## 🚀 快速开始

```go
package main

import (
	"context"
	"log"

	"github.com/ppd324/dashvector-go/dashvector"
)


func main() {
	client := dashvector.New(
		"YOUR_API_KEY",
		"YOUR_ENDPOINT",
	)
	log.Println("dashvector client initialized")

	if client == nil {
		log.Fatal("failed to create dashvector client")
	}
	coll := client.Collection("quickstart")
	id := "9129292831698944"
	var r []Resource
	err = coll.Search(context.Background(), dashvector.SearchQuery{
		Id: &id,
	}, dashvector.WithIncludeVector(true), dashvector.WithTopK(10),
	).All(&r)
	
	
}

```