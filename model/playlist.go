package model

import "github.com/pkg/errors"

type Playlist struct {
	id       uint
	Name     string
	Entries  []Page
	position uint
}

func (p Playlist) ID() uint {
	return p.id
}

type playlist struct {
	id       uint
	name     string
	entries  []page
	position uint
}

func (p *playlist) GetPage() Page {
	entries := make([]Page, len(p.entries))
	for i, pg := range p.entries {
		entries[i] = pg.GetPage()
	}
	return Playlist{p.id, p.name, entries, p.position}
}

func (p *playlist) Add(other page) error {
	if p == other {
		return errors.Errorf("cannot add %s to %s, would create loop", p, p)
	}

	if otherPlaylist, ok := other.(*playlist); ok {
		if otherPlaylist.contains(p) {
			return errors.Errorf("cannot add %s to %s, would create loop", otherPlaylist, p)
		}
	}

	p.entries = append(p.entries, other)

	return nil
}

func (p *playlist) contains(other page) bool {
	for _, e := range p.entries {
		if e == other {
			return true
		}
		if ePlaylist, ok := e.(*playlist); ok {
			if ePlaylist.contains(other) {
				return true
			}
		}
	}
	return false
}

func (p *playlist) currentPage() page {
	if len(p.entries) == 0 {
		return nil
	}

	pg := p.entries[p.position]
	if pgPlaylist, ok := pg.(*playlist); ok {
		return pgPlaylist.currentPage()
	}

	return pg
}

func (p *playlist) advance() {
	if len(p.entries) == 0 {
		return
	}

	p.position = (p.position + 1) % uint(len(p.entries))
	if subPlaylist, ok := p.entries[p.position].(*playlist); ok {
		subPlaylist.advance()
	}
}
