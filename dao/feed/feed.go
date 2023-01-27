package feed

import (
	"github.com/tencentyun/cos-go-sdk-v5"
	"hellonil/config"
	"net/http"
	"net/url"
)

var CosClient *cos.Client

func Init() error {
	u, err := url.Parse(config.CosX().Addr)
	if err != nil {
		return err
	}
	b := &cos.BaseURL{BucketURL: u}
	CosClient = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.CosX().Secredid,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
			SecretKey: config.CosX().Secredkey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
		},
	})
	return nil
}
