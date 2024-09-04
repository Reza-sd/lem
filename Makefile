hello:
	echo "Hello, world!"
test-internal:
	find ./internal/ -type d -exec go test -v {} \;
pretty:
	gofmt -l -w . && goimports -d -w .
