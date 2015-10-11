package tvmaze

import (
	"fmt"
	"time"

	"github.com/simplereach/timeutils"
)

// ShowResponse wraps a TV Maze search response
type ShowResponse struct {
	Score float64
	Show  Show
}

// Show wraps a TV Maze show object
type Show struct {
	ID        int
	Name      string
	Type      string
	Genres    []string
	Status    string
	Runtime   int
	Premiered timeutils.Time
	Summary   string
	Network   network
	Embeds    struct {
		Episodes []Episode
	} `json:"_embedded"`
	Remotes map[string]int `json:"externals"`
}

// GetTitle return the show title
func (s Show) GetTitle() string {
	return s.Name
}

// GetDescription returns a summary of the show
func (s Show) GetDescription() string {
	return s.Summary
}

// GetNetwork returns the network that currently broadcasts the show
func (s Show) GetNetwork() string {
	return s.Network.Name
}

// GetFirstAired return the time the first episode was aired
func (s Show) GetFirstAired() time.Time {
	return s.Premiered.Time
}

// GetTVRageID returns the show's ID on tvrage.com
func (s Show) GetTVRageID() int {
	return s.Remotes["tvrage"]
}

// FindShows finds all matches for a given search string
func (c Client) FindShows(name string) (s []ShowResponse, err error) {
	path := baseURLWithPathQuery("search/shows", "q", name)

	if err := c.get(path, &s); err != nil {
		return nil, err
	}

	return s, nil
}

// GetShow finds all matches for a given search string
func (c Client) GetShow(name string) (*Show, error) {
	path := baseURLWithPathQuery("singlesearch/shows", "q", name)

	show := &Show{}
	if err := c.get(path, show); err != nil {
		return nil, err
	}

	return show, nil
}

// RefreshShow refreshes a show from the server
func (c Client) RefreshShow(show *Show) (err error) {
	url := baseURLWithPath(fmt.Sprintf("shows/%d", show.ID))

	if err := c.get(url, &show); err != nil {
		return err
	}

	return nil
}
