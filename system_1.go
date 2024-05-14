package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	Monday = 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

type ScheduleEntry struct {
	Month           int
	Date            int
	Day             int
	Event_name      string
	Start_time_hour int
	Start_time_min  int
	End_time_hour   int
	End_time_min    int
	Memo            string
	Record          string
}

type Schedule struct {
	Id         int
	Event_data *ScheduleEntry
}

func dayOfWeek(day string) int {
	switch day {
	case "月":
	case "Mon":
	case "mon":
		return Monday

	case "火":
	case "Tue":
	case "tue":
		return Tuesday

	case "水":
	case "Wed":
	case "wed":
		return Wednesday

	case "木":
	case "Thr":
	case "thr":
		return Thursday

	case "金":
	case "Fri":
	case "fri":
		return Friday

	case "土":
	case "Sat":
	case "sat":
		return Saturday

	case "日":
	case "sun":
	case "Sun":
		return Sunday

	}
	return -1
}

// constructor and initializer of ScheduleEntry
func NewScheduleEntry(month int, date int, day string, event_name string, start_time_hour int, start_time_min int, end_time_hour int, end_time_min int, memo string, record string) *ScheduleEntry {
	s := new(ScheduleEntry)
	s.Month = month
	s.Date = date
	s.Event_name = event_name
	s.Start_time_hour = start_time_hour
	s.Start_time_min = start_time_min
	s.End_time_hour = end_time_hour
	s.End_time_min = end_time_min
	s.Memo = memo
	s.Record = record

	s.Day = dayOfWeek(day)
	if s.Day == -1 {
		fmt.Printf("day of the week is wrong\n")
		fmt.Printf("	Monday.. 月/Mon/mon\n")
		fmt.Printf("	Tuesday.. 火/Tue/tue\n")
		fmt.Printf("	Wednesday.. 水/Wed/wed\n")
		fmt.Printf("	Thursday.. 木/Thr/thr\n")
		fmt.Printf("	Friday.. 金/Fri/fri\n")
		fmt.Printf("	Saturday.. 土/Sat/sat\n")
		fmt.Printf("	Sunday.. 日/Sun/sun\n")
	}

	return s
}

func NewSchedule(id int, schedule_entry *ScheduleEntry) *Schedule {
	s := new(Schedule)
	s.Id = id
	s.Event_data = schedule_entry
	return s
}

func MakeId(str_date string, str_start_time string) int {
	str_id := str_date + str_start_time

	id, err := strconv.Atoi(str_id)
	if err != nil {
		fmt.Printf("cannot make id\n")
	}
	return id
}

var scheduleList []*Schedule

func mainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	if r.Method == "POST" {

		// get input
		str_date := r.FormValue("date") // ex. 20240409
		day := r.FormValue("day")       // mon, tue, wed, thr, fri
		event_name := r.FormValue("event")
		str_start_time := r.FormValue("start_time")
		str_end_time := r.FormValue("end_time")
		memo := r.FormValue("memo")
		record := r.FormValue("record")

		// check input type
		date, err := strconv.Atoi(str_date)
		if err != nil {
			fmt.Printf("date must be number\n")
		}
		start_time, err := strconv.Atoi(str_start_time)
		if err != nil {
			fmt.Printf("start_time must be number\n")
		}
		end_time, err := strconv.Atoi(str_end_time)
		if err != nil {
			fmt.Printf("end_time must be number\n")
		}

		// make schedule entry
		schedule_entry := NewScheduleEntry(date, day, event_name, start_time, end_time, memo, record)
		schedule := NewSchedule(MakeId(str_date, str_start_time), schedule_entry)

		// append list
		scheduleList = append(scheduleList, schedule)
		sort.SliceStable(scheduleList, func(i, j int) bool {
			return scheduleList[i].Id < scheduleList[j].Id
		})

		// check sort
		for i, v := range scheduleList {
			fmt.Println(i, v)
			// fmt.Printf("date: %s, day: %s, event_name: %s, start_time: %s, end_time: %s, memo: %s, record: %s\n", str_date, day, event_name, str_start_time, str_end_time, memo, record)
		}
	}

	tmpl.Execute(w, nil)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler)
	http.Handle("/", r)
	fmt.Println("boot server")
	http.ListenAndServe(":8080", nil)
}
