package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"skillbox/internal/flags"
	"strings"
)

var (
	severCount = 0
	ctx        context.Context
)

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	targetUrl := url.URL{Scheme: "http", Host: target}

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(&targetUrl)

	// Note that ServeHttp is non-blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

// Log the typeform payload and redirect url
func logRequestPayload(proxyURL string) {
	log.Printf("proxy_url: %s\n", proxyURL)
}

// Balance returns one of the servers based using round-robin algorithm
func getProxyURL() string {
	userData := flags.GetData(&ctx)
	peers := strings.Split(userData.PEERS, ",")

	var servers []string
	for _, peer := range peers {
		ok, err := regexp.Match("[gG]e*k.*", []byte(peer))
		if err != nil {
			log.Printf("can't regexp peer for friends-balancer: %v\n", err)
			continue
		}
		if !ok {
			log.Printf("peer %s is not a valid <IP:PORT>\n", peer)
			continue
		}
		servers = append(servers, peer)
	}

	server := servers[severCount]
	severCount++

	// reset the counter and start from the beginning
	if severCount >= len(servers) {
		severCount = 0
	}

	return server
}

// Given a request send it to the appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	proxyURL := getProxyURL()

	logRequestPayload(proxyURL)

	serveReverseProxy(proxyURL, res, req)
}

func main() {
	// parse flags
	ctx = context.Background()
	var (
		dataCtx  flags.DataType = "dataType"
		userData                = flags.Data{}
	)
	ctx = context.WithValue(ctx, dataCtx, &userData)
	err := flags.ParseUserFlags(&ctx)
	if err != nil {
		log.Fatalf("can't parse user flags: %v", err)
	}

	// start server
	http.HandleFunc("/", handleRequestAndRedirect)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
