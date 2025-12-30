package initialize

import (
	"context"
	"zpi/server/cmd/storage/config"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinIOClient *minio.Client

// InitMinIO 初始化MinIO客户端
func InitMinIO() {
	c := config.GlobalServerConfig.MinIOInfo
	var err error
	MinIOClient, err = minio.New(c.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.AccessKey, c.SecretKey, ""),
		Secure: c.UseSSL,
	})
	if err != nil {
		klog.Fatalf("init minio client failed: %s", err.Error())
	}

	// 检查bucket是否存在，不存在则创建
	ctx := context.Background()
	exists, err := MinIOClient.BucketExists(ctx, c.Bucket)
	if err != nil {
		klog.Fatalf("check bucket exists failed: %s", err.Error())
	}
	if !exists {
		err = MinIOClient.MakeBucket(ctx, c.Bucket, minio.MakeBucketOptions{})
		if err != nil {
			klog.Fatalf("create bucket failed: %s", err.Error())
		}
		klog.Infof("bucket %s created successfully", c.Bucket)
	}
}
