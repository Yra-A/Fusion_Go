package oss

import (
    "bytes"
    conf "github.com/Yra-A/Fusion_Go/pkg/configs/oss"
)

// UploadFile 上传图片文件
func UploadFile(filename string, data []byte) error {
    reader := bytes.NewReader(data)
    objectName := conf.BaseURL + filename
    err := bucket.PutObject(objectName, reader)
    if err != nil {
        return err
    }
    return nil
}
