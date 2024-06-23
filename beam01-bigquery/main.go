package main

import (
	"context"
	"flag"
	"fmt"
	"reflect"

	// "runtime/debug"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"

	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/bigqueryio"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/local"

	// "github.com/apache/beam/sdks/v2/go/pkg/beam/options/gcpopts"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/runners/prism"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/debug"
)

type bqres struct {
	// export fields by capitalizing field names **ALWAYS** or get err "type has unexported field"
	State_name           string `bigquery:"state_name"`
	Total_confirmed_case int64  `bigquery:"total_confirmed_case"`
}

func init() {
	beam.RegisterType(reflect.TypeOf((*bqres)(nil)))
}

func main() {
	flag.Parse()
	beam.Init()
	// ctx := context.Background()
	p, s := beam.NewPipelineWithRoot()
	// lines := textio.Read(s, "protocol://path/file*.txt")
	// lines := textio.Read(s, "test.txt")
	project_id := "bluebirz-playground"
	// project_id := gcpopts.GetProject(ctx)
	query_str := `
    SELECT
      state_name,
      SUM(confirmed_cases) AS total_confirmed_case
    FROM
      bigquery-public-data.covid19_nyt.us_states
    GROUP BY
      1
  `
	query_results := bigqueryio.Query(s, project_id, query_str, reflect.TypeOf(bqres{}), bigqueryio.UseStandardSQL())
	debug.Print(s, query_results)
	target_table := fmt.Sprintf("%s:%s", project_id, "test_dataset.test_tb01")
	// target_rows, err := json.Marshal(query_results)
	// if err != nil {
	// fmt.Println(err)
	// }
	bigqueryio.Write(s, project_id, target_table, query_results)

	if _, err := prism.Execute(context.Background(), p); err != nil {
		fmt.Printf("Pipeline failed: %v", err)
	}
}
