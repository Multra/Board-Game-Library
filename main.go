package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Game struct {
	Name string `xml:"name"`
	Stat Stats  `xml:"stats"`
	//	Rate Rating `xml:,innerxml"`
}

type Obj struct {
	Bg []Game `xml:"item"`
}

type Stats struct {
	Ptime      int    `xml:"maxplaytime,attr"`
	Maxplayers int    `xml:"maxplayers,attr"`
	Minplayers int    `xml:"minplayers,attr"`
	Rate       Rating `xml:"rating"`
}

type Rating struct {
	Bayrating float32 `xml:",any"`
}

func main() {
	var acctName string
	fmt.Println("User:")
	fmt.Scan(&acctName)
	retrieve(acctName)

}

func retrieve(acctName string) {
	url := "https://www.boardgamegeek.com/xmlapi/collection/" + acctName + "?stats=1"
	res, _ := http.Get(url)
	dat, _ := ioutil.ReadAll(res.Body)
	g := Obj{}
	xml.Unmarshal(dat, &g)
	for _, bg := range g.Bg {
		fmt.Println("Name:", bg.Name)
		fmt.Println("Play time:", bg.Stat.Ptime)
		fmt.Println("Max Players:", bg.Stat.Maxplayers)
		fmt.Println("Average Rating:", bg.Stat.Rate.Bayrating)
		fmt.Println()
	}
}
