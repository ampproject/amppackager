package bunny

import (
	"context"
	"fmt"
)

// VideoLibrary represents the response of the the List and Get Video Library API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/videolibrarypublic_index2 https://docs.bunny.net/reference/videolibrarypublic_index
type VideoLibrary struct {
	ID *int64 `json:"Id,omitempty"`

	Name               *string  `json:"Name,omitempty"`
	VideoCount         *int64   `json:"VideoCount,omitempty"`
	TrafficUsage       *int64   `json:"TrafficUsage,omitempty"`
	StorageUsage       *int64   `json:"StorageUsage,omitempty"`
	DateCreated        *string  `json:"DateCreated,omitempty"`
	ReplicationRegions []string `json:"ReplicationRegions,omitempty"`
	APIKey             *string  `json:"ApiKey,omitempty"`
	ReadOnlyAPIKey     *string  `json:"ReadOnlyApiKey,omitempty"`
	HasWatermark       *bool    `json:"HasWatermark,omitempty"`

	WatermarkPositionLeft *int32  `json:"WatermarkPositionLeft,omitempty"`
	WatermarkPositionTop  *int32  `json:"WatermarkPositionTop,omitempty"`
	WatermarkWidth        *int32  `json:"WatermarkWidth,omitempty"`
	PullZoneID            *int64  `json:"PullZoneId,omitempty"`
	StorageZoneID         *int64  `json:"StorageZoneId,omitempty"`
	WatermarkHeight       *int32  `json:"WatermarkHeight,omitempty"`
	EnabledResolutions    *string `json:"EnabledResolutions,omitempty"`

	ViAiPublisherID                  *string  `json:"ViAiPublisherId,omitempty"`
	VastTagURL                       *string  `json:"VastTagUrl,omitempty"`
	WebhookURL                       *string  `json:"WebhookUrl,omitempty"`
	CaptionsFontSize                 *int32   `json:"CaptionsFontSize,omitempty"`
	CaptionsFontColor                *string  `json:"CaptionsFontColor,omitempty"`
	CaptionsBackground               *string  `json:"CaptionsBackground,omitempty"`
	UILanguage                       *string  `json:"UILanguage,omitempty"`
	AllowEarlyPlay                   *bool    `json:"AllowEarlyPlay,omitempty"`
	PlayerTokenAuthenticationEnabled *bool    `json:"PlayerTokenAuthenticationEnabled,omitempty"`
	AllowedReferrers                 []string `json:"AllowedReferrers,omitempty"`
	BlockedReferrers                 []string `json:"BlockedReferrers,omitempty"`
	BlockNoneReferrer                *bool    `json:"BlockNoneReferrer,omitempty"`
	EnableMP4Fallback                *bool    `json:"EnableMP4Fallback,omitempty"`
	KeepOriginalFiles                *bool    `json:"KeepOriginalFiles,omitempty"`
	AllowDirectPlay                  *bool    `json:"AllowDirectPlay,omitempty"`
	EnableDRM                        *bool    `json:"EnableDRM,omitempty"`
	Bitrate240p                      *int32   `json:"Bitrate240p,omitempty"`
	Bitrate360p                      *int32   `json:"Bitrate360p,omitempty"`
	Bitrate480p                      *int32   `json:"Bitrate480p,omitempty"`
	Bitrate720p                      *int32   `json:"Bitrate720p,omitempty"`
	Bitrate1080p                     *int32   `json:"Bitrate1080p,omitempty"`
	Bitrate1440p                     *int32   `json:"Bitrate1440p,omitempty"`
	Bitrate2160p                     *int32   `json:"Bitrate2160p,omitempty"`
	APIAccessKey                     *string  `json:"ApiAccessKey,omitempty"`
	ShowHeatmap                      *bool    `json:"ShowHeatmap,omitempty"`
	EnableContentTagging             *bool    `json:"EnableContentTagging,omitempty"`
	PullZoneType                     *int32   `json:"PullZoneType,omitempty"`
	CustomHTML                       *string  `json:"CustomHTML,omitempty"`
	Controls                         *string  `json:"Controls,omitempty"`
	PlayerKeyColor                   *string  `json:"PlayerKeyColor,omitempty"`
	FontFamily                       *string  `json:"FontFamily,omitempty"`
}

// VideoLibraryGetOpts represents optional query parameters available when Getting or Listing Video Libraries
type VideoLibraryGetOpts struct {
	IncludeAccessKey bool `url:"includeAccessKey"`
}

// Get retrieves the Video Library with the given id.
//
// Bunny.net API docs: https://docs.bunny.net/reference/videolibrarypublic_index2
func (s *VideoLibraryService) Get(
	ctx context.Context,
	id int64,
	opts *VideoLibraryGetOpts,
) (*VideoLibrary, error) {
	path := fmt.Sprintf("videolibrary/%d", id)
	return resourceGet[VideoLibrary](ctx, s.client, path, opts)
}
