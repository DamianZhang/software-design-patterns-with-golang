package structs

import "fmt"

type ChannelSubscriber struct {
	name            string
	youtubeChannels []*YoutubeChannel
	polytype        IChannelSubscriber
}

type IChannelSubscriber interface {
	GetNotifiedOfNewVideo(newVideo *Video, channelSubscriber *ChannelSubscriber)
}

func NewChannelSubscriber(name string, polytype IChannelSubscriber) *ChannelSubscriber {
	return &ChannelSubscriber{
		name:            name,
		youtubeChannels: make([]*YoutubeChannel, 0),
		polytype:        polytype,
	}
}

func (c *ChannelSubscriber) Subscribe(youtubeChannel *YoutubeChannel) {
	c.youtubeChannels = append(c.youtubeChannels, youtubeChannel)
	youtubeChannel.AddChannelSubscriber(c)

	fmt.Printf("%s 訂閱了 %s。\n", c.Name(), youtubeChannel.Name())
}

func (c *ChannelSubscriber) GetNotifiedOfNewVideo(newVideo *Video) {
	c.polytype.GetNotifiedOfNewVideo(newVideo, c)
}

func (c *ChannelSubscriber) Unsubscribe(youtubeChannel *YoutubeChannel) {
	for i, y := range c.youtubeChannels {
		if y == youtubeChannel {
			c.youtubeChannels = append(c.youtubeChannels[:i], c.youtubeChannels[i+1:]...)
			break
		}
	}
	youtubeChannel.RemoveChannelSubscriber(c)

	fmt.Printf("%s 解除訂閱了 %s。\n", c.Name(), youtubeChannel.Name())
}

func (c *ChannelSubscriber) Name() string {
	return c.name
}
