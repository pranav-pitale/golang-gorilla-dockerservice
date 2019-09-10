package model

import "time"

// Velociraptor type stores velociraptor Timestamp and total number
type Velociraptor struct {
	TimeStamp         time.Time
	TotalVelociraptor int
}

// YearStore type is a map which stores year of Timestamp and its values points to Month store
type YearStore struct {
	Year map[int]*MonthStore 
}

// MonthStore type is a map which stores month of TimeStamp and its value points to Array of Velociraptors
type MonthStore struct {
	Month map[int]*Velociraptors 
}

// Velociraptors type is array of Velociraptor type
type Velociraptors struct {
	Velociraptor []Velociraptor
}
