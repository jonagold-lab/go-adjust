package adjust

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestKPIService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	t.Log("Setup Done")
	defer teardown()

	mux.HandleFunc("/kpis/v1/random.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusOK)
		w.Write(loadFixture("kpis.json"))
	})
	opt := Options{}
	got, _, err := client.KPI.List(context.Background(), &opt)
	if err != nil {
		t.Errorf("KPI.List returned error: %v", err)
	}

	want := &KPI{}
	responseToInterface(loadFixture("kpis.json"), &want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("KPI.List = %+v, want %+v", got, want)
	}
}

func TestKPIService_ListWithOptions(t *testing.T) {
	client, mux, _, teardown := setup()
	t.Log("Setup Done")
	defer teardown()

	mux.HandleFunc("/kpis/v1/random.json", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusOK)
		w.Write(loadFixture("kpis_options.json"))
	})
	opt := Options{
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
	got, _, err := client.KPI.List(context.Background(), &opt)
	if err != nil {
		t.Errorf("KPI.List returned error: %v", err)
	}

	want := &KPI{}
	responseToInterface(loadFixture("kpis_options.json"), &want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("KPI.List = %+v, want %+v", got, want)
	}
}
