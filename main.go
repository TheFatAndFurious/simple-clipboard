package main

import (
	"database/sql"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
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

	query := `SELECT content FROM entries ORDER BY created_at DESC`
	rows, err := db.Query(query)
	if err!= nil {
		log.Fatal(err)
    }
	defer rows.Close()




	var data []string
	for rows.Next() {
		var content string
		if err := rows.Scan(&content); err != nil {
			log.Fatal(err)
		}
		data = append(data, content)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	myApp := app.New()
	myWindow := myApp.NewWindow("List Widget")

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

	myWindow.SetContent(list)
	myWindow.ShowAndRun()

	


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
		data = append(data, content)
		list.Refresh()
	})

	s := hook.Start()
	<-hook.Process(s)
}

