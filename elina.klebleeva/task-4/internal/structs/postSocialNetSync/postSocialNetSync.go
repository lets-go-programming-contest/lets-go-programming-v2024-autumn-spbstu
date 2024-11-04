package postSocialNetSync

import (
	"strings"
	"sync"
)

type PostSocialNetSync struct {
	ID          int
	TextContent string
	myMutex     sync.RWMutex
}

func (post *PostSocialNetSync) GetTextContent() string {
	post.myMutex.RLock()
	defer post.myMutex.RUnlock()
	return post.TextContent
}

func (post *PostSocialNetSync) AddTextContent(newText string) {
	post.myMutex.Lock()
	defer post.myMutex.Unlock()
	post.TextContent += newText
}

func (post *PostSocialNetSync) RemoveTextContent(oldText string) {
	post.myMutex.Lock()
	defer post.myMutex.Unlock()
	post.TextContent = strings.Replace(post.TextContent, oldText, "", 1)
}
