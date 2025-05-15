package tautulli

// MovieExport represents the contents of a Tautulli export file for a Movie
// Current supported fields include Metadata Level 1 and Media Info Level 1 fields
type MovieExport struct {
	AddedAt               string   `json:"addedAt" csv:"addedAt"`
	AudienceRating        float32  `json:"audienceRating" csv:"audienceRating"`
	AudienceRatingImage   string   `json:"audienceRatingImage" csv:"audienceRatingImage"`
	ContentRating         string   `json:"contentRating" csv:"contentRating"`
	Duration              int      `json:"duration" csv:"duration"`
	DurationHuman         string   `json:"durationHuman" csv:"durationHuman"`
	EditionTitle          string   `json:"editionTitle" csv:"editionTitle"`
	Guid                  string   `json:"guid" csv:"guid"`
	HasCreditsMarker      bool     `json:"hasCreditsMarker" csv:"hasCreditsMarker"`
	Locations             []string `json:"locations" csv:"locations"`
	OriginalTitle         string   `json:"originalTitle" csv:"originalTitle"`
	OriginallyAvailableAt string   `json:"originallyAvailableAt" csv:"originallyAvailableAt"`
	Rating                float32  `json:"rating" csv:"rating"`
	RatingImage           string   `json:"ratingImage" csv:"ratingImage"`
	RatingKey             int      `json:"ratingKey" csv:"ratingKey"`
	Studio                string   `json:"studio" csv:"studio"`
	Summary               string   `json:"summary" csv:"summary"`
	Tagline               string   `json:"tagline" csv:"tagline"`
	Title                 string   `json:"title" csv:"title"`
	TitleSort             string   `json:"titleSort" csv:"titleSort"`
	Type                  string   `json:"type" csv:"type"`
	UserRating            *float32 `json:"userRating" csv:"userRating"`
	Year                  int      `json:"year" csv:"year"`
	MediaFiles            []Media  `json:"media" csv:"media"`
}

// Media represents the media files for each title
type Media struct {
	AspectRatio        float32 `json:"aspectRatio" csv:"aspectRatio"`
	AudioChannels      int     `json:"audioChannels" csv:"audioChannels"`
	AudioCodec         string  `json:"audioCodec" csv:"audioCodec"`
	AudioProfile       *string `json:"audioProfile" csv:"audioProfile"`
	Bitrate            int     `json:"bitrate" csv:"bitrate"`
	Container          string  `json:"container" cvs:"container"`
	Duration           int     `json:"duration" csv:"duration"`
	Hdr                bool    `json:"hdr" csv:"hdr"`
	Height             int     `json:"height" csv:"height"`
	IsOptimizedVersion bool    `json:"isOptimizedVersion" csv:"isOptimizedVersion"`
	VideoCodec         string  `json:"videoCodec" csv:"videoCodec"`
	VideoFrameRate     string  `json:"videoFrameRate" csv:"videoFrameRate"`
	VideoProfile       string  `json:"videoProfile" csv:"videoProfile"`
	VideoResolution    string  `json:"videoResolution" csv:"videoResolution"`
	Width              int     `json:"width" csv:"width"`
}
