package structs

import "fmt"

type WaterBallChannelSubscriber struct{}

func NewWaterBallChannelSubscriber() *WaterBallChannelSubscriber {
	return &WaterBallChannelSubscriber{}
}

func (w *WaterBallChannelSubscriber) GetNotifiedOfNewVideo(newVideo *Video, channelSubscriber *ChannelSubscriber) {
	if newVideo.Length().Minute() >= 3 {
		newVideo.GainLike()
		fmt.Printf("%s 對影片 \"%s\" 按讚。\n", channelSubscriber.Name(), newVideo.Title())
	}
}
