build:
	go build -a --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo .
