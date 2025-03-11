package sagemcomfast5359

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"strconv"

	_ "embed"
)

//go:embed user-agent.txt
var userAgent string

// Action represents the request payload structure.
type Action struct {
	Service    string            `json:"service"`
	Method     string            `json:"method"`
	Parameters map[string]string `json:"parameters"`
}

// ActionOptions allows customization of the DoAction request.
type ActionOptions struct {
	Client  *http.Client
	Host    string
	Action  Action
	Header  http.Header
	Cookies []*http.Cookie
	Debug   bool
	logger  *slog.Logger
}

// DoAction executes an HTTP request with the provided options and headers.
func DoAction(ctx context.Context, opts ActionOptions) (*http.Response, error) {
	if opts.Client == nil {
		opts.Client = http.DefaultClient
	}
	body, err := json.Marshal(opts.Action)
	if err != nil {
		return nil, fmt.Errorf("error marshalling json: %w", err)
	}
	url := fmt.Sprintf("http://%s/ws/NeMo/Intf/lan:getMIBs", opts.Host)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	// Apply the specified headers
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "nl,en;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-sah-ws-4-call+json; charset=utf-8")
	req.Header.Set("Origin", fmt.Sprintf("http://%s", opts.Host))
	req.Header.Set("Referer", fmt.Sprintf("http://%s/", opts.Host))
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Length", strconv.Itoa(len(body)))
	if opts.Header != nil {
		for key, values := range opts.Header {
			for _, value := range values {
				req.Header.Add(key, value)
			}
		}
	}
	for _, cookie := range opts.Cookies {
		req.AddCookie(cookie)
	}
	// debug request
	if opts.Debug {
		bytes, _ := httputil.DumpRequestOut(req, true)
		opts.logger.InfoContext(ctx, string(bytes))
	}
	resp, err := opts.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	// debug response
	if opts.Debug {
		bytes, _ := httputil.DumpResponse(resp, true)
		opts.logger.InfoContext(ctx, string(bytes))
	}
	return resp, nil
}
