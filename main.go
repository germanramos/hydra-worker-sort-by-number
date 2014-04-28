package main

import (
	"encoding/json"
	"os"
	"sort"

	worker "github.com/innotech/hydra-worker-pong/vendors/github.com/innotech/hydra-worker-lib"
)

const (
	DECR int = 0
	INCR int = 1
)

type lessByAttr func(i, j int) bool

type ByAttr []map[string]interface{}

func (a ByAttr) Len() int           { return len(a) }
func (a ByAttr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAttr) Less(i, j int) bool { return lessByAttr(i, j) }

func main() {
	if len(os.Args) < 3 {
		panic("Invalid number of arguments, you need to add at least the arguments for the server address and the service name")
	}
	serverAddr := os.Args[1]  // e.g. "tcp://localhost:5555"
	serviceName := os.Args[2] // e.g. sort-by-number
	verbose := len(os.Args) >= 4 && os.Args[3] == "-v"

	// New Worker connected to Hydra Load Balancer
	mapAndSortWorker := worker.NewWorker(serverAddr, serviceName, verbose)
	fn := func(instances []map[string]interface{}, args map[string]string) []interface{} {
		order := args["order"]
		attr := args["sortAttr"]

		if order == DECR {
			lessByAttr = func(i, j int) bool {
				return a[i][attr].(float64) > a[j][attr].(float64)
			}
		} else {
			lessByAttr = func(i, j int) bool {
				return a[i][attr].(float64) < a[j][attr].(float64)
			}
		}

		sort.Sort(ByAttr(instances))

		return instances
	}
	mapAndSortWorker.Run(fn)
}
