.PHONY: default
default:
	@echo Running developer server...
	go run *.go;
build_client:
	@echo Start build client
	cd client && npx parcel build --out-dir ../public/static/ index.html
