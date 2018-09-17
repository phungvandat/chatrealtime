bin: 
	go build -o bin/migrator ./cmd/migrator
local-db: bin
	eval "docker-compose -f localdb-docker-compose.yaml down"
	eval "docker-compose -f localdb-docker-compose.yaml up -d"
local-env: local-db
	@cat .env_migrator.yaml.chatsex > .env_migrator.yaml
	@cat .env.chatsex > .env
	@echo "Waiting for database connection..."
	@while ! docker exec chat-sex-v10_db_1 pg_isready -h localhost -p 5432 > /dev/null; do \
		sleep 1; \
	done
	bin/migrator up
runner: bin
	go build -o server cmd/server/*.go
run: runner
	ENV=local ./server; rm server
build:
	cd cmd/server && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o api .

