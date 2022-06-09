/**
 * @Author: boyyang
 * @Date: 2022-06-06 09:49:04
 * @LastEditTime: 2022-06-09 08:36:58
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\colly\colly.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

// "fmt"
// "io/ioutil"
// "net/http"
// "strconv"
// "time"
// "github.com/gocolly/colly"

var wg sync.WaitGroup

func collyInit() {
	c := colly.NewCollector()

	c.OnHTML("div[class='m-img-wrap']", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		if strings.HasSuffix(link, ".html") {
			// fmt.Println(link)
			c.Visit(e.Response.Request.AbsoluteURL(link))
		}
		// wg.Add(1)
		// go download(url, name, num)
	})

	// g-page-content g-page-main f-pd-2
	c.OnHTML("div[class='g-page-content g-page-main f-pd-2']", func(e *colly.HTMLElement) {
		name := e.DOM.Find("h1").Text()
		link, _ := e.DOM.Children().Find("a").Attr("href")
		if strings.HasSuffix(link, ".jpg") {
			fmt.Println(link)
			wg.Add(1)
			go download("https:"+link, name+".jpg")
		}
	})

	c.Visit("https://m.woyaogexing.com/shouji/index_5.html")
	//https://m.woyaogexing.com/shouji/hot/
	//https://m.woyaogexing.com/shouji/new/
	//https://m.woyaogexing.com/tupian/qinglv/
	//https://m.woyaogexing.com/tupian/weimei/
	//https://img2.woyaogexing.com/2022/06/05/aa0d3990d5b5a6ed!400x400.jpg
}

func download(url string, name string) {
	defer wg.Done()
	fmt.Println(name, url)
	res, err := http.Get(url)
	if err == nil {
		content, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		if err == nil {
			ioutil.WriteFile(fmt.Sprintf("assets/%s", name), content, 0644)
		}
	}
}

func main() {
	// collyInit()
	// wg.Wait()
	var test int = 15
	var age string
	if test > 10 {
		age = "大于10"
	} else {
		age = "小于10"
	}
	fmt.Println(&test)
	fmt.Println(age)

}
