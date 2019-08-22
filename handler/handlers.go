package handler

import (

	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	m "velociraptorgo/datastore"
	n "velociraptorgo/models"
	"velociraptorgo/utils"
)


// GetTotalVelociraptor returns total number of Velociraptor between +/- 3 years of timestamp
func GetTotalVelociraptor(response http.ResponseWriter, request *http.Request) {
	// Fetching requested timestamp from get
	vars := mux.Vars(request)
	timestamp := vars["timestamp"]

	// Fetching +/- 3years date range, min and max date range
	var daterange = utils.CreateYearRangeBetweenTimeStamp(timestamp)
	var yearrange = daterange.YearRange

	// Fetching Date from store
	var data = m.YearStore
	
	sum := 0
	// Iterating through +/- 3 year range
	for _, year := range yearrange {

		if yearKey, ok := data.Year[year]; ok {
			// If year is present
			monthStore := yearKey

			for _, monthStorevalue := range monthStore.Month {
				velociraptors := monthStorevalue.Velociraptor

				for _, velociraptor := range velociraptors {
					// Checking min and max date range
					if utils.InTimeSpan(daterange.MinDate, daterange.MaxDate, velociraptor.TimeStamp) {
						sum += velociraptor.TotalVelociraptor
					}
				}
			}
		}
	}
	bytes, err := json.Marshal(sum)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.Write(bytes)
}

// UpdateVelociraptor is updates total number of Velociraptor
func UpdateVelociraptor(response http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	veloc := new(n.Velociraptor)
	// Unmarshalling request body
	err = json.Unmarshal(body, veloc)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	m.UpdateDataStore(veloc)

	bytes, err := json.Marshal("Success")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	response.Write(bytes)
}
