package datastore

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"math"
	"sort"

	n "velociraptorgo/models"
	"velociraptorgo/utils"
)

// YearStore needs to be accessed by handlers
var YearStore = new(n.YearStore)

// LoadDataStore creates loads the data from file and creates in memory data store
func LoadDataStore() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	csvfile, err := os.Open("data/timeline_data.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()
	r := csv.NewReader(csvfile)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("error reading all lines: %v", err)
	}

	YearStore.Year = make(map[int]*n.MonthStore)

	// Iterating line by line through file
	for i, record := range lines {
		if i == 0 {
			// skip header line
			continue
		}

		v := new(n.Velociraptor)
		t, _ := time.Parse(time.RFC3339, record[0])
		year := t.Year()
		month := int(t.Month())

		v.TimeStamp = t
		v.TotalVelociraptor, err = strconv.Atoi(record[1])

		if yearKey, ok := YearStore.Year[year]; ok {
			// If year is present in year store
			monthStore := yearKey
			// If month is present in month store
			if monthKey, ok := monthStore.Month[month]; ok {
				var velociraptors = monthKey
				velociraptors.Velociraptor = append(velociraptors.Velociraptor, *v)

			} else {
				// If year is present and month is not present
				monthStore := new(n.MonthStore)
				monthStore.Month = make(map[int]*n.Velociraptors)
				velociraptors := new(n.Velociraptors)
				velociraptors.Velociraptor = append(velociraptors.Velociraptor, *v)
				monthStore.Month[month] = velociraptors
				YearStore.Year[year].Month[month] = monthStore.Month[month]
			}
		} else {
			// If year is not present, create year store
			velociraptors := new(n.Velociraptors)
			velociraptors.Velociraptor = append(velociraptors.Velociraptor, *v)
			monthStore := new(n.MonthStore)
			monthStore.Month = make(map[int]*n.Velociraptors)
			monthStore.Month[month] = velociraptors
			YearStore.Year[year] = monthStore
		}
	}
	// fmt.Println(YearStore.Year[2013].Month[12].Velociraptor[0].Total)
	return
}
// UpdateDataStore updates velociraptor data store 
func UpdateDataStore(veloc *n.Velociraptor){
	year := veloc.TimeStamp.Year()
	month := int(veloc.TimeStamp.Month())

	// Update flag checks whether total is updated by exact timestamp logic
	updateflag := false
	// Fetching data store
	var data = YearStore
	// Exact time stamp match logic
	// Checking whether year is present in Year store
	if yearKey, ok := data.Year[year]; ok {
		
		monthStore := yearKey
		// Checking whether month is present in month store
		if monthKey, ok := monthStore.Month[month]; ok {
			var velociraptors = monthKey
			for index, velociraptor := range velociraptors.Velociraptor {
				if int(velociraptor.TimeStamp.Sub(veloc.TimeStamp)) == 0 {
					velociraptors.Velociraptor[index].TotalVelociraptor = veloc.TotalVelociraptor
					updateflag = true
				}
			}
		}
	}
	if updateflag == false {
		// Generating +/- 3 years year range and min and max daterange if there is no exact match
		daterange := utils.CreateYearRangeBetweenTimeStamp(veloc.TimeStamp.Format(time.RFC3339))
		yearrange := daterange.YearRange
		// Creating maps which will store dates in data store between min max date range and its value will date diff between dates and given timestamp
		closestdistance := make(map[string]int)
		// Creating maps which will store dates in data store between min max date range and its value will pointer to Velociraptors
		closeststore := make(map[string]*n.Velociraptors)
		// Creating maps which will store dates in data store between min max date range and its value will be index of Velociraptor
		closestindex := make(map[string]int)

		// Iterating for +/- 3 years range
		for _, year := range yearrange {
			if _, ok := data.Year[year]; ok {
				monthStore := data.Year[year]

				for _, month := range monthStore.Month {
					velociraptors := month.Velociraptor
					for index, velociraptor := range velociraptors {
						if utils.InTimeSpan(daterange.MinDate, daterange.MaxDate, velociraptor.TimeStamp) {
							// Calculating date diff between time stamp and request time stamp and storing it in map
							diff := float64(int(velociraptor.TimeStamp.Sub(veloc.TimeStamp).Hours()) / 24)
							// Adding dates against datediff in map
							closestdistance[velociraptor.TimeStamp.String()] = int(math.Abs(diff))
							// Adding dates against *Velociraptor to map
							closeststore[velociraptor.TimeStamp.String()] = month
							// Adding dates against index of Velociraptor to map
							closestindex[velociraptor.TimeStamp.String()] = index
						}
					}
				}
			}
		}
		// Fetching minimumn closest time stamp 
		keys := make([]string, 0, len(closestdistance))
		for k := range closestdistance {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		// getting minimum closest date from timestamp
		mindistancekey := keys[len(keys)-1]
		minindex := closestindex[mindistancekey]
		if _, ok := closeststore[mindistancekey]; ok {
			// Updating total
			closeststore[mindistancekey].Velociraptor[minindex].TotalVelociraptor = veloc.TotalVelociraptor
		}
	}
}