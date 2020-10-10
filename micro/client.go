package micro

import (
	data "testGin/datasharems/share"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
)

// GetDataShareClient 获取dataShare client
func GetDataShareClient() data.ShareDataService {
	return data.NewShareDataService("go.micro.srv.share.consumer", client.NewClient(client.Registry(etcd.NewRegistry(registry.Addrs("192.168.8.39")))))
}
