BINARY="NewToDo"
VERSION=1.0.0


default:
	@go build -o ${BINARY}.exe 

run:
	@./${BINARY}.exe 

.PHONY: default run