package postSocialNet

type PostSocialNet interface {
	AddTextContent(newText string)
	RemoveTextContent(oldText string)
	GetTextContent() string
}
