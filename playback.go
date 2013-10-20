package main

import (
	"log"
	"time"
)

type Track struct {
	id       string
	duration time.Duration
	link     string
	artist   string
	title    string
}

var songQueue = make([]*Track, 0)
var nowPlaying *Track

func enqueueSong(track *Track) {
	songQueue = append(songQueue, track)
	tickPlayer()
}

func nowPlayingMessage(track Track, progress time.Duration) hash {
	return hash{
		"type":     "play",
		"track":    track.id,
		"link":     track.link,
		"artist":   track.artist,
		"title":    track.title,
		"progress": 1000 * progress.Seconds(),
		"duration": 1000 * track.duration.Seconds(),
	}
}

func dequeueSong() {
	if len(songQueue) == 0 {
		nowPlaying = nil
		return
	}
	nowPlaying = songQueue[0]
	songQueue = songQueue[1:]
}

var playerTicks = make(chan empty, 1)

func tickPlayer() {
	playerTicks <- empty{}
}

func player() {
	for {
		<-playerTicks
		if nowPlaying == nil {
			dequeueSong()
			go playTrack()
		}
	}
}

var progressNotifier = make(chan *Client)
var nothingPlaying = nowPlayingMessage(Track{}, time.Duration(0))

func (c *Client) notifyProgress() {
	if nowPlaying != nil {
		progressNotifier <- c
	} else {
		c.sendQueue <- nothingPlaying
	}
}

var skipChan = make(chan empty)

func playTrack() {
	if nowPlaying == nil {
		broadcast(nothingPlaying)
		return
	}
	broadcast(nowPlayingMessage(*nowPlaying, time.Duration(0)))
	trackStart := time.Now()
	log.Println("start playback of track", nowPlaying.id)
	// Add 5 seconds to compensate for song load time:  SC is slow (~1s), but
	// mobile is slower (~3s!).  We could do a range of compensation times and
	// have the client notify us when it's ready for the next track, but this is
	// simpler for now (long-term desyncs, ew).
	trackEnd := time.After(nowPlaying.duration + 5*time.Second)
	for {
		select {
		case <-trackEnd:
			log.Println("end playback of track", nowPlaying.id)
			nowPlaying = nil
			tickPlayer()
			return
		case <-skipChan:
			log.Println("skip playback of track", nowPlaying.id)
			nowPlaying = nil
			tickPlayer()
			return
		case c := <-progressNotifier:
			progress := time.Now().Sub(trackStart)
			c.sendQueue <- nowPlayingMessage(*nowPlaying, progress)
		}
	}
}
