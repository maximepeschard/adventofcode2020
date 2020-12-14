day ?= $(shell ls -r | egrep -m 1 '\d{2}_.*')
day_clean = $(shell echo $(day) | tr -d '/')
name = $(shell echo $(day_clean) | cut -d '_' -f2)
input = $(shell ls $(day_clean)/input.txt 2> /dev/null)

build:
	go build -o bin/$(name) ./$(day_clean)

run: build
	./bin/$(name) $(input)

test:
	go test ./$(day_clean)/...
