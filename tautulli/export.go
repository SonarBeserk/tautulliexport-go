package tautulli

import (
	"encoding/json"
	"fmt"
	"os"
)

// MovieExport represents the contents of a Tautulli export file for a Movie
// Current supported fields include Metadata Level 1 and Media Info Level 1 fields
type MovieExport struct {
	AddedAt               string   `json:"addedAt"`
	AudienceRating        float32  `json:"audienceRating"`
	AudienceRatingImage   string   `json:"audienceRatingImage"`
	ContentRating         string   `json:"contentRating"`
	Duration              int      `json:"duration"`
	DurationHuman         string   `json:"durationHuman"`
	EditionTitle          *string  `json:"editionTitle"`
	Guid                  string   `json:"guid"`
	HasCreditsMarker      bool     `json:"hasCreditsMarker"`
	Locations             []string `json:"locations"`
	OriginalTitle         *string  `json:"originalTitle"`
	OriginallyAvailableAt string   `json:"originallyAvailableAt"`
	Rating                float32  `json:"rating"`
	RatingImage           string   `json:"ratingImage"`
	RatingKey             int      `json:"ratingKey"`
	Studio                string   `json:"studio"`
	Summary               string   `json:"summary"`
	Tagline               string   `json:"tagline"`
	Title                 string   `json:"title"`
	TitleSort             string   `json:"titleSort"`
	Type                  string   `json:"type"`
	UserRating            *float32 `json:"userRating"`
	Year                  int      `json:"year"`
	MediaFiles            []Media  `json:"media"`
}

// Media represents the media files for each title
type Media struct {
	AspectRatio        float32 `json:"aspectRatio"`
	AudioChannels      int     `json:"audioChannels"`
	AudioCodec         string  `json:"audioCodec"`
	AudioProfile       *string `json:"audioProfile"`
	Bitrate            int     `json:"bitrate"`
	Container          string  `json:"container"`
	Duration           int     `json:"duration"`
	Hdr                bool    `json:"hdr"`
	Height             int     `json:"height"`
	IsOptimizedVersion bool    `json:"isOptimizedVersion"`
	VideoCodec         string  `json:"videoCodec"`
	VideoFrameRate     string  `json:"videoFrameRate"`
	VideoProfile       string  `json:"videoProfile"`
	VideoResolution    string  `json:"videoResolution"`
	Width              int     `json:"width"`
}

func ParseJsonMovieExport(path string) ([]MovieExport, error) {
	var export []MovieExport
	exportBytes, err := os.ReadFile(path)
	if err != nil {
		return export, fmt.Errorf("error reading file: %w", err)
	}

	err = json.Unmarshal(exportBytes, &export)
	if err != nil {
		fmt.Println("error:", err)
	}

	return export, nil
}
