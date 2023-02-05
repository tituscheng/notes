package notes

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"notes/model"
	"notes/util"
	"os"
	"path/filepath"
	"sort"
)

const (
	CACHE = ".notes"
)

var (
	cacheFolder         *string
	selectedFolderIndex = 0
	notes               []model.Notes
	selectedNoteIndex   = 0
)

func Init() {
	userDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	if cacheFolder == nil {
		if path := filepath.Join(userDir, CACHE); len(path) > 0 {
			cacheFolder = &path
		}
		if _, err := os.Stat(*cacheFolder); errors.Is(err, fs.ErrNotExist) {
			err := os.Mkdir(*cacheFolder, 0755)
			if err != nil {
				panic(err)
			}
		}
	}
}

func CreateFolder(folder string) error {
	if cacheFolder == nil {
		Init()
	}
	newFolderpath := filepath.Join(*cacheFolder, folder)
	if _, err := os.Stat(newFolderpath); errors.Is(err, fs.ErrNotExist) {
		err := os.Mkdir(newFolderpath, 0755)
		return err
	} else {
		return nil
	}
}

func GetNotes() []model.Notes {
	return notes
}

func Read() {
	if cacheFolder == nil {
		Init()
	}
	folders, err := ioutil.ReadDir(*cacheFolder)
	if err != nil {
		panic(err)
	}
	notes = []model.Notes{}
	for _, folder := range folders {
		if folder.IsDir() {
			noteFolder := model.Notes{Name: folder.Name()}
			noteFolderPath := filepath.Join(*cacheFolder, folder.Name())
			files, err := ioutil.ReadDir(noteFolderPath)
			if err != nil {
				panic(err)
			}
			for _, file := range files {
				note := model.Note{Id: file.Name()}
				notePath := filepath.Join(noteFolderPath, file.Name())
				note.Path = notePath
				data, err := ioutil.ReadFile(notePath)
				if err != nil {
					panic(err)
				}
				note.ParseTime(file.Name())
				note.Content = string(data)
				note.Title = util.GetTitle(note.Content)
				noteFolder.Items = append(noteFolder.Items, note)
			}
			notes = append(notes, noteFolder)
		}
	}
}

func CreateNewNote(folderName string, note *model.Note) error {
	if cacheFolder == nil {
		Init()
	}
	newNotePath := filepath.Join(*cacheFolder, folderName, note.FileName())
	return ioutil.WriteFile(newNotePath, []byte(note.Content), 0755)
}

func SetSelectedFolderIndex(folderIndex int) {
	selectedFolderIndex = folderIndex
}

func GetSelectedFolder() model.Notes {
	if len(notes) == 0 {
		return model.Notes{}
	}
	sort.Slice(notes[selectedFolderIndex].Items, func(i, j int) bool {
		return notes[selectedFolderIndex].Items[i].CreatedDate.After(notes[selectedFolderIndex].Items[j].CreatedDate)
	})
	return notes[selectedFolderIndex]
}

func SaveNote(m model.Note) {
	err := ioutil.WriteFile(m.Path, []byte(m.Content), 0755)
	if err != nil {
		panic(err)
	}
}

func SetSelectedNoteIndex(noteIndex int) {
	selectedNoteIndex = noteIndex
}

func GetSelectedNoteIndex() int {
	if selectedNoteIndex >= 0 {
		return selectedNoteIndex
	} else {
		return -1
	}
}
