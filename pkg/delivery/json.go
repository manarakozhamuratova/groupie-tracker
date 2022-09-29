package delivery

import (
	"encoding/json"
	"net/http"
)

const (
	ArtistURL   = "https://groupietrackers.herokuapp.com/api/artists"
	RelationURL = "https://groupietrackers.herokuapp.com/api/relation"
)

func parseJson(url string, dataBox interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return json.NewDecoder(response.Body).Decode(dataBox)
}

func (h *Handler) jsonInfo() error {
	if len(h.Artists) != 0 {
		return nil
	}
	if err := parseJson(ArtistURL, &h.Artists); err != nil {
		return err
	}

	if err := parseJson(RelationURL, &h.Relation); err != nil {
		return err
	}

	for i, w := range h.Relation.Index {
		h.Artists[i].DatesLocations = w.DatesLocations
	}
	return nil
}
