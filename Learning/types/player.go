package movie

import "strings"

type Movie struct {
	Name     string
	Director string
	Year     string
}

func (m *Movie) CreateUpperInfo() *Movie {

	m.Year = strings.ToUpper(m.Year)
	m.Name = strings.ToUpper(m.Name)
	m.Director = strings.ToUpper(m.Director)
	return m
}
