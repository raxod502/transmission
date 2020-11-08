.PHONY: backend-dev
backend-dev:
	watchexec -w backend/src -e go -r -n go run backend/src/main.go

.PHONY: backend-build
backend-build:
	mkdir -p backend/out
	go build -o backend/out/main backend/src/main.go

.PHONY: backend-prod
backend-prod:
	backend/out/main

.PHONY: frontend-dev
frontend-dev:
	npx rollup -c -w

.PHONY: frontend-build
frontend-build:
	npx rollup -c

.PHONY: image-build
image-build:
	docker build . -t transmission

.PHONY: image-run
image-run:
	docker run -it --rm -p 127.0.0.1:3455:3455 transmission

.PHONY: deploy
deploy:
	heroku container:push web
	heroku container:release web
