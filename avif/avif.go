package avif

import (
	"github.com/zhangyiming748/FastMediaInfo"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
)

/*
* 将mp3文件转为aac文件
接收源文件的绝对路径
*/
func ToAvif(fp string) {
	ext := path.Ext(fp)
	out := strings.Replace(fp, ext, ".avif", 1)
	mi := FastMediaInfo.GetStandMediaInfo(fp)
	width, _ := strconv.Atoi(mi.Image.Width)
	height, _ := strconv.Atoi(mi.Image.Height)
	crf := FastMediaInfo.GetCRF("avif", width, height)
	if crf == "" {
		crf = "31"
	}
	cmd := exec.Command("ffmpeg", "-y", "-i", fp, "-c:v", "libaom-av1", "-crf", crf, "-still-picture", "1", out)

	log.Printf("开始执行命令%v\n", cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("命令%v执行失败%v\n", cmd.String(), err)
		return
	} else {
		log.Printf("命令成功执行%v\n", string(output))
		os.Remove(fp)
	}
}
