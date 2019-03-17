package stat

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/mrwinstead/knv/cluster"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"path"
)

type ClusterMember struct {
	Identity cluster.ReportedIdentity
}

func GetClusterMembers(ctx context.Context, client *clientv3.Client,
	pa cluster.PathAdvisor) ([]ClusterMember, error) {
	members := make([]ClusterMember, 0)

	membersPrefix := pa.ExpandPath(cluster.PathMembers)

	logrus.Println(membersPrefix)

	results, getErr := client.Get(ctx, membersPrefix, clientv3.WithPrefix())
	if nil != getErr {
		err := errors.Wrap(getErr, "could not fetch cluster members")
		return nil, err
	}

	for _, result := range results.Kvs {
		memberID := cluster.ReportedIdentity(path.Base(string(result.Key)))
		members = append(members, ClusterMember{memberID})
		logrus.Println(string(result.Key))
	}

	return members, nil
}
