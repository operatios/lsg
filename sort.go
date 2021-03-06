package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func sortFiles(files []File, sortType string, reverse bool) {
	switch strings.ToLower(sortType) {
	case "s", "size":
		sort.Slice(files, func(i, j int) bool {
			return files[i].size() > files[j].size()
		})

	case "t", "time":
		sort.Slice(files, func(i, j int) bool {
			return files[i].modTime() > files[j].modTime()
		})

	case "x", "extension":
		sort.Slice(files, func(i, j int) bool {
			return cmpCaseInsensitive(files[i].ext(), files[j].ext())
		})

	case "c", "category":
		sort.Slice(files, func(i, j int) bool {
			return files[i].category() < files[j].category()
		})

	case "":
		sort.Slice(files, func(i, j int) bool {
			return cmpCaseInsensitive(files[i].name(), files[j].name())
		})

	default:
		fmt.Fprintf(os.Stderr, "Invalid sorting parameter: %s\n", sortType)
		os.Exit(1)
	}

	if reverse {
		for i, j := 0, len(files)-1; i < j; i, j = i+1, j-1 {
			files[i], files[j] = files[j], files[i]
		}
	}
}

func cmpCaseInsensitive(a, b string) bool {
	return strings.ToLower(a) < strings.ToLower(b)
}
