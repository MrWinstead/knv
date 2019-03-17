package cmd

import (
	"github.com/mrwinstead/knv/stat"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/mrwinstead/knv/cluster"
	"github.com/mrwinstead/knv/configuration"
)

var (
	clusterStatCmd = &cobra.Command{
		Use:     "stat",
		Aliases: []string{"cluster-stat"},
		Run:     clusterStatCmdEntrypoint,
	}
)

func init() {
	rootCmd.AddCommand(clusterStatCmd)
}

func clusterStatCmdEntrypoint(cmd *cobra.Command, args []string) {
	etcdClient := NewEtcdClientOrFatal(configuredViper)
	pathAdvisor := cluster.NewEtcdPathAdvisor(configuredViper.GetString(
		configuration.KeyClusterInformationPathPrefix))

	members, membershipListErr := stat.GetClusterMembers(rootContext,
		etcdClient, pathAdvisor)
	if nil != membershipListErr {
		logrus.Fatal("could not list members", membershipListErr)
	}
	logrus.WithFields(logrus.Fields{"members": members}).
		Println("current cluster members")
}
