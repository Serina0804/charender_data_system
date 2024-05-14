package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

type Schedule struct {
	Id        int
	Date      int
	Day       string
	EventName string
	StartTime int
	EndTime   int
	Memo      string
	Record    string
}

func MakeId(date int, start_time int) int {
	str_id := strconv.Itoa(date) + strconv.Itoa(start_time)

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

		scheduleList = append(scheduleList, &Schedule{
			Id:        MakeId(date, start_time),
			Date:      date,
			Day:       day,
			EventName: event_name,
			StartTime: start_time,
			EndTime:   end_time,
			Memo:      memo,
			Record:    record,
		})

		sort.SliceStable(scheduleList, func(i, j int) bool {
			return scheduleList[i].Id < scheduleList[j].Id
		})

		// check sort
		for i, v := range scheduleList {
			fmt.Println(i, v)
			// fmt.Printf("date: %s, day: %s, event_name: %s, start_time: %s, end_time: %s, memo: %s, record: %s\n", str_date, day, event_name, str_start_time, str_end_time, memo, record)
		}
	}

	tmpl.Execute(w, scheduleList)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler)
	// r.HandleFunc("/accept", acceptHandler)
	http.Handle("/", r)
	fmt.Println("boot server")
	http.ListenAndServe(":8080", nil)
}
