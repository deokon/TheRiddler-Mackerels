package main

import (
	"fmt"
	"strings"
)

var states = [50]string{
	"alabama",
	"alaska",
	"arizona",
	"arkansas",
	"california",
	"colorado",
	"connecticut",
	"delaware",
	"florida",
	"georgia",
	"hawaii",
	"idaho",
	"illinois",
	"indiana",
	"iowa",
	"kansas",
	"kentucky",
	"louisiana",
	"maine",
	"maryland",
	"massachusetts",
	"michigan",
	"minnesota",
	"mississippi",
	"missouri",
	"montana",
	"nebraska",
	"nevada",
	"newhampshire",
	"newjersey",
	"newmexico",
	"newyork",
	"northcarolina",
	"northdakota",
	"ohio",
	"oklahoma",
	"oregon",
	"pennsylvania",
	"rhodeisland",
	"southcarolina",
	"southdakota",
	"tennessee",
	"texas",
	"utah",
	"vermont",
	"virginia",
	"washington",
	"westvirginia",
	"wisconsin",
	"wyoming",
}

func statesNotInCommon(w string) []string {
	uncommon := []string{}

	for _, s := range states {
		if !strings.ContainsAny(s, w) {
			uncommon = append(uncommon, s)
			// fmt.Printf("%s != %s\n", w, s)
		}
	}

	return uncommon
}

func main() {
	stateCounts := map[string]int{}
	for _, s := range states {
		stateCounts[s] = 0
	}

	dict := loadDictionary("word.list")

	longest := ""
	for w := range dict.words {
		uncommon := statesNotInCommon(w)
		dict.words[w] = len(uncommon)
		if len(uncommon) == 1 {
			stateCounts[uncommon[0]]++
			// fmt.Printf("%s => %s\n", w, uncommon[0])

			if len(w) >= len(longest) {
				fmt.Printf("%s => %s\n", w, uncommon[0])
				longest = w
			}
		}
	}

	for s, c := range stateCounts {
		fmt.Printf("%s => %d\n", s, c)
	}
}
