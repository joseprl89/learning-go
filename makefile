green:
	go fmt ./...
	go test ./... && git commit -am "GREEN: `date '+%d/%m/%Y %H:%M'`" || make revert

red:
	go fmt ./...
	go test ./... && make revert || git commit -am "RED: `date '+%d/%m/%Y %H:%M'`"

revert:
	git reset --hard

coverage:
	go test ./... -race -coverprofile=profile.out -covermode=atomic $d
	rm profile.out
