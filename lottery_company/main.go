package main

import (
	"fmt"
	"lottery_company/route"
)

func main() {
	r := route.SetRoute()
	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
