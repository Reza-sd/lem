hello:
	echo "Hello, world!"
test-internal:
	find ./internal/ -type d -exec go test -v {} \;
pretty:
	gofmt -l -w . && goimports -d -w . && go mod tidy
update:
	sudo apt upgrade -y && sudo apt update -y
lint:
	golint ./... && echo "========vet============" && go vet ./... && echo "======================="
git:
	git add . && git commit -m "." && git push
# If you only want to undo changes in the working directory without affecting the staging area
re:
	git restore .



