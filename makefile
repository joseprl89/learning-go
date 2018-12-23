verify:
	git add
	go test && make commit || make revert

revert:
	git reset --hard

commit:
	git add .
	git commit -m "`date '+%d/%m/%Y %H:%M'`: Tests passed!"
