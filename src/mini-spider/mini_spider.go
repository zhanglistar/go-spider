package main

import (
	"deque"
	"fetcher"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"saver"
	"time"
	"encoding/json"
	"regexp"
	"runtime"
	"strings"

	"code.google.com/p/gcfg"
	"code.google.com/p/log4go"
	"github.com/docopt/docopt.go"
	"www.baidu.com/golang-lib/log"

	"working_queue"
)

type CrawlInfo struct {
	Depth       int
	Url         string
	ResultQueue *deque.Deque
	Fetcher     fetcher.Fetcher
	Wg          *sync.WaitGroup
	Cfg         *Config
}

type Config struct {
	Spider struct {
			   UrlListFile     string
			   OutputDirectory string
			   MaxDepth        int
			   CrawlInterval   int
			   CrawlTimeout    int
			   TargetUrl       string
			   ThreadCount     int
		   }
	Re     *regexp.Regexp
	Log    *log4go.Logger
}

func crawl_worker(args interface{}) {
	arg := args.(CrawlInfo)
	defer arg.Wg.Done()
	if arg.Depth < 0 {
		return
	}
	resultChan := make(chan fetcher.FetchResult)
	go func() {
		resultChan <- arg.Fetcher.Fetch(arg.Url)
	}()
	select {
	case <-time.After(time.Duration(arg.Cfg.Spider.CrawlTimeout) * time.Second):
		arg.Cfg.Log.Error("Crawl %s timeout", arg.Url)
		return
	case result := <-resultChan:
		if result.GetErr() != nil {
			return
		}
		for _, v := range (result.GetUrls()) {
			if arg.Depth < 1 {
				continue
			}
			arg.ResultQueue.Push(CrawlInfo{
				Depth: arg.Depth - 1,
				Url: v,
				ResultQueue: arg.ResultQueue,
				Fetcher: arg.Fetcher,
				Wg: arg.Wg,
				Cfg: arg.Cfg,
			})
		}

		if arg.Cfg.Re.Match([]byte(arg.Url)) {
			f := saver.FileSaver{}
			f.Save([]byte(result.GetBody()),
				arg.Cfg.Spider.OutputDirectory + "/" + strings.Replace(arg.Url, "/", "_", -1))
		}
	}
}

func Crawl(urls []string, cfg *Config, fetch fetcher.Fetcher) {
	depth := cfg.Spider.MaxDepth
	if depth <= 0 {
		return
	}

	var wg sync.WaitGroup
	toCrawlQueue := deque.NewDeque()
	for _, v := range (urls) {
		toCrawlQueue.Push(CrawlInfo{
			Depth: depth,
			Url: v,
			ResultQueue: toCrawlQueue,
			Fetcher: fetch,
			Wg: &wg,
			Cfg: cfg,
		})
	}

	request := make(chan working_queue.WorkRequest)
	dispatcher := working_queue.NewDispatcher(cfg.Spider.ThreadCount)
	dispatcher.Start(request)
	filter := make(map[string]bool)
	for {
		for toCrawlQueue.Len() > 0 {
			item := toCrawlQueue.Pop()
			itemValue := item.(CrawlInfo)
			if _, ok := filter[itemValue.Url]; ok {
				continue
			}
			wg.Add(1)
			filter[itemValue.Url] = true
			//			fmt.Println("Crawl:", itemValue.Url)
			cfg.Log.Info("Crawl:%s", itemValue.Url)
			request <- working_queue.WorkRequest{Args: itemValue, Handle: crawl_worker}
			time.Sleep(time.Duration(cfg.Spider.CrawlInterval) * time.Second)
		}
		wg.Wait()
		if toCrawlQueue.Len() <= 0 {
			break
		}
	}
}

func CheckConfig(cfg *Config, log *log4go.Logger) bool {
	var err error
	cfg.Re, err = regexp.Compile(cfg.Spider.TargetUrl)
	if err != nil {
		log.Error("TargetUrl not valid regrex!")
		return false
	}
	if cfg.Spider.CrawlInterval <= 0 {
		log.Error("CrawlInterval <= 0, should > 0!")
		return false
	}
	if cfg.Spider.CrawlTimeout <= 0 {
		log.Error("CrawlTimeout <= 0, should > 0!")
		return false
	}
	if cfg.Spider.MaxDepth <= 0 {
		log.Error("MaxDepth <= 0, should > 0!")
		return false
	}
	if cfg.Spider.ThreadCount <= 0 {
		log.Error("ThreadCount <= 0, should > 0!")
		return false
	}
	if len(cfg.Spider.OutputDirectory) == 0 {
		log.Error("OutputDirectory empty, should be not!")
		return false
	}
	if len(cfg.Spider.TargetUrl) == 0 {
		log.Error("TargetUrl empty, should be not!")
		return false
	}
	if len(cfg.Spider.UrlListFile) == 0 {
		log.Error("UrlListFile empty, should be not!")
		return false
	}
	return true
}

func ParseOptions(argv []string) map[string]interface{} {
	usage := `Usage: ./mini-spider [options]

  Options:
	-c CONF_FILE  set config directory
	-l LOG_DIR set log direcotry`
	opts, err := docopt.Parse(usage, argv, true, "1.0.0", false)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if opts["-c"] == nil || opts["-l"] == nil {
		fmt.Println(usage)
		os.Exit(-1)
	}
	return opts
}

func InitLog(logdir string) {
	err := log.Init("mini_spider", "INFO", logdir, false, "M", 2)
	if err != nil {
		fmt.Println("Log init failed")
		os.Exit(-1)
	}
	log.Logger.Info("Open log done.")
}

func InitConfig(confdir string) *Config {
	log.Logger.Info("Read conf %s", confdir)
	var cfg Config
	cfg.Log = &log.Logger
	err := gcfg.ReadFileInto(&cfg, confdir + "/spider.conf")
	if err != nil {
		log.Logger.Error("Read %s/spider.conf failed", confdir)
		os.Exit(-1)
	}
	return &cfg
}

func GetRootUrls(cfg *Config) []string {
	log.Logger.Info("UrllistFile file %s", cfg.Spider.UrlListFile)
	file, err := os.Open(cfg.Spider.UrlListFile)
	if err != nil {
		log.Logger.Error("Open UrlListFile %s failed", cfg.Spider.UrlListFile)
		os.Exit(-1)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Logger.Error("Read UrlListFile %s failed", cfg.Spider.UrlListFile)
		os.Exit(-1)
	}

	var json_value []interface{}
	err = json.Unmarshal([]byte(content), &json_value)
	if err != nil {
		log.Logger.Error("json.Unmarshal failed")
		os.Exit(-1)
	}

	urls := make([]string, len(json_value))
	for i, v := range (json_value) {
		urls[i] = v.(string)
	}
	return urls
}

func StartCrawl(urls []string, cfg *Config) {
	log.Logger.Info("Begin crawling, total num %d", len(urls))
	runtime.GOMAXPROCS(cfg.Spider.ThreadCount)
	fetch := fetcher.HttpFetcher{}
	Crawl(urls, cfg, fetch)
	log.Logger.Info("All done")
}

func main() {
	opts := ParseOptions(os.Args[1:])
	InitLog(opts["-l"].(string))
	cfg := InitConfig(opts["-c"].(string))
	if !CheckConfig(cfg, &log.Logger) {
		log.Logger.Error("Config invalid")
		os.Exit(-1)
	}
	urls := GetRootUrls(cfg)
	StartCrawl(urls, cfg)
}
