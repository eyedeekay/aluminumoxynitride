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

var workdir, err = os.Getwd()

func main() {
	directory := flag.String("directory", filepath.Join(workdir, "spinel"), "The directory to store aluminumoxynitride data in")
	flag.Parse()
	if err != nil {
		log.Fatal(err)
	}
	rundir, err := filepath.Abs(*directory)
	if err != nil {
		log.Fatal(err)
	}
	profile, err := filepath.Abs(filepath.Join(rundir, "profile"))
	if err != nil {
		log.Fatal(err)
	}
	os.MkdirAll(profile, 0755)
	ARGS = append(ARGS, "--user-data-dir="+profile)
	ARGS = append(ARGS, flag.Args()...)

	if err != nil {
		log.Fatal(err)
	}
	if err := os.Chdir(rundir); err != nil {
		log.Fatal(err)
	}

	if I2PDaemon, err := StartI2P(rundir); err != nil {
		log.Fatal(err)
	} else {
		if I2PDaemon != nil {
			defer I2PDaemon.Stop()
		}
	}
	if err = WriteOutExtensions(profile); err != nil {
		log.Fatal(err)
	}
	CHROMIUM, ERROR = SecureExtendedChromium("spinel", false, extensionPaths(profile), EXTENSIONHASHES, ARGS...)
	//CHROMIUM, ERROR = ExtendedChromium("spinel", false, extensionPaths("extensions"), ARGS...)
	if ERROR != nil {
		log.Fatal(ERROR)
	}
	defer CHROMIUM.Close()
	<-CHROMIUM.Done()
}
