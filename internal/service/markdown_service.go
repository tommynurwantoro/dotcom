package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"tommynurwantoro/dotcom/internal/domain"

	"github.com/gomarkdown/markdown"
)

type MarkdownService struct{}

func (m *MarkdownService) GetAllBlogs() []domain.MyFile {
	files, err := ioutil.ReadDir("markdown")
	if err != nil {
		log.Fatal(err)
	}

	var arr []domain.MyFile
	for _, file := range files {
		split := strings.Split(file.Name(), "|")
		arr = append(arr, domain.MyFile{FileName: file.Name(), Number: split[0], Date: split[1], Title: strings.Split(split[2], ".")[0]})
	}

	return arr
}

func (m *MarkdownService) GetBlogByID(id string) (content []byte) {
	content = make([]byte, 0)
	files := m.GetAllBlogs()

	for _, file := range files {
		split := strings.Split(file.FileName, "|")
		if split[0] == id {
			data, err := ioutil.ReadFile(fmt.Sprintf("markdown/%s", file.FileName))
			if err != nil {
				log.Fatal(err)
			}

			content = markdown.ToHTML(data, nil, nil)
		}
	}

	return
}

func NewMarkdownService() domain.MarkdownService {
	return &MarkdownService{}
}
