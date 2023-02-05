package gui

import (
	"fmt"
	"notes/model"
	"notes/notes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	notesFolderView *NotesFolderView
)

type NotesFolderView struct {
	Container      *fyne.Container
	folderListView *widget.List
	notesFolders   []model.Notes
}

func (notesFolderView *NotesFolderView) Refresh() {
	notes.Read()
	notesFolderView.notesFolders = notes.GetNotes()
	notesFolderView.folderListView.Refresh()
}

func (notesFolderView *NotesFolderView) Build() {
	notes.Read()
	notesFolderView.notesFolders = notes.GetNotes()
	notesFolderView.folderListView = widget.NewList(
		func() int {
			return len(notesFolderView.notesFolders)
		},
		func() fyne.CanvasObject {
			c := container.NewHBox(widget.NewIcon(theme.FolderIcon()), widget.NewLabel("template"), layout.NewSpacer(), widget.NewLabel("template"))
			return c
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*fyne.Container).Objects[1].(*widget.Label).SetText(notesFolderView.notesFolders[i].Name)
			o.(*fyne.Container).Objects[3].(*widget.Label).SetText(fmt.Sprintf("%d", len(notesFolderView.notesFolders[i].Items)))
		})
	if len(notesFolderView.notesFolders) > 0 {
		notesFolderView.folderListView.Select(0)
	}
	notesFolderView.folderListView.OnSelected = func(id widget.ListItemID) {
		notes.SetSelectedFolderIndex(id)
		notesView.Refresh()
	}

	addFolderButton := widget.NewButtonWithIcon("New Folder", theme.DocumentCreateIcon(), func() {
		folderNameData := binding.NewString()
		formDialog := dialog.NewForm("New Folder", "OK", "Cancel", []*widget.FormItem{
			{
				Text:   "Name",
				Widget: widget.NewEntryWithData(folderNameData),
			},
		}, func(b bool) {
			if b {
				if folderName, _ := folderNameData.Get(); len(folderName) > 0 {
					notes.CreateFolder(folderName)
					notesFolderView.Refresh()
				}
			}
		}, mainWindow)
		formDialog.Show()
	})
	notesFolderListViewContainer := container.NewBorder(nil, addFolderButton, nil, nil, notesFolderView.folderListView)
	notesFolderView.Container = notesFolderListViewContainer
}

func NewNotesFolderView() *NotesFolderView {
	nfv := &NotesFolderView{}
	nfv.Build()
	return nfv
}
