package util

import (
	"github.com/h2non/filetype"
	"github.com/zhangyiming748/avif-image-encoder/constant"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

/*
获取当前文件夹和子文件夹下的全部文件的绝对路径 并放在字符串切片下
*/
func GetAllfiles(p *constant.Param) []string {
	log.Println("开始获取文件路径...")
	var allFiles []string
	filepath.Walk(p.GetRoot(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			allFiles = append(allFiles, path)
		}
		return nil
	})
	currentFiles := getFilesByMate(allFiles)
	return currentFiles
}

/*
根据扩展名判断路径切片中符合的文件
*/
func getFilesByExt(fps []string) (allFiles []string) {
	audio := []string{
		".jpg",
		".png",
		".gif",
		".webp",
		".cr2",
		".tif",
		".bmp",
		".heif",
		".ico",
		".dwg",
	}
	for _, fp := range fps {
		for _, item := range audio {
			if filepath.Ext(fp) == item {
				allFiles = append(allFiles, fp)
			}
		}
	}
	return allFiles
}

/*
根据元数据判断路径切片中符合的文件
*/

func getFilesByMate(fps []string) (allFiles []string) {
	for _, fp := range fps {
		f, err := os.ReadFile(fp)
		if err != nil {
			log.Fatalf("Error opening file:%v\n", err)
		}
		if filetype.IsImage(f) && !strings.Contains(path.Ext(fp), "avif") && !strings.Contains(path.Ext(fp), "gif") && !strings.Contains(path.Ext(fp), "avif") {
			allFiles = append(allFiles, fp)
		}
	}
	return allFiles
}
