package fetcher

type Fetcher interface {
	Fetch(url string) FetchResult
}

type FetchResult struct {
	body string
	urls []string
	err error
}

func NewFetchResult() *FetchResult {
	return &FetchResult{body:"", urls:nil, err:nil}
}

func (fr *FetchResult) GetBody() string {
	return fr.body
}

func (fr *FetchResult) GetUrls() []string {
	return fr.urls
}

func (fr *FetchResult) GetErr() error {
	return fr.err
}
