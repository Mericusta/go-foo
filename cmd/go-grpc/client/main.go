package main

import (
	"context"
	"flag"
	"go-foo/cmd/go-grpc/myservice"
	"log"
	"math/rand/v2"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func randomPoint() *myservice.Point {
	lat := (rand.Int32N(180) - 90) * 1e7
	long := (rand.Int32N(360) - 180) * 1e7
	return &myservice.Point{Latitude: lat, Longitude: long}
}

// runRecordRoute sends a sequence of points to server and expects to get a RouteSummary from server.
func runRecordRoute(client myservice.RouteGuideClient) {
	// Create a random number of random points
	// pointCount := int(rand.Int32N(100)) + 2 // Traverse at least two points
	pointCount := 60
	var points []*myservice.Point
	for i := 0; i < pointCount; i++ {
		points = append(points, randomPoint())
	}
	log.Printf("Traversing %d points.", len(points))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.RecordRoute(ctx)
	if err != nil {
		log.Fatalf("client.RecordRoute failed: %v", err)
	}
	for _, point := range points {
		if err := stream.Send(point); err != nil {
			log.Fatalf("client.RecordRoute: stream.Send(%v) failed: %v", point, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("client.RecordRoute failed: %v", err)
	}
	log.Printf("Route summary: %v", reply)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			// *caFile = data.Path("x509/ca_cert.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := myservice.NewRouteGuideClient(conn)

	// // Looking for a valid feature
	// printFeature(client, &pb.Point{Latitude: 409146138, Longitude: -746188906})

	// // Feature missing.
	// printFeature(client, &pb.Point{Latitude: 0, Longitude: 0})

	// // Looking for features between 40, -75 and 42, -73.
	// printFeatures(client, &pb.Rectangle{
	// 	Lo: &pb.Point{Latitude: 400000000, Longitude: -750000000},
	// 	Hi: &pb.Point{Latitude: 420000000, Longitude: -730000000},
	// })

	// RecordRoute
	count := 100
	wg := &sync.WaitGroup{}
	wg.Add(count)
	for index := 0; index != count; index++ {
		go func() {
			runRecordRoute(client)
			wg.Done()
		}()
	}

	wg.Wait()

	// // RouteChat
	// runRouteChat(client)
}
