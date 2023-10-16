package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
)

type SystemMinioService struct{}

func (s *SystemMinioService) UpdateFile(req system.UpdateFile, file *multipart.FileHeader, fileName string) {
	var utilsMinio utils.MinioOperate
	size := file.Size
	f, err := file.Open()
	if err != nil {
		return
	}
	minioInfo, err1 := utilsMinio.SaveFileToMinio(req.BucketName, req.ObjectName, fileName, f, size)
	if err1 != nil {
		return
	}
	fmt.Println(minioInfo)
}

func (s *SystemMinioService) GetFile(req system.GetFile) (minioObject *minio.Object, err error) {
	var utilsMinio utils.MinioOperate
	data, err1 := utilsMinio.GetFileFromMinio(req.BucketName, req.ObjectName)
	if err1 != nil {
		global.GVA_LOG.Error("minio下载文件失败")
		return nil, err1
	}
	return data, nil
}

func (s *SystemMinioService) DeleteFile(req system.DeleteFile) error {
	var utilsMinio utils.MinioOperate
	err1 := utilsMinio.RemoveFile(req.BucketName, req.ObjectName)
	if err1 != nil {
		global.GVA_LOG.Error("删除minio文件失败")
		return err1
	}
	return nil
}
