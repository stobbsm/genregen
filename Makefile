server-dev:
	go build -o tmp/genregen server.go

frontend-dev:
	cd vue; \
	yarn build --mode development

server-prod:
	go build -o bin/genregen server.go

frontend-prod:
	cd vue; \
	yarn build --mode production

dev: frontend-dev server-dev

prod: frontend-prod server-prod