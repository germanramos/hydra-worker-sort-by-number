package main

import (
	// "encoding/json"
	"os"
	"sort"
	"strconv"

	worker "github.com/innotech/hydra-worker-sort-by-number/vendors/github.com/innotech/hydra-worker-lib"
)

const (
	decr string = "0"
	incr string = "1"
)

var order, sortAttr string

type Instances []map[string]interface{}

func (a Instances) Len() int      { return len(a) }
func (a Instances) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (s Instances) Less(i, j int) bool {
	var less bool
	if order == decr {
		a1, _ := strconv.ParseFloat(s[i]["Info"].(map[string]interface{})[sortAttr].(string), 64)
		a2, _ := strconv.ParseFloat(s[j]["Info"].(map[string]interface{})[sortAttr].(string), 64)
		less = a1 > a2
	} else {
		a1, _ := strconv.ParseFloat(s[i]["Info"].(map[string]interface{})[sortAttr].(string), 64)
		a2, _ := strconv.ParseFloat(s[j]["Info"].(map[string]interface{})[sortAttr].(string), 64)
		less = a1 < a2
	}
	return less
}

func main() {
	if len(os.Args) < 3 {
		panic("Invalid number of arguments, you need to add at least the arguments for the server address and the service name")
	}
	serverAddr := os.Args[1]  // e.g. "tcp://localhost:5555"
	serviceName := os.Args[2] // e.g. sort-by-number
	verbose := len(os.Args) >= 4 && os.Args[3] == "-v"

	// New Worker connected to Hydra Load Balancer
	sortByNumberWorker := worker.NewWorker(serverAddr, serviceName, verbose)
	fn := func(instances []interface{}, args map[string]interface{}) []interface{} {
		var finalInstances []map[string]interface{}
		finalInstances = make([]map[string]interface{}, 0)
		for _, instance := range instances {
			finalInstances = append(finalInstances, instance.(map[string]interface{}))
		}

		sortAttr = args["sortAttr"].(string)
		order = args["order"].(string)
		sort.Sort(Instances(finalInstances))

		var finalInstances2 []interface{}
		finalInstances2 = make([]interface{}, 0)
		for _, instance := range finalInstances {
			finalInstances2 = append(finalInstances2, instance)
		}
		return finalInstances2
	}
	sortByNumberWorker.Run(fn)
}
