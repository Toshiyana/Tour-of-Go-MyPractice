// 理解できていない
// reference:
// * https://gist.github.com/harryhare/6a4979aa7f8b90db6cbc74400d0beb49
// * https://www.ninton.co.jp/archives/5486

// Web Crawlerの演習問題
// Web Crawler: サイトの情報を収集するプログラム

// 同じURLを2度とることなく並列してURLを取ってくる、Crawl関数の実装
// 実装のポイント
// * mapを使った重複防止
// * goroutineを使った並列処理

// Web Crawlerの演習問題だが、実際にクロールはせず、fakeFetcherというものを定義する。

package main

import (
	"fmt"
	"sync"
)

type mutexCache struct {
	mux   sync.Mutex
	store map[string]bool
}

func (cache *mutexCache) setVisited(name string) bool {
	cache.mux.Lock()
	defer cache.mux.Unlock()

	if cache.store[name] {
		return true
	}

	cache.store[name] = true

	return false
}

var cacheInstance = mutexCache{store: make(map[string]bool)}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.

	// interface型のメソッドを定義
	Fetch(url string) (body string, urls []string, err error)
}

func crawlInner(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth <= 0 {
		return
	}

	if cacheInstance.setVisited(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go crawlInner(u, depth-1, fetcher, wg)
	}
	return
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	waitGroup := &sync.WaitGroup{}

	waitGroup.Add(1)

	go crawlInner(url, depth, fetcher, waitGroup)

	waitGroup.Wait()
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)

}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

// fakeFetcher型のレシーバ引数を持つFetchメソッド(interface型Fetcherの要素)
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// 実際にクロールはせずにfake。
// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language", // body of URL
		[]string{ // slice of URLs found on that page
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
