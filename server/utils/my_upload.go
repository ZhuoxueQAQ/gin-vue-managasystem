package utils

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: MakeFileChunk
//@description: 创建切片文件
//@param: content []byte, fileName string, contentNumber int, fileMd5 string)
//@return: error, string

func MakeFileChunk(content []byte, fileName string, contentNumber int, fileMd5 string) (error, string) {
	path := filepath.Join(breakpointDir, fileMd5)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err, path
	}
	err, pathc := makeFileContent(content, fileName, path, contentNumber)
	return err, pathc
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: MergeChunk
//@description: 合并切片文件
//@param: fileName string, FileMd5 string
//@return: error, string

func MergeChunk(fileName string, FileMd5 string, projectID string, fileTypeID string) (error, string) {
	// 待合并文件目录
	targetDir := filepath.Join(finishDir, projectID, fileTypeID)
	// 切片所在目录
	chuckDir := filepath.Join(breakpointDir, FileMd5)
	_ = os.MkdirAll(targetDir, os.ModePerm)
	savePath := filepath.Join(targetDir, fileName)
	// 读取切片文件
	rd, err := ioutil.ReadDir(chuckDir)
	if err != nil {
		return err, savePath
	}
	fd, err := os.OpenFile(savePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		return err, savePath
	}
	defer fd.Close()
	for k := range rd {
		content, _ := ioutil.ReadFile(filepath.Join(chuckDir, fileName+"_"+strconv.Itoa(k)))
		_, err = fd.Write(content)
		if err != nil {
			_ = os.Remove(savePath)
			return err, savePath
		}
	}
	// 删除缓存分片文件
	if err := os.RemoveAll(chuckDir); err != nil {
		return err, chuckDir
	}
	return nil, savePath
}

func DeleteProjectFile(record autocode.ProjectFileRecord) error {
	path := GetProjectFilePath(record)
	err := os.Remove(path)
	return err
}

// DeleteProjectFiles 删除一个项目的所有附件
func DeleteProjectFiles(projectID string) (err error) {
	if err = os.RemoveAll(filepath.Join(finishDir, projectID)); err != nil {
		return err
	}
	return nil
}

func GetProjectFilePath(record autocode.ProjectFileRecord) string {
	return filepath.Join(finishDir, record.ProjectID, record.FileTypeID, record.FileName)
}
