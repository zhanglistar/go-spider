package fetcher
import (
	"net/http"
	"io/ioutil"
	"parser"
)

type HttpFetcher struct {}

func (_ HttpFetcher) Fetch(url string) FetchResult {
	response, err := http.Get("http://www.baidu.com/")
	if err != nil {
		return FetchResult{"", nil, err}
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return FetchResult{"", nil, err}
		}
		urls := parser.ParseUrls(url, contents)
		return FetchResult{string(contents), urls, nil}
	}
}
