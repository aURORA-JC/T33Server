package proxy

import (
	"github.com/siddontang/go-log/log"
	"io"
	"net/http"
	"strings"
)

func T33ProxyService(w http.ResponseWriter, r *http.Request) {
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	resp.Header.Add("Proxy-Server", "T33Server 0.1.1-dev [proxy]")
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)

	log.Infoln("[proxy]", r.Method, r.RequestURI, strings.Split(r.RemoteAddr, ":")[0], r.UserAgent())
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
