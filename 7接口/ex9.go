// no need for me to code it, so I copied it from seikichi/gopl

package main

import (
	"log"
	"net/http"
	"net/url"
	"sort"
	"text/template"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func newTracks() []*Track {
	return []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var trackTable = template.Must(template.New("tracktable").Parse(`
<h1>Tracks</h1>
<table>
<tr style='text-align: left'>
  <th><a href='{{.NewURL "Title"}}'>Title</a></th>
  <th><a href='{{.NewURL "Artist"}}'>Artist</a></th>
  <th><a href='{{.NewURL "Album"}}'>Album</a></th>
  <th><a href='{{.NewURL "Year"}}'>Year</a></th>
  <th><a href='{{.NewURL "Length"}}'>Length</a></th>
</tr>
{{range .Tracks}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

type data struct {
	Tracks []*Track
	r      *http.Request
}

func (d *data) NewURL(sortKey string) *url.URL {
	u := *d.r.URL
	q := u.Query()
	q.Add("sort[]", sortKey)
	u.RawQuery = q.Encode()
	return &u
}

type multiSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (s multiSort) Len() int {
	return len(s.t)
}

func (s multiSort) Less(i, j int) bool {
	return s.less(s.t[i], s.t[j])
}

func (s multiSort) Swap(i, j int) {
	s.t[i], s.t[j] = s.t[j], s.t[i]
}

func (s multiSort) byTitle() multiSort {
	return multiSort{s.t, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		return s.less(x, y)
	}}
}

func (s multiSort) byArtist() multiSort {
	return multiSort{s.t, func(x, y *Track) bool {
		if x.Artist != y.Artist {
			return x.Artist < y.Artist
		}
		return s.less(x, y)
	}}
}

func (s multiSort) byAlbum() multiSort {
	return multiSort{s.t, func(x, y *Track) bool {
		if x.Album != y.Album {
			return x.Album < y.Album
		}
		return s.less(x, y)
	}}
}

func (s multiSort) byYear() multiSort {
	return multiSort{s.t, func(x, y *Track) bool {
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		return s.less(x, y)
	}}
}

func (s multiSort) byLength() multiSort {
	return multiSort{s.t, func(x, y *Track) bool {
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return s.less(x, y)
	}}
}

func newMultiSort(t []*Track) multiSort {
	return multiSort{t, func(_, _ *Track) bool { return false }}
}

func sortByQuery(t []*Track, q url.Values) {
	s := newMultiSort(t)
	for _, key := range q["sort[]"] {
		switch key {
		case "Title":
			s = s.byTitle()
		case "Artist":
			s = s.byArtist()
		case "Album":
			s = s.byAlbum()
		case "Year":
			s = s.byYear()
		case "Length":
			s = s.byLength()
		}
	}
	sort.Sort(s)
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		t := newTracks()
		sortByQuery(t, r.URL.Query())
		trackTable.Execute(w, &data{t, r})
	}

	http.HandleFunc("/", handler)
	log.Println("Running server in localhost:8000 ...")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
