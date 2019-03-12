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
	//session
	s, err := realtor.NewSession()
	if err != nil {
		fmt.Println("NewSession: " + err.Error())
		return
	}
	//place
	p, err := s.FindPlace("Pittsburgh, PA")
	if err != nil {
		fmt.Println("FindPlace: " + err.Error())
		return
	}
	fmt.Printf("%+v", p)
	//query
	req := realtor.ListingRequest{
		Facets: realtor.Facets{
			BedsMin:                   4,
			IncludePendingContingency: false,
		},
		PageSize:    200,
		BoundingBox: "-80.3646522705942,40.252238699938296,-79.77860032967624,40.75813885384505",
		Pos:         "40.252239,-80.364652,40.758139,-79.7786,11",
	}
	r, err := s.Search(req)
	if err != nil {
		fmt.Println("FindPlace: " + err.Error())
		return
	}
	fmt.Printf("%+v", r)
}
