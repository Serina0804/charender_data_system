package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
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

func MakeTwoDigit(num int) string {
	if num < 10 {
		return "0" + strconv.Itoa(num)
	}
	return strconv.Itoa(num)
}

func MakeId(month int, date int, start_hour int, start_min int) int {
	str_month := MakeTwoDigit(month)
	str_date := MakeTwoDigit(date)
	str_start_hour := MakeTwoDigit(start_hour)
	str_start_min := MakeTwoDigit(start_min)

	str_id := str_month + str_date + str_start_hour + str_start_min

	id, err := strconv.Atoi(str_id)
	if err != nil {
		fmt.Printf("cannot make id\n")
	}
	return id
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	var errorMsgs []string
	var scheduleList []Schedule

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

		db, err := sql.Open("sqlite3", "database.db")
		if err != nil {
			fmt.Printf("cannot open database: %s\n", err)
		}
		defer db.Close()

		_, err = db.Exec(`INSERT INTO schedule (id, month, date, day, eventName, startHour, startMin, endHour, endMin, memo, record) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`,
			MakeId(month, date, start_hour, start_min),
			month,
			date,
			day,
			event_name,
			start_hour,
			start_min,
			end_hour,
			end_min,
			memo,
			record)
		if err != nil {
			fmt.Printf("cannot insert entry into database: %s\n", err)
		}

		rows, err := db.Query(`SELECT id, month, date, day, eventName, startHour, startMin, endHour, endMin, memo, record FROM schedule ORDER BY id;`)
		if err != nil {
			fmt.Printf("cannot query database: %s\n", err)
		}
		defer rows.Close()

		scheduleList = []Schedule{}
		for rows.Next() {
			var s Schedule
			err := rows.Scan(&s.Id, &s.Month, &s.Date, &s.Day, &s.EventName, &s.StartHour, &s.StartMin, &s.EndHour, &s.EndMin, &s.Memo, &s.Record)
			if err != nil {
				fmt.Printf("cannot scan row: %s\n", err)
			}
			scheduleList = append(scheduleList, s)
		}
	}

	for i, v := range scheduleList {
		fmt.Println(i, v)
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

func initializeDatabase() error {
	db, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		return err
	}
	defer db.Close()

	// query that make schedule table
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS schedule (
            id INTEGER PRIMARY KEY,
    		month     INTEGER,
			date      INTEGER,
			day       TEXT,
			eventName TEXT,
			startHour INTEGER,
			startMin  INTEGER,
			endHour   INTEGER,
			endMin    INTEGER,
			memo      TEXT,
			record    TEXT
    	);
    `)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler)

	// initialize database
	if err := initializeDatabase(); err != nil {
		fmt.Println("initializeDatabase: ", err)
		return
	}

	http.Handle("/", r)
	fmt.Println("boot server")
	http.ListenAndServe(":8080", nil)
}
