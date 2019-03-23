/*
package apisrv is a basic gRPC serving harness from which you can serve an MMF.
The gRPC service is defined in [REPO_ROOT]/api/protobuf-spec/function.proto
Most of the documentation for what these calls should do is in that file!

Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

package apisrv

import (
	"context"
	"net"
	"time"

	log "github.com/sirupsen/logrus"

	api "github.com/GoogleCloudPlatform/open-match/internal/pb"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"

	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"google.golang.org/grpc"
)

// Logrus structured logging setup
var (
	fnLogFields = log.Fields{
		"app":       "openmatch",
		"component": "function_service",
		"mmfname":   "openmatch-mmf-golang-serving-mmlogic-simple",
	}
	fnLog = log.WithFields(fnLogFields)
)

// FunctionServer implements api.FunctionServer, the server generated by compiling
// the protobuf, by fulfilling the api.FunctionServer interface.
type FunctionServer struct {
	grpc    *grpc.Server
	cfg     *viper.Viper
	pool    *redis.Pool
	mmlogic api.MmLogicClient
	fnName  string
}
type functionServer FunctionServer

// New returns an instantiated service
func New(cfg *viper.Viper, pool *redis.Pool, client api.MmLogicClient) *FunctionServer {

	// Initialize the server state
	s := FunctionServer{
		pool:    pool,
		grpc:    grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{})),
		cfg:     cfg,
		mmlogic: client,
		fnName:  "openmatch-mmf-golang-serving-mmlogic-simple",
	}

	// Register gRPC server
	api.RegisterFunctionServer(s.grpc, (*functionServer)(&s))
	fnLog.Info("Successfully registered gRPC server")
	return &s
}

// Open starts the api grpc service listening on the configured port.
func (s *FunctionServer) Open() error {
	ln, err := net.Listen("tcp", ":"+s.cfg.GetString("api.functions.port"))
	if err != nil {
		fnLog.WithFields(log.Fields{
			"error": err.Error(),
			"port":  s.cfg.GetInt("api.functions.port"),
		}).Error("net.Listen() error")
		return err
	}
	fnLog.WithFields(log.Fields{"port": s.cfg.GetInt("api.functions.port")}).Info("TCP net listener initialized")

	go func() {
		err := s.grpc.Serve(ln)
		if err != nil {
			fnLog.WithFields(log.Fields{"error": err.Error()}).Error("gRPC serve() error")
		}
		fnLog.Info("serving gRPC endpoints")
	}()

	// Initialize OpenCensus count metrics
	stats.Record(context.Background(), FnGrpcErrors.M(0))
	stats.Record(context.Background(), FnGrpcRequests.M(0))

	return nil
}

// Run is this service's implementation of the gRPC call defined in
// api/protobuf-spec/function.proto
// This doesn't include any of the MMLogic API boilerplate harness code, as
// that is all optional - MMFs can always be written to directly read/write
// to/from Redis.
func (s *functionServer) Run(ctx context.Context, fnArgs *api.Arguments) (*api.Result, error) {

	// Set up tagging for OpenCensus
	funcName := "Run"
	_ = funcName
	fnCtx, _ := tag.New(ctx, tag.Insert(KeyMethod, funcName))
	fnCtx, _ = tag.New(fnCtx, tag.Insert(KeyFnName, s.fnName))

	// Everything the mmf needs to know about its specific
	// run (where to look in state storage for its profile, where to write
	// its results to state storage) needs to be in the fnArgs message.
	start := time.Now()
	err := mmfRun(fnCtx, fnArgs, s.cfg, s.mmlogic)
	stats.Record(fnCtx, FnGrpcLatencySecs.M(time.Since(start).Seconds()))

	if err != nil {
		fnLog.WithFields(log.Fields{"error": err.Error()}).Error("function.Run error")
		stats.Record(fnCtx, FnGrpcErrors.M(1))
		return &api.Result{Success: false, Error: err.Error()}, err
	}
	stats.Record(fnCtx, FnGrpcRequests.M(1))
	return &api.Result{Success: true}, nil
}
