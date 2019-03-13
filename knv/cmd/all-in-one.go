package cmd

import (
	"github.com/mrwinstead/knv"
	"github.com/mrwinstead/knv/configuration"
	"github.com/mrwinstead/knv/observation"
	"github.com/mrwinstead/knv/service"
	"github.com/mrwinstead/knv/storage"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var allInOneCmd = &cobra.Command{
	Use:  "all-in-one",
	RunE: allInOneCmdEntrypointE,
}

func init() {
	rootCmd.AddCommand(allInOneCmd)
}

func allInOneCmdEntrypointE(cmd *cobra.Command, args []string) error {
	manager := observation.NewChannelManager(uint(
		configuredViper.GetInt(configuration.KeyObservationBacklogMax)))
	store, storageBuildErr := storage.NewBadgerBackedStore(
		configuredViper.GetString(configuration.KeyDirectoryNameRootDatabase),
		nil,
		manager)
	if nil != storageBuildErr {
		log.Fatalln("could not create backing storage", storageBuildErr)
	}

	storageService := service.New(store, manager)

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

	serveErr := grpcServer.Serve(listener)
	if nil != serveErr {
		log.Fatalln("could not run service", serveErr)
	}

	return nil
}
