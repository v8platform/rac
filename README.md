# rac
Исполнитель `rac` от 1С Предприятие


### Пример использования
```go 

package main

import (
	"github.com/k0kubun/pp"
	"github.com/v8platform/rac"
	"log"
)

func main() {

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

```