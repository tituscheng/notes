package gui

import (
	"fmt"
	"notes/notes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

var (
	mainWindow fyne.Window
)

const (
	HEIGHT = 1100
	WIDTH  = 600
)

func Start() {
	notes.Init()
	a := app.New()
	mainWindow = a.NewWindow("Notes")
	mainWindow.SetContent(widget.NewMultiLineEntry())
	mainWindow.Resize(fyne.NewSize(HEIGHT, WIDTH))
	mainWindow.SetCloseIntercept(func() {
		mainWindow.Hide()
	})
	mainWindow.CenterOnScreen()

	notes := NewNotes()

	mainWindow.SetContent(notes.Container)

	// textNote := widget.NewMultiLineEntry()

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("MyApp",
			fyne.NewMenuItem("Show", func() {
				fmt.Println("Show")
				mainWindow.Show()
			}))
		desk.SetSystemTrayMenu(m)
	}
	mainWindow.ShowAndRun()
	a.Run()
}

type Notes struct {
	Container      *fyne.Container
	folderListView *widget.List
}

func (notes *Notes) Build() {
	notesFolderView = NewNotesFolderView()
	notesView = NewNotesView()
	notesMainContainer := container.NewHSplit(notesFolderView.Container, notesView.Container)
	notesMainContainer.SetOffset(0.2)
	notes.Container = container.NewMax(notesMainContainer)
}

func NewNotes() *Notes {
	n := &Notes{}
	n.Build()
	return n
}
