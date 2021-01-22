package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func main() {
	systray.Run(onReady, onExit)

}

func getTopResourcesConsumption() string {

	result, _ := exec.Command("bash", "-c", "ps", "-eo", "pid,ppid,cmd,%mem,%cpu --sort=-%mem | head").Output()
	return string(result)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Pretty awesome")
	systray.SetTemplateIcon(icon.Data, icon.Data)
	menuData := getTopResourcesConsumption()
	fmt.Println(menuData)
	menuDataSlice := strings.Split(menuData, "\n")

	var menuMap map[string]*systray.MenuItem
	menuMap = make(map[string]*systray.MenuItem)

	for _, s := range menuDataSlice {
		var variable = systray.AddMenuItem(s, "test description")
		menuMap[s] = variable

	}

	a := menuMap[menuDataSlice[0]]

	b := menuMap[menuDataSlice[1]]

	c := menuMap[menuDataSlice[2]]

	d := menuMap[menuDataSlice[3]]

	print("\n")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	//mQuit.AddSubMenuItemCheckbox("SubMenuBottom - Toggle Panic!", "SubMenu Test (bottom) - Hide/Show Panic!", false)
	systray.AddMenuItem("New Button", "Cool un bouton !")

	for {
		select {
		case <-mQuit.ClickedCh:
			systray.Quit()
			fmt.Println("Quit2 now...")
			return

		case <-a.ClickedCh:
			fmt.Println("Processus tué a.")
			//return

		case <-b.ClickedCh:
			fmt.Println("Processus tué b.")
			//return

		case <-c.ClickedCh:
			fmt.Println("Processus tué c.")
			//return

		case <-d.ClickedCh:
			fmt.Println("Processus tué d.")
			//	return
		}
	}

}

func onExit() {
	fmt.Println("Merci de nous envoyer une correction, car le projet est vraiment bien merci")
}
