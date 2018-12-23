test_commit_or_revert:
	dt=$(date '+%d/%m/%Y %H:%M:%S');
	go test && make commit || make revert

revert:
	git reset --hard

commit:
	dt=$(date '+%d/%m/%Y %H:%M:%S');
	git add .
	git commit -m "${dt}: Tests passed!"
