package structs

type FireBallChannelSubscriber struct{}

func NewFireBallChannelSubscriber() *FireBallChannelSubscriber {
	return &FireBallChannelSubscriber{}
}

func (f *FireBallChannelSubscriber) GetNotifiedOfNewVideo(newVideo *Video, channelSubscriber *ChannelSubscriber) {
	if newVideo.Length().Minute() <= 1 {
		channelSubscriber.Unsubscribe(newVideo.Owner())
	}
}
