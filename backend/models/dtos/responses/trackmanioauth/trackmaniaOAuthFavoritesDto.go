package trackmanioauth

type TrackmaniaTracksDto struct {
	Uid          string `json:"uid"`
	Name         string `json:"name"`
	Author       string `json:"author"`
	ThumbnailUrl string `json:"thumbnailUrl"`
}
