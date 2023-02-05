package gui

import (
	"fmt"
	"log"
	"notes/model"
	"notes/notes"
	"notes/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	notesView *NotesView
)

type NotesView struct {
	Container   *fyne.Container
	notesList   *widget.List
	editor      *Editor
	notesFolder model.Notes
	data        binding.String
}

func (notesView *NotesView) Refresh() {
	notes.Read()
	notesView.notesFolder = notes.GetSelectedFolder()
	if notesView.notesList != nil {
		notesView.notesList.Refresh()
	}
}

func (notesView *NotesView) Build() {
	// variable initiation
	notesView.data = binding.NewString()
	notesView.notesFolder = notes.GetSelectedFolder()

	// editor
	notesView.editor = NewEditor()
	notesView.editor.Bind(notesView.data)
	notesView.editor.OnChanged = func(s string) {
		title := util.GetTitle(s)
		fmt.Printf("Title: %s\n", title)
		notesView.notesFolder.Items[notes.GetSelectedNoteIndex()].Title = title
		notesView.notesList.Refresh()
		notesView.data.Set(s)
	}

	notesView.editor.OnSave = func() {
		text, _ := notesView.data.Get()
		notesView.notesFolder.Items[notes.GetSelectedNoteIndex()].SetContent(text)
		notes.SaveNote(notesView.notesFolder.Items[notes.GetSelectedNoteIndex()])
	}
	notesView.editor.OnCursorChanged = func() {
		// text, _ := notesView.data.Get()
		// notesView.notesFolder.Items[notes.GetSelectedNoteIndex()].SetContent(text)
		// notes.SaveNote(notesView.notesFolder.Items[notes.GetSelectedNoteIndex()])
	}
	ctrlSTab := &desktop.CustomShortcut{KeyName: fyne.KeyS, Modifier: fyne.KeyModifierControl}
	notesView.editor.TypedShortcut(ctrlSTab)

	// Notes tool bar
	notesToolBar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			newNote := model.NewNote()
			notes.CreateNewNote(notesView.notesFolder.Name, newNote)
			notesView.Refresh()
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}),
	)
	notesView.notesList = widget.NewList(
		func() int {
			return len(notesView.notesFolder.Items)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(notesView.notesFolder.Items[i].Title)
		})
	notesView.notesList.OnSelected = func(id widget.ListItemID) {
		notes.SetSelectedNoteIndex(id)
		note := notesView.notesFolder.Items[id]
		notesView.data.Set(note.Content)
	}

	if len(notesView.notesFolder.Items) > 0 {
		notesView.notesList.Select(0)
		notes.SetSelectedNoteIndex(0)
		notesView.data.Set(notesView.notesFolder.Items[0].Content)
	}
	notesViewHSplit := container.NewHSplit(notesView.notesList, notesView.editor)
	notesViewHSplit.SetOffset(0.3)
	notesView.Container = container.NewBorder(notesToolBar, nil, nil, nil, notesViewHSplit)
}

func NewNotesView() *NotesView {
	nv := &NotesView{}
	nv.Build()
	return nv
}
