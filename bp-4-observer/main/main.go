package main

import (
	"bp-4-observer/structs"
	"fmt"
)

func main() {
	fmt.Println("subscription mechanism simulator is starting...")
	waterBall := structs.NewChannelSubscriber("水球", structs.NewWaterBallChannelSubscriber())
	fireBall := structs.NewChannelSubscriber("火球", structs.NewFireBallChannelSubscriber())

	pewDiePie := structs.NewYoutubeChannel("PewDiePie")
	waterballSA := structs.NewYoutubeChannel("水球軟體學院")

	C1M1S2 := structs.NewVideo("C1M1S2", "這個世界正是物件導向的呢！", structs.NewTime(240), waterballSA)
	HelloGuys := structs.NewVideo("Hello guys", "Clickbait", structs.NewTime(30), pewDiePie)
	C1M1S3 := structs.NewVideo("C1M1S3", "物件 vs. 類別", structs.NewTime(60), waterballSA)
	Minecraft := structs.NewVideo("Minecraft", "Let’s play Minecraft", structs.NewTime(1800), pewDiePie)

	waterBall.Subscribe(waterballSA)
	waterBall.Subscribe(pewDiePie)
	fireBall.Subscribe(waterballSA)
	fireBall.Subscribe(pewDiePie)

	waterballSA.UploadNewVideoAndNotifyChannelSubscribers(C1M1S2)
	pewDiePie.UploadNewVideoAndNotifyChannelSubscribers(HelloGuys)
	waterballSA.UploadNewVideoAndNotifyChannelSubscribers(C1M1S3)
	pewDiePie.UploadNewVideoAndNotifyChannelSubscribers(Minecraft)
}
