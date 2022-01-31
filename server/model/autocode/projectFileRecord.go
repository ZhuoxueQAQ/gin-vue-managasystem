package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type ProjectFileRecord struct {
	global.GVA_MODEL
	FileName   string `json:"fileName" form:"fileName" gorm:"column:file_name;comment:文件名;size:100;"`
	FileMD5    string `json:"fileMD5" form:"fileMD5" gorm:"column:file_md5;comment:文件md5;size:150;"`
	ProjectID  string `json:"projectID" form:"projectID" gorm:"column:project_id;comment:所属项目ID;size:10;"`
	FileTypeID string `json:"fileTypeID" form:"fileTypeID" gorm:"column:file_type_id;comment:所属类型ID;size:10;"`
}

// TableName Project 表名
func (ProjectFileRecord) TableName() string {
	return "project_file"
}
