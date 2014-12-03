//
//  UDP ping command
//  Model 3, uses abstract network interface
//

package main

import (
	"github.com/innotech/hydra-worker-sort-by-number/vendors/github.com/innotech/hydra-worker-lib/vendors/github.com/pebbe/zmq4/examples/intface"

	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	iface := intface.New()
	for {
		msg, err := iface.Recv()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%q\n", msg)
	}
}
