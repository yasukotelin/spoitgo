package spoitgo

import (
	"image"
	// JPEGを読み込めるようにインポート
	_ "image/jpeg"
	// PNGを読み込めるようにインポート
	_ "image/png"
	"io"
	"os/user"
	"path/filepath"

	"io/ioutil"
	"os"
)

const hdImageH = 1080
const hdImageW = 1920

// ImagePath 画像のフルパスとファイル名を保持する構造体
type ImagePath struct {
	FullPath string
	Name     string
}

// ReadOnlyHdImagePaths HD画質の画像ファイルパスのみをローカルSpotligthディレクトリから取得します
func ReadOnlyHdImagePaths() ([]ImagePath, error) {
	paths, err := readAllSpotlightDir()
	if err != nil {
		return nil, err
	}

	hdPaths, err := filterOnlyHdImage(paths)
	if err != nil {
		return nil, err
	}

	return hdPaths, nil
}

// CloneSpotlightImage ローカルに保存されているSpotlight画像を指定ディレクトリにCloneします。
// 新規Clone件数が返却され、Cloneが成功した場合はerrorにはnilが格納されます
// IOエラー発生時にはerrorに値が格納され返却されます
func CloneSpotlightImage(savedir string) (int, error) {
	// Spotlight画像を全取得
	paths, err := ReadOnlyHdImagePaths()
	if err != nil {
		return 0, err
	}

	newCount := 0

	// Clone処理
	for _, v := range paths {
		// 保存先ファイルパス
		srcFilePath := filepath.Join(savedir, v.Name+".jpg")

		// 保存先ファイルパスで一旦Open
		// Openできた場合は既にファイルがあるのでClone処理を行わない
		_, err := os.Open(srcFilePath)
		if err == nil {
			continue
		}

		// 保存先ファイルをCreate
		dstFile, err := os.Create(srcFilePath)
		if err != nil {
			return 0, err
		}
		// 保存元ファイルをOpen
		srcFile, err := os.Open(v.FullPath)
		if err != nil {
			return 0, err
		}

		// コピー処理
		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return 0, nil
		}

		newCount++
	}

	return newCount, nil
}

// readAllSpotlightDir Spotlightローカル保存ディレクトリを読み込みます
func readAllSpotlightDir() ([]ImagePath, error) {

	u, err := user.Current()
	if err != nil {
		return nil, err
	}

	spotLightDir := filepath.Join(u.HomeDir, "AppData", "Local",
		"Packages", "Microsoft.Windows.ContentDeliveryManager_cw5n1h2txyewy", "LocalState", "Assets")

	files, err := ioutil.ReadDir(spotLightDir)
	if err != nil {
		return nil, err
	}

	var imagePaths []ImagePath
	for _, v := range files {
		imagePath := ImagePath{
			FullPath: filepath.Join(spotLightDir, v.Name()),
			Name:     v.Name(),
		}
		imagePaths = append(imagePaths, imagePath)
	}

	return imagePaths, nil
}

// filterOnlyHdImage HD画質の画像ファイルパスのみを抽出して返却します
func filterOnlyHdImage(filePaths []ImagePath) ([]ImagePath, error) {
	var hdImagePaths []ImagePath

	for _, v := range filePaths {
		f, err := os.Open(v.FullPath)
		if err != nil {
			return nil, err
		}
		defer f.Close()

		config, _, err := image.DecodeConfig(f)
		if err != nil {
			// 画像ファイルでない場合はcontinueする
			continue
		}

		if config.Width >= hdImageW && config.Height >= hdImageH {
			hdImagePaths = append(hdImagePaths, v)
		}

		f.Close()
	}

	return hdImagePaths, nil
}
