package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// type ToDoReading struct {
// 	UserId    int    `json:"userid"`
// 	Id        int    `json:"id"`
// 	Title     string `json:"title"`
// 	Completed bool   `json:"completed"`
// }

func insertReq(db *sql.DB, route string, number string) {
	var ins *sql.Stmt
	ins, errIns := db.Prepare("INSERT into `proxy_requests`.`requests`(`request_url`,`time`) VALUES (?,?)")
	if errIns != nil {
		log.Fatal(errIns)
	}
	defer ins.Close()
	dt := time.Now()

	_, errRes := ins.Exec(route+number, dt.Format("01-02-2006 15:04:04"))
	if errRes != nil {
		log.Fatal(errRes)
	}
}

func insertResp(db *sql.DB, host string, number string, body string) {

	var insResponse *sql.Stmt
	insResponse, errIns := db.Prepare("INSERT into `proxy_requests`.`responses`(`Response_Body`,`Time`,`Host`) VALUES (?,?,?)")
	if errIns != nil {
		log.Fatal(errIns)
	}
	defer insResponse.Close()
	dt2 := time.Now()

	_, errRes := insResponse.Exec(body, dt2.Format("01-02-2006 15:04:04"), host+number)
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

	//todos
	for i := 1; i <= 200; i++ {
		number := strconv.Itoa(i)
		http.HandleFunc("/todos/"+number, func(rw http.ResponseWriter, r *http.Request) {

			insertReq(db, "http://127.0.0.1:8080/todos/", number)
			insertReq(db, "https://jsonplaceholder.typicode.com/todos/", number)

			response, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + number)
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

			fmt.Printf("%+v\n", mp)

			mp["foo"] = "bar"
			NewBytes, err := json.Marshal(mp)
			if err != nil {
				log.Fatal(err)
			}

			insertResp(db, "https://jsonplaceholder.typicode.com/todos/", number, JsonString)

			rw.Header().Set("Content-Type", "application/json")
			rw.Write(NewBytes)

			insertResp(db, "http://127.0.0.1:8080/todos/", number, JsonString)

		})
	}

	//posts
	for i := 1; i <= 100; i++ {
		number := strconv.Itoa(i)
		http.HandleFunc("/posts/"+number, func(rw http.ResponseWriter, r *http.Request) {

			insertReq(db, "http://127.0.0.1:8080/posts/", number)
			insertReq(db, "https://jsonplaceholder.typicode.com/posts/", number)

			response, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + number)
			if err != nil {
				log.Fatal(err)
			}

			bytes, errRead := io.ReadAll(response.Body)
			if errRead != nil {
				log.Fatal(errRead)
			}
			JsonString := string(bytes)

			insertResp(db, "https://jsonplaceholder.typicode.com/posts/", number, JsonString)

			var mp map[string]interface{}
			err = json.Unmarshal(bytes, &mp)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%+v\n", mp)

			mp["foo"] = "bar"
			NewBytes, err := json.Marshal(mp)
			if err != nil {
				log.Fatal(err)
			}

			rw.Header().Set("Content-Type", "application/json")
			rw.Write(NewBytes)

			insertResp(db, "http://127.0.0.1:8080/posts/", number, JsonString)
		})
	}

	//comments
	for i := 1; i <= 500; i++ {
		number := strconv.Itoa(i)
		http.HandleFunc("/comments/"+number, func(rw http.ResponseWriter, r *http.Request) {

			insertReq(db, "http://127.0.0.1:8080/comments/", number)
			insertReq(db, "https://jsonplaceholder.typicode.com/comments/", number)

			response, err := http.Get("https://jsonplaceholder.typicode.com/comments/" + number)
			if err != nil {
				log.Fatal(err)
			}

			bytes, errRead := io.ReadAll(response.Body)
			if errRead != nil {
				log.Fatal(errRead)
			}
			JsonString := string(bytes)

			insertResp(db, "https://jsonplaceholder.typicode.com/comments/", number, JsonString)

			var mp map[string]interface{}
			err = json.Unmarshal(bytes, &mp)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%+v\n", mp)

			mp["foo"] = "bar"
			NewBytes, err := json.Marshal(mp)
			if err != nil {
				log.Fatal(err)
			}
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(NewBytes)

			insertResp(db, "http://127.0.0.1:8080/comments/", number, JsonString)
		})
	}

	//albums
	for i := 1; i <= 100; i++ {
		number := strconv.Itoa(i)
		http.HandleFunc("/albums/"+number, func(rw http.ResponseWriter, r *http.Request) {

			insertReq(db, "http://127.0.0.1:8080/albums/", number)
			insertReq(db, "https://jsonplaceholder.typicode.com/albums/", number)

			response, err := http.Get("https://jsonplaceholder.typicode.com/albums/" + number)
			if err != nil {
				log.Fatal(err)
			}

			bytes, errRead := io.ReadAll(response.Body)
			if errRead != nil {
				log.Fatal(errRead)
			}
			JsonString := string(bytes)

			insertResp(db, "https://jsonplaceholder.typicode.com/albums/", number, JsonString)

			var mp map[string]interface{}
			err = json.Unmarshal(bytes, &mp)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%+v\n", mp)

			mp["foo"] = "bar"
			NewBytes, err := json.Marshal(mp)
			if err != nil {
				log.Fatal(err)
			}
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(NewBytes)

			insertResp(db, "http://127.0.0.1:8080/albums/", number, JsonString)
		})
	}

	//photos
	for i := 1; i <= 5000; i++ {
		number := strconv.Itoa(i)
		http.HandleFunc("/photos/"+number, func(rw http.ResponseWriter, r *http.Request) {

			insertReq(db, "http://127.0.0.1:8080/photos/", number)
			insertReq(db, "https://jsonplaceholder.typicode.com/photos/", number)

			response, err := http.Get("https://jsonplaceholder.typicode.com/photos/" + number)
			if err != nil {
				log.Fatal(err)
			}

			bytes, errRead := io.ReadAll(response.Body)
			if errRead != nil {
				log.Fatal(errRead)
			}
			JsonString := string(bytes)

			insertResp(db, "https://jsonplaceholder.typicode.com/photos/", number, JsonString)

			var mp map[string]interface{}
			err = json.Unmarshal(bytes, &mp)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%+v\n", mp)

			mp["foo"] = "bar"
			NewBytes, err := json.Marshal(mp)
			if err != nil {
				log.Fatal(err)
			}
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(NewBytes)

			insertResp(db, "http://127.0.0.1:8080/photos/", number, JsonString)

		})
	}

	//users
	for i := 1; i <= 10; i++ {
		number := strconv.Itoa(i)
		http.HandleFunc("/users/"+number, func(rw http.ResponseWriter, r *http.Request) {

			insertReq(db, "http://127.0.0.1:8080/users/", number)
			insertReq(db, "https://jsonplaceholder.typicode.com/users/", number)

			response, err := http.Get("https://jsonplaceholder.typicode.com/users/" + number)
			if err != nil {
				log.Fatal(err)
			}

			bytes, errRead := io.ReadAll(response.Body)
			if errRead != nil {
				log.Fatal(errRead)
			}
			JsonString := string(bytes)

			insertResp(db, "https://jsonplaceholder.typicode.com/users/", number, JsonString)

			var mp map[string]interface{}
			err = json.Unmarshal(bytes, &mp)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%+v\n", mp)

			mp["foo"] = "bar"
			NewBytes, err := json.Marshal(mp)
			if err != nil {
				log.Fatal(err)
			}
			rw.Header().Set("Content-Type", "application/json")
			rw.Write(NewBytes)

			insertResp(db, "http://127.0.0.1:8080/users/", number, JsonString)
		})
	}

	log.Println("App listening to port 8080")
	// dt := time.Now()
	// log.Println(dt.Format("01-02-2006 15:04:04"))
	// log.Println(reflect.TypeOf(dt.Format("01-02-2006 15:04:04")))
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
