package main

import (
	"log"
	"os"
	"strconv"

	"github.com/avocatl/mmk/bm"
)

func main() {
	mmk, err := bm.NewClient(nil)
	if err != nil {
		log.Fatal(err)
	}

	cid, err := strconv.Atoi(os.Getenv("MMK_COMPANY_ID"))
	if err != nil {
		log.Fatal(err)
	}

	opts := bm.AvailabilityOptions{
		CompanyID: cid,
	}

	res, err := mmk.Availability.GetAvailability(2022, &opts)
	if err != nil {
		log.Fatal(err)
	}

	for i, vv := range res {
		if i%200 == 0 {
			log.Printf("boat: %d\t| from: %s\t| to: %s\t| status: %d\t| option expires:%s\t|\n", vv.YachtID, vv.DateFrom, vv.DateTo, vv.Status, vv.OptionExpirationDate)
		}
	}

	log.Printf("retrieved %d availability records from mmk", len(res))
}
