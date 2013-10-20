package main

import (
	"errors"
	"github.com/jmcvetta/napping"
	"net/http"
	"time"
)

const CLIENT_ID = "c89c631231435d7320d3d93f2dca04b8"
const SC_API = "http://api.soundcloud.com/"

var apiClient napping.Session

func init() {
	apiClient.Client = http.DefaultClient
}

type ApiError struct {
	Errors []struct {
		Error_Message string `json:"error_message"`
	} `json:"errors"`
}

func (err ApiError) String() string {
	if len(err.Errors) > 0 {
		return err.Errors[0].Error_Message
	} else {
		return "unknown error"
	}
}

type ApiTrack struct {
	Duration int    `json:"duration"`
	Link     string `json:"permalink_url"`
	Title    string `json:"title"`
	User     struct {
		Username string `json:"username"`
	} `json:"user"`
	Streamable bool `json:"streamable"`
}

func trackInfo(trackId string) (*Track, error) {
	result := new(ApiTrack)
	result.Duration = -1
	apiError := new(ApiError)
	resp, err := apiClient.Get(
		SC_API+"tracks/"+trackId+".json",
		&napping.Params{"client_id": CLIENT_ID},
		result, apiError,
	)
	if err != nil {
		return nil, err
	}
	if resp.HttpResponse().StatusCode != 200 {
		return nil, errors.New(apiError.String())
	}
	if !result.Streamable {
		return nil, errors.New("track not streamable")
	}
	return &Track{
		trackId,
		time.Millisecond * time.Duration(result.Duration),
		result.Link,
		result.User.Username,
		result.Title,
	}, nil
}
