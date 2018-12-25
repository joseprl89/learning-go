green:
	go fmt ./...
	make test && make commit_green || make revert
	git push

commit_green:
	git add .
	git commit -am "GREEN: `date '+%d/%m/%Y %H:%M'`"

red:
	go fmt ./...
	make test && make revert || make commit_red
	git push

commit_red:
	git add .
	git commit -am "RED: `date '+%d/%m/%Y %H:%M'`"

test:
	go test ./...

revert:
	git reset --hard

run-server:
	go run euler/server

coverage:
	go test ./... -race -coverprofile=profile.out -covermode=atomic $d
	rm profile.out
