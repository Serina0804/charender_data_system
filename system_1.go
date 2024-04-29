package main

import (
	"fmt"
	"math"
	"net/http"

	"github.com/gorilla/mux"
)

type ScheduleEntry struct {
	Day        int
	Event_name string
	Start_time int
	End_time   int
	Memo       string
	Record     string
}

type Schedule struct {
	Id         int
	Event_data *ScheduleEntry
}

type ScheduleList struct {
	Schedule *Schedule
}

const start_time_digit = 2

// constructor and initializer of ScheduleEntry
func NewScheduleEntry(day int, event_name string, start_time int, end_time int, memo string, record string) *ScheduleEntry {
	s := new(ScheduleEntry)
	s.Day = day
	s.Event_name = event_name
	s.Start_time = start_time
	s.End_time = end_time
	s.Memo = memo
	s.Record = record
	return s
}

func NewSchedule(id int, schedule_entry *ScheduleEntry) *Schedule {
	s := new(Schedule)
	s.Id = id
	s.Event_data = schedule_entry
	return s
}

func MakeId(day int, start_time int) int {
	id := day*int(math.Pow10(start_time_digit)) + start_time
	return id
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RegisterHandler)

	http.Handle("/", r)

	fmt.Println("boot server")
	http.ListenAndServe(":8080", nil)

	// get usr input

	day := 20240429
	event_name := "dev"
	start_time := 1000
	end_time := 1700
	memo := "TODO"
	record := ""

	schedule_entry := NewScheduleEntry(day, event_name, start_time, end_time, memo, record)
	schedule := NewSchedule(MakeId(day, start_time), schedule_entry)
	schedule_list := new(ScheduleList)

	schedule_list.append(schedule)
}
