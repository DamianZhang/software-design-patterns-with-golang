package structs

import "fmt"

type YoutubeChannel struct {
	name               string
	videos             []*Video
	channelSubscribers []*ChannelSubscriber
}

func NewYoutubeChannel(name string) *YoutubeChannel {
	return &YoutubeChannel{
		name:               name,
		videos:             make([]*Video, 0),
		channelSubscribers: make([]*ChannelSubscriber, 0),
	}
}

func (y *YoutubeChannel) uploadNewVideo(newVideo *Video) {
	y.videos = append(y.videos, newVideo)

	fmt.Printf("頻道 %s 上架了一則新影片 \"%s\"。\n", y.Name(), newVideo.Title())
}

func (y *YoutubeChannel) notifyChannelSubscribers(newVideo *Video) {
	for _, channelSubscriber := range y.channelSubscribers {
		channelSubscriber.GetNotifiedOfNewVideo(newVideo)
	}
}

func (y *YoutubeChannel) UploadNewVideoAndNotifyChannelSubscribers(newVideo *Video) {
	y.uploadNewVideo(newVideo)
	y.notifyChannelSubscribers(newVideo)
}

func (y *YoutubeChannel) AddChannelSubscriber(channelSubscriber *ChannelSubscriber) {
	y.channelSubscribers = append(y.channelSubscribers, channelSubscriber)
}

func (y *YoutubeChannel) RemoveChannelSubscriber(channelSubscriber *ChannelSubscriber) {
	for i, c := range y.channelSubscribers {
		if c == channelSubscriber {
			y.channelSubscribers = append(y.channelSubscribers[:i], y.channelSubscribers[i+1:]...)
			break
		}
	}
}

func (y *YoutubeChannel) Name() string {
	return y.name
}
