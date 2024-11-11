package postSocialNetUnsync

import (
	"strings"
)

type PostSocialNetUnsync struct {
	ID          int
	TextContent string
}

func (post *PostSocialNetUnsync) GetTextContent() string {
	return post.TextContent
}

func (post *PostSocialNetUnsync) AddTextContent(newText string) {
	post.TextContent += newText
}

func (post *PostSocialNetUnsync) RemoveTextContent(oldText string) {
	post.TextContent = strings.Replace(post.TextContent, oldText, "", 1)
}
