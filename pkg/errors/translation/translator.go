package translation

// Object : This is the translation object used to translate messages
type Translator interface {
	GetTranslation(item string, lang string) string
}
