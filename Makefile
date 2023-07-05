cover:
	go test -v -coverprofile cover.out ./awards

test:
	go test ./awards