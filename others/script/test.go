package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"script/rime"
	"strings"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "s":
			goto SORT
		case "t":
			rime.Temp()
			return
		case "tp":
			rime.Pinyin(filepath.Join(rime.RimeDir, "cn_dicts/temp.txt"))
			return
		}
	}

	// 从 others/cn_en.txt 更新中英混输词库
	rime.CnEn()
	fmt.Println("--------------------------------------------------")

	areYouOK()

SORT:
	// 排序，顺便去重
}

func areYouOK() {
	fmt.Println("Are you OK:")
	var isOK string
	_, _ = fmt.Scanf("%s", &isOK)
	isOK = strings.ToLower(isOK)
	if isOK != "ok" && isOK != "y" && isOK != "yes" {
		os.Exit(123)
	}
}
