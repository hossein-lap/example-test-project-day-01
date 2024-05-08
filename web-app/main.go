package main

import (
	// "encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	// "reflect"

	"database/sql"
	// "github.com/go-sql-driver/mysql"
)

type RequestHandler struct {}

// main func {{{
func main() {

	mux := http.NewServeMux()

	mux.Handle("/filename/", &RequestHandler{})
	mux.Handle("/redis/", &RequestHandler{})
	mux.Handle("/database/", &RequestHandler{})

	http.ListenAndServe(":1224", mux) // server
}
// }}}

// http server {{{
func (h *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

 	// vars
	request_type := r.Method
	section := strings.SplitN(r.URL.String(), "/", 4)[1]
	key := strings.SplitN(r.URL.String(), "/", 4)[2]
	params := r.URL.Query()
	value := params.Get(key)

	if request_type == http.MethodGet {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("<%s> FATAL IO %s\n", section, err.Error())
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
			return
		case section == "filestorage":
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

// database handler {{{
type Json struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type DBModel struct {
	ID uint64
	MyFieldName Json `json:"my_field_name"`
}

func DatabaseHandler(table string) {
	db, err := sql.Open("mysql", fmt.Sprintf("root:password@tcp(127.0.0.1)/%s", table))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query(fmt.Sprintf("SELECT * FROM %s", table))

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var row DBModel
		err := results.Scan(&row.ID, &row.MyFieldName)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(row.MyFieldName.Key)
	}
}
// }}}

// // get json {{{
// func getJSON(sqlString string) (string, error) {
//     rows, err := db.Query(sqlString)
//     if err != nil {
//         return "", err
//     }
//     defer rows.Close()
//     columns, err := rows.Columns()
//     if err != nil {
//         return "", err
//     }
//     count := len(columns)
//     tableData := make([]map[string]interface{}, 0)
//     values := make([]interface{}, count)
//     valuePtrs := make([]interface{}, count)
//     for rows.Next() {
//         for i := 0; i < count; i++ {
//           valuePtrs[i] = &values[i]
//         }
//         rows.Scan(valuePtrs...)
//         entry := make(map[string]interface{})
//         for i, col := range columns {
//             var v interface{}
//             val := values[i]
//             b, ok := val.([]byte)
//             if ok {
//                 v = string(b)
//             } else {
//                 v = val
//             }
//             entry[col] = v
//         }
//         tableData = append(tableData, entry)
//     }
//     jsonData, err := json.Marshal(tableData)
//     if err != nil {
//         return "", err
//     }
//     fmt.Println(string(jsonData))
//     return string(jsonData), nil 
// }
// // }}}


