package api

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

// GetSavedTracks returns a list of all the songs the user has saved
func GetSavedTracks(opts *SearchOptions) (result *SavedTracksPaged, err error) {
	v, err := query.Values(opts)

	if err != nil {
		return result, err
	}

	t := getAccessToken()

	r := buildRequest("GET", apiURLBase+"me/tracks", v, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &result)

	return result, err
}

// GetNextSavedTracks takes in the Next fields from the paging objects returned from Saved and moves forward through the results
func GetNextSavedTracks(url string) (sr *SavedTracksPaged, err error) {
	t := getAccessToken()

	r, err := http.NewRequest("GET", url, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, &sr)

	return sr, err
}

// SaveTrack takes in a TrackID and saves it to the users library
func SaveTrack(trackID string) (err error) {
	v, err := query.Values(nil)

	if err != nil {
		return err
	}

	v.Add("ids", trackID)

	t := getAccessToken()

	r := buildRequest("PUT", apiURLBase+"me/tracks", v, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, nil)

	return err
}

// RemoveSavedTrack takes in a TrackID and removes it from the users library
func RemoveSavedTrack(trackID string) (err error) {
	v, err := query.Values(nil)

	if err != nil {
		return err
	}

	v.Add("ids", trackID)

	t := getAccessToken()

	r := buildRequest("DELETE", apiURLBase+"me/tracks", v, nil)
	r.Header.Add("Authorization", "Bearer "+t)

	err = makeRequest(r, nil)

	return err
}