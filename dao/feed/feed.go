package feed

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"hellonil/config"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Cos struct {
	*cos.Client
}

var ctx = context.Background()

// 新建一个cos客户端
func NewCosClient() (*Cos, error) {
	u, err := url.Parse(config.CosUrl)
	if err != nil {
		return nil, err
	}
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.CosSecretID,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
			SecretKey: config.CosSecretkey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
		},
	})

	return &Cos{client}, nil
}

// 根据视频名字或者图片名字上传视频，targetName:上传服务器时的目标名字 path:需要上传的文件的路径
// 返回生成之后的url和错误
func (cos *Cos) Upload(targetName string, path string) (Url string, err error) {
	f, err := os.Open(path)
	defer f.Close() //延迟关闭文件
	if err != nil {
		//日志
		return "", err
	}
	var r io.Reader
	r = f //转化成io.Reader
	_, err = cos.Object.Put(ctx, targetName, r, nil)
	if err != nil {
		//日志
		return "", err
	}
	//删除本地文件
	err = os.Remove(path)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return fmt.Sprintf("%s/%s", config.CosUrl, targetName), nil
}
