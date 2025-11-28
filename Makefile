day01-test:
	go test ./days/day01/
day01-bench:
	go test -bench . -run '^$\' -benchmem ./days/day01