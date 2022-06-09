/**
 * @Author: boyyang
 * @Date: 2022-06-06 09:49:04
 * @LastEditTime: 2022-06-09 16:32:07
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\colly\colly.go
 * @[如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package creaw

import (
	"blog/global"
	"blog/models"
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

var Wg sync.WaitGroup

func CollyInit() {
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
			fmt.Println(link, name)
			Wg.Add(1)
			go download("https:"+link, name+".jpg")
		}
	})

	c.Visit("https://m.woyaogexing.com/shouji/index_4.html")
	//https://m.woyaogexing.com/shouji/hot/
	//https://m.woyaogexing.com/shouji/new/
	//https://m.woyaogexing.com/tupian/qinglv/
	//https://m.woyaogexing.com/tupian/weimei/
	//https://img2.woyaogexing.com/2022/06/05/aa0d3990d5b5a6ed!400x400.jpg
}

func download(url string, name string) {
	defer Wg.Done()
	res, err := http.Get(url)
	if err == nil {
		content, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		if err == nil {
			ioutil.WriteFile(fmt.Sprintf("assets/%s", name), content, 0644)
			path := fmt.Sprintf("/%d/%s/%s", 1, "images", name)
			global.Client.Object.Put(context.Background(), path, bytes.NewReader(content), nil)
			upload := models.Upload{
				Url:      path,
				FileName: name,
				UserID:   1,
			}
			err := global.DB.Create(&upload)
			if err == nil {
				fmt.Println("💞🎈图片上传成功", name, path)
			}
		}
	}
}
