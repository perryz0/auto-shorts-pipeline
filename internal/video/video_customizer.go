package video

// VideoCustomizer defines the interface for customizing video inputs

// Current imp supports generation of short-form quizzes and sports highlights vids
type VideoCustomizer interface {
	ParseInput(input string) error
	GenerateScript() (string, error)
	GetVideoType() string
}
