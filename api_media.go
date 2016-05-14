package godingtalk

//MediaResponse is
type MediaResponse struct {
	OAPIResponse
	Type  string
	MediaID int `json:"media_id"`
}

//UploadMedia is to upload media file to DingTalk
func (c *DingTalkClient) UploadMedia() (media MediaResponse, err error) {
	return media, err
}
