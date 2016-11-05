package football

import (
// "fmt"
)

type Highlight struct {
}

type Football struct {
	Id         string
	highlights []Highlight
}

type FootballReposity struct {
	// 	projects []*Project
}

func NewFootballReposity() *FootballReposity {
	return &FootballReposity{}
}

func (repos *FootballReposity) GetHighlights() []Highlight {
	// m.projects = append(m.projects, project)
	return nil
}
