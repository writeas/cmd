package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/writeas/nerds/store"
)

var (
	outDir    string
	indexPage []byte
	debugging bool
	db        *sql.DB
)

func poster(w http.ResponseWriter, r *http.Request) {
	post := r.FormValue("w")

	if post == "" {
		fmt.Fprintf(w, "%s", indexPage)
		return
	}

	var friendlyId string
	var err error
	if outDir != "" {
		// Using file-based storage
		friendlyId, err = store.SavePost(outDir, []byte(post))
	} else {
		// Using database storage
		friendlyId = store.GenerateFriendlyRandomString(store.FriendlyIdLen)
		editToken := store.Generate62RandomString(32)

		_, err = db.Exec("INSERT INTO posts (id, content, modify_token, text_appearance) VALUES (?, ?, ?, 'mono')", friendlyId, post, editToken)
	}
	if err != nil {
		fmt.Printf("Error saving: %s\n", err)
		fmt.Fprint(w, "Couldn't save :(\n")
		return
	}

	if debugging {
		fmt.Printf("Saved new post %s\n", friendlyId)
	}

	fmt.Fprintf(w, "https://write.as/%s", friendlyId)

	if !strings.Contains(r.UserAgent(), "Android") {
		fmt.Fprint(w, "\n")
	}
}

func main() {
	outDirPtr := flag.String("o", "", "Directory where text files will be stored.")
	staticDirPtr := flag.String("s", "./static", "Directory where required static files exist.")
	portPtr := flag.Int("p", 8080, "Port to listen on.")
	debugPtr := flag.Bool("debug", false, "Enables garrulous debug logging.")
	flag.Parse()

	outDir = *outDirPtr
	debugging = *debugPtr

	fmt.Print("Initializing...")
	var err error
	indexPage, err = ioutil.ReadFile(*staticDirPtr + "/index.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("DONE")

	// Connect to database
	dbUser := os.Getenv("WA_USER")
	dbPassword := os.Getenv("WA_PASSWORD")
	dbName := os.Getenv("WA_DB")
	dbHost := os.Getenv("WA_HOST")

	if outDir == "" && (dbUser == "" || dbPassword == "" || dbName == "") {
		// Ensure parameters needed for storage (file or database) are given
		fmt.Println("Database user, password, or database name not set.")
		return
	}

	if outDir == "" {
		fmt.Print("Connecting to database...")
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4", dbUser, dbPassword, dbHost, dbName))
		if err != nil {
			fmt.Printf("\n%s\n", err)
			return
		}
		defer db.Close()
		fmt.Println("CONNECTED")
	}

	fmt.Printf("Serving on http://localhost:%d\n", *portPtr)

	http.HandleFunc("/", poster)
	http.ListenAndServe(fmt.Sprintf(":%d", *portPtr), nil)
}
