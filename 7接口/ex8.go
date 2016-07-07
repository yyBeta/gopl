package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"Pokemon Go", "NianticLabs", "Games", 2012, length("2m33s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	fmt.Println("[*] original:")
	printTracks(tracks)
	var click int
	var isAsc [5]bool
	// Ascending or descending order of Title, Artist, Year and Length, isAsc[0] is useless
	var sortFirst, sortSecond int
	for true {
		fmt.Println("\n | input '1' to click on 'Title', '2' to click on 'Artist',")
		fmt.Printf(" | '3' to click on 'Year', '4' to click on 'Length', others to exit: ")
		fmt.Scanln(&click)
		if click > 0 && click < 5 {
			isAsc[click] = !isAsc[click] // change the Ascending or descending order
		} else {
			os.Exit(0)
		}

		if sortFirst != click { // change the sorting order
			if sortFirst == 0 {
				sortFirst = click
			} else {
				sortSecond = sortFirst
				sortFirst = click
			}
		}

		sort.Sort(customSort{tracks, func(x, y *Track) bool { // sort interface
			switch sortFirst {
			case 1:
				if x.Title != y.Title {
					return (x.Title < y.Title) == isAsc[1] // return negtive bool value if Descending
				}
			case 2:
				if x.Artist != y.Artist {
					return (x.Artist < y.Artist) == isAsc[2]
				}
			case 3:
				if x.Year != y.Year {
					return (x.Year < y.Year) == isAsc[3]
				}
			case 4:
				if x.Length != y.Length {
					return (x.Length < y.Length) == isAsc[4]
				}
			}

			switch sortSecond {
			case 1:
				if x.Title != y.Title {
					return (x.Title < y.Title) == isAsc[1]
				}
			case 2:
				if x.Artist != y.Artist {
					return (x.Artist < y.Artist) == isAsc[2]
				}
			case 3:
				if x.Year != y.Year {
					return (x.Year < y.Year) == isAsc[3]
				}
			case 4:
				if x.Length != y.Length {
					return (x.Length < y.Length) == isAsc[4]
				}
			}
			return false
		}})

		var order1, order2 []byte // show ascending or descending order
		order1, order2 = []byte("↓"), []byte("↓")
		if isAsc[sortFirst] {
			order1 = []byte("↑")
		}
		if isAsc[sortSecond] {
			order2 = []byte("↑")
		}
		fields := [5]string{"None", "Title", "Artist", "Year", "Length"}
		fmt.Println("\n[*] sort by :", fields[sortFirst], string(order1), " then:", fields[sortSecond], string(order2))
		printTracks(tracks)
	}
}
