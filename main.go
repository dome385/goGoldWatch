package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Config struct {
	App            fyne.App
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
	MainWindow     fyne.Window
	PriceContainer *fyne.Container
}

var myApp Config

func main() {
	// create a fyne application
	fyneApp := app.NewWithID("de.gocode.goldwatch.preferences")
	myApp.App = fyneApp

	// create our loggers
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// open a connection to the database

	// create a database repository

	// create and size a fyne window
	myApp.MainWindow = fyneApp.NewWindow("GoGoldWatch")
	myApp.MainWindow.Resize(fyne.NewSize(770, 410))
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()

	myApp.makeUI()
	// show and run
	myApp.MainWindow.ShowAndRun()
}
