//go:build generate
// +build generate

package main

//import "github.com/zserge/lorca"

import (
	crx3 "github.com/mediabuyerbot/go-crx3"
	"log"
	"os"
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
		// open extension-hash.go and write hash to it
		hashfile, err := os.OpenFile("extension-hash.go", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			return err
		}
		defer hashfile.Close()
		if _, err := hashfile.WriteString("\t\"" + hash + "\",\n"); err != nil {
			return err
		}
		zipsfile, err := os.OpenFile("extension-zips.go", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			return err
		}
		defer zipsfile.Close()
		if _, err := zipsfile.WriteString("\t\"" + outdir + "\",\n"); err != nil {
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
	extzips := "package main\n"
	extzips += "\n"
	extzips += "var EXTENSIONS = []string{" + "\n"
	exthash := "package main\n"
	exthash += "\n"
	exthash += "var EXTENSIONHASHES = []string{" + "\n"
	ioutil.WriteFile("extension-zips.go", []byte(extzips), 0644)
	ioutil.WriteFile("extension-hash.go", []byte(exthash), 0644)

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
	// open extension-hash.go and write hash to it
	hashfile, err := os.OpenFile("extension-hash.go", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer hashfile.Close()
	if _, err := hashfile.WriteString("}\n"); err != nil {
		log.Fatal(err)
	}
	zipsfile, err := os.OpenFile("extension-zips.go", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer zipsfile.Close()
	if _, err := zipsfile.WriteString("}\n"); err != nil {
		log.Fatal(err)
	}
	//lorca.Embed("main", "assets.go", "i2pchrome.js")
	//lorca.Embed("i2pchrome", "lib/assets.go", "i2pchrome.js")
}
