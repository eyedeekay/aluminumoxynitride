//go:build generate
// +build generate

package main

//import "github.com/zserge/lorca"

import (
	crx3 "github.com/mediabuyerbot/go-crx3"
	"log"
	//	"os"
	//hashdir "github.com/sger/go-hashdir"
	"github.com/eyedeekay/go-ccw"
	"io/ioutil"
	"strings"
)

func DownloadExtension(extensionID, filepath string) (string, error) {
	if err := crx3.DownloadFromWebStore(extensionID, filepath); err != nil {
		return "", err
	}
	return filepath, nil
}

func UnpackAndMoveExtension(filepath, outdir string) error {
	if err := crx3.Unpack(filepath); err != nil {
		return err
	}
	return nil
}

func DownloadAndUnpackAndMoveExtension(extensionID, filepath, outdir string) error {
	path, err := DownloadExtension(extensionID, filepath)
	if err != nil {
		return err
	}
	if err := UnpackAndMoveExtension(path, outdir); err != nil {
		return err
	}
	unzpath := strings.TrimRight(filepath, ".crx")
	if hash, err := ccw.ZipAndHashDir(outdir); err != nil {
		return err
	} else {
		log.Println("Extension hash for "+unzpath+":", hash)
		err := ioutil.WriteFile(unzpath+".hash", []byte(hash), 0644)
		if err != nil {
			return err
		}
	}

	if outdir == unzpath || outdir == "" {
		return nil
	}
	// check if the directory exists
	/*if _, err := os.Stat(outdir); err == nil {
		// if it does, delete it
		os.RemoveAll(outdir)
	}
	// if it doesn't, copy the contents of the zip to the directory
	if err := os.Rename(path, outdir); err != nil {
		return err
	}*/
	return nil
}

func main() {
	// You can also run "npm build" or webpack here, or compress assets, or
	// generate manifests, or do other preparations for your assets.
	if err := DownloadAndUnpackAndMoveExtension("njdfdhgcmkocbgbhcioffdbicglldapd", "localcdn.crx", "localcdn"); err != nil {
		log.Fatal(err)
	}
	if err := DownloadAndUnpackAndMoveExtension("fockhhgebmfjljjmjhbdgibcmofjbpca", "onionbrowser.crx", "onionbrowser"); err != nil {
		log.Fatal(err)
	}
	if err := DownloadAndUnpackAndMoveExtension("oiigbmnaadbkfbmpbfijlflahbdbdgdf", "scriptsafe.crx", "scriptsafe"); err != nil {
		log.Fatal(err)
	}
	if err := DownloadAndUnpackAndMoveExtension("cjpalhdlnbpafiamejdnhcphjbkeiagm", "ublockorigin.crx", "ublockorigin"); err != nil {
		log.Fatal(err)
	}
	if err := DownloadAndUnpackAndMoveExtension("ikdjcmomgldfciocnpekfndklkfgglpe", "i2pchrome.js.crx", "i2pchrome.js"); err != nil {
		log.Fatal(err)
	}

	//lorca.Embed("main", "assets.go", "i2pchrome.js")
	//lorca.Embed("i2pchrome", "lib/assets.go", "i2pchrome.js")
}
