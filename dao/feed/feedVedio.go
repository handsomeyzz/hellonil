package feed

import (
	"context"
	"fmt"
	"hellonil/config"
	"io"
)

var ctx = context.Background()

func UploadVideo(name string, reader io.Reader) (string, error) {
	nameMp4 := fmt.Sprintf("%s.mp4", name)
	_, err := CosClient.Object.Put(ctx, nameMp4, reader, nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return fmt.Sprintf("%s/%s", config.CosX().Addr, nameMp4), nil
}
