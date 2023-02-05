package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

type Editor struct {
	widget.Entry
	OnSave func()
}

func (e *Editor) TypedShortcut(s fyne.Shortcut) {
	if _, ok := s.(*desktop.CustomShortcut); !ok {
		e.Entry.TypedShortcut(s)
		return
	}
	if s.ShortcutName() == "CustomDesktop:Command+S" {
		e.OnSave()
	}
}

func NewEditor() *Editor {
	editor := &Editor{}
	editor.MultiLine = true
	editor.Wrapping = fyne.TextTruncate
	editor.ExtendBaseWidget(editor)
	return editor
}
