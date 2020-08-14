// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/sync/errgroup"

	drghs_v1 "github.com/GoogleCloudPlatform/devrel-services/drghs/v1"

	"github.com/GoogleCloudPlatform/devrel-services/drghs-worker/pkg/googlers"

	"cloud.google.com/go/errorreporting"
	"golang.org/x/build/maintner"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/instrumentation/grpctrace"
)

var (
	listen     = flag.String("listen", "0.0.0.0:6343", "listen address")
	intListen  = flag.String("intListen", "0.0.0:6344", "listen for internal service")
	sloAddress = flag.String("sloServer", "0.0.0:3009", "address of slo service")
	verbose    = flag.Bool("verbose", false, "enable verbose debug output")
	bucket     = flag.String("bucket", "cdpe-maintner", "Google Cloud Storage bucket to use for log storage")
	token      = flag.String("token", "", "Token to Access GitHub with")
	projectID  = flag.String("gcp-project", "", "The GCP Project this is using")
	owner      = flag.String("owner", "", "The owner of the GitHub repository")
	repo       = flag.String("repo", "", "The repository to track")
)

var (
	corpus          = &maintner.Corpus{}
	googlerResolver googlers.Resolver
	errorClient     *errorreporting.Client
)

func main() {
	// Set log to Stdout. Default for log is Stderr
	log.SetOutput(os.Stdout)
	flag.Parse()

	ctx := context.Background()

	if *owner == "" {
		err := fmt.Errorf("must provide --owner")
		// logAndPrintError(err)
		log.Fatal(err)
	}

	if *repo == "" {
		err := fmt.Errorf("must provide --repo")
		// logAndPrintError(err)
		log.Fatal(err)
	}
	group, ctx := errgroup.WithContext(context.Background())

	group.Go(
		// Get SLO rules for the repo
		func() error {
			parent := fmt.Sprintf("owners/%s/repositories/%s", *owner, *repo)

			ticker := time.NewTicker(10 * time.Second)

			for t := range ticker.C {
				log.Printf("Slo sync at %v", t)

				_, err := getSlos(ctx, parent)
				if err != nil {
					// logAndPrintError(err)
					log.Printf("Slo sync err: %v", err)
				}
			}
			return nil
		})

	err := group.Wait()
	log.Fatal(err)
}

func getSlos(ctx context.Context, parent string) ([]*drghs_v1.SLO, error) {

	conn, err := grpc.Dial(
		*sloAddress,
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				grpctrace.UnaryClientInterceptor(global.Tracer("maintner-leif")),
				buildRetryInterceptor(),
			),
		),
	)

	if err != nil {
		return nil, fmt.Errorf("Error connecting to SLO server: %v", err)
	}
	defer conn.Close()

	sloClient := drghs_v1.NewSLOServiceClient(conn)

	response, err := sloClient.ListSLOs(ctx, &drghs_v1.ListSLOsRequest{Parent: parent})
	if err != nil {
		return nil, fmt.Errorf("Error getting SLOs: %v", err)
	}

	slos := response.GetSlos()
	nextPage := response.GetNextPageToken()

	for nextPage != "" {
		response, err = sloClient.ListSLOs(ctx, &drghs_v1.ListSLOsRequest{Parent: parent, PageToken: nextPage})
		if err != nil {
			// logAndPrintError(err)
			log.Printf("Error getting SLOs: %v", err)
			continue
		}

		slos = append(slos, response.GetSlos()...)
		nextPage = response.GetNextPageToken()
	}
	return slos, nil
}

func unaryInterceptorLog(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	log.Printf("Starting RPC: %v at %v", info.FullMethod, start)

	// Used for Debugging incoming context and metadata issues
	// md, _ := metadata.FromIncomingContext(ctx)
	// log.Printf("RPC: %v. Metadata: %v", info.FullMethod, md)
	m, err := handler(ctx, req)
	if err != nil {
		errorClient.Report(errorreporting.Entry{
			Error: err,
		})
		log.Printf("RPC: %v failed with error %v", info.FullMethod, err)
	}

	log.Printf("Finishing RPC: %v. Took: %v", info.FullMethod, time.Now().Sub(start))
	return m, err
}

func buildRetryInterceptor() grpc.UnaryClientInterceptor {
	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(500 * time.Millisecond)),
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted, codes.Unavailable, codes.DeadlineExceeded),
		grpc_retry.WithMax(5),
	}
	return grpc_retry.UnaryClientInterceptor(opts...)
}
