/**
 * @Author: boyyang
 * @Date: 2022-03-27 10:36:33
 * @LastEditTime: 2022-06-09 16:15:56
 * @LastEditors: boyyang
 * @Description:
 * @FilePath: \blog\setupClient\client.go
 * [如果痛恨所处的黑暗，请你成为你想要的光。 --塞尔维亚的天空]
 */

package setupClient

import (
	"blog/global"
	"blog/setting"
	"fmt"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func SetupClient() *cos.Client {
	u, _ := url.Parse(setting.ClientUrl)
	b := &cos.BaseURL{BucketURL: u}
	global.Client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，
			// https://console.cloud.tencent.com/cam/capi
			SecretID: setting.ClientSecretID,
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，
			// https://console.cloud.tencent.com/cam/capi
			SecretKey: setting.ClientSecretKey,
		},
	})
	fmt.Println("💞🎈🎐对象储存初始化成功")
	return global.Client
}
