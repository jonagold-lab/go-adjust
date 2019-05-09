package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/jonagold-lab/go-adjust/adjust"
)

func main() {
	pw := flag.String("password", "", "Adjust Password")
	email := flag.String("email", "", "Adjust Email")
	appID := flag.String("app_id", "", "Adjust app ID")
	flag.Parse()
	if *pw == "" || *email == "" || *appID == "" {
		panic("please set password, email and app_id flag")
	}
	client, err := adjust.NewClient(nil, *email, *pw, *appID)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	opt := adjust.Options{
		AttributionSource: "dynamic",
		AttributionType:   "click",
		StartDate:         "2019-05-01",
		EndDate:           "2019-05-09",
		UTCOffset:         "00:00",
		EventKpis:         "tam0sc_events,xxwrmf_events,wk9hmc_events,tucsbx_events",
		Kpis:              "installs",
		Sandbox:           false,
		HumanReadableKpis: true,
		Grouping:          "day,networks,campaigns,adgroups,creatives",
		Reattributed:      "all",
	}
	list, _, err := client.KPI.List(context.Background(), &opt)
	if err != nil {
		log.Fatalf("Kpis List error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(&list)
	fmt.Println(string(res))
	fmt.Println("----------------")
}
