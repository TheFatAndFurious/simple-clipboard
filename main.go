package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

func main() {
	hook.Register(hook.KeyDown, []string{"ctrl", "c"}, func(e hook.Event) {
		fmt.Println("[Event] Ctrl+Shift+X detected!")
		// hook.End()
	})

	s := hook.Start()
	<-hook.Process(s)
}

