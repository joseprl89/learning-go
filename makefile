verify:
	go fmt
	go test testing && make commit || make revert

revert:
	git reset --hard

commit:
	git commit -am "`date '+%d/%m/%Y %H:%M'`: Tests passed!"

coverage:
	go test ./... -race -coverprofile=profile.out -covermode=atomic $d
	rm profile.out
