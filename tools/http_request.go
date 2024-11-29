package tools

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

type ERROR_TYPE string

const (
	TimeOut ERROR_TYPE = "time out"
	BadUrl  ERROR_TYPE = "invalid url"
	LiveUrl ERROR_TYPE = "valid url"
)

type ResponseResult struct {
	StatusCode   int
	ResponeBody  string
	ResponseTime int64
	Flag         ERROR_TYPE
}

// 60s =60000ms
var Max_TimeOut int64 = 60 * 1000

func Post(url string, data []byte) ResponseResult {
	rr := ResponseResult{StatusCode: 200, ResponeBody: "", ResponseTime: Max_TimeOut, Flag: LiveUrl}
	tp := newTransport()
	client := &http.Client{Transport: tp}
	contentType := "application/json; charset:utf-8"
	dataByte := bytes.NewReader(data)
	resp, err := client.Post(url, contentType, dataByte)
	if err != nil {
		rr.ResponeBody = fmt.Sprintf("request error: %s", err.Error())
		rr.StatusCode = 500
		rr.Flag = BadUrl
		return rr
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		rr.ResponeBody = fmt.Sprintf("pase body error: %s", err.Error())
		rr.StatusCode = 504
		rr.Flag = TimeOut
		return rr
	}
	rr.ResponeBody = string(body)
	rr.StatusCode = resp.StatusCode
	rr.ResponseTime = tp.Duration().Milliseconds()
	return rr
}
func GetWithHeader(url string, header map[string]string) ResponseResult {

	rr := ResponseResult{StatusCode: 200, ResponeBody: "", ResponseTime: Max_TimeOut, Flag: LiveUrl}
	tp := newTransport()
	client := &http.Client{Transport: tp}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		rr.StatusCode = 500
		return rr
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		rr.ResponeBody = fmt.Sprintf("request error: %s", err.Error())
		rr.StatusCode = 500
		rr.Flag = BadUrl
		return rr
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		rr.ResponeBody = fmt.Sprintf("pase body error: %s", err.Error())
		rr.StatusCode = 504
		rr.Flag = TimeOut
		return rr
	}
	rr.ResponeBody = string(body)
	rr.StatusCode = resp.StatusCode
	rr.ResponseTime = tp.Duration().Milliseconds()
	return rr
}
func Get(url string) ResponseResult {
	rr := ResponseResult{StatusCode: 200, ResponeBody: "", ResponseTime: Max_TimeOut, Flag: LiveUrl}
	tp := newTransport()
	client := &http.Client{Transport: tp}
	resp, err := client.Get(url)
	if err != nil {
		rr.ResponeBody = fmt.Sprintf("request error: %s", err.Error())
		rr.StatusCode = 500
		rr.Flag = BadUrl
		return rr
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		rr.ResponeBody = fmt.Sprintf("pase body error: %s", err.Error())
		rr.StatusCode = 504
		rr.Flag = TimeOut
		return rr
	}
	rr.ResponeBody = string(body)
	rr.StatusCode = resp.StatusCode
	rr.ResponseTime = tp.Duration().Milliseconds()
	return rr
}

type customTransport struct {
	rtp       http.RoundTripper
	dialer    *net.Dialer
	connStart time.Time
	connEnd   time.Time
	reqStart  time.Time
	reqEnd    time.Time
}

func newTransport() *customTransport {

	tr := &customTransport{
		dialer: &net.Dialer{
			Timeout:   120 * time.Second,
			KeepAlive: 120 * time.Second,
		},
	}
	tr.rtp = &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		Dial:                tr.dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	return tr
}

func (tr *customTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	tr.reqStart = time.Now()
	resp, err := tr.rtp.RoundTrip(r)
	tr.reqEnd = time.Now()
	return resp, err
}

func (tr *customTransport) dial(network, addr string) (net.Conn, error) {
	tr.connStart = time.Now()
	cn, err := tr.dialer.Dial(network, addr)
	tr.connEnd = time.Now()
	return cn, err
}

func (tr *customTransport) ReqDuration() time.Duration {
	return tr.Duration() - tr.ConnDuration()
}

func (tr *customTransport) ConnDuration() time.Duration {
	return tr.connEnd.Sub(tr.connStart)
}

func (tr *customTransport) Duration() time.Duration {
	return tr.reqEnd.Sub(tr.reqStart)
}

func Warning(warningUrl, args string, messageId int64) {
	requestUrl := fmt.Sprintf("%s?message_id=%d&notice_args=%s", warningUrl, messageId, args)
	fmt.Printf(">>>>>>>>>>>warning url:%s", requestUrl)
	response := Get(requestUrl)
	fmt.Printf("send message result:%s", response.ResponeBody)
}
