package filesystem

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var _ IFilesystem = (*LocalFilesystem)(nil)

type LocalFilesystem struct {
	config LocalSystemConfig
}

// 在 Go 语言中，os.Stat 函数用于获取文件或目录的状态信息。os.MkdirAll 方法用于创建多级目录。
func (l LocalFilesystem) Write(bucketName string, objectName string, stream []byte) error {
	filePath := l.Path(bucketName, objectName)
	dir := path.Dir(filePath)
	if len(dir) > 0 && !isDirExist(dir) {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}
	return nil
}

func (l LocalFilesystem) Copy(bucketName string, srcObjectName, objectName string) error {
	return l.WriteLocal(bucketName, l.Path(bucketName, srcObjectName), objectName)
}

func (l LocalFilesystem) CopyObject(srcBucketName string, srcObjectName, dstBucketName string, dstObjectName string) error {
	//TODO implement me
	panic("implement me")
}

func (l LocalFilesystem) Delete(bucketName string, objectName string) error {
	//TODO implement me
	panic("implement me")
}

func (l LocalFilesystem) GetObject(bucketName string, objectName string) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (l LocalFilesystem) PublicUrl(bucketName, objectName string) string {
	//TODO implement me
	panic("implement me")
}

func (l LocalFilesystem) PrivateUrl(bucketName, objectName string, filename string, expire time.Duration) string {
	//TODO implement me
	panic("implement me")
}

func (l LocalFilesystem) InitiateMultipartUpload(bucketName, objectName string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (l LocalFilesystem) PutObjectPart(bucketName, objectName string, uploadID string, index int, data io.Reader, size int64) (ObjectPart, error) {
	//TODO implement me
	panic("implement me")
}

func (l LocalFilesystem) CompleteMultipartUpload(bucketName, objectName, uploadID string, parts []ObjectPart) error {
	//TODO implement me
	panic("implement me")
}

func (l LocalFilesystem) AbortMultipartUpload(bucketName, objectName, uploadID string) error {
	//TODO implement me
	panic("implement me")
}

func NewLocalFileSystem(config LocalSystemConfig) *LocalFilesystem {
	return &LocalFilesystem{config}
}

func (l LocalFilesystem) Driver() string {
	return LocalDriver
}
func (l LocalFilesystem) BucketPublicName() string {
	return l.config.BucketPublic
}
func (l LocalFilesystem) BucketPrivateName() string {
	return l.config.BucketPrivate
}
func NewLocalFilesystem(config LocalSystemConfig) *LocalFilesystem {
	return &LocalFilesystem{config}
}
func (l LocalFilesystem) Stat(bucketName string, objectName string) (*FileStatInfo, error) {
	info, err := os.Stat(l.Path(bucketName, objectName))
	if err != nil {
		return nil, err
	}
	return &FileStatInfo{
		Name:        filepath.Base(objectName),
		Size:        info.Size(),
		Ext:         filepath.Ext(objectName),
		MimeType:    "",
		LastModTime: info.ModTime(),
	}, nil
}
func (l LocalFilesystem) WriteLocal(bucketName string, localFile string, objectName string) error {
	scrFile, err := os.Open(localFile)
	if err != nil {
		return err
	}
	defer scrFile.Close()
	objectName = l.Path(bucketName, objectName)
	dir := path.Dir(objectName)
	if len(dir) > 0 && !isDirExist(dir) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil

}
func (l LocalFilesystem) Path(bucketName string, objectName string) string {
	return fmt.Sprintf(
		"%s/%s/%s",
		strings.TrimRight(l.config.Root, "/"),
		bucketName,
		strings.TrimLeft(objectName, "/"),
	)
}
