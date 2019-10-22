
.PHONY:README
README:
	go run tools/generateReadMe.go

.PHONY:clean
clean:
	-rm README.md