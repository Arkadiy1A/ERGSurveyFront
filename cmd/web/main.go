package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type SurveyModel struct {
	Survey   Survey
	Question Question
	Url      string
}

func main() {
	port := "8080"

	surv := CreateDummySurvey()
	url := os.Getenv("BASE_URL")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "templates/survey.component.gohtml", SurveyModel{Question: *surv.CurrentQuestion(), Url: url})
	})

	http.HandleFunc("/table", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Questions = %v\n", surv)
		render(w, "templates/table.component.gohtml", SurveyModel{Survey: surv})
	})

	http.HandleFunc("/question", func(w http.ResponseWriter, r *http.Request) {
		render(w, "templates/question.component.gohtml", SurveyModel{Url: url})
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			fmt.Printf("Method submit has been called\n")
			res := &ResponsePayload{}
			err := readJSON(w, r, res)
			if err != nil {
				fmt.Printf("Falsed to parse JSON: %v\n", err)
			}
			ip := readUserIP(r)
			fmt.Printf("Request from: %s\n", ip)
			surv.Increment(res.Id, ip)
		}
	})

	http.HandleFunc("/newQuestion", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			fmt.Printf("Method submit has been called\n")
			res := &NewQuestionPayload{}
			err := readJSON(w, r, res)
			if err != nil {
				fmt.Printf("Falsed to parse JSON: %v\n", err)
			}

			if res.Pin == "31415926" {
				surv.AddQuestion(res.Name, res.Q1, res.Q2, res.Q3)
			}
		}
	})

	http.HandleFunc("/latest", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Printf("Method latest has been called\n")
			data, err := json.Marshal(*surv.CurrentQuestion())
			if err != nil {
				fmt.Printf("failed to encode the object to JSON: %v", err)
			}
			w.Write(data)
		}
	})

	http.HandleFunc("/setQuestion", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Get the query parameters from the request
			params := r.URL.Query()

			// Get the "name" parameter value and print it
			pin := params.Get("pin")
			num := params.Get("num")
			//fmt.Printf("Name parameter value is: %s\n", name)

			if pin == "31415926" {
				numInt, _ := strconv.Atoi(num)
				surv.SetQuestion(numInt)
			}
		}
	})

	fmt.Printf("Starting survey frontend on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panic(err)
	}
}
