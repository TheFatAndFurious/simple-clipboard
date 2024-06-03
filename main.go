package main

import (
	"fmt"

	"github.com/go-vgo/robotgo/clipboard"
	hook "github.com/robotn/gohook"
)

func main() {
	hook.Register(hook.KeyDown, []string{"ctrl", "c"}, func(e hook.Event) {
		fmt.Println("[Event] Ctrl+Shift+X detected!")
		fmt.Println(clipboard.ReadAll())
		// hook.End()
	})

	s := hook.Start()
	<-hook.Process(s)
}

