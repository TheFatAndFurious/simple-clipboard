package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-vgo/robotgo/clipboard"
	_ "github.com/mattn/go-sqlite3"
	hook "github.com/robotn/gohook"
)

func main() {

	db, err := sql.Open("sqlite3", "entries.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("db has been created")
	defer db.Close()

	sqlSmt := `
	CREATE TABLE IF NOT EXISTS entries (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
        content TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL 
    );`

	_, err = db.Exec(sqlSmt)
	if err != nil{
		log.Fatal(err)
	}

	


	hook.Register(hook.KeyDown, []string{"ctrl", "c"}, func(e hook.Event) {
		fmt.Println("[Event] Ctrl+Shift+X detected!")
		content, err := clipboard.ReadAll()
		if err!= nil {
			log.Fatal(err)
        }
		_, err = db.Exec("INSERT INTO entries (content) VALUES (?)", content)
		if err != nil {
			log.Fatal(err)
		}		
		// hook.End()
	})

	s := hook.Start()
	<-hook.Process(s)
}

