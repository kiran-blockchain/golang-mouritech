package main

import (
	"fmt"
	"time"

	"problem-01/mockstream"
)

func producer(stream mockstream.Stream) (tweets []*mockstream.Tweet) {
	for {
		tweet, err := stream.Next()
		if err == mockstream.ErrEOF {
			return tweets
		}

		tweets = append(tweets, tweet)
	}
}

func consumer(tweets []*mockstream.Tweet) {
	for _, t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := mockstream.GetMockStream()

	// Producer
	tweets := producer(stream)

	// Consumer
	consumer(tweets)

	fmt.Printf("Process took %s\n", time.Since(start))
}
