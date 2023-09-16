package utils

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// yourBucketName填写存储空间名称。
var BucketName = "douyin-duu"

// 创建OSSClient实例。
// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
var OSSClient, _ = oss.New("endpoint", "accessKeyID", "accessKeySecret")

// 获取存储空间。
var Bucket, _ = OSSClient.Bucket(BucketName)

// 初始化VideoList
var VideoList = []*model.Video{}
