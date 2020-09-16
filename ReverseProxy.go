package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
)

var addr = "127.0.0.1:2002"

func NewSingleHostReverseProxy(target *url.URL) *httputil.ReverseProxy {
	targetQuery := target.RawQuery
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}
	return &httputil.ReverseProxy{Director: director,ModifyResponse: modify}
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func modify(rsp *http.Response) error {
	if rsp.StatusCode != 200 {
		// 获取内容
	fmt.Println("1111111")
		oldPayload, err := ioutil.ReadAll(rsp.Body)
		if err == nil {
			// 追加内容
			newPayload := []byte("StatusCode error :" + string(oldPayload))
			rsp.Body = ioutil.NopCloser(bytes.NewBuffer(newPayload))
			rsp.ContentLength = int64(len(newPayload))
			rsp.Header.Set("Content-Length",strconv.FormatInt(int64(len(newPayload)),10))
		}
		return nil
	}
	return nil
}


func main() {
	//127.0.0.1:2002/xxx
	//127.0.0.1:2003/base/xxx
	rs1 := "http://127.0.0.1:2003/base"
	url1, err1 := url.Parse(rs1)
	if err1 != nil {
		log.Println(err1)
	}
	proxy := NewSingleHostReverseProxy(url1)
	log.Println("Starting httpserver at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}