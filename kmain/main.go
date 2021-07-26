package main

import (
	"io"
	_ "net/http/pprof"
	"runtime"

	"github.com/icexin/eggos/app/sh"
	"github.com/icexin/eggos/cga/fbcga"
	"github.com/icexin/eggos/console"
	"github.com/icexin/eggos/debug"
	"github.com/icexin/eggos/fs"
	"github.com/icexin/eggos/inet"
	"github.com/icexin/eggos/pci"
	"github.com/icexin/eggos/ps2/mouse"

	_ "github.com/icexin/eggos/e1000"
	"github.com/icexin/eggos/kbd"
	"github.com/icexin/eggos/kernel"
	"github.com/icexin/eggos/uart"
	"github.com/icexin/eggos/vbe"
)

func main() {
	// trap and syscall threads use two Ps,
	// and the remainings are for other goroutines
	runtime.GOMAXPROCS(6)

	kernel.Init()
	uart.Init()
	kbd.Init()
	mouse.Init()
	console.Init()

	fs.Init()
	vbe.Init()
	fbcga.Init()
	pci.Init()

	err := inet.Init()
	if err != nil {
		panic(err)
	}

	debug.Logf("[runtime] go version:%s", runtime.Version())
	w := console.Console()
	io.WriteString(w, "\nwelcome to eggos\n")
	sh.Bootstrap()
}
