package bunny

import (
	"context"
	"fmt"
)

// VideoLibraryUpdateOptions represents the request parameters for the Update Storage
// Zone API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_updatepullzone
type VideoLibraryUpdateOptions struct {
	Name                             *string `json:"Name,omitempty"`
	CustomHTML                       *string `json:"CustomHTML,omitempty"`
	PlayerKeyColor                   *string `json:"PlayerKeyColor,omitempty"`
	EnableTokenAuthentication        *bool   `json:"EnableTokenAuthentication,omitempty"`
	EnableTokenIPVerification        *bool   `json:"EnableTokenIPVerification,omitempty"`
	ResetToken                       *bool   `json:"ResetToken,omitempty"`
	WatermarkPositionLeft            *int32  `json:"WatermarkPositionLeft,omitempty"`
	WatermarkPositionTop             *int32  `json:"WatermarkPositionTop,omitempty"`
	WatermarkWidth                   *int32  `json:"WatermarkWidth,omitempty"`
	WatermarkHeight                  *int32  `json:"WatermarkHeight,omitempty"`
	EnabledResolutions               *string `json:"EnabledResolutions,omitempty"`
	ViAiPublisherID                  *string `json:"ViAiPublisherId,omitempty"`
	VastTagURL                       *string `json:"VastTagUrl,omitempty"`
	WebhookURL                       *string `json:"WebhookUrl,omitempty"`
	CaptionsFontSize                 *int32  `json:"CaptionsFontSize,omitempty"`
	CaptionsFontColor                *string `json:"CaptionsFontColor,omitempty"`
	CaptionsBackground               *string `json:"CaptionsBackground,omitempty"`
	UILanguage                       *string `json:"UILanguage,omitempty"`
	AllowEarlyPlay                   *bool   `json:"AllowEarlyPlay,omitempty"`
	PlayerTokenAuthenticationEnabled *bool   `json:"PlayerTokenAuthenticationEnabled,omitempty"`
	BlockNoneReferrer                *bool   `json:"BlockNoneReferrer,omitempty"`
	EnableMP4Fallback                *bool   `json:"EnableMP4Fallback,omitempty"`
	KeepOriginalFiles                *bool   `json:"KeepOriginalFiles,omitempty"`
	AllowDirectPlay                  *bool   `json:"AllowDirectPlay,omitempty"`
	EnableDRM                        *bool   `json:"EnableDRM,omitempty"`
	Controls                         *string `json:"Controls,omitempty"`
	Bitrate240p                      *int32  `json:"Bitrate240p,omitempty"`
	Bitrate360p                      *int32  `json:"Bitrate360p,omitempty"`
	Bitrate480p                      *int32  `json:"Bitrate480p,omitempty"`
	Bitrate720p                      *int32  `json:"Bitrate720p,omitempty"`
	Bitrate1080p                     *int32  `json:"Bitrate1080p,omitempty"`
	Bitrate1440p                     *int32  `json:"Bitrate1440p,omitempty"`
	Bitrate2160p                     *int32  `json:"Bitrate2160p,omitempty"`
	ShowHeatmap                      *bool   `json:"ShowHeatmap,omitempty"`
	EnableContentTagging             *bool   `json:"EnableContentTagging,omitempty"`
	FontFamily                       *string `json:"FontFamily,omitempty"`
}

// Update changes the configuration the Video Library with the given ID.
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_updatepullzone
func (s *VideoLibraryService) Update(ctx context.Context, id int64, opts *VideoLibraryUpdateOptions) (*VideoLibrary, error) {
	path := fmt.Sprintf("videolibrary/%d", id)
	return resourcePostWithResponse[VideoLibrary](
		ctx,
		s.client,
		path,
		opts,
	)
}
