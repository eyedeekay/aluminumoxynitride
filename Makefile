
VERSION=0.0.1
DESC="A wrapper for Chrome-based browsers that auto-configures them to use I2P."
APP=aluminumoxynitride

echo:
	echo gothub release -u eyedeekay -r $(APP) -t $(VERSION) -n $(VERSION) -d $(DESC)

build:
	go build --tags "netgo osusergo" -o $(APP)-$(GOOS)-$(GOARCH)

version:
	gothub release --pre-release -u eyedeekay -r $(APP) -t $(VERSION) -n $(VERSION) -d $(DESC)

release: all version upload-all

all:
	GOOS=linux GOARCH=amd64 make build
	GOOS=linux GOARCH=arm make build
	GOOS=linux GOARCH=arm64 make build
	GOOS=darwin GOARCH=amd64 make build
	GOOS=darwin GOARCH=arm64 make build
	GOOS=windows GOARCH=amd64 make build
	GOOS=windows GOARCH=386 make build
	GOOS=freebsd GOARCH=amd64 make build
	GOOS=freebsd GOARCH=386 make build
	GOOS=freebsd GOARCH=arm make build
	GOOS=freebsd GOARCH=arm64 make build
	GOOS=openbsd GOARCH=amd64 make build
	GOOS=openbsd GOARCH=386 make build
	GOOS=openbsd GOARCH=arm make build
	GOOS=openbsd GOARCH=arm64 make build
	GOOS=netbsd GOARCH=amd64 make build
	GOOS=netbsd GOARCH=386 make build
	GOOS=netbsd GOARCH=arm make build
	GOOS=netbsd GOARCH=arm64 make build

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
