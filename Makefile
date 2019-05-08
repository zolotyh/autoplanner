SHELL = /bin/sh
.PHONY: run server client open

run: server client open
	@echo Running developer server...

open:
	@echo openning browser
	open 'http://localhost:8000'


client:
	@echo Start build client
	cd client && npm start
server:
	@echo Start build server
	go run *.go;
