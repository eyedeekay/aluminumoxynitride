//go:generate go run -tags generate gen.go

package main

import (
	"embed"
	"flag"
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

func extensionPaths(outpath string) []string {
	var paths []string
	for _, extension := range EXTENSIONS {
		paths = append(paths, outpath+"/"+extension)
	}
	return paths
}

func WriteOutExtensions(outdir string) error {
	// Walk the contents of extensionContent and write the files out to disk
	os.MkdirAll(outdir, 0755)
	return fs.WalkDir(extensionContent, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		outpath := filepath.Join(outdir, path)
		if d.IsDir() {
			if err := os.MkdirAll(outpath, 0755); err != nil {
				//log.Println(err)
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
	directory := flag.String("directory", "", "The directory to store aluminumoxynitride data in")
	flag.Parse()
	profile := filepath.Join(*directory, "profile")
	os.MkdirAll(profile, 0755)
	ARGS = append(ARGS, "--user-data-dir="+*directory+"/profile")
	ARGS = append(ARGS, flag.Args()...)
	var workdir string
	var err error
	if directory != nil && *directory == "" {
		workdir, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		workdir = *directory
		if err := os.Chdir(workdir); err != nil {
			log.Fatal(err)
		}
	}
	if err = StartI2P(workdir); err != nil {
		log.Fatal(err)
	}
	if err = WriteOutExtensions("i2pchromium-browser"); err != nil {
		log.Fatal(err)
	}
	CHROMIUM, ERROR = SecureExtendedChromium("i2pchromium-browser", false, extensionPaths("i2pchromium-browser"), EXTENSIONHASHES, ARGS...)
	//CHROMIUM, ERROR = ExtendedChromium("i2pchromium-browser", false, extensionPaths("extensions"), ARGS...)
	if ERROR != nil {
		log.Fatal(ERROR)
	}
	defer CHROMIUM.Close()
	<-CHROMIUM.Done()
}
