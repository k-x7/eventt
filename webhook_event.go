package eventt

import "time"

type WebhookEvent struct {
	EventType string `json:"eventType"`
}

type eventType interface {
	eventName() string
}

// WebhookEventTypes
// see: https://github.com/Sonarr/Sonarr/blob/v3.0.9.1549/src/NzbDrone.Core/Notifications/Webhook/WebhookEventType.cs#L9
const (
	// see GrabEvent
	grab = "Grab"

	// see DownloadEvent
	download = "Download"

	// see RenameEvent
	rename = "Rename"

	// see EpisodeFileDeleteEvent
	episodeFileDelete = "EpisodeFileDelete"

	// see SeriesDeleteEvent
	seriesDelete = "SeriesDelete"

	// see HealthIssueEvent
	health = "Health"

	// see ApplicationUpdateEvent
	applicationUpdate = "ApplicationUpdate"

	// see TestEvent
	test = "Test"
)

// GrabEvent webhook grab payload
// see: https://github.com/Sonarr/Sonarr/blob/v3.0.9.1549/src/NzbDrone.Core/Notifications/Webhook/Webhook.cs#L23
type GrabEvent struct {
	Series struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Path     string `json:"path"`
		TvdbID   int    `json:"tvdbId"`
		TvMazeID int    `json:"tvMazeId"`
		ImdbID   string `json:"imdbId"`
		Type     string `json:"type"`
	} `json:"series"`
	Episodes []struct {
		ID            int       `json:"id"`
		EpisodeNumber int       `json:"episodeNumber"`
		SeasonNumber  int       `json:"seasonNumber"`
		Title         string    `json:"title"`
		AirDate       string    `json:"airDate"`
		AirDateUtc    time.Time `json:"airDateUtc"`
	} `json:"episodes"`
	Release struct {
		Quality        string `json:"quality"`
		QualityVersion int    `json:"qualityVersion"`
		ReleaseGroup   string `json:"releaseGroup"`
		ReleaseTitle   string `json:"releaseTitle"`
		Indexer        string `json:"indexer"`
		Size           int    `json:"size"`
	} `json:"release"`
	DownloadClient     string `json:"downloadClient"`
	DownloadClientType string `json:"downloadClientType"`
	DownloadID         string `json:"downloadId"`
	EventType          string `json:"eventType"`
}

func (e GrabEvent) eventName() string {
	return grab
}

// DownloadEvent webhook download payload
// see: https://github.com/Sonarr/Sonarr/blob/v3.0.9.1549/src/NzbDrone.Core/Notifications/Webhook/Webhook.cs#L42
type DownloadEvent struct {
	Series struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Path     string `json:"path"`
		TvdbID   int    `json:"tvdbId"`
		TvMazeID int    `json:"tvMazeId"`
		ImdbID   string `json:"imdbId"`
		Type     string `json:"type"`
	} `json:"series"`
	Episodes []struct {
		ID            int       `json:"id"`
		EpisodeNumber int       `json:"episodeNumber"`
		SeasonNumber  int       `json:"seasonNumber"`
		Title         string    `json:"title"`
		AirDate       string    `json:"airDate"`
		AirDateUtc    time.Time `json:"airDateUtc"`
	} `json:"episodes"`
	EpisodeFile struct {
		ID             int    `json:"id"`
		RelativePath   string `json:"relativePath"`
		Path           string `json:"path"`
		Quality        string `json:"quality"`
		QualityVersion int    `json:"qualityVersion"`
		ReleaseGroup   string `json:"releaseGroup"`
		Size           int    `json:"size"`
	} `json:"episodeFile"`
	IsUpgrade bool   `json:"isUpgrade"`
	EventType string `json:"eventType"`
}

func (e DownloadEvent) eventName() string {
	return download
}

// RenameEvent webhook rename payload
// see: https://github.com/Sonarr/Sonarr/blob/v3.0.9.1549/src/NzbDrone.Core/Notifications/Webhook/Webhook.cs#L71
type RenameEvent struct {
	Series struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Path     string `json:"path"`
		TvdbID   int    `json:"tvdbId"`
		TvMazeID int    `json:"tvMazeId"`
		ImdbID   string `json:"imdbId"`
		Type     string `json:"type"`
	} `json:"series"`
	RenamedEpisodeFiles []struct {
		PreviousRelativePath string `json:"previousRelativePath"`
		PreviousPath         string `json:"previousPath"`
		ID                   int    `json:"id"`
		RelativePath         string `json:"relativePath"`
		Quality              string `json:"quality"`
		QualityVersion       int    `json:"qualityVersion"`
		ReleaseGroup         string `json:"releaseGroup"`
		SceneName            string `json:"sceneName"`
		Size                 int    `json:"size"`
	} `json:"renamedEpisodeFiles"`
	EventType string `json:"eventType"`
}

func (e RenameEvent) eventName() string {
	return rename
}

// EpisodeFileDeleteEvent webhook episode file delete payload
// see: https://github.com/Sonarr/Sonarr/blob/v3.0.9.1549/src/NzbDrone.Core/Notifications/Webhook/Webhook.cs#L83
type EpisodeFileDeleteEvent struct {
	Series struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Path     string `json:"path"`
		TvdbID   int    `json:"tvdbId"`
		TvMazeID int    `json:"tvMazeId"`
		ImdbID   string `json:"imdbId"`
		Type     string `json:"type"`
	} `json:"series"`
	Episodes []struct {
		ID            int       `json:"id"`
		EpisodeNumber int       `json:"episodeNumber"`
		SeasonNumber  int       `json:"seasonNumber"`
		Title         string    `json:"title"`
		AirDate       string    `json:"airDate"`
		AirDateUtc    time.Time `json:"airDateUtc"`
	} `json:"episodes"`
	EpisodeFile struct {
		SeriesID     int       `json:"seriesId"`
		SeasonNumber int       `json:"seasonNumber"`
		RelativePath string    `json:"relativePath"`
		Path         string    `json:"path"`
		Size         int       `json:"size"`
		DateAdded    time.Time `json:"dateAdded"`
		ReleaseGroup string    `json:"releaseGroup"`
		Quality      struct {
			Quality struct {
				ID         int    `json:"id"`
				Name       string `json:"name"`
				Source     string `json:"source"`
				Resolution int    `json:"resolution"`
			} `json:"quality"`
			Revision struct {
				Version  int  `json:"version"`
				Real     int  `json:"real"`
				IsRepack bool `json:"isRepack"`
			} `json:"revision"`
		} `json:"quality"`
		MediaInfo struct {
			ContainerFormat                    string  `json:"containerFormat"`
			VideoFormat                        string  `json:"videoFormat"`
			VideoCodecID                       string  `json:"videoCodecID"`
			VideoProfile                       string  `json:"videoProfile"`
			VideoCodecLibrary                  string  `json:"videoCodecLibrary"`
			VideoBitrate                       int     `json:"videoBitrate"`
			VideoBitDepth                      int     `json:"videoBitDepth"`
			VideoMultiViewCount                int     `json:"videoMultiViewCount"`
			VideoColourPrimaries               string  `json:"videoColourPrimaries"`
			VideoTransferCharacteristics       string  `json:"videoTransferCharacteristics"`
			VideoHdrFormat                     string  `json:"videoHdrFormat"`
			VideoHdrFormatCompatibility        string  `json:"videoHdrFormatCompatibility"`
			Width                              int     `json:"width"`
			Height                             int     `json:"height"`
			AudioFormat                        string  `json:"audioFormat"`
			AudioCodecID                       string  `json:"audioCodecID"`
			AudioCodecLibrary                  string  `json:"audioCodecLibrary"`
			AudioAdditionalFeatures            string  `json:"audioAdditionalFeatures"`
			AudioBitrate                       int     `json:"audioBitrate"`
			RunTime                            string  `json:"runTime"`
			AudioStreamCount                   int     `json:"audioStreamCount"`
			AudioChannelsContainer             int     `json:"audioChannelsContainer"`
			AudioChannelsStream                int     `json:"audioChannelsStream"`
			AudioChannelPositions              string  `json:"audioChannelPositions"`
			AudioChannelPositionsTextContainer string  `json:"audioChannelPositionsTextContainer"`
			AudioChannelPositionsTextStream    string  `json:"audioChannelPositionsTextStream"`
			AudioProfile                       string  `json:"audioProfile"`
			VideoFps                           float64 `json:"videoFps"`
			AudioLanguages                     string  `json:"audioLanguages"`
			Subtitles                          string  `json:"subtitles"`
			ScanType                           string  `json:"scanType"`
			SchemaRevision                     int     `json:"schemaRevision"`
		} `json:"mediaInfo"`
		Episodes struct {
			Value []struct {
				SeriesID                   int       `json:"seriesId"`
				TvdbID                     int       `json:"tvdbId"`
				EpisodeFileID              int       `json:"episodeFileId"`
				SeasonNumber               int       `json:"seasonNumber"`
				EpisodeNumber              int       `json:"episodeNumber"`
				Title                      string    `json:"title"`
				AirDate                    string    `json:"airDate"`
				AirDateUtc                 time.Time `json:"airDateUtc"`
				Overview                   string    `json:"overview"`
				Monitored                  bool      `json:"monitored"`
				AbsoluteEpisodeNumber      int       `json:"absoluteEpisodeNumber"`
				SceneAbsoluteEpisodeNumber int       `json:"sceneAbsoluteEpisodeNumber"`
				SceneSeasonNumber          int       `json:"sceneSeasonNumber"`
				SceneEpisodeNumber         int       `json:"sceneEpisodeNumber"`
				UnverifiedSceneNumbering   bool      `json:"unverifiedSceneNumbering"`
				Ratings                    struct {
					Votes int     `json:"votes"`
					Value float64 `json:"value"`
				} `json:"ratings"`
				Images []struct {
					CoverType string `json:"coverType"`
					URL       string `json:"url"`
				} `json:"images"`
				EpisodeFile struct {
					IsLoaded bool `json:"isLoaded"`
				} `json:"episodeFile"`
				HasFile bool `json:"hasFile"`
				ID      int  `json:"id"`
			} `json:"value"`
			IsLoaded bool `json:"isLoaded"`
		} `json:"episodes"`
		Series struct {
			Value struct {
				TvdbID            int       `json:"tvdbId"`
				TvRageID          int       `json:"tvRageId"`
				TvMazeID          int       `json:"tvMazeId"`
				ImdbID            string    `json:"imdbId"`
				Title             string    `json:"title"`
				CleanTitle        string    `json:"cleanTitle"`
				SortTitle         string    `json:"sortTitle"`
				Status            string    `json:"status"`
				Overview          string    `json:"overview"`
				AirTime           string    `json:"airTime"`
				Monitored         bool      `json:"monitored"`
				QualityProfileID  int       `json:"qualityProfileId"`
				LanguageProfileID int       `json:"languageProfileId"`
				SeasonFolder      bool      `json:"seasonFolder"`
				LastInfoSync      time.Time `json:"lastInfoSync"`
				Runtime           int       `json:"runtime"`
				Images            []struct {
					CoverType string `json:"coverType"`
					URL       string `json:"url"`
				} `json:"images"`
				SeriesType        string `json:"seriesType"`
				Network           string `json:"network"`
				UseSceneNumbering bool   `json:"useSceneNumbering"`
				TitleSlug         string `json:"titleSlug"`
				Path              string `json:"path"`
				Year              int    `json:"year"`
				Ratings           struct {
					Votes int     `json:"votes"`
					Value float64 `json:"value"`
				} `json:"ratings"`
				Genres []string `json:"genres"`
				Actors []struct {
					Name      string        `json:"name"`
					Character string        `json:"character"`
					Images    []interface{} `json:"images"`
				} `json:"actors"`
				Certification  string    `json:"certification"`
				Added          time.Time `json:"added"`
				FirstAired     time.Time `json:"firstAired"`
				QualityProfile struct {
					Value struct {
						Name           string `json:"name"`
						UpgradeAllowed bool   `json:"upgradeAllowed"`
						Cutoff         int    `json:"cutoff"`
						Items          []struct {
							Quality struct {
								ID         int    `json:"id"`
								Name       string `json:"name"`
								Source     string `json:"source"`
								Resolution int    `json:"resolution"`
							} `json:"quality,omitempty"`
							Items   []interface{} `json:"items"`
							Allowed bool          `json:"allowed"`
							ID      int           `json:"id,omitempty"`
							Name    string        `json:"name,omitempty"`
						} `json:"items"`
						ID int `json:"id"`
					} `json:"value"`
					IsLoaded bool `json:"isLoaded"`
				} `json:"qualityProfile"`
				LanguageProfile struct {
					Value struct {
						Name      string `json:"name"`
						Languages []struct {
							Language struct {
								ID   int    `json:"id"`
								Name string `json:"name"`
							} `json:"language"`
							Allowed bool `json:"allowed"`
						} `json:"languages"`
						UpgradeAllowed bool `json:"upgradeAllowed"`
						Cutoff         struct {
							ID   int    `json:"id"`
							Name string `json:"name"`
						} `json:"cutoff"`
						ID int `json:"id"`
					} `json:"value"`
					IsLoaded bool `json:"isLoaded"`
				} `json:"languageProfile"`
				Seasons []struct {
					SeasonNumber int  `json:"seasonNumber"`
					Monitored    bool `json:"monitored"`
					Images       []struct {
						CoverType string `json:"coverType"`
						URL       string `json:"url"`
					} `json:"images"`
				} `json:"seasons"`
				Tags []int `json:"tags"`
				ID   int   `json:"id"`
			} `json:"value"`
			IsLoaded bool `json:"isLoaded"`
		} `json:"series"`
		Language struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"language"`
		ID int `json:"id"`
	} `json:"episodeFile"`
	DeleteReason string `json:"deleteReason"`
	EventType    string `json:"eventType"`
}

func (e EpisodeFileDeleteEvent) eventName() string {
	return episodeFileDelete
}

// SeriesDeleteEvent webhook series delete payload
// see: https://github.com/Sonarr/Sonarr/blob/v3.0.9.1549/src/NzbDrone.Core/Notifications/Webhook/Webhook.cs#L97
type SeriesDeleteEvent struct {
	Series struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Path     string `json:"path"`
		TvdbID   int    `json:"tvdbId"`
		TvMazeID int    `json:"tvMazeId"`
		ImdbID   string `json:"imdbId"`
		Type     string `json:"type"`
	} `json:"series"`
	DeletedFiles bool   `json:"deletedFiles"`
	EventType    string `json:"eventType"`
}

func (e SeriesDeleteEvent) eventName() string {
	return seriesDelete
}

// Health webhook health payload
// see: https://github.com/Sonarr/Sonarr/blob/v3.0.9.1549/src/NzbDrone.Core/Notifications/Webhook/Webhook.cs#L109
type HealthEvent struct {
	Level     string `json:"level"`
	Message   string `json:"message"`
	Type      string `json:"type"`
	WikiURL   string `json:"wikiUrl"`
	EventType string `json:"eventType"`
}

func (e HealthEvent) eventName() string {
	return health
}

// ApplicationUpdateEvent webhook application update payload
// see: https://github.com/Sonarr/Sonarr/blob/v3.0.9.1549/src/NzbDrone.Core/Notifications/Webhook/Webhook.cs#L123
type ApplicationUpdateEvent struct {
	Message         string `json:"message"`
	PreviousVersion string `json:"previousVersion"`
	NewVersion      string `json:"newVersion"`
	EventType       string `json:"eventType"`
}

func (e ApplicationUpdateEvent) eventName() string {
	return applicationUpdate
}

// TestEvent webhook test payload
// see: https://github.com/Sonarr/Sonarr/blob/v3.0.9.1549/src/NzbDrone.Core/Notifications/Webhook/Webhook.cs#L153
type TestEvent struct {
	Series struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Path     string `json:"path"`
		TvdbID   int    `json:"tvdbId"`
		TvMazeID int    `json:"tvMazeId"`
		Type     string `json:"type"`
	} `json:"series"`
	Episodes []struct {
		ID            int    `json:"id"`
		EpisodeNumber int    `json:"episodeNumber"`
		SeasonNumber  int    `json:"seasonNumber"`
		Title         string `json:"title"`
	} `json:"episodes"`
	EventType string `json:"eventType"`
}

func (e TestEvent) eventName() string {
	return test
}

// UnknownEvent parse any unknown events, this could happened if Sonarr update or add
// new webhook events or change them, like what happened when OnImport and OnDownload.
type UnknownEvent map[string]interface{}

func (e UnknownEvent) eventName() string {
	return "Unknown"
}
