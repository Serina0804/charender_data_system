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
	Month     int
	Date      int
	Day       string
	EventName string
	StartHour int
	StartMin  int
	EndHour   int
	EndMin    int
	Memo      string
	Record    string
}

func MakeId(month int, str_date string, str_start_hour string, str_start_min string) int {
	str_id := strconv.Itoa(month) + str_date + str_start_hour + str_start_min

	id, err := strconv.Atoi(str_id)
	if err != nil {
		fmt.Printf("cannot make id\n")
	}
	return id
}

var scheduleList []Schedule

func mainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	var errorMsgs []string

	if r.Method == "POST" {

		// get input
		str_month := r.FormValue("month") // 05
		str_date := r.FormValue("date")   // 02
		day := r.FormValue("day")         // mon, tue, wed, thr, fri
		event_name := r.FormValue("event")
		str_start_hour := r.FormValue("start_hour") // 10
		str_start_min := r.FormValue("start_min")   // 00
		str_end_hour := r.FormValue("end_hour")     // 20
		str_end_min := r.FormValue("end_min")       //00
		memo := r.FormValue("memo")
		record := r.FormValue("record")

		// check input type
		month, err := strconv.Atoi(str_month)
		if err != nil {
			errorMsgs = append(errorMsgs, "month must be number")
		}
		date, err := strconv.Atoi(str_date)
		if err != nil {
			errorMsgs = append(errorMsgs, "date must be number")
		}
		start_hour := 0
		if str_start_hour != "" {
			start_hour, err = strconv.Atoi(str_start_hour)
			if err != nil {
				errorMsgs = append(errorMsgs, "start hour must be number")
			}
		}
		start_min := 0
		if str_start_min != "" {
			start_min, err = strconv.Atoi(str_start_min)
			if err != nil {
				errorMsgs = append(errorMsgs, "start minute must be number")
			}
		}

		end_hour := 0
		if str_end_hour != "" {
			end_hour, err = strconv.Atoi(str_end_hour)
			if err != nil {
				errorMsgs = append(errorMsgs, "end hour must be number")
			}
		}

		end_min := 0
		if str_end_min != "" {
			end_min, err = strconv.Atoi(str_end_min)
			if err != nil {
				errorMsgs = append(errorMsgs, "end minute must be number")
			}
		}

		// Check if there are any error messages
		if len(errorMsgs) > 0 {
			// If there are errors, render the template with the error messages
			tmpl.Execute(w, struct {
				ScheduleList []Schedule
				ErrorMsgs    []string
			}{
				ScheduleList: scheduleList,
				ErrorMsgs:    errorMsgs,
			})
			return
		}

		scheduleList = append(scheduleList, Schedule{
			Id:        MakeId(month, str_date, str_start_hour, str_start_min),
			Month:     month,
			Date:      date,
			Day:       day,
			EventName: event_name,
			StartHour: start_hour,
			StartMin:  start_min,
			EndHour:   end_hour,
			EndMin:    end_min,
			Memo:      memo,
			Record:    record,
		})

		sort.SliceStable(scheduleList, func(i, j int) bool {
			return scheduleList[i].Id < scheduleList[j].Id
		})

		// check sort
		for i, v := range scheduleList {
			fmt.Println(i, v)
		}
	}

	// Render the template with the schedule list and error messages
	tmpl.Execute(w, struct {
		ScheduleList []Schedule
		ErrorMsgs    []string
	}{
		ScheduleList: scheduleList,
		ErrorMsgs:    nil,
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler)

	http.Handle("/", r)
	fmt.Println("boot server")
	http.ListenAndServe(":8080", nil)
}
