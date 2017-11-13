package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/Patagonicus/go-beamer/assets"
	"github.com/Patagonicus/go-beamer/model"
	"github.com/pkg/errors"
)

type pageDescription struct {
	ID   uint
	Name string
}

func main() {
	m := model.New()

	http.HandleFunc("/admin", getAdminHandler(m))
	http.HandleFunc("/new/info", getNewInfoHandler(m))
	http.HandleFunc("/new/info/create", getCreateInfoHandler(m))
	http.HandleFunc("/new/playlist", getNewPlaylistHandler(m))
	http.HandleFunc("/new/playlist/create", getCreatePlaylistHandler(m))
	http.HandleFunc("/edit/playlist/add/", getPlaylistAddHandler(m))
	http.HandleFunc("/edit/playlist/", getEditPlaylistHandler(m))
	http.HandleFunc("/view", getViewHandler(m))
	http.HandleFunc("/", getRootHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getEditHandler(m *model.Model) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "edit")
	}
}

func getAdminHandler(m *model.Model) func(http.ResponseWriter, *http.Request) {
	tmpl := mustTemplate("admin.tpl")
	return func(w http.ResponseWriter, r *http.Request) {
		type page struct {
			ID   uint
			Name string
			Edit string
		}

		var data struct {
			Pages []page
		}
		allPages := m.GetAllPages()
		for i, p := range getPageDescriptions(allPages) {
			var edit string
			switch allPages[i].(type) {
			case model.Infopage:
				edit = fmt.Sprintf("/edit/info/%d", p.ID)
			case model.Playlist:
				edit = fmt.Sprintf("/edit/playlist/%d", p.ID)
			default:
				log.Fatalf("unknown page type: %v", allPages[i])
			}
			data.Pages = append(data.Pages, page{p.ID, p.Name, edit})
		}

		err := tmpl.Execute(w, data)
		if err != nil {
			log.Printf("error rendering admin template: %s", err)
		}
	}
}

func getNewInfoHandler(m *model.Model) func(http.ResponseWriter, *http.Request) {
	tmpl := mustTemplate("new_info.tpl")
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			log.Printf("error rendering new info template: %s", err)
		}
	}
}

func getCreateInfoHandler(m *model.Model) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		desc := r.PostFormValue("description")
		id := m.NewInfopage(desc)
		log.Printf("Add infopage %d: %s\n", id, desc)

		http.Redirect(w, r, "/admin", http.StatusFound)
	}
}

func getNewPlaylistHandler(m *model.Model) func(http.ResponseWriter, *http.Request) {
	tmpl := mustTemplate("new_playlist.tpl")
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			log.Printf("error rendering new playlist template: %s", err)
		}
	}
}

func getCreatePlaylistHandler(m *model.Model) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		name := r.PostFormValue("name")
		id := m.NewPlaylist(name)
		log.Printf("Add playlist %d: %s\n", id, name)

		http.Redirect(w, r, fmt.Sprintf("/edit/playlist/%d", id), http.StatusFound)
	}
}

func getEditPlaylistHandler(m *model.Model) func(http.ResponseWriter, *http.Request) {
	tmpl := mustTemplate("edit_playlist.tpl")
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getIDFromPath("/edit/playlist/", r.URL.Path)
		if err != nil {
			log.Printf("invalid playlist id: %s", err)
			fmt.Fprintf(w, "Invalid playlist id: %s", id, err)
			return
		}

		playlist, err := getPlaylist(m, id)
		if err != nil {
			log.Printf("failed to get playlist %d: %s", id, err)
			fmt.Fprintf(w, "Error loading playlist %d: %s", id, err)
			return
		}

		pages := getPageDescriptions(m.GetAllPages())

		data := struct {
			Playlist model.Playlist
			Entries  []pageDescription
			Pages    []pageDescription
		}{
			playlist,
			getPageDescriptions(playlist.Entries),
			pages,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Printf("error rendering edit playlist template: %s", err)
		}
	}
}

func getPlaylistAddHandler(m *model.Model) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		re := regexp.MustCompile("^/edit/playlist/add/([0-9]*)/([0-9]*)$")
		parts := re.FindStringSubmatch(r.URL.Path)
		if parts == nil {
			log.Printf("invalid request: %s", r.URL.Path)
			fmt.Fprintf(w, "invalid request: %s", r.URL.Path)
			return
		}

		playlistID, err := getIDFromString(parts[1])
		if err != nil {
			log.Printf("invalid playlist id: %s", parts[1])
			fmt.Fprintf(w, "invalid playlist id: %s", parts[1])
			return
		}

		otherID, err := getIDFromString(parts[2])
		if err != nil {
			log.Printf("invalid page id: %s", parts[2])
			fmt.Fprintf(w, "invalid page id: %d", parts[2])
			return
		}

		err = m.AddToPlaylist(playlistID, otherID)
		if err != nil {
			log.Printf("could not add %d to playlist %d: %s", otherID, playlistID, err)
			fmt.Fprintf(w, "could not add %d to playlist %d: %s", otherID, playlistID, err)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/edit/playlist/%d", playlistID), http.StatusFound)
	}
}

func getViewHandler(m *model.Model) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		current := m.GetCurrentPage()
		fmt.Fprintf(w, "%s", current)
		m.Advance()
	}
}

func getRootHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/view", http.StatusFound)
	}
}

func mustTemplate(name string) *template.Template {
	return template.Must(template.New(name).Parse(string(assets.MustAsset(name))))
}

func getIDFromPath(prefix string, path string) (uint, error) {
	idString := path[len(prefix):]
	return getIDFromString(idString)
}

func getIDFromString(s string) (uint, error) {
	id, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func getPlaylist(m *model.Model, id uint) (model.Playlist, error) {
	page, ok := m.GetPage(id)
	if !ok {
		return model.Playlist{}, errors.Errorf("playlist %d does not exist", id)
	}

	playlist, ok := page.(model.Playlist)
	if !ok {
		return model.Playlist{}, errors.Errorf("page %d is not a playlist", id)
	}

	return playlist, nil
}

func getPageDescriptions(pages []model.Page) []pageDescription {
	result := make([]pageDescription, len(pages))
	for i, p := range pages {
		result[i] = getPageDescription(p)
	}
	return result
}

func getPageDescription(page model.Page) pageDescription {
	var result pageDescription
	result.ID = page.ID()
	switch p := page.(type) {
	case model.Infopage:
		result.Name = p.Description
	case model.Playlist:
		result.Name = p.Name
	default:
		log.Printf("unknown page type: %v", p)
		result.Name = "unknown page type"
	}
	return result
}
