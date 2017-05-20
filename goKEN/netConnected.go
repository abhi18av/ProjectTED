package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/fatih/color"
)

func main() {

	/*
				"github.com/gizak/termui"
				"github.com/spf13/cobra"
				"github.com/fatih/color"
				"github.com/gosuri/uiprogress"
				"github.com/briandowns/spinner"
		https://github.com/HouzuoGuo/tiedot
		https://github.com/boltdb/bolt
		https://github.com/cznic/ql
		https://github.com/patrickmn/go-cache
		https://github.com/syndtr/goleveldb
		https://github.com/cayleygraph/cayley
		https://github.com/dgraph-io/dgraph
		https://github.com/mattn/go-sqlite3
		https://github.com/gdamore/tcell
		https://github.com/kyokomi/emoji
		go get github.com/gonum/plot/...
		go get -u github.com/zieckey/gochart

	*/
	// Make a get request
	rs, err := http.Get("https://google.com")
	// Process response
	if err != nil {
		color.Red("Not connected")
		//panic("Not connected to the net") // More idiomatic way would be to print the error and die unless it's a serious error

		// Learn about exit status in Golang
		os.Exit(1)
	}

	defer rs.Body.Close()

	bodyBytes, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		panic(err)

	}

	bodyString := string(bodyBytes)

	println(bodyString)
}
