package datastore

import (
	"fmt"
	"testing"
	 "time"
	 n "velociraptorgo/models"
)

func TestUpdateDataStoreforExactMatch(t *testing.T) {

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

	// Updating data store
	UpdateDataStore(veloc)

	if YearStore.Year[2014].Month[11].Velociraptor[0].TotalVelociraptor !=4 {
		fmt.Println("Total is ",YearStore.Year[2014].Month[11].Velociraptor[0].TotalVelociraptor )
		t.Errorf("Test failed to update")
	}
}

func TestUpdateDataStoreClosestDateRangeUpdate(t *testing.T) {
	// Creating data store
	timerange,_ := time.Parse(time.RFC3339, "2014-11-08T19:02:17-08:00")
	v := new(n.Velociraptor)
	v.TimeStamp = timerange
	v.TotalVelociraptor = 2
	
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

	veloc := new(n.Velociraptor)
	veloc.TimeStamp,_ = time.Parse(time.RFC3339, "2014-11-09T19:02:17-08:00")
	veloc.TotalVelociraptor = 4

	UpdateDataStore(veloc)

	if YearStore.Year[2014].Month[11].Velociraptor[0].TotalVelociraptor !=4{
		fmt.Println("Total is ",YearStore.Year[2014].Month[11].Velociraptor[0].TotalVelociraptor )
		t.Errorf("Test failed to update")
	}
}