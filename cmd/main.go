package main

import (
	"fmt"

	"github.com/ysbrothersk/spoitgo"
	"github.com/ysbrothersk/spoitgo/logo"
)

const spotLightDir = "C:\\Users\\yasu\\AppData\\Local\\Packages\\Microsoft.Windows.ContentDeliveryManager_cw5n1h2txyewy\\LocalState\\Assets"

func main() {
	logo.Print()

	fileNames, err := spoitgo.ReadOnlyHdImagePaths()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range fileNames {
		fmt.Println(v)
	}
}
