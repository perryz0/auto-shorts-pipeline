package video

import (
	"errors"
	"fmt"
	"strings"
)

// SportsHighlightsCustomizer handles customization for sports highlight videos
type SportsHighlightsCustomizer struct {
	VideoType string
	Sport     string
	Focus     string
}

// ParseInput parses the input string and sets the values for the sports highlight video
func (shc *SportsHighlightsCustomizer) ParseInput(input string) error {
	parts := strings.Split(input, ",")
	if len(parts) < 3 {
		return errors.New("insufficient input arguments")
	}
	shc.VideoType = strings.TrimSpace(parts[0])
	shc.Sport = strings.TrimSpace(parts[1])
	shc.Focus = strings.TrimSpace(parts[2])
	return nil
}

// GenerateScript generates a script based on the sports highlights video inputs
func (shc *SportsHighlightsCustomizer) GenerateScript() (string, error) {
	// Custom logic for script generation
	return fmt.Sprintf("Generating a %s highlight video focusing on %s in %s.", shc.VideoType, shc.Focus, shc.Sport), nil
}

// GetVideoType returns the type of video being generated
func (shc *SportsHighlightsCustomizer) GetVideoType() string {
	return shc.VideoType
}
