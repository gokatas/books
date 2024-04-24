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

type Book struct {
	Title string
	Authors
	Year int
}

type Authors []string

func (as Authors) String() string {
	return strings.Join(as, ", ")
}

func printBooks(books []Book) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	format := "%v\t%v\t%v\n"
	fmt.Fprintf(tw, format, "Title", "Authors", "Year")
	fmt.Fprintf(tw, format, "-----", "-------", "----")
	for _, book := range books {
		fmt.Fprintf(tw, format, book.Title, book.Authors, book.Year)
	}
	tw.Flush()
}

type byYear []Book

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	var books = []Book{ // []*book might be faster with many elements
		{"The Go Programming Language", Authors{"Kernighan", "Donovan"}, 2015},
		{"The Phoenix Project", Authors{"Kim", "Behr", "Spafford"}, 2013},
		{"The Lord of the Rings", Authors{"Tolkien"}, 1954},
	}
	sort.Sort(byYear(books))
	printBooks(books)
}
