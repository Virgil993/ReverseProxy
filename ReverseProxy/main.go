package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// type ToDoReading struct {
// 	UserId    int    `json:"userid"`
// 	Id        int    `json:"id"`
// 	Title     string `json:"title"`
// 	Completed bool   `json:"completed"`
// }

func insertReq(db *sql.DB, route string) {
	var ins *sql.Stmt
	ins, errIns := db.Prepare("INSERT into `proxy_requests`.`requests`(`request_url`,`time`) VALUES (?,?)")
	if errIns != nil {
		log.Fatal(errIns)
	}
	defer ins.Close()
	dt := time.Now()

	_, errRes := ins.Exec(route, dt.Format("01-02-2006 15:04:04"))
	if errRes != nil {
		log.Fatal(errRes)
	}
}

func insertResp(db *sql.DB, host string, body string) {

	var insResponse *sql.Stmt
	insResponse, errIns := db.Prepare("INSERT into `proxy_requests`.`responses`(`Response_Body`,`Time`,`Host`) VALUES (?,?,?)")
	if errIns != nil {
		log.Fatal(errIns)
	}
	defer insResponse.Close()
	dt2 := time.Now()

	_, errRes := insResponse.Exec(body, dt2.Format("01-02-2006 15:04:04"), host)
	if errRes != nil {
		log.Fatal(errRes)
	}
}

func main() {

	db, err := sql.Open("mysql", "root:1wD3Fsx&$*(fg@tcp(localhost:3306)/proxy_requests")
	if err != nil {
		log.Println("error validating sql.Open arguments")
		panic(err.Error())
	}
	defer db.Close()

	routes := []string{"/todos/", "/comments/", "/albums/", "/photos/", "/posts/", "/users/"}
	for i := 0; i < len(routes); i++ {
		http.HandleFunc(routes[i], func(rw http.ResponseWriter, r *http.Request) {

			url := string(r.Host) + r.URL.String()
			url2 := "https://jsonplaceholder.typicode.com" + r.URL.String()
			insertReq(db, url)
			insertReq(db, url2)
			// log.Println(r.Header)
			response, err := http.Get("https://jsonplaceholder.typicode.com" + r.URL.String())
			if err != nil {
				log.Fatal(err)
			}

			bytes, errRead := io.ReadAll(response.Body)
			if errRead != nil {
				log.Fatal(errRead)
			}

			JsonString := string(bytes)
			var mp map[string]interface{}
			err = json.Unmarshal(bytes, &mp)
			if err != nil {
				log.Fatal(err)
			}

			// fmt.Printf("%+v\n", mp)

			mp["foo"] = "bar"
			NewBytes, err := json.Marshal(mp)
			if err != nil {
				log.Fatal(err)
			}

			insertResp(db, url2, JsonString)

			rw.Header().Set("Content-Type", "application/json")
			rw.Write(NewBytes)

			JsonFinalResp := string(NewBytes)
			insertResp(db, url, JsonFinalResp)

		})
	}

	log.Println("App listening to port 8080")
	// dt := time.Now()
	// log.Println(dt.Format("01-02-2006 15:04:04"))
	// log.Println(reflect.TypeOf(dt.Format("01-02-2006 15:04:04")))
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
