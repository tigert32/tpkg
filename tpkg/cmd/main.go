package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"tpkg/pkgprocessor"
)

func main() {

	singleFile := flag.String("f", "", "Single file to process")
	dirPath := flag.String("d", "", "Directory to process")
	outputDir := flag.String("o", "", "Output directory")

	flag.Parse()

	if *singleFile != "" { //单个文件模式
		pkgprocessor.ProcessFile(*singleFile, *outputDir)
	} else if *dirPath != "" { //文件夹内含有多个.pkg
		pkgprocessor.ProcessDirectory(*dirPath, *outputDir)
	} else {
		fmt.Println("Usage: -d 支持目录下多个pkg文件 -f 支持目录内单个文件 -o 设定输出目录")
		fmt.Println("Usage: " + filepath.Base(os.Args[0]) + " -f sss.pkg | -d d:/appid_123/ver_123 [-o <output directory>]")
	}
}
