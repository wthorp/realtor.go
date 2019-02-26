package scrape

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"realtor.go/pkg/realtor"
)

func main() {
	var place
	//get command line options
	flag.StringVar(&wms, "place", "Pittsburgh, PA", "The name of the area to search")
	flag.Parse()
	if place == "" {
		flag.PrintDefaults()
		return
	}

	s := realtor.NewSession()
	p, err := s.FindPlace("Pittsburgh, PA")
	if err {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%+v", p)
}