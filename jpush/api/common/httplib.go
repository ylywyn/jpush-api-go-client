package common

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

//  Get returns *HttpRequest with GET method.
func Get(url string) *HttpRequest {
	var req http.Request
	req.Method = "GET"
	req.Header = http.Header{}
	return &HttpRequest{url, &req, map[string]string{}, 60 * time.Second, 60 * time.Second, nil, nil, nil}
}

// Post returns *HttpRequest with POST method.
func Post(url string) *HttpRequest {
	var req http.Request
	req.Method = "POST"
	req.Header = http.Header{}
	return &HttpRequest{url, &req, map[string]string{}, 60 * time.Second, 60 * time.Second, nil, nil, nil}
}

// HttpRequest provides more useful methods for requesting one url than http.Request.
type HttpRequest struct {
	url              string
	req              *http.Request
	params           map[string]string
	connectTimeout   time.Duration
	readWriteTimeout time.Duration
	tlsClientConfig  *tls.Config
	proxy            func(*http.Request) (*url.URL, error)
	transport        http.RoundTripper
}

// SetTimeout sets connect time out and read-write time out for Request.
func (b *HttpRequest) SetTimeout(connectTimeout, readWriteTimeout time.Duration) *HttpRequest {
	b.connectTimeout = connectTimeout
	b.readWriteTimeout = readWriteTimeout
	return b
}

// SetTLSClientConfig sets tls connection configurations if visiting https url.
func (b *HttpRequest) SetTLSClientConfig(config *tls.Config) *HttpRequest {
	b.tlsClientConfig = config
	return b
}

// Header add header item string in request.
func (b *HttpRequest) Header(key, value string) *HttpRequest {
	b.req.Header.Set(key, value)
	return b
}

// Set the protocol version for incoming requests.
// Client requests always use HTTP/1.1.
func (b *HttpRequest) SetProtocolVersion(vers string) *HttpRequest {
	if len(vers) == 0 {
		vers = "HTTP/1.1"
	}

	major, minor, ok := http.ParseHTTPVersion(vers)
	if ok {
		b.req.Proto = vers
		b.req.ProtoMajor = major
		b.req.ProtoMinor = minor
	}

	return b
}

// SetCookie add cookie into request.
func (b *HttpRequest) SetCookie(cookie *http.Cookie) *HttpRequest {
	b.req.Header.Add("Cookie", cookie.String())
	return b
}

// Set transport to
func (b *HttpRequest) SetTransport(transport http.RoundTripper) *HttpRequest {
	b.transport = transport
	return b
}

// Set http proxy
// example:
//
//	func(req *http.Request) (*url.URL, error) {
// 		u, _ := url.ParseRequestURI("http://127.0.0.1:8118")
// 		return u, nil
// 	}
func (b *HttpRequest) SetProxy(proxy func(*http.Request) (*url.URL, error)) *HttpRequest {
	b.proxy = proxy
	return b
}

// Param adds query param in to request.
// params build query string as ?key1=value1&key2=value2...
func (b *HttpRequest) Param(key, value string) *HttpRequest {
	b.params[key] = value
	return b
}

// Body adds request raw body.
// it supports string and []byte.
func (b *HttpRequest) Body(data interface{}) *HttpRequest {
	switch t := data.(type) {
	case string:
		bf := bytes.NewBufferString(t)
		b.req.Body = ioutil.NopCloser(bf)
		b.req.ContentLength = int64(len(t))
	case []byte:
		bf := bytes.NewBuffer(t)
		b.req.Body = ioutil.NopCloser(bf)
		b.req.ContentLength = int64(len(t))
	}
	return b
}

func (b *HttpRequest) getResponse() (*http.Response, error) {
	var paramBody string
	if len(b.params) > 0 {
		var buf bytes.Buffer
		for k, v := range b.params {
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
			buf.WriteByte('&')
		}
		paramBody = buf.String()
		paramBody = paramBody[0 : len(paramBody)-1]
	}

	if b.req.Method == "GET" && len(paramBody) > 0 {
		if strings.Index(b.url, "?") != -1 {
			b.url += "&" + paramBody
		} else {
			b.url = b.url + "?" + paramBody
		}
	} else if b.req.Method == "POST" && b.req.Body == nil && len(paramBody) > 0 {
		b.Header("Content-Type", "application/x-www-form-urlencoded")
		b.Body(paramBody)
	}

	url, err := url.Parse(b.url)
	if url.Scheme == "" {
		b.url = "http://" + b.url
		url, err = url.Parse(b.url)
	}
	if err != nil {
		return nil, err
	}

	b.req.URL = url
	trans := b.transport

	if trans == nil {
		// create default transport
		trans = &http.Transport{
			TLSClientConfig: b.tlsClientConfig,
			Proxy:           b.proxy,
			Dial:            TimeoutDialer(b.connectTimeout, b.readWriteTimeout),
		}
	} else {
		// if b.transport is *http.Transport then set the settings.
		if t, ok := trans.(*http.Transport); ok {
			if t.TLSClientConfig == nil {
				t.TLSClientConfig = b.tlsClientConfig
			}
			if t.Proxy == nil {
				t.Proxy = b.proxy
			}
			if t.Dial == nil {
				t.Dial = TimeoutDialer(b.connectTimeout, b.readWriteTimeout)
			}
		}
	}

	client := &http.Client{
		Transport: trans,
	}

	resp, err := client.Do(b.req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// String returns the body string in response.
// it calls Response inner.
func (b *HttpRequest) String() (string, error) {
	data, err := b.Bytes()
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Bytes returns the body []byte in response.
// it calls Response inner.
func (b *HttpRequest) Bytes() ([]byte, error) {
	resp, err := b.getResponse()
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ToFile saves the body data in response to one file.
// it calls Response inner.
func (b *HttpRequest) ToFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	resp, err := b.getResponse()
	if err != nil {
		return err
	}
	if resp.Body == nil {
		return nil
	}
	defer resp.Body.Close()
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// ToJson returns the map that marshals from the body bytes as json in response .
// it calls Response inner.
func (b *HttpRequest) ToJson(v interface{}) error {
	data, err := b.Bytes()
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}

// ToXml returns the map that marshals from the body bytes as xml in response .
// it calls Response inner.
func (b *HttpRequest) ToXML(v interface{}) error {
	data, err := b.Bytes()
	if err != nil {
		return err
	}
	err = xml.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}

// Response executes request client gets response mannually.
func (b *HttpRequest) Response() (*http.Response, error) {
	return b.getResponse()
}

// TimeoutDialer returns functions of connection dialer with timeout settings for http.Transport Dial field.
func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
}
