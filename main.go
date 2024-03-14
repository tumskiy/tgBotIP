package main

import (
	"fmt"
)

func main() {
	initDB()
	createUser("150951163", "tumsk1y", "tumsk1y")
	site, err := request()
	if err != nil {
		return
	}
	fmt.Println(site.IP + site.City + site.AsnOrg)
}
