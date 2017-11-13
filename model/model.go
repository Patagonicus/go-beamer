package model

import (
	"sync"

	"github.com/pkg/errors"
)

type Page interface {
	ID() uint
}

type page interface {
	GetPage() Page
}

type Model struct {
	pages    map[uint]page
	next     uint
	playlist *playlist
	lock     *sync.RWMutex
}

func New() *Model {
	m := &Model{
		make(map[uint]page),
		0,
		nil,
		&sync.RWMutex{},
	}
	plalistID := m.NewPlaylist("default")
	m.playlist = m.pages[plalistID].(*playlist)
	return m
}

func (m *Model) NewInfopage(description string) uint {
	m.lock.Lock()
	defer m.lock.Unlock()

	id := m.next
	m.next++

	m.pages[id] = &infopage{id, description}

	return id
}

func (m *Model) NewPlaylist(name string) uint {
	m.lock.Lock()
	defer m.lock.Unlock()

	id := m.next
	m.next++

	m.pages[id] = &playlist{id, name, nil, 0}

	return id
}

func (m *Model) GetAllPages() []Page {
	m.lock.RLock()
	defer m.lock.RUnlock()

	pages := make([]Page, 0, len(m.pages))
	for _, p := range m.pages {
		pages = append(pages, p.GetPage())
	}

	return pages
}

func (m *Model) GetPage(id uint) (Page, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	page, ok := m.pages[id]
	if !ok {
		return nil, false
	}
	return page.GetPage(), true
}

func (m *Model) AddToPlaylist(pid, id uint) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	playlistPage, ok := m.pages[pid]
	if !ok {
		return errors.Errorf("playlist %d not found", pid)
	}
	playlist, ok := playlistPage.(*playlist)
	if !ok {
		return errors.Errorf("%d is not a playlist", pid)
	}

	other, ok := m.pages[id]
	if !ok {
		return errors.Errorf("page %d not found", id)
	}

	return playlist.Add(other)
}

func (m *Model) GetCurrentPage() Page {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if m.playlist == nil {
		return nil
	}

	pg := m.playlist.currentPage()
	if pg == nil {
		return nil
	}

	return pg.GetPage()
}

func (m *Model) Advance() {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.playlist == nil {
		return
	}

	m.playlist.advance()
}
