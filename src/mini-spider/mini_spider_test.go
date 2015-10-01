package main
import (
	"testing"
	"www.baidu.com/golang-lib/log"
)

func Test_CheckConfig_should_success_when_everythin_ok(t *testing.T) {
	var cfg Config
	cfg.Spider.UrlListFile = "test"
	cfg.Spider.TargetUrl= ".html"
	cfg.Spider.OutputDirectory= "output"
	cfg.Spider.CrawlInterval = 1
	cfg.Spider.CrawlTimeout = 10
	cfg.Spider.MaxDepth = 2
	cfg.Spider.ThreadCount = 8

	err := log.Init("mini_spider", "INFO", "test", false, "M", 2)
	if err != nil {
		t.Error("Log init failed")
	}

	err1 := CheckConfig(&cfg, &log.Logger)
	if !err1 {
		t.Error("CheckConfig failed")
	}
}

func Test_CheckConfig_should_fail_when_somthing_not_ok(t *testing.T) {
	var cfg Config
	cfg.Spider.UrlListFile = ""
	cfg.Spider.TargetUrl= ".html"
	cfg.Spider.OutputDirectory= "output"
	cfg.Spider.CrawlInterval = 1
	cfg.Spider.CrawlTimeout = 10
	cfg.Spider.MaxDepth = 2
	cfg.Spider.ThreadCount = 8

//	err := log.Init("test", "INFO", "./log", false, "M", 2)
//	if err != nil {
//		fmt.Println(err)
//		t.Error("Log init failed")
//	}

	err2 := CheckConfig(&cfg, &log.Logger)
	if err2 != false {
		t.Error("CheckConfig failed")
	}
}

