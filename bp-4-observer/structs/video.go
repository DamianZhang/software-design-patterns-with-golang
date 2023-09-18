package structs

type Video struct {
	title       string
	description string
	length      Time
	like        int
	owner       *YoutubeChannel
}

func NewVideo(title, description string, length Time, owner *YoutubeChannel) *Video {
	return &Video{
		title:       title,
		description: description,
		length:      length,
		like:        0,
		owner:       owner,
	}
}

func (v *Video) GainLike() {
	v.like++
}

func (v *Video) Title() string {
	return v.title
}

func (v *Video) Length() Time {
	return v.length
}

func (v *Video) Owner() *YoutubeChannel {
	return v.owner
}
