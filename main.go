//go:generate go run -tags generate gen.go

package main

import (
	"embed"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	//"fmt"
	"path/filepath"

	. "github.com/eyedeekay/go-ccw"
)

//go:embed i2pchrome.js/*
//go:embed localcdn/*
//go:embed onionbrowse/*
//go:embed scriptsafe/*
//go:embed ublockorigin/*
var extensionContent embed.FS

var EXTENSIONS = []string{
	"localcdn",
	"onionbrowse",
	"scriptsafe",
	"ublockorigin",
	"i2pchrome.js",
}
var EXTENSIONHASHES = []string{
	"b9f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8",
	"b9f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8",
	"b9f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8",
	"b9f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8f8",
	"60ca1a760703fb3ba6aceaa2fe146cb4e915b0c22f04469a6d86404f4dbafdb0",
}

func extensionPaths(outpath string) []string {
	var paths []string
	for _, extension := range EXTENSIONS {
		paths = append(paths, outpath+"/"+extension)
	}
	return paths
}

func WriteOutExtensions(outdir string) error {
	// Walk the contents of extensionContent and write the files out to disk
	return fs.WalkDir(extensionContent, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		outpath := filepath.Join(outdir, path)
		if d.IsDir() {
			if err := os.MkdirAll(outpath, 0755); err != nil {
				log.Println(err)
			}
		}
		bytes, err := extensionContent.ReadFile(path)
		if err != nil {
			log.Println(err)
		}
		if err := ioutil.WriteFile(outpath, bytes, 0644); err != nil {
			log.Println(err)
		}
		return nil
	})
}

var ARGS = []string{
	"--safebrowsing-disable-download-protection",
	"--disable-client-side-phishing-detection",
	"--disable-3d-apis",
	"--disable-accelerated-2d-canvas",
	"--disable-remote-fonts",
	"--disable-sync-preferences",
	"--disable-sync",
	"--disable-speech",
	"--disable-webgl",
	"--disable-reading-from-canvas",
	"--disable-gpu",
	"--disable-32-apis",
	"--disable-auto-reload",
	"--disable-background-networking",
	"--disable-d3d11",
	"--disable-file-system",
}

func main() {
	WriteOutExtensions("extensions")
	//CHROMIUM, ERROR = SecureExtendedChromium("i2pchromium-browser", false, extensionPaths("extensions"), EXTENSIONHASHES, ARGS...)
	CHROMIUM, ERROR = ExtendedChromium("i2pchromium-browser", false, extensionPaths("extensions"), ARGS...)
	if ERROR != nil {
		log.Fatal(ERROR)
	}
	defer CHROMIUM.Close()
	<-CHROMIUM.Done()
}
