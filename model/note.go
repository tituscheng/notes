package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/lithammer/shortuuid/v4"
)

type Note struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	Title       string    `json:"title"`
	CreatedDate time.Time `json:"createdDate"`
	Content     string    `json:"content"`
}

func (n *Note) SetContent(content string) {
	n.Content = content
}

func (n *Note) ParseTime(fileName string) {
	comps := strings.Split(fileName, "_")
	createdDate, _ := time.Parse(time.RFC850, comps[1])
	n.CreatedDate = createdDate
}

func (n *Note) FileName() string {
	return fmt.Sprintf("%s_%s", n.Id, time.Now().Format(time.RFC850))
}

func NewNote() *Note {
	n := &Note{Id: shortuuid.New(), Title: "New Notes"}
	return n
}

type Notes struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	Items []Note `json:"items"`
}
