package config

var (
	Extensions   = []string{".yaml", ".yml"}
	SpecsSuffix  = []string{".spec.yaml", ".spec.yml"}
	ScreenSuffix = ".screen.py"
)

type Config struct {
	Platform string
}
