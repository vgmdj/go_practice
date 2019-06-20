package proxy

import "testing"

func TestProxy_Request(t *testing.T) {
	proxy := new(Proxy)
	proxy.Request()
}
