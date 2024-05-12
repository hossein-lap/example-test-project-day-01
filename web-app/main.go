package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	// "reflect"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	// "context"
    "github.com/redis/go-redis/v9"
)

const (
	port_redis string = "6349"
	port_databse string = "3360"
	port_webapp string = "9090"

	minikube_ip string = "192.168.46.2"
	database_password string = "test1234"
)

var (
)

type RequestHandler struct {}

// main func {{{
func main() {

	mux := http.NewServeMux()

	mux.Handle("/filestorage/", &RequestHandler{})
	mux.Handle("/redis/", &RequestHandler{})
	mux.Handle("/database/", &RequestHandler{})

	test("database-mysql")

	http.ListenAndServe(fmt.Sprintf(":%s", port_webapp), mux) // server
}
// }}}

// http server {{{
func (h *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

 	// vars
	request_type := r.Method
	tmp := strings.SplitN(r.URL.String(), "/", 3)
	section := tmp[1]
	key := tmp[2]
	params := r.URL.Query()
	value := params.Get(key)

	if request_type == http.MethodGet {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("<%s> FATAL IO %s\n", section, err.Error())
		}

		if key == "" {
			log.Printf("[%s] key is empty", section)
			return
		}

		if len(body) != 0 {

			// bytes := json.RawMessage(body)
			// fmt.Printf("[type]: %s\n" , reflect.TypeOf(bytes))
			// fmt.Printf("%s\n", bytes)
			// // for k, v := range bytes {
			// // 	fmt.Printf("key: %d, value: %s\n", k, string(v))
			// // }
			// fmt.Printf("[bytes] %s\n", bytes)
			// fmt.Printf("[section: %s] [method: %s] key: %s, body: %s\n", section, request_type, key, body) // type: []uint8

		} else {
			fmt.Printf("[section: %s] [method: %s] key: %s\n", section, request_type, key) // type: []uint8
		}

		switch {
		case section == "redis":
			// search
			if value != "" {
				fmt.Printf("[search value]> [section: %s] [value: %s]\n", section, value)
				// do the search
			} else {
				log.Printf("<%s> [search]: error: empty value\n", section)
			}
			return
		case section == "database":
			fmt.Printf("[%s] path: %s\n", section, key)
			return
		case section == "filestorage":
			fmt.Printf("[%s] path: %s\n", section, key)
			return
		default:
			fmt.Printf("Default\n")
		}

	} else if request_type == http.MethodPost || request_type == http.MethodPut {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("<%s> FATAL IO %s\n", section, err.Error())
		}

		switch {
		case section == "redis":
			params := r.URL.Query()
			value := params.Get(key)

			if value == "" {
				log.Printf("[%s] Invalid key: %s\n", section, key)
				return
			}


			if body != nil {
				fmt.Printf("[section: %s] [method: %s] key: %s, value: %s\n", section, request_type, key, value) // type: []uint8
			} else {
				fmt.Printf("[section: %s] [method: %s] no body, url %s\n", section, request_type, value) // type: []uint8
			}

			// // search
			// if value != "" {
			// 	fmt.Printf("Has body/data: %s\n", value)
			// 	fmt.Printf("[search value]> [section: %s] [value: %s]\n", section, value)
			// 	// do the search
			// } else {
			// 	fmt.Printf("No body/data:\n")
			// 	// log.Printf("<%s> [search]: error: empty value\n", section)
			// }

			return
		case section == "database":
			return
		case section == "filestorage":
			fmt.Printf("[%s] path: %s\n", section, value)
			test(value)
			return
		default:
			fmt.Printf("Default\n")
		}

	}

}
// }}}

// get env variable {{{
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
// }}}

func test(tmp string) {

	// database{{{

	sql_url := fmt.Sprintf("root:%s@tcp(%s:%s)/test", database_password, minikube_ip, port_databse)
	db, err := sql.Open("mysql", sql_url)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Printf("[test] <%s> Success!\n", tmp)

	// }}}

	// cache{{{

	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", minikube_ip, port_redis),
		Password: "", // no password
		DB: 0,  // default DB
	})
	fmt.Println(client)

	// }}}

}
