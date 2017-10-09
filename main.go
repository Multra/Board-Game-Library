package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
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
	//	fmt.Println("User:")
	//	fmt.Scan(&acctName)
	http.HandleFunc("/process", process)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe("127.0.0.1:8080", nil)
	//	a := retrieve(acctName)

	//	gNum := selector(len(a.Bg))
	//	for _, n := range gNum {
	//		fmt.Println(a.Bg[n])

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

func selector(n int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	m := 0
	if n < 5 {
		m = n
	} else {
		m = 5
	}
	s := make([]int, m)
	var rn int
	for c := 0; c < 5; c++ {
		rn = rand.Intn(n)
		fmt.Println(rn)
		for p, a := range s {
			if a == rn && c > 0 {
				c = p
				break
			}
		}
		s[c] = rn
		fmt.Println(s)
	}
	return s
}

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//	fmt.Println(r.PostForm["name"])
	bggname := strings.Join(r.PostForm["name"], "")
	a := retrieve(bggname)
	gNum := selector(len(a.Bg))

	for _, n := range gNum {
		//		fmt.Println(a.Bg[n])
		io.WriteString(w, a.Bg[n].Name+"\n")
	}
	//	io.WriteString(w, string(r.PostForm["name"]))
}

func toString(g Game) string {
	return fmt.Sprintf("%v\n", g)
}
