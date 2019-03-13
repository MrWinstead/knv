package service

import (
	"context"

	"github.com/mrwinstead/knv"
	"github.com/mrwinstead/knv/observation"
	"github.com/mrwinstead/knv/storage"
)

const (
	sinkNameOperationDelete = "ServiceOperationDelete"
	sinkNameOperationGet    = "ServiceOperationGet"
	sinkNameOperationPut    = "ServiceOperationPut"

	operationNameDelete = "Delete"
	operationNameGet    = "Get"
	operationNamePut    = "Put"
)

var (
	_ knv.DatabaseServer = &database{}
)

type Observable struct {
	Operation string

	Error string

	Request  knv.DatabaseGetRequest
	Response knv.DatabasePutResponse
}

type database struct {
	valueStore storage.MultiKeyValue

	sinkDelete observation.Sink
	sinkGet    observation.Sink
	sinkPut    observation.Sink
}

func New(valueStore storage.MultiKeyValue, mgr observation.Manager) knv.DatabaseServer {
	created := &database{
		valueStore: valueStore,
	}
	populateSinks(created, mgr)
	return created
}

func populateSinks(d *database, mgr observation.Manager) {
	d.sinkDelete = mgr.Sink(sinkNameOperationDelete)
	d.sinkGet = mgr.Sink(sinkNameOperationGet)
	d.sinkPut = mgr.Sink(sinkNameOperationPut)
}

func (d *database) Get(ctx context.Context, req *knv.DatabaseGetRequest) (*knv.DatabaseGetResponse, error) {
	obs := &Observable{
		Operation: operationNameGet,
	}
	defer d.sinkGet.Submit(nil, obs)

	resp := &knv.DatabaseGetResponse{}
	return resp, nil
}

func (d *database) Delete(ctx context.Context, req *knv.DatabaseDeleteRequest) (*knv.DatabaseDeleteResponse, error) {
	obs := &Observable{
		Operation: operationNameDelete,
	}
	defer d.sinkGet.Submit(nil, obs)

	resp := &knv.DatabaseDeleteResponse{}
	return resp, nil
}

func (d *database) Put(ctx context.Context, req *knv.DatabasePutRequest) (*knv.DatabasePutResponse, error) {
	obs := &Observable{
		Operation: operationNamePut,
	}
	defer d.sinkGet.Submit(nil, obs)

	resp := &knv.DatabasePutResponse{}
	return resp, nil
}
