green:
	go fmt ./...
	make test && make commit_green || make revert

commit_green:
	git add .
	git commit -am "GREEN: `date '+%d/%m/%Y %H:%M'`"
	git push

red:
	go fmt ./...
	make test && make revert || make commit_red

commit_red:
	git add .
	git commit -am "RED: `date '+%d/%m/%Y %H:%M'`"
	git push

test:
	go test ./...

revert:
	git reset --hard

run-server:
	go run euler/server

coverage:
	go test ./... -race -coverprofile=profile.out -covermode=atomic $d
	rm profile.out
