package handler

import (
	"bytes"
	"time"
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/gorilla/mux"

	n "velociraptorgo/models"
	m "velociraptorgo/datastore"
)

func TestGetVelociraptorHandler(t *testing.T) {

	timerange,_ := time.Parse(time.RFC3339, "2014-11-08T19:02:17-08:00")
	v := new(n.Velociraptor)
	v.TimeStamp = timerange
	v.TotalVelociraptor = 2

	var YearStore = new(n.YearStore)
	YearStore.Year = make(map[int]*n.MonthStore)

	velociraptors := new(n.Velociraptors)
	velociraptors.Velociraptor = append(velociraptors.Velociraptor, *v)
	monthStore := new(n.MonthStore)
	monthStore.Month = make(map[int]*n.Velociraptors)
	YearStore.Year = make(map[int]*n.MonthStore)
	monthStore.Month[11] = velociraptors
	YearStore.Year[2014] = monthStore

	timerangeNew,_ := time.Parse(time.RFC3339, "2014-09-08T19:02:17-08:00")
	ve := new(n.Velociraptor)
	ve.TimeStamp = timerangeNew
	ve.TotalVelociraptor = 2

	monthStoreNew := new(n.MonthStore)
	monthStoreNew.Month = make(map[int]*n.Velociraptors)
	velociraptorsNew := new(n.Velociraptors)
	velociraptorsNew.Velociraptor = append(velociraptorsNew.Velociraptor, *ve)
	monthStoreNew.Month[9] = velociraptorsNew
	YearStore.Year[2014].Month[9] = monthStoreNew.Month[9]

	m.YearStore = YearStore

	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api").Subrouter()
	sub.Methods("GET").Path("/gettotalvelociraptor/{timestamp}").HandlerFunc(GetTotalVelociraptor)
	req, _ := http.NewRequest("GET", "api/gettotalvelociraptor/2014-10-08T19:02:17-08:00", nil)

	vars := map[string]string{
        "timestamp": "2014-10-08T19:02:17-08:00",
    }

    req = mux.SetURLVars(req, vars)
	res := httptest.NewRecorder()

	GetTotalVelociraptor(res, req)
	result := res.Body.String()
	
	if result != "4"  {
		t.Error("Test failed")
	}
	if res.Code !=200{
			t.Error("Test failed")
	}

}

func TestUpdateVelociraptorHandler(t *testing.T) {

	var YearStore = new(n.YearStore)
	YearStore.Year = make(map[int]*n.MonthStore)

	timerange,_ := time.Parse(time.RFC3339, "2014-11-08T19:02:17-08:00")
	v := new(n.Velociraptor)
	v.TimeStamp = timerange
	v.TotalVelociraptor = 2
	// Creating data store
	velociraptors := new(n.Velociraptors)
	velociraptors.Velociraptor = append(velociraptors.Velociraptor, *v)
	monthStore := new(n.MonthStore)
	monthStore.Month = make(map[int]*n.Velociraptors)
	YearStore.Year = make(map[int]*n.MonthStore)
	monthStore.Month[11] = velociraptors
	YearStore.Year[2014] = monthStore

	// Creating exact timestamp match for update
	veloc := new(n.Velociraptor)
	veloc.TimeStamp = timerange
	veloc.TotalVelociraptor = 4
	m.YearStore = YearStore
	
	payload := []byte(`{"TimeStamp":"2014-11-01T19:02:17-08:00","TotalVelociraptor": 50}`)

	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api").Subrouter()
	sub.Methods("POST").Path("/updatevelociraptor").HandlerFunc(UpdateVelociraptor)
	req, _ := http.NewRequest("POST", "/updatevelociraptor", bytes.NewBuffer(payload))
	res := httptest.NewRecorder()
	UpdateVelociraptor(res,req)

	if YearStore.Year[2014].Month[11].Velociraptor[0].TotalVelociraptor != 50  {
		t.Error("Test failed")
	}

	if res.Code !=200{
			t.Error("Test failed")
	}
}