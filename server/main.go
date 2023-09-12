package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"solver/dpll"
)

const PORT int = 3000

func scanRequest(request string) [][]string {
	var lines [][]string

	for _, clause := range strings.Split(request, "\n") {
		clause = strings.TrimSpace(clause)
		if clause != "" {
			re := regexp.MustCompile(`\s+`)
			literalsList := re.Split(clause, -1)
			lines = append(lines, literalsList)
		}
	}
	return lines
}

func writeResponse(w http.ResponseWriter, jsonResponse []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	_, err := w.Write(jsonResponse)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	bodyStr := string(body)

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	lines := scanRequest(bodyStr)
	solution := dpll.RunDPLL(lines)
	jsonResponse, err := json.Marshal(solution)

	if err != nil {
		panic(err.Error())
	}

	writeResponse(w, jsonResponse)

}

func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server on port", PORT)
	if err := http.ListenAndServe(fmt.Sprint(":", PORT), nil); err != nil {
		println(err.Error())
	}
}
