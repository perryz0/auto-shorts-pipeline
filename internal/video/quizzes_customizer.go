package video

import (
	"errors"
	"fmt"
	"strings"
)

// QuizzesCustomizer handles customization for quiz videos
type QuizzesCustomizer struct {
	VideoType  string
	Topic      string
	Difficulty string
}

// ParseInput parses the input string and sets the values for the quiz video
func (qc *QuizzesCustomizer) ParseInput(input string) error {
	parts := strings.Split(input, ",")
	if len(parts) < 3 {
		return errors.New("insufficient input arguments")
	}
	qc.VideoType = strings.TrimSpace(parts[0])
	qc.Topic = strings.TrimSpace(parts[1])
	qc.Difficulty = strings.TrimSpace(parts[2])
	return nil
}

// GenerateScript generates a script based on the quiz video inputs
func (qc *QuizzesCustomizer) GenerateScript() (string, error) {
	// Custom logic for script generation
	return fmt.Sprintf("Creating a %s quiz on %s with difficulty %s.", qc.VideoType, qc.Topic, qc.Difficulty), nil
}

// GetVideoType returns the type of video being generated
func (qc *QuizzesCustomizer) GetVideoType() string {
	return qc.VideoType
}
