package main

import (
	"fmt"

	"github.com/Kingfish219/PlaNet/internal/startup"
)

func main() {
	startup := startup.New()
	err := startup.Prepare()
	if err != nil {
		fmt.Println(err)

		return
	}

	err = startup.Initialize()
	if err != nil {
		fmt.Println(err)

		return
	}
}
