package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/pubsubio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/options/gcpopts"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/debug"
)

var (
	// project_id     = flag.String("project", "bluebirz-playground", "project id")
	topic_1        = flag.String("topic_1", "", "topic 1 id")
	subscription_1 = flag.String("subscription_1", "", "subscription 1 id")
	topic_2        = flag.String("topic_2", "", "topic 2 id")
	subscription_2 = flag.String("subscription_2", "", "subscription 2 id")
)

//
// func init() {
// }

func run_beam() {
	flag.Parse()
	// project_id := "bluebirz-playground"
	// topic := "test-topic1"
	// subscription := "test-sub1"
	// if topic == nil {
	// 	log.Fatal("topic can't be nil")
	// }

	beam.Init()
	ctx := context.Background()
	p, s := beam.NewPipelineWithRoot()

	project_id := gcpopts.GetProject(ctx)
	lines := pubsubio.Read(s, project_id, *topic_1, &pubsubio.ReadOptions{WithAttributes: true, Subscription: *subscription_1})
	debug.Print(s, lines)

	pubsubio.Write(s, project_id, *topic_2, lines)
	if err := beamx.Run(ctx, p); err != nil {
		fmt.Printf("error: %v", err)
	}
}

func main() {
	run_beam()
}
