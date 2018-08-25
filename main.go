package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/ysbrothersk/spoitgo/logo"
	"github.com/ysbrothersk/spoitgo/spoitgo"
)

var u, uerror = user.Current()
var saveDir = filepath.Join(u.HomeDir, "Pictures", "SpoitGo")

func init() {
	if uerror != nil {
		panic(uerror)
	}
}

func main() {
	logo.Print()

	isCreate, err := createDirIfNothing(saveDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	if isCreate {
		fmt.Println("create new directory:", saveDir)
	}

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

func isExist(dir string) bool {
	_, err := os.Stat(dir)
	return err == nil
}

// getTotalCount ディレクトリ内ファイルの総数を返却します
func getTotalCount(dir string) (int, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return 0, err
	}

	return len(files), nil
}

// createDirIfNothing dirパスにディレクトリがない場合ディレクトリを生成します
// ディレクトリを生成した場合はtrue、しなかった場合はfalseを返却します。
// エラーが発生した場合はerrを返却します
func createDirIfNothing(dir string) (bool, error) {
	if isExist(dir) {
		return false, nil
	}

	// ディレクトリの生成処理
	// TODO: この権限のbyte値をよく調べる
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return false, err
	}
	return true, nil
}
