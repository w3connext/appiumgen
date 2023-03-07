package appiumgen

type Generator interface {
	Generate(content []byte) (string, error)
}
