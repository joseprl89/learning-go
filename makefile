green:
	go fmt ./...
	make test && make commit_green || make revert

commit_green:
	git add .
	git commit -am ":white_check_mark: `date '+%d/%m/%Y %H:%M'`"
	git push

red:
	go fmt ./...
	make test && make revert || make commit_red

commit_red:
	git add .
	git commit -am ":red_circle: `date '+%d/%m/%Y %H:%M'`"
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
