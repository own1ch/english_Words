package serverCommand

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func parseHtml(word string) [4]string {
	var text string
	var result [4]string
	doc, _ := goquery.NewDocument("http://nordmine.ru/dic/" + word)
	temp := doc
	temp.Find("div").Each(func(i int, s *goquery.Selection) {
		text, _ = s.Attr("class")
		if text == "col-lg-6" {
			if s.Find("h2").Text() == "Предложения" {
				text = s.Find("p").Contents().Text()
				result = parseText(text)
			}
		}
	})
	return result
}

func parseText(text string) [4]string {
	var temp []string
	var result [4]string
	var k int
	text = strings.Replace(text, "?", "?\n", len(text))
	text = strings.Replace(text, ".", ".\n", len(text))
	text = text[:strings.LastIndex(text, "\n")]
	temp = strings.Split(text, "\n")
	if len(temp) < 3 {
		k = 2
	} else {
		k = 4
	}
	fmt.Println(temp)
	for i := 0; i < k; i++ {
		if temp[i] != "" {
			result[i] = temp[i]
		}
	}
	return result
}
