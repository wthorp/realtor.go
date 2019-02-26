package main

import (
	"flag"
	"fmt"

	"realtor.go/pkg/realtor"
)

func main() {
	var place string
	//get command line options
	flag.StringVar(&place, "place", "Pittsburgh, PA", "The name of the area to search")
	flag.Parse()
	if place == "" {
		flag.PrintDefaults()
		return
	}

	s, err := realtor.NewSession()
	if err != nil {
		fmt.Println("NewSession: " + err.Error())
		return
	}
	p, err := s.FindPlace("Pittsburgh, PA")
	if err != nil {
		fmt.Println("FindPlace: " + err.Error())
		return
	}
	fmt.Printf("%+v", p)
}
