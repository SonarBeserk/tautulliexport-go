package tautulli_test

import (
	"testing"

	"github.com/SonarBeserk/tautulliexport-go/tautulli"
	"github.com/google/go-cmp/cmp"
)

func TestMovieExamples(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		file            string
		expectedExports []tautulli.MovieExport
	}{
		// the table itself
		{"../examples/json/Movie - The Polar Express [842].json", []tautulli.MovieExport{
			{
				AddedAt:             "2024-03-09T19:34:13",
				AudienceRating:      6.4,
				AudienceRatingImage: "rottentomatoes://image.rating.upright",
				ContentRating:       "G",
				Duration:            5994400,
				DurationHuman:       "1 hr 40 mins",
				EditionTitle:        nil,
				Guid:                "plex://movie/5d77682e6f4521001ea9ac27",
				HasCreditsMarker:    true,
				Locations: []string{
					"/mnt/media/Jellyfin/Movies/The Polar Express (2004)/The Polar Express (2004) - 2160p.mkv",
				},
				OriginalTitle:         nil,
				OriginallyAvailableAt: "2004-11-10",
				Rating:                5.5,
				RatingImage:           "rottentomatoes://image.rating.rotten",
				RatingKey:             842,
				Studio:                "Golden Mean",
				Summary:               "This is the story of a young hero boy who boards on a powerful magical train that's headed to the North Pole and Santa Claus' home on Christmas Eve night. During this ride, he embarks on a journey of self-discovery which shows him that the wonder of life never fades for those who believe.",
				Tagline:               "This holiday season... believe.",
				Title:                 "The Polar Express",
				TitleSort:             "Polar Express",
				Type:                  "movie",
				UserRating:            nil,
				Year:                  2004,
				MediaFiles: []tautulli.Media{
					{
						AspectRatio:        2.35,
						AudioChannels:      6,
						AudioCodec:         "dca",
						AudioProfile:       toStringPtr("dts"),
						Bitrate:            11322,
						Container:          "mkv",
						Duration:           5994400,
						Hdr:                false,
						Height:             1608,
						Width:              3840,
						IsOptimizedVersion: false,
						VideoCodec:         "hevc",
						VideoFrameRate:     "24p",
						VideoProfile:       "main",
						VideoResolution:    "4k",
					},
				},
			},
		}},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.file, func(t *testing.T) {
			export, err := tautulli.ParseJsonMovieExport(tt.file)
			if err != nil {
				t.Errorf("error occurred: %s", err)
			}

			diff := cmp.Diff(export, tt.expectedExports)
			if diff != "" {
				t.Errorf("export values do not match (-want +got):\n%s", diff)
			}
		})
	}
}

func toStringPtr(v string) *string { return &v }
