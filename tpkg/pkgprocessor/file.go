package pkgprocessor

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path"
)

type Entry struct {
	Name   string
	Offset uint32
	Size   uint32
}

func ReadEntry(file *os.File, seek *int64) (*Entry, error) {
	// 跳转到偏移量
	_, err := file.Seek(*seek, io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("跳转到偏移量失败: %v", err)
	}

	// 读取条目名称大小
	nameSizeBytes := make([]byte, 4)
	_, err = file.Read(nameSizeBytes)
	if err != nil {
		return nil, err
	}
	nameSize := binary.LittleEndian.Uint32(nameSizeBytes)

	// 读取条目名称
	nameBytes := make([]byte, nameSize)
	_, err = file.Read(nameBytes)
	if err != nil {
		return nil, err
	}
	name := string(nameBytes)

	// 读取偏移量
	offsetBytes := make([]byte, 4)
	_, err = file.Read(offsetBytes)
	if err != nil {
		return nil, err
	}
	offset := binary.LittleEndian.Uint32(offsetBytes)

	// 读取大小
	sizeBytes := make([]byte, 4)
	_, err = file.Read(sizeBytes)
	if err != nil {
		return nil, err
	}
	size := binary.LittleEndian.Uint32(sizeBytes)

	// 更新文件指针位置
	currentPos, err := file.Seek(0, io.SeekCurrent)
	if err != nil {
		return nil, fmt.Errorf("获取当前文件指针位置失败: %v", err)
	}
	*seek = currentPos

	ret := &Entry{
		Name:   name,
		Offset: offset,
		Size:   size,
	}
	fmt.Println(ret)
	return ret, nil
}

func CreateEntryFile(outFilePath string, entry *Entry, file *os.File) error {
	_, err := file.Seek(int64(entry.Offset), io.SeekStart)
	if err != nil {
		return fmt.Errorf("跳转到偏移量失败: %v", err)
	}

	data := make([]byte, entry.Size)
	_, err = file.Read(data)
	if err != nil {
		return fmt.Errorf("读取条目数据失败: %v", err)
	}

	tdir := outFilePath + "/" + path.Dir(entry.Name)
	fmt.Println(tdir)
	os.MkdirAll(tdir, 0755)
	spath, filename := path.Split(entry.Name)
	fmt.Println(spath, filename)
	fmt.Println(tdir + "/" + filename)
	outFile, err := os.Create(tdir + "/" + filename)
	if err != nil {
		return fmt.Errorf("创建条目文件失败: %v", err)
	}
	defer outFile.Close()

	_, err = outFile.Write(data)
	if err != nil {
		return fmt.Errorf("写入条目数据失败: %v", err)
	}

	return nil
}
