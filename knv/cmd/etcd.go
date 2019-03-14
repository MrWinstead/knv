package cmd

import (
	"time"

	"github.com/etcd-io/etcd/clientv3"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/mrwinstead/knv/configuration"
)

func NewEtcdClientOrFatal(v *viper.Viper) *clientv3.Client {
	endpoints := v.GetStringSlice(configuration.KeyBackingServiceEtcd)
	client, buildErr := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 10 * time.Second,
	})
	if nil != buildErr {
		err := errors.Wrap(buildErr, "could not create etcdv3 client")
		panic(err)
	}
	return client
}
