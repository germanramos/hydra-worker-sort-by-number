package main

import (
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
		if _, ok := s[i]["Info"].(map[string]interface{})[sortAttr]; !ok {
			less = true
		} else if _, ok := s[j]["Info"].(map[string]interface{})[sortAttr]; !ok {
			less = false
		} else {
			a1, _ := strconv.ParseFloat(s[i]["Info"].(map[string]interface{})[sortAttr].(string), 64)
			a2, _ := strconv.ParseFloat(s[j]["Info"].(map[string]interface{})[sortAttr].(string), 64)
			less = a1 > a2
		}
	} else {
		if _, ok := s[i]["Info"].(map[string]interface{})[sortAttr]; !ok {
			less = true
		} else if _, ok := s[j]["Info"].(map[string]interface{})[sortAttr]; !ok {
			less = false
		} else {
			a1, _ := strconv.ParseFloat(s[i]["Info"].(map[string]interface{})[sortAttr].(string), 64)
			a2, _ := strconv.ParseFloat(s[j]["Info"].(map[string]interface{})[sortAttr].(string), 64)
			less = a1 < a2
		}
	}
	return less
}

func main() {
	// New Worker connected to Hydra Load Balancer
	sortByNumberWorker, _ := worker.NewWorker(os.Args)
	fn := func(instances []interface{}, clientParams map[string][]string, args map[string]interface{}) []interface{} {
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
