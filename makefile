green:
	go fmt ./...
	go test ./... && git commit -am "GREEN: `date '+%d/%m/%Y %H:%M'`" || make revert
	git push

red:
	go fmt ./...
	go test ./... && make revert || git commit -am "RED: `date '+%d/%m/%Y %H:%M'`"
	git push

revert:
	git reset --hard

run-server:
	go run euler/server

coverage:
	go test ./... -race -coverprofile=profile.out -covermode=atomic $d
	rm profile.out
