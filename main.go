package main

import (
	"github.com/zhangyiming748/avif-image-encoder/avif"
	"github.com/zhangyiming748/avif-image-encoder/constant"
	"github.com/zhangyiming748/avif-image-encoder/log"
	"github.com/zhangyiming748/avif-image-encoder/util"
	"os"
)

func main() {
	p := new(constant.Param)
	p.SetRoot("D:\\.telegram")
	if root := os.Getenv("root"); root != "" {
		p.SetRoot(root)
	}
	log.SetLog(p)
	files := util.GetAllfiles(p)
	for _, f := range files {
		avif.ToAvif(f)
	}
}
