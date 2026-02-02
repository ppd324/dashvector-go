/**
 * @Author: peidong pei
 * @Description: 描述
 * @File: client
 * @Date: 2026/1/16 15:11
 * @Software : GoLand
 */

package dashvector

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/http/httptrace"
	"time"

	"github.com/ppd324/dashvector-go/schema"
)

const (
	defaultReadTimeout  = 5 * time.Second
	defaultWriteTimeout = 5 * time.Second
	defaultDialTimeout  = 10 * time.Second
	defaultConnPoolSize = 10
)

type Client struct {
	endpoint     string
	apiKey       string
	apiVersion   string
	readTimeout  time.Duration
	writeTimeout time.Duration
	dialTimeout  time.Duration
	connPoolSize int
	*http.Client
	httpTrace bool
	enableLog bool
}

func New(endpoint string, apiKey string, opts ...Option) *Client {
	cli := &Client{
		endpoint:     endpoint,
		apiKey:       apiKey,
		readTimeout:  defaultReadTimeout,
		writeTimeout: defaultWriteTimeout,
		dialTimeout:  defaultDialTimeout,
		connPoolSize: defaultConnPoolSize,
		enableLog:    false,
		apiVersion:   "v1",
	}
	for _, opt := range opts {
		opt(cli)
	}
	cli.Client = &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: cli.dialTimeout,
			}).Dial,
			MaxIdleConnsPerHost: cli.connPoolSize,
			IdleConnTimeout:     cli.dialTimeout,
		},
		Timeout: cli.readTimeout + cli.writeTimeout,
	}
	return cli
}

func (c *Client) Collection(name string) Collection {
	return newDvCollection(c, name)
}

func (c *Client) Close() error {
	c.Client.CloseIdleConnections()
	return nil
}

func (c *Client) DropCollection(name string) error {
	return nil
}

func (c *Client) CreateCollection(name string, schema *schema.Schema) error {
	return nil

}

func (c *Client) ListCollections(ctx context.Context) ([]string, error) {
	colls, err := call[any, []string](c, ctx, "GET", "collections", nil)
	if err != nil {
		return nil, err
	}
	if colls == nil {
		return []string{}, nil
	}
	return *colls, nil
}

func call[I, O any](client *Client, ctx context.Context, method, path string, body *I) (*O, error) {
	var reader io.Reader
	hasBody := false
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		if client.enableLog {
			slog.Info("request body", "body", string(bodyBytes))
		}
		hasBody = true
		reader = bytes.NewReader(bodyBytes)
	}
	callUrl := fmt.Sprintf("https://%s/%s/%s", client.endpoint, client.apiVersion, path)
	req, err := http.NewRequest(method, callUrl, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("dashvector-auth-token", client.apiKey)
	if hasBody {
		req.Header.Set("Content-Type", "application/json")
	}
	if client.httpTrace {
		var start, dnsStart, connStart, tlsStart time.Time
		trace := &httptrace.ClientTrace{
			DNSStart:             func(info httptrace.DNSStartInfo) { dnsStart = time.Now() },
			DNSDone:              func(info httptrace.DNSDoneInfo) { fmt.Println("DNS耗时:", time.Since(dnsStart)) },
			ConnectStart:         func(_, _ string) { connStart = time.Now() },
			ConnectDone:          func(_, _ string, err error) { fmt.Println("TCP连接耗时:", time.Since(connStart)) },
			TLSHandshakeStart:    func() { tlsStart = time.Now() },
			TLSHandshakeDone:     func(state tls.ConnectionState, err error) { fmt.Println("TLS耗时:", time.Since(tlsStart)) },
			GotFirstResponseByte: func() { fmt.Println("TTFB耗时:", time.Since(start)) },
		}
		ctx = httptrace.WithClientTrace(ctx, trace)
	}

	req = req.WithContext(ctx)
	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("call api  error", "error", err)
		return nil, err
	}
	if client.httpTrace {
		slog.Info("call api", "url", callUrl, "method", method, "status", resp.StatusCode, "took", time.Since(start))
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if resp.StatusCode != http.StatusOK {
		slog.Error("call api error", "statusCode", resp.StatusCode)
		return nil, fmt.Errorf("call api error,statusCode:%d", resp.StatusCode)
	}
	recv := new(commonResponse[O])
	err = json.NewDecoder(resp.Body).Decode(&recv)
	if err != nil {
		slog.Error("decode response error", "error", err)
		return nil, err
	}
	if recv.Code != 0 {
		slog.Error("call api error", "code", recv.Code, "message", recv.Message)
		return nil, fmt.Errorf("call api error,code:%d,message:%s", recv.Code, recv.Message)
	}
	if client.enableLog {
		slog.Info("call api success", "requestId", recv.RequestId, "output", recv.Output)
	}
	return &recv.Output, nil
}
