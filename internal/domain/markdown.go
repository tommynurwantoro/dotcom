package domain

type MarkdownService interface {
	GetAllBlogs() []MyFile
	GetBlogByID(id string) []byte
}
