package spoitgo

import (
	"image"
	_ "image/jpeg"
	_ "image/png"

	"io/ioutil"
	"os"
)

const hdImageH = 1080
const hdImageW = 1920

const spotLightDir = "C:\\Users\\yasu\\AppData\\Local\\Packages\\Microsoft.Windows.ContentDeliveryManager_cw5n1h2txyewy\\LocalState\\Assets"

// ReadAllSpotlightDir Spotlightローカル保存ディレクトリを読み込みます
func ReadAllSpotlightDir() ([]string, error) {
	files, err := ioutil.ReadDir(spotLightDir)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, v := range files {
		fileNames = append(fileNames, spotLightDir+"\\"+v.Name())
	}

	return fileNames, nil
}

// ReadOnlyHdImagePaths HD画質の画像ファイルパスのみをローカルSpotligthディレクトリから取得します
func ReadOnlyHdImagePaths() ([]string, error) {
	paths, err := ReadAllSpotlightDir()
	if err != nil {
		return nil, err
	}

	hdPaths, err := filterOnlyHdImage(paths)
	if err != nil {
		return nil, err
	}

	return hdPaths, nil
}

// filterOnlyHdImage HD画質の画像ファイルパスのみを抽出して返却します
func filterOnlyHdImage(filePaths []string) ([]string, error) {
	var hdImagePaths []string

	for _, v := range filePaths {
		f, err := os.Open(v)
		if err != nil {
			return nil, err
		}

		defer f.Close()
		config, _, err := image.DecodeConfig(f)
		if err != nil {
			return nil, err
		}

		if config.Width >= hdImageW && config.Height >= hdImageH {
			hdImagePaths = append(hdImagePaths, v)
		}
	}

	return hdImagePaths, nil
}
