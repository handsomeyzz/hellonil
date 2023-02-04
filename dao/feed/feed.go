package feed

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
	"hellonil/setting"
	"io"
	"net/http"
	"net/url"
	"os"
)

var cosClient *cos.Client

var ctx = context.Background()

// 新建一个cos客户端
func Init(cfg *setting.CosConfig) error {
	u, err := url.Parse(cfg.Host)
	if err != nil {
		return err
	}
	b := &cos.BaseURL{BucketURL: u}
	cosClient = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cfg.SecretId,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
			SecretKey: cfg.SecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
		},
	})
	return nil
}

// 根据视频名字或者图片名字上传视频，targetName:上传服务器时的目标名字 path:需要上传的文件的路径
// 返回生成之后的url和错误
func Upload(targetName string, path string) (url string, err error) {
	f, err := os.Open(path)
	defer f.Close() //延迟关闭文件
	if err != nil {
		//日志
		return
	}
	var r io.Reader
	r = f //转化成io.Reader
	_, err = cosClient.Object.Put(ctx, targetName, r, nil)
	if err != nil {
		//日志
		return "", nil
	}
	//删除本地文件
	u, ok := viper.Get("cos.host").(string)
	if !ok {
		fmt.Println("viper类型断言出错了！")
		return "", nil
	}
	url = fmt.Sprintf("%s/%s", u, targetName)
	err = os.Remove(path)
	if err != nil {
		return "", err
	}

	return url, nil
}
