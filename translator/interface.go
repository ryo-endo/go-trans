package translator

type Translator interface {
	Trans(s string, from string, to string) (string, error)
}
