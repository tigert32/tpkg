package pkgprocessor

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"tpkg/utils"
)

func ProcessFile(filePath, outputDir string) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Printf("获取文件绝对路径失败: %v\n", err)
		return
	}
	fmt.Println(absPath)

	fileName := filepath.Base(absPath)
	baseName := strings.Split(fileName, ".")[0]
	outFilePath := baseName
	if outputDir != "" {
		outFilePath = filepath.Join(outputDir, baseName)
	}
	fmt.Println(outFilePath)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("无法打开文件: %v\n", err)
		return
	}
	defer file.Close()

	// 读取前16字节
	header := make([]byte, 16)
	_, err = file.Read(header)
	if err != nil {
		fmt.Printf("读取文件头部失败: %v\n", err)
		return
	}

	// 校验 TPKG
	expectedHeader := []byte{0x54, 0x50, 0x4B, 0x47}
	if !utils.CheckHeader(header[:4], expectedHeader) {
		fmt.Println("文件头不匹配")
		return
	}

	// 获取条目数量
	entryCount := binary.LittleEndian.Uint32(header[12:16])
	fmt.Printf("条目数量: %d\n", entryCount)

	// 读取和解析条目
	var seek int64 = 16
	for i := 0; i < int(entryCount); i++ {
		entry, err := ReadEntry(file, &seek)
		if err != nil {
			fmt.Printf("读取条目失败: %v\n", err)
			return
		}
		err = CreateEntryFile(outFilePath, entry, file)
		if err != nil {
			fmt.Printf("创建条目文件失败: %v\n", err)
			return
		}
	}
}

func ProcessDirectory(directoryPath, outputDir string) {
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".pkg") {
			outputSubDir := directoryPath
			if outputDir != "" {
				outputSubDir = filepath.Join(directoryPath, outputDir)
			}
			//处理单个文件
			ProcessFile(path, outputSubDir)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("遍历目录失败: %v\n", err)
	}
}
