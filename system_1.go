package charender_data_system

import (
	//"database/sql"
	"fmt"
	"html/template"
	"net/http"
	// "log"      // 追加
	// "io/ioutil" // 追加

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type Timetable struct {
	Day     string
	Period  string
	Course  string
	Teacher string
}

var (
	timetable []Timetable
	store     = sessions.NewCookieStore([]byte("secret"))
	users     = map[string]string{
		"alice": "password123",
		"bob":   "password456",
	}
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username, ok := session.Values["username"].(string)
	if !ok || username == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, timetable)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		//if authenticateUser(username, password) {
		//	session.Values["username"] = username
		//	session.Save(r, w)
		//
		//	http.Redirect(w, r, "/", http.StatusFound)
		//	return
		//} else {
		//	http.Error(w, "認証エラー", http.StatusUnauthorized)
		//	return
		//}
	}

	tmpl := template.Must(template.ParseFiles("login.html"))
	tmpl.Execute(w, nil)
}

//func createUser(username, password string) error {
//	db, err := sql.Open("sqlite3", "database.db")
//	if err != nil {
//		return err
//	}
//	defer db.Close()
//
//	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

//func authenticateUser(username, password string) bool {
//	db, err := sql.Open("sqlite3", "database.db")
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//	defer db.Close()
//
//	var storedPassword string
//	err = db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//
//	return storedPassword == password
//}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// ユーザー名の重複をチェック
		//exists, err := checkUserExists(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if exists {
			http.Error(w, "ユーザー名が既に存在します", http.StatusBadRequest)
			return
		}

		// データベースにユーザーを追加
		err = createUser(username, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// 登録完了後、ログインページにリダイレクト
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("register.html"))
	tmpl.Execute(w, nil)
}

//func checkUserExists(username string) (bool, error) {
//	db, err := sql.Open("sqlite3", "database.db")
//	if err != nil {
//		return false, err
//	}
//	defer db.Close()
//
//	var count int
//	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
//	if err != nil {
//		return false, err
//	}
//
//	return count > 0, nil
//}

//// データベースに時間割のエントリを追加する関数
//func addTimetableEntry(day, period, course, teacher string) error {
//	db, err := sql.Open("sqlite3", "database.db")
//	if err != nil {
//		return err
//	}
//	defer db.Close()
//
//	_, err = db.Exec("INSERT INTO timetable (day, period, course, teacher) VALUES (?, ?, ?, ?)", day, period, course, teacher)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

// データベースから時間割のエントリを取得する関数
//func getTimetableEntries() ([]Timetable, error) {
//	var timetable []Timetable
//
//	db, err := sql.Open("sqlite3", "database.db")
//	if err != nil {
//		return timetable, err
//	}
//	defer db.Close()
//
//	rows, err := db.Query("SELECT day, period, course, teacher FROM timetable")
//	if err != nil {
//		return timetable, err
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var entry Timetable
//		if err := rows.Scan(&entry.Day, &entry.Period, &entry.Course, &entry.Teacher); err != nil {
//			return timetable, err
//		}
//		timetable = append(timetable, entry)
//	}
//
//	if err := rows.Err(); err != nil {
//		return timetable, err
//	}
//
//	return timetable, nil
//}

// 時間割のエントリを追加するハンドラ
//func addTimetableEntryHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method != http.MethodPost {
//		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//		return
//	}
//
//	day := r.FormValue("day")
//	period := r.FormValue("period")
//	course := r.FormValue("course")
//	teacher := r.FormValue("teacher")
//
//	// データベースに時間割のエントリを追加
//	err := addTimetableEntry(day, period, course, teacher)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	// 成功した場合は、リダイレクトまたは適切な応答を返す
//	http.Redirect(w, r, "/", http.StatusFound)
//}
//
//// 時間割を表示するハンドラ
//func timetableHandler(w http.ResponseWriter, r *http.Request) {
//	// データベースから時間割のエントリを取得
//	timetable, err := getTimetableEntries()
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	// テンプレートにデータを渡して表示
//	tmpl := template.Must(template.ParseFiles("timetable.html"))
//	tmpl.Execute(w, timetable)
//}

// データベースの初期化
//func initializeDatabase() error {
//	db, err := sql.Open("sqlite3", "database.db")
//	if err != nil {
//		return err
//	}
//	defer db.Close()
//
//	// timetableテーブルを作成するクエリ
//	_, err = db.Exec(`
//        CREATE TABLE IF NOT EXISTS timetable (
//            id INTEGER PRIMARY KEY AUTOINCREMENT,
//            day TEXT NOT NULL,
//            period TEXT NOT NULL,
//            course TEXT NOT NULL,
//            teacher TEXT NOT NULL
//        );
//    `)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", loginHandler)
	r.HandleFunc("/index", indexHandler)
	// r.HandleFunc("/course", courseHandler)
	//r.HandleFunc("/add", addTimetableEntryHandler) // 時間割のエントリを追加するハンドラ
	r.HandleFunc("/register", registerHandler) // 新規登録ハンドラを追加
	//r.HandleFunc("/timetable", timetableHandler) // 時間割を表示するハンドラ
	// r.HandleFunc("/add", addTimetableEntryHandler) // 時間割のエントリを追加するハンドラ

	http.Handle("/", r)
	// データベースの初期化
	//if err := initializeDatabase(); err != nil {
	//	fmt.Println("データベースの初期化エラー:", err)
	//	return
	//}

	fmt.Println("サーバーを起動しています...")
	http.ListenAndServe(":8080", nil)
}
