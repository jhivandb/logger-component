package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var logRows = [][]string{
	{
		"2025-06-24T07:08:43.327Z",
		"ERROR",
		"An error occurred while syncing the employee information.",
		`{"module":"wso2/par_app.utils","error":{"causes":[],"message":"Error while executing SQL query: 
    DELETE FROM REDACTED WHERE REDACTED =  ? ;
. DELETE command denied to user 'REDACTED'@'00.0.0.00' for table 'REDACTED'.","detail":{"errorCode":1142,"sqlState":"42000"},"stackTrace":[]}}`,
		"v1.0",
		"3140f2f9-3ed4-4476-b1f4-c40e740682c9",
	},
	{
		"2025-06-24T06:53:21.577Z",
		"ERROR",
		"An error occurred while syncing the employee information.",
		`{"module":"wso2/par_app.utils","error":{"causes":[],"message":"Error while executing SQL query: 
    DELETE FROM REDACTED WHERE REDACTED =  ? ;
. DELETE command denied to user 'REDACTED'@'00.0.0.00' for table 'REDACTED'.","detail":{"errorCode":1142,"sqlState":"42000"},"stackTrace":[]}}`,
		"v1.0",
		"3140f2f9-3ed4-4476-b1f4-c40e740682c9",
	},
	{
		"2025-06-24T06:50:34.277Z",
		"ERROR",
		"An error occurred while syncing the employee information.",
		`{"module":"wso2/par_app.utils","error":{"causes":[],"message":"Error while executing SQL query: 
    DELETE FROM REDACTED WHERE REDACTED =  ? ;
. DELETE command denied to user 'REDACTED'@'00.0.0.00' for table 'REDACTED'.","detail":{"errorCode":1142,"sqlState":"42000"},"stackTrace":[]}}`,
		"v1.0",
		"3140f2f9-3ed4-4476-b1f4-c40e740682c9",
	},
	{
		"2025-06-24T05:33:16.920Z",
		"ERROR",
		"An error occurred while syncing the employee information.",
		`{"module":"wso2/par_app.utils","error":{"causes":[],"message":"Error while executing SQL query: 
    DELETE FROM REDACTED WHERE REDACTED =  ? ;
. DELETE command denied to user 'REDACTED'@'00.0.0.00' for table 'REDACTED'.","detail":{"errorCode":1142,"sqlState":"42000"},"stackTrace":[]}}`,
		"v1.0",
		"3140f2f9-3ed4-4476-b1f4-c40e740682c9",
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	randomRow := logRows[rand.Intn(len(logRows))]

	logEntry := fmt.Sprintf(`{"timestamp":"%s","level":"%s","message":"%s","details":%s,"version":"%s","id":"%s"}`,
		randomRow[0], randomRow[1], randomRow[2], randomRow[3], randomRow[4], randomRow[5])

	fmt.Println(logEntry)

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "logging")
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", handler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
