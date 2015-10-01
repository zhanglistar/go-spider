package parser

import (
	"regexp"
	"strings"
)

func ParseUrls(url string, content []byte) (urls []string) {
	re := regexp.MustCompile("href\\s*=\\s*['\"]?\\s*([^'\"\\s]+)\\s*['\"]?")
	allsubmatch := re.FindAllSubmatch(content, -1)
	for _, v2 := range allsubmatch {
		for k, v := range v2 {
			if k == 0 {
				// k == 0 是表示匹配的全部元素
				continue
			}
			v := string(v)
			if strings.HasPrefix(v, "javascript") {
				continue
			}
			s := strings.TrimRight(v, "/ \\")
			if len(s) == 0 {
				continue
			}
			if strings.HasPrefix(s, "http") {
				urls = append(urls, s)
				continue
			}
			if strings.HasPrefix(s, "//") {
				if strings.HasPrefix(url, "https") {
					s = "https:" + s
				} else {
					s = "http:" + s
				}
				urls = append(urls, s)
				continue
			}
			if strings.HasPrefix(s, "/") {
				s = url + s
			} else {
				s = url + "/" + s
			}
			urls = append(urls, s)
		}
	}
	return
}

