build:
	@gomobile bind -o frameworks/Provider.framework -target=ios -ldflags="-s -w" ./provider
