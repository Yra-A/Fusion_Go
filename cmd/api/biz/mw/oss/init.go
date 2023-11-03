package oss

import (
    "fmt"
    conf "github.com/Yra-A/Fusion_Go/pkg/configs/oss"
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
    client *oss.Client
    bucket *oss.Bucket
)

func Init() {
    // 创建OSSClient实例
    client, err := oss.New(conf.EndPoint, conf.AccessKeyID, conf.AccessKeySecret)
    if err != nil {
        fmt.Errorf("oss init error: %v", err)
    }
    // 获取存储空间。
    bucket, err = client.Bucket(conf.OssBucketName)
    if err != nil {
        fmt.Errorf("oss init error: %v", err)
    }

}
