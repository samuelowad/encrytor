package gui

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/samuelowad/encryptor/pkg"
	"log"
	"strconv"
	"time"
)

func Gui() {
	myApp := app.New()
	myWindow := myApp.NewWindow("locker")
	myWindow.Resize(fyne.NewSize(800, 300))

	NUMBER_OF_TRIALS := 3
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter passkey...")

	encrypt := widget.NewButtonWithIcon("encrypt", theme.LoginIcon(), func() {
		pkg.Encrypt("./testData")
	})
	decrypt := container.NewVBox(input, widget.NewButton("Save", func() {
		if len(input.Text) < 1 {
			if NUMBER_OF_TRIALS < 1 {
				dialog.NewInformation("TRIAL EXCEEDED", "trial too much exiting now", myWindow).Show()
				time.Sleep(time.Second * 3)
				myApp.Quit()
			}
			NUMBER_OF_TRIALS--
			dialog.ShowError(errors.New("text is empty "+strconv.Itoa(NUMBER_OF_TRIALS)+" trials left"), myWindow)
		}
		log.Println("Content was:", input.Text)
	}), encrypt)

	myWindow.SetContent(decrypt)
	myWindow.ShowAndRun()
}
