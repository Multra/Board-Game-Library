package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Game struct {
	Name  string `xml:"name"`
	Stat  Stats  `xml:"stats"`
	Ident int    `xml:"objectid,attr"`
}

type Obj struct {
	Bg []Game `xml:"item"`
}

type Stats struct {
	Ptime      int    `xml:"maxplaytime,attr"`
	Maxplayers int    `xml:"maxplayers,attr"`
	Minplayers int    `xml:"minplayers,attr"`
	Rate       Rating `xml:"rating>bayesaverage"`
}

type Rating struct {
	Bayrating float32 `xml:"value,attr"`
}

func main() {
	var acctName string
	fmt.Println("User:")
	fmt.Scan(&acctName)
	a := retrieve(acctName)
	for _, bg := range a.Bg {
		fmt.Println("Name:", bg.Name)
		fmt.Println("ID:", bg.Ident)
		fmt.Println("Play time:", bg.Stat.Ptime)
		fmt.Println("Max Players:", bg.Stat.Maxplayers)
		fmt.Printf("Average Rating: %.2f\n", bg.Stat.Rate.Bayrating)
		fmt.Println()
	}

}

func retrieve(acctName string) Obj {
	res, _ := http.Get("https://www.boardgamegeek.com/xmlapi/collection/" + acctName + "?stats=1")
	if res.StatusCode != 200 {
		log.Fatal("Error code " + string(res.StatusCode) + " returned, try again later")
	}
	dat, _ := ioutil.ReadAll(res.Body)
	g := Obj{}
	xml.Unmarshal(dat, &g)
	return g
}
