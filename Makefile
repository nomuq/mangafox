help: # this command
	# [generating help from tasks header]
	@egrep '^[A-Za-z0-9_-]+:' Makefile

osx-build: # Creates Mac OSX
	@GOOS=darwin GOARCH=amd64 go build -o build/mangafox-server ./cmd/server

windows-build: # Creates Windows
	@GOOS=windows GOARCH=amd64 go build -o build/mangafox-server.exe ./cmd/server

linux-build: # Creates Linux
	@GOOS=linux GOARCH=amd64 go build -o build/mangafox-server ./cmd/server

linux-arm-build: # Creates Linux ARM
	@GOOS=linux GOARCH=arm go build -o build/mangafox-server-linux-arm ./cmd/server

linux-arm64-build: # Creates Linux ARM64
	@GOOS=linux GOARCH=arm64 go build -o build/mangafox-server-linux-arm64 ./cmd/server



osx-mangareader-indexer-build: # Creates Mac OSX
	@GOOS=darwin GOARCH=amd64 go build -o build/mangareader-indexer ./cmd/mangareader

windows-mangareader-indexer-build: # Creates Windows
	@GOOS=windows GOARCH=amd64 go build -o build/mangareader-indexer.exe ./cmd/mangareader

linux-mangareader-indexer-build: # Creates Linux
	@GOOS=linux GOARCH=amd64 go build -o build/mangareader-indexer ./cmd/mangareader

linux-arm-mangareader-indexer-build: # Creates Linux ARM
	@GOOS=linux GOARCH=arm go build -o build/mangareader-indexer-linux-arm ./cmd/mangareader

linux-arm64-mangareader-indexer-build: # Creates Linux ARM64
	@GOOS=linux GOARCH=arm64 go build -o build/mangareader-indexer-linux-arm64 ./cmd/mangareader

builds: # Creates executables for OSX/Windows/Linux
	@make osx-build
	@make windows-build
	@make linux-build
	@make windows-gui-build
	@make osx-mangareader-indexer-build
	@make windows-mangareader-indexer-build
	@make linux-mangareader-indexer-build
	@make linux-arm-mangareader-indexer-build

remove-builds: # Remove executables
	@rm -rf build/