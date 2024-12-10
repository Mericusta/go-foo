package main

import (
	"fmt"
	"go-foo/cmd/go-grpc/myservice"
	"io"
	"sync/atomic"
	"time"
)

type routeGuideServer struct {
	myservice.UnimplementedRouteGuideServer
}

// RecordRoute records a route composited of a sequence of points.
//
// It gets a stream of points, and responds with statistics about the "trip":
// number of points,  number of known features visited, total distance traveled, and
// total time spent.
func (s *routeGuideServer) RecordRoute(stream myservice.RouteGuide_RecordRouteServer) error {
	var pointCount int32
	// startTime := time.Now()
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			// endTime := time.Now()
			return stream.SendAndClose(nil)
			// return stream.SendAndClose(&pb.RouteSummary{
			// 	PointCount:   pointCount,
			// 	FeatureCount: featureCount,
			// 	Distance:     distance,
			// 	ElapsedTime:  int32(endTime.Sub(startTime).Seconds()),
			// })
		}
		if err != nil {
			return err
		}
		pointCount++
		c <- point
	}
}

var (
	c       chan *myservice.Point
	counter atomic.Int64
)

func init() {
	go func() {
		timer := time.NewTimer(time.Second * 10)
		c = make(chan *myservice.Point, 1024)
		for {
			select {
			case <-timer.C:
				fmt.Printf("counter: %v\n", counter.Load())
				return
			case <-c:
				// time.Sleep(time.Second)
				// fmt.Printf("v = %+v\n", v)
				counter.Add(1)
			}
		}
	}()
}
