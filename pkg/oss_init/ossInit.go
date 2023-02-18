package oss_init

import (
	"dousheng/conf"
	g "dousheng/pkg/global"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudwego/kitex/pkg/klog"
	"sync"
)

var bucket *oss.Bucket
var once sync.Once

// OSSInit 初始化
func OSSInit() {
	once.Do(func() {
		// 连接OSS账户
		client, err := oss.New(conf.EndPoint, conf.AccessKeyID, conf.AccessKeySecret)

		if err != nil {
			klog.Error("连接OSS账户失败" + err.Error())
		} else { // OSS账户连接成功

			// 连接存储空间
			bucket, err = client.Bucket(conf.BucketName)
			if err != nil {
				klog.Error("连接存储空间失败" + err.Error())
			} else { // 存储空间连接成功
				klog.Error("OSS初始化完成")
			}
		}
	})
	g.OssBucket = bucket

}
