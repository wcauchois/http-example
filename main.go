package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/wcauchois/http-example/index"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Config struct {
	Port   int    `yaml:"port"`
	User   string `yaml:"user"`
	DBName string `yaml:"dbname"`
}

func readConfig(filename string) (out *Config, err error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	var c Config
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return
	}
	out = &c
	return
}

func main() {
	config, _ := readConfig("conf.yml")
	connStr := fmt.Sprintf("user=%v dbname=%v sslmode=disable", config.User, config.DBName)
	db, err := sql.Open("postgres", connStr)
	/*
		rows, err := db.Query("select * from posts")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var name string
			var body string
			rows.Scan(&id, &name, &body)
			fmt.Println(id, name, body)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
	*/
	indexHandler, err := index.New(db)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", indexHandler)
	go func() {
		log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), nil))
	}()
	log.Printf("Listening on port %v", config.Port)
	select {} // Sleep forever
}
