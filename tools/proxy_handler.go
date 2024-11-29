package tools

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/jerryTJ/controller/internal/app"
)

type ProxyHandler struct {
	TargetUrl   string
	ServerInfos map[string]app.ServerInfo
}

// ServeHTTP方法，绑定DefaultHandler
func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 解析目标URL
	targetURL, err := url.Parse(ph.TargetUrl)
	if err != nil {
		http.Error(w, "Invalid target URL", http.StatusBadRequest)
		return
	}

	// 创建一个新的请求，复制原始请求的数据
	proxyReq, err := http.NewRequest(r.Method, targetURL.ResolveReference(r.URL).String(), r.Body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	domain := strings.Split(r.Host, ":")[0]
	serverInfo := ph.ServerInfos[domain]
	// 复制请求头
	proxyReq.Header = r.Header
	proxyReq.Header.Set("x-color", serverInfo.Color)
	proxyReq.Header.Set("x-chain", serverInfo.ChainID)
	proxyReq.Header.Set("x-version", serverInfo.Version)

	// 发送代理请求
	resp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		http.Error(w, "Failed to reach target server", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// 将响应头写回给客户端
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// 设置响应状态码
	w.WriteHeader(resp.StatusCode)

	// 将响应体写回给客户端
	if _, err := io.Copy(w, resp.Body); err != nil {
		log.Println("Failed to copy response body:", err)
	}
}
