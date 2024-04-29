package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// フォームからのデータを取得
		date := r.FormValue("date")
		day := r.FormValue("day")
		event := r.FormValue("event")
		startTime := r.FormValue("start_time")
		endTime := r.FormValue("end_time")
		memo := r.FormValue("memo")
		record := r.FormValue("record")

		// 取得したデータを出力
		fmt.Printf("Received data: date=%s, day=%s, event=%s, start_time=%s, end_time=%s, memo=%s, record=%s\n",
			date, day, event, startTime, endTime, memo, record)
	}

	// フォームのHTMLを読み込む
	tmpl := template.Must(template.ParseFiles("form.html"))

	// HTMLをレンダリングしてクライアントに返す
	tmpl.Execute(w, nil)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// POSTメソッドのみを許可
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// フォームデータを読み取る
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusInternalServerError)
		return
	}

	// フォームからの入力データを取得
	date := r.Form.Get("date")
	day := r.Form.Get("day")
	event := r.Form.Get("event")
	startTime := r.Form.Get("start_time")
	endTime := r.Form.Get("end_time")
	memo := r.Form.Get("memo")
	record := r.Form.Get("record")

	// 取得したデータをターミナルに表示
	fmt.Printf("Received data: date=%s, day=%s, event=%s, start_time=%s, end_time=%s, memo=%s, record=%s\n",
		date, day, event, startTime, endTime, memo, record)

	// クライアントにレスポンスを返す
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Received data successfully")
}

func main() {
	// フォームのHTMLを表示するハンドラを登録
	http.HandleFunc("/form", FormHandler)
	http.HandleFunc("/register", RegisterHandler)

	fmt.Println("Server is running...")
	http.ListenAndServe(":8080", nil)
}
