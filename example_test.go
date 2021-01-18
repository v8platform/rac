package rac_test

import (
	"github.com/k0kubun/pp"
	"github.com/v8platform/rac"
	"log"
)

func ExampleManager_InfobasesList() {

	manager, err := rac.NewManager("localhost")

	if err != nil {
		log.Fatal(err)
	}

	clusters, err := manager.Clusters(rac.ClustersList{})

	if err != nil {
		log.Fatal(err)
	}

	pp.Println(clusters)

	infobases, err := manager.InfobasesList()

	if err != nil {
		log.Fatal(err)
	}

	pp.Println(infobases)

}
