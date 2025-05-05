package ports

type UpdateFileArgs struct {
	Key        string
	BucketName string
	Content    []byte
}

type BucketPort interface {
	UploadFile(key string, BucketName string) (string, error)
	DownloadFile(key string, BucketName string) ([]byte, error)
	DeleteFile(key string, BucketName string) error
}
