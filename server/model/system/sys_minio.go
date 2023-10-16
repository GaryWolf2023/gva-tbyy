package system

type UpdateFile struct {
	BucketName string `json:"bucketName"`
	ObjectName string `json:"objectName"`
}

type GetFile struct {
	BucketName string `json:"bucketName"`
	ObjectName string `json:"objectName"`
}

type DeleteFile struct {
	BucketName string `json:"bucketName"`
	ObjectName string `json:"objectName"`
}
