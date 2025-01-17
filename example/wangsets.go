package main

import (
	"fmt"
	"os"

	"github.com/lafriks/go-tiled"
)

const (
	tmxPath = "../go-tiled/assets/test_wangsets_map.tmx"
)

func main() {
	var tiledMap *tiled.Map
	var err error
	tiledMap, err = tiled.LoadFromFile(tmxPath)
	if err != nil {
		fmt.Printf("error parsing gameMap: %s", err.Error())
		os.Exit(2)
	}

	fmt.Print(tiledMap.Tilesets[0].WangSets)
}
