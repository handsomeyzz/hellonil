package feed

import (
	"bytes"
	"fmt"
	cmdchain "github.com/rainu/go-command-chain"
	"strings"
)

// 获取当前视频的编码格式,时长，和程序运行时错误
func encodeAndTime(path string) (encode string, duration string, err error) {
	output := &bytes.Buffer{}
	cmd := "D:\\programming\\ffmpeg-2023-01-22-git-9d5e66942c-full_build\\ffmpeg-2023-01-22-git-9d5e66942c-full_build\\bin\\ffprobe.exe" //上线linux需要修改
	//通过管道拼接
	err = cmdchain.Builder().
		Join(cmd, "-show_streams", path).
		Join("grep", "-E", "'codec_name|duration='").
		Finalize().WithOutput(output).Run()
	if err != nil {
		panic(err)
		return "", "", err
	}
	en := strings.Split(output.String(), "\n")
	encodeVideo, dn := en[2], en[3]
	encode = strings.Split(encodeVideo, "=")[1]
	duration = strings.Split(dn, "=")[1]
	return encode, duration, nil
}

// 转码任何格式转为h.264并降低码率
func lowerBitRate(path string) (h264 string, err error) {
	//对视频文件路径进行预处理
	p := strings.Split(path, ".")[0]
	dir := strings.Split(p, "\\")                   //linux下换成/
	dirreal := strings.Join(dir[:len(dir)-1], "\\") //linux下换成/
	name := fmt.Sprintf("%s\\%s", dirreal, dir[len(dir)-1]+"_h264_lowerbit.mp4")

	output := &bytes.Buffer{}
	cmd := "D:\\programming\\ffmpeg-2023-01-22-git-9d5e66942c-full_build\\ffmpeg-2023-01-22-git-9d5e66942c-full_build\\bin\\ffmpeg.exe" //上线linux需要修改
	err = cmdchain.Builder().
		Join(cmd, "-i", path, "-vcodec", "h264", "-b:v", "6000k", "-bufsize", "6000k", name).
		Finalize().WithOutput(output).Run()
	if err != nil {
		//日志
		return "", err
	}
	return name, nil
}

// 保存图片
func savePic(path string) (picPath string, err error) {
	p := strings.Split(path, ".")[0]
	dir := strings.Split(p, "\\")                   //linux下换成/
	dirreal := strings.Join(dir[:len(dir)-1], "\\") //linux下换成/
	name := fmt.Sprintf("%s\\%s", dirreal, dir[len(dir)-1]+".jpg")
	output := &bytes.Buffer{}

	cmd := "D:\\programming\\ffmpeg-2023-01-22-git-9d5e66942c-full_build\\ffmpeg-2023-01-22-git-9d5e66942c-full_build\\bin\\ffmpeg.exe" //上线linux需要修改
	err = cmdchain.Builder().
		Join(cmd, "-i", path, "-ss", "00:00:01", "-frames:v", "1", name).
		Finalize().WithOutput(output).Run()
	if err != nil {
		//日志
		return "", err
	}
	return name, nil
}

// 处理视频的函数，传入视频路径
func DealVideo(path string) (videoPath string, picPath string, err error) {
	encode, _, err := encodeAndTime(path)
	videopath := path
	if encode != "h264" {
		//降低码率
		videopath, err = lowerBitRate(path)
		if err != nil {
			//日志
			return "", "", err
		}
	}
	//保存图片
	picPath, err = savePic(videopath)
	if err != nil {
		//日志
		return "", "", err
	}

	return videopath, picPath, nil
}
