package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type ProjectFile struct {
	global.GVA_MODEL
	Name       string `json:"name" form:"name" gorm:"column:name;comment:文件名;size:100;"`
	FileMD5    string `json:"fileMD5" form:"fileMD5" gorm:"column:file_md5;comment:文件md5;size:150;"`
	ProjectID  int    `json:"projectID" form:"projectID" gorm:"column:project_id;comment:所属项目ID;"`
	FileTypeID int    `json:"fileTypeID" form:"fileTypeID" gorm:"column:file_type_id;comment:所属类型ID;"`
}

// TableName Project 表名
func (ProjectFile) TableName() string {
	return "project_file"
}
