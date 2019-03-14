package cmd

import (
	"context"
	"github.com/mrwinstead/knv"
	"github.com/mrwinstead/knv/cluster"
	"github.com/mrwinstead/knv/configuration"
	"github.com/mrwinstead/knv/observation"
	"github.com/mrwinstead/knv/service"
	"github.com/mrwinstead/knv/storage"
	"github.com/oklog/ulid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

var allInOneCmd = &cobra.Command{
	Use: "all-in-one",
	Run: allInOneCmdEntrypoint,
}

func init() {
	rootCmd.AddCommand(allInOneCmd)
}

func allInOneCmdEntrypoint(cmd *cobra.Command, args []string) {
	observationManager := observation.NewChannelManager(uint(
		configuredViper.GetInt(configuration.KeyObservationBacklogMax)))
	go func() {
		obsMgrErr := observationManager.Run(rootContext)
		if nil != obsMgrErr && context.Canceled != obsMgrErr {
			logrus.Fatalln("could not run observation manager",
				obsMgrErr)
		}
	}()
	stdoutObserver := observation.NewLogrusObserver(os.Stdout)
	observationManager.Register(stdoutObserver)

	serverIdentity := ulid.MustNew(ulid.Now(), monotonicIDSource)

	store, storageBuildErr := storage.NewBadgerBackedStore(
		configuredViper.GetString(configuration.KeyDirectoryNameRootDatabase),
		nil,
		observationManager)
	if nil != storageBuildErr {
		log.Fatalln("could not create backing storage", storageBuildErr)
	}

	etcdClient := NewEtcdClientOrFatal(configuredViper)
	pathAdvisor := cluster.NewEtcdPathAdvisor(configuredViper.GetString(
		configuration.KeyClusterInformationPathPrefix))

	membershipMgr, mgrBuildErr := cluster.NewEtcdMembershipReporter(
		etcdClient, observationManager,
		cluster.ReportedIdentity(serverIdentity.String()), pathAdvisor,
		configuredViper.GetDuration(configuration.KeyMembershipLeaseTime))
	if nil != mgrBuildErr {
		logrus.Fatal("unable to build membership reporter",
			mgrBuildErr)
	}
	go func() {
		mgrRunErr := membershipMgr.Run(rootContext)
		if nil != mgrRunErr && context.Canceled != mgrRunErr {
			logrus.Fatal("error running cluster membership observationManager",
				mgrRunErr)
		}
	}()

	storageService := service.New(store, observationManager)

	listener, listenErr := net.Listen("tcp",
		":"+configuredViper.GetString(configuration.KeyPortGRPC))
	if nil != listenErr {
		log.Fatalln("could not listen on ",
			configuredViper.GetString(configuration.KeyPortGRPC), listenErr)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	knv.RegisterDatabaseServer(grpcServer, storageService)
	reflection.Register(grpcServer)
	go func() {
		serveErr := grpcServer.Serve(listener)
		if nil != serveErr {
			log.Fatalln("could not run service", serveErr)
		}
	}()
	go func() {
		select {
		case <-rootContext.Done():
			grpcServer.Stop()
		}
	}()

	<-rootContext.Done()
}
