
VERSION=0.0.3
DESC="A wrapper for Chrome-based browsers that auto-configures them to use I2P."
APP=aluminumoxynitride

echo:
	echo gothub release -u eyedeekay -r $(APP) -t $(VERSION) -n $(VERSION) -d $(DESC)

delete-release:
	gothub delete -u eyedeekay -r $(APP) -t $(VERSION)

build:
	go build --tags "netgo osusergo" -o $(APP)-$(GOOS)-$(GOARCH)

version:
	gothub release --pre-release -u eyedeekay -r $(APP) -t $(VERSION) -n $(VERSION) -d $(DESC)

release: all version upload-all

docker:
	sudo rm -rf $(PWD).docker-build 
	cp -rv $(PWD) $(PWD).docker-build
	docker run -it --rm \
		--env GOOS=$(GOOS) \
		--env GOARCH=$(GOARCH) \
		-w /go/src/github.com/eyedeekay/aluminumoxynitride \
		-v $(PWD).docker-build:/go/src/github.com/eyedeekay/aluminumoxynitride \
		-v $(GOPATH)/src/github.com/eyedeekay/go-I2P-jpackage:/go/src/github.com/eyedeekay/go-I2P-jpackage \
		eyedeekay/i2p.plugins.tor-manager make build
	#go build
	cp -v $(PWD).docker-build/aluminumoxynitride* $(PWD)
	sudo chown $(USER):$(USER) $(PWD)/aluminumoxynitride*

all: linux darwin windows freebsd openbsd netbsd

backup-embed:
	mkdir -p ../../../github.com/eyedeekay/go-I2P-jpackage.bak
	cp ../../../github.com/eyedeekay/go-I2P-jpackage/* ../../../github.com/eyedeekay/go-I2P-jpackage.bak -r;true
	rm -f ../../../github.com/eyedeekay/go-I2P-jpackage/*.tar.xz
	tar -cvJf ../../../github.com/eyedeekay/go-I2P-jpackage/build.windows.I2P.tar.xz README.md LICENSE
	tar -cvJf ../../../github.com/eyedeekay/go-I2P-jpackage/build.linux.I2P.tar.xz README.md LICENSE

unbackup-embed:
	cp ../../../github.com/eyedeekay/go-I2P-jpackage.bak/*.tar.xz ../../../github.com/eyedeekay/go-I2P-jpackage/; true

unembed-windows:
	mv ../../../github.com/eyedeekay/go-I2P-jpackage/build.windows.I2P.tar.xz ../../../github.com/eyedeekay/
	tar -cvJf ../../../github.com/eyedeekay/go-I2P-jpackage/build.windows.I2P.tar.xz README.md LICENSE

unembed-linux:
	mv ../../../github.com/eyedeekay/go-I2P-jpackage/build.linux.I2P.tar.xz ../../../github.com/eyedeekay/
	tar -cvJf ../../../github.com/eyedeekay/go-I2P-jpackage/build.linux.I2P.tar.xz README.md LICENSE

linux: unbackup-embed backup-embed unembed-windows
	GOOS=linux GOARCH=amd64 make docker
	GOOS=linux GOARCH=arm make docker
	GOOS=linux GOARCH=arm64 make docker

darwin: unembed-linux unembed-windows
	GOOS=darwin GOARCH=amd64 make build
	GOOS=darwin GOARCH=arm64 make build

windows: unbackup-embed backup-embed unembed-linux
	GOOS=windows GOARCH=amd64 make build
	GOOS=windows GOARCH=386 make build

freebsd: unembed-linux unembed-windows
	GOOS=freebsd GOARCH=amd64 make build
	GOOS=freebsd GOARCH=386 make build
	GOOS=freebsd GOARCH=arm make build
	GOOS=freebsd GOARCH=arm64 make build

openbsd: unembed-linux unembed-windows
	#GOOS=openbsd GOARCH=amd64 make build
	#GOOS=openbsd GOARCH=386 make build
	#GOOS=openbsd GOARCH=arm make build
	#GOOS=openbsd GOARCH=arm64 make build

netbsd: unembed-linux unembed-windows
	#GOOS=netbsd GOARCH=amd64 make build
	#GOOS=netbsd GOARCH=386 make build
	#GOOS=netbsd GOARCH=arm make build
	#GOOS=netbsd GOARCH=arm64 make build

upload-all:
	GOOS=linux GOARCH=amd64 make upload
	GOOS=linux GOARCH=arm make upload
	GOOS=linux GOARCH=arm64 make upload
	GOOS=darwin GOARCH=amd64 make upload
	GOOS=darwin GOARCH=arm64 make upload
	GOOS=windows GOARCH=amd64 make upload
	GOOS=windows GOARCH=386 make upload
	GOOS=freebsd GOARCH=amd64 make upload
	GOOS=freebsd GOARCH=386 make upload
	GOOS=freebsd GOARCH=arm make upload
	GOOS=freebsd GOARCH=arm64 make upload
	
upload-disabled:
	GOOS=openbsd GOARCH=amd64 make upload
	GOOS=openbsd GOARCH=386 make upload
	GOOS=openbsd GOARCH=arm make upload
	GOOS=openbsd GOARCH=arm64 make upload
	GOOS=netbsd GOARCH=amd64 make upload
	GOOS=netbsd GOARCH=386 make upload
	GOOS=netbsd GOARCH=arm make upload
	GOOS=netbsd GOARCH=arm64 make upload

upload:
	gothub upload -R -u eyedeekay -r $(APP) -t $(VERSION) -n $(APP)-$(GOOS)-$(GOARCH) -f $(APP)-$(GOOS)-$(GOARCH)

clean:
	git clean -fd
	rm -rf basic extensions i2pchrome.js i2pchromium-browser localcdn onionbrowse scriptsafe ublockorigin aluminumoxynitride*