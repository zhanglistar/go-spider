package parser

import (
	"testing"
	"fmt"
)

const URL string = "http://www.baidu.com"
func Test_GetAllUrlFromPage_http_string(t *testing.T) {
	content := "href =\"http://www.baidu.com\"";
	result := ParseUrls(URL, []byte(content))
	if len(result) != 1 {
		t.Error("fail to get url from page!")
	}
	fmt.Println(result)
}

func Test_GetAllUrlFromPage_relative_path_string(t *testing.T) {
	content := "href =\"t2.baidu.com\"";
	result := ParseUrls(URL, []byte(content))
	if len(result) != 1 {
		t.Error("fail to get url from page!")
	}
	fmt.Println(result)
}

func Test_GetAllUrlFromPage_no_string(t *testing.T) {
	content := "href =\"\"";
	result := ParseUrls(URL, []byte(content))
	if len(result) != 0 {
		t.Error("fail to get url from page!")
	}
	fmt.Println(result)
}

func Test_GetAllUrlFromPage_no_string_in_single_quote(t *testing.T) {
	content := "href =''";
	result := ParseUrls(URL, []byte(content))
	if len(result) != 0 {
		t.Error("fail to get url from page!")
	}
	fmt.Println(result)
}

func Test_GetAllUrlFromPage_double_slash(t *testing.T) {
	content := "href =\"//test.abc\"";
	result := ParseUrls(URL, []byte(content))
	if result[0] != "http://test.abc" {
		fmt.Println(result[0])
		t.Error("fail to get url from page!")
	}
	fmt.Println(result)
}

func Test_GetAllUrlFromPage_singal_slash(t *testing.T) {
	content := "href =\"/abc\"";
	result := ParseUrls(URL, []byte(content))
	if len(result) != 1 {
		t.Error("fail to get url from page!")
	}
	if result[0] != URL + "/abc" {
		t.Error("fail to get url from page!")
	}
	fmt.Println(result)
}

func Test_GetAllUrlFromPage_has_http_prefix(t *testing.T) {
	content := "href =\"http://a.b.c\"";
	result := ParseUrls(URL, []byte(content))
	if len(result) == 0 {
		t.Error("fail to get url from page!")
	}
	if result[0] != "http://a.b.c" {
		t.Error("fail to get url from page!")
	}
	fmt.Println(result)
}

func Test_GetAllUrlFromPage_only_javascript(t *testing.T) {
	content := "href =\"javascript:;\"";
	result := ParseUrls(URL, []byte(content))
	if len(result) != 0 {
		t.Error("fail to get url from page!")
	}
	fmt.Println(result)
}
