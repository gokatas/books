// Books sorts and prints a collection of books. Based on
// https://github.com/adonovan/gopl.io/blob/master/ch7/sorting.
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
)

type book struct {
	title string
	authors
	year int
}

type authors []string

func (as authors) String() string {
	return strings.Join(as, ", ")
}

func printBooks(books []book) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	format := "%v\t%v\t%v\n"
	fmt.Fprintf(w, format, "Title", "Authors", "Year")
	fmt.Fprintf(w, format, "-----", "-------", "----")
	for _, book := range books {
		fmt.Fprintf(w, format, book.title, book.authors, book.year)
	}
	w.Flush()
}

type byYear []book

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].year < x[j].year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	var books = []book{ // []*book might be faster with many elements
		{"The Lord of the Rings", authors{"Tolkien"}, 1954},
		{"The Phoenix Project", authors{"Kim", "Behr", "Spafford"}, 2013},
		{"The Go Programming Language", authors{"Kernighan", "Donovan"}, 2015},
	}
	sort.Sort(byYear(books))
	printBooks(books)
}
