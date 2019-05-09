package adjust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var kpiTest = []struct {
	name    string
	path    string
	fixture string
	opts    *Options
}{
	{"No Options", "/kpis/v1/random.json", "kpis.json", nil},
	{"With Options", "/kpis/v1/random.json", "kpis_options.json", &Options{
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
		Reattributed:      "all"},
	},
}

func TestKPIService_List(t *testing.T) {
	for _, tt := range kpiTest {
		t.Run(tt.name, func(t *testing.T) {
			client, mux, _, teardown := setup()
			t.Log(fmt.Sprintf("Setup Done %s", tt.name))
			defer teardown()
			mux.HandleFunc(tt.path, func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				w.WriteHeader(http.StatusOK)
				w.Write(loadFixture(tt.fixture))
			})
			got, _, err := client.KPI.List(context.Background(), tt.opts)
			if err != nil {
				t.Errorf("KPI.List returned error: %v", err)
			}
			want := &KPI{}
			responseToInterface(loadFixture(tt.fixture), &want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("KPI.List = %+v, want %+v", got, want)
			}
		})
	}
}
