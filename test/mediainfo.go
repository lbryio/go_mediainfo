package main

import (
	"fmt"
	mediainfo "github.com/lbryio/go_mediainfo"
	"os"
)

func main() {
	info, err := mediainfo.GetMediaInfo(os.Args[1])
	if err != nil {
		fmt.Printf("open failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%v\n", info)
}
