package utils

import (
	"time"
	"testing"

)

func TestCreateYearRangeBetweenTimeStampMinDateCheck(t *testing.T) {
	testdatestring:="2014-10-08T19:02:17-08:00"
	daterange:=CreateYearRangeBetweenTimeStamp("2014-10-08T19:02:17-08:00")
	mindate,_:=time.Parse(time.RFC3339,"2011-10-08T19:02:17-08:00")

	if daterange.MinDate.Sub(mindate) !=0 {
		t.Errorf("Min date(%q) == %q, want %q", testdatestring, daterange.MinDate , mindate)
	}
}
func TestCreateYearRangeBetweenTimeStampMaxDateCheck(t *testing.T) {
	testdatestring:="2014-10-08T19:02:17-08:00"
	daterange:=CreateYearRangeBetweenTimeStamp("2014-10-08T19:02:17-08:00")
	maxdate,_:=time.Parse(time.RFC3339,"2017-10-08T19:02:17-08:00")

	if daterange.MaxDate.Sub(maxdate) !=0 {
		t.Errorf("Min date(%q) == %q, want %q", testdatestring, daterange.MaxDate , maxdate)
	}
}
func TestCreateYearRangeBetweenTimeStampYearRangeLength(t *testing.T) {
	testdatestring:="2014-10-08T19:02:17-08:00"
	daterange:=CreateYearRangeBetweenTimeStamp(testdatestring)
	
	if len(daterange.YearRange) !=7 {
		t.Errorf("Failed to create range of years")
	}
}
func TestInTimeSpan(t *testing.T) {
	mindate,_:= time.Parse(time.RFC3339,"2011-10-08T19:02:17-08:00")
	maxdate, _ := time.Parse(time.RFC3339,"2017-10-08T19:02:17-08:00")
	intime,_ := time.Parse(time.RFC3339,"2015-10-08T19:02:17-08:00")
	if InTimeSpan(mindate,maxdate,intime) == false{
		t.Errorf("Test failed to check InTime Span")
	}
}