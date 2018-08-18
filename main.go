package main

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"

	"github.com/ysbrothersk/spoitgo/logo"
	"github.com/ysbrothersk/spoitgo/spoitgo"
)

func main() {
	logo.Print()

	u, err := user.Current()
	if err != nil {
		fmt.Println(err)
		return
	}
	saveDir := filepath.Join(u.HomeDir, "Pictures", "SpoitGo")

	newCount, err := spoitgo.CloneSpotlightImage(saveDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("New File:", newCount)

	total, err := getTotalCount(saveDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total:", total)
}

// getTotalCount ディレクトリ内ファイルの総数を返却します
func getTotalCount(dir string) (int, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return 0, err
	}

	return len(files), nil
}
