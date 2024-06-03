package common

import (
	"fmt"
	"sync"

	"github.com/Kingfish219/PlaNet/network/dns"
)

var lock = &sync.Mutex{}

type PlanetConf struct {
	ActiveDns *dns.Dns
}

var instance *PlanetConf

func GetConfigurations() *PlanetConf {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Creating single instance now.")
			instance = &PlanetConf{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return instance
}
