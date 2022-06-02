package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/getlantern/systray"
)

var badge *VpnBadge

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println(os.ErrInvalid)
		return
	}
	name := flag.Arg(0)

	badge = NewVpnBadge(name)

	systray.Run(onReady, onExit)
}

func onReady() {
	iconOn := getIcon("vpn-on.ico")
	iconOff := getIcon("vpn-off.ico")

	systray.SetIcon(iconOff)
	systray.SetTitle("VPN Badge")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()

	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			found, err := badge.ScanNetwork()
			if err != nil {
				fmt.Println("ScanNetwork error occuerd")
				systray.Quit()
			}

			if found {
				systray.SetIcon(iconOn)
			} else {
				systray.SetIcon(iconOff)
			}
		}
	}
}

func onExit() {
	// clean up here
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	return b
}
