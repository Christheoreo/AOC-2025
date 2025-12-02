day01-test:
	go test ./days/day01/
day01-bench:
	go test -bench . -run '^$\' -benchmem ./days/day01
day02-test:
	go test ./days/day02/
day02-bench:
	go test -bench . -run '^$\' -benchmem ./days/day02