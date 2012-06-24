package lastfm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiBaseURL = "http://ws.audioscrobbler.com/2.0/?&api_key=%s&format=json&method=%s"
)

type LastFM struct {
	ApiKey string
}

type TopArtistsResponse struct {
	TopArtists TopArtists
}

type TopArtists struct {
	Artists []Artist `json:"artist"`
}

type Artist struct {
	Name      string
	PlayCount string
	Url       string
}

type TopTracksResponse struct {
	TopTracks TopTracks
}

type TopTracks struct {
	Tracks []Track `json:"track"`
}

type Track struct {
	Name      string
	Duration  string
	PlayCount string
	Url       string
	Artist    Artist
}

type TopAlbumsResponse struct {
	TopAlbums TopAlbums
}

type TopAlbums struct {
	Albums []Album `json:"album"`
}

type Album struct {
	Name      string
	PlayCount string
	Artist    Artist
}

type RecentTracksResponse struct {
	RecentTracks RecentTracks
}

type RecentTracks struct {
	Tracks []RecentTrack `json:"track"`
}

type RecentTrack struct {
	Artist struct {
		Name string `json:"#text"`
	}

	Name string

	Album struct {
		Name string `json:"#text"`
	}

	Url string

	Date struct {
		Date string `json:"#text"`
	}
}

func (lastfm *LastFM) GetTopArtists(user string) ([]Artist, error) {
	url := fmt.Sprintf(apiBaseURL, lastfm.ApiKey, "user.gettopartists") + "&user=" + user

	body, err := getResponseBody(url)
	if err != nil {
		return nil, err
	}

	var topArtistsResponse TopArtistsResponse
	err = json.Unmarshal(body, &topArtistsResponse)
	if err != nil {
		return nil, err
	}

	return topArtistsResponse.TopArtists.Artists, nil
}

func (lastfm *LastFM) GetTopTracks(user string) ([]Track, error) {
	url := fmt.Sprintf(apiBaseURL, lastfm.ApiKey, "user.gettoptracks") + "&user=" + user

	body, err := getResponseBody(url)
	if err != nil {
		return nil, err
	}

	var topTracksResponse TopTracksResponse
	err = json.Unmarshal(body, &topTracksResponse)
	if err != nil {
		return nil, err
	}

	return topTracksResponse.TopTracks.Tracks, nil
}

func (lastfm *LastFM) GetTopAlbums(user string) ([]Album, error) {
	url := fmt.Sprintf(apiBaseURL, lastfm.ApiKey, "user.gettopalbums") + "&user=" + user

	body, err := getResponseBody(url)
	if err != nil {
		return nil, err
	}

	var topAlbumsResponse TopAlbumsResponse
	err = json.Unmarshal(body, &topAlbumsResponse)
	if err != nil {
		return nil, err
	}

	return topAlbumsResponse.TopAlbums.Albums, nil
}

func (lastfm *LastFM) GetRecentTracks(user string) ([]RecentTrack, error) {
	url := fmt.Sprintf(apiBaseURL, lastfm.ApiKey, "user.getrecenttracks") + "&user=" + user

	body, err := getResponseBody(url)
	if err != nil {
		return nil, err
	}

	var recentTracksResponse RecentTracksResponse
	err = json.Unmarshal(body, &recentTracksResponse)
	if err != nil {
		return nil, err
	}

	return recentTracksResponse.RecentTracks.Tracks, nil
}

func getResponseBody(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
