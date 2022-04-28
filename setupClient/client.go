/**
 * @Author: boyyang
 * @Date: 2022-03-27 10:36:33
 * @LastEditTime: 2022-04-09 23:37:57
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\client\client.go
 * [如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package client

import (
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

var Client *cos.Client

func SetupClient() *cos.Client {
	u, _ := url.Parse("https://7072-prod-2g489qm8208c3cfd-1301921121.cos.ap-shanghai.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	Client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// appid 1301921121
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: "AKID85lAJBm6AyVljQr0sSmpfKn0vZw3sQhh",
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: "kDPLzNUKtBk23nGRM0MREkzxd3C4f03p",
		},
	})
	// opt := &cos.BucketPutCORSOptions{
	// 	Rules: []cos.BucketCORSRule{
	// 		{
	// 			AllowedOrigins: []string{"*"},
	// 			AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "UPDATE"},
	// 			AllowedHeaders: []string{"Authorization", "Content-Length", "X-CSRF-Token", "Token", "session"},
	// 			MaxAgeSeconds:  172800,
	// 			ExposeHeaders:  []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
	// 		},
	// 	},
	// }
	// _, err := Client.Bucket.PutCORS(context.Background(), opt)
	// if err != nil {
	// 	panic(err)
	// }
	return Client
}
