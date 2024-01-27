p=$(shell pwd)
DATABASE_URI=mysql://user:user@/bochat

gen:
	go run $(CURDIR)/tools/gen_code/gen.go

server:
	go run main.go server --app --platform --migrate_sql

migrate:
	go run main.go migrate

test:
	make test.delivery
	make test.service
	make test.repository


check:
	golangci-lint run -v --tests=false --skip-dirs vendor,tools --timeout 4m  
	find ./ -iname 'sqlite.db' -print -exec rm -f '{}' \; 
	CGO_ENABLED=1 PROJ_DIR=$p go test bochat/pkg/...  

test.delivery:
	PROJ_DIR=$p go test -v bochat/pkg/delivery/...

test.service:
	PROJ_DIR=$p go test -v bochat/pkg/service/...

test.repository:
	PROJ_DIR=$p go test -v bochat/pkg/repository/...

gen.graphql:
	-gqlgen generate --verbose --config $(CURDIR)/platform_gqlgen.yml
	make gen.graphql.format
	
gen.graphql.format:
	go fmt $(CURDIR)/pkg/delivery/graph/platform/*.go

gen.mock:
	mockgen -source pkg/iface/repository.go -destination internal/mock/repository_mock.go -package mock
	mockgen -source pkg/iface/service.go -destination internal/mock/service_mock.go -package mock
	# mockgen -source ./vendor/github.com/bsm/redislock/redislock.go -destination internal/mock/redis_lock_mock.go -package mock
	sed -i 's/bochat\/vendor\///g' $(CURDIR)/internal/mock/*_mock.go

gen.dto.easyjson:
	easyjson -all ./pkg/model/dto/message.go

gen.migrate:
	goose -dir "./deployment/database" create $(n) sql 

migrate.db.up:
	goose mysql "user:user@/bochat?parseTime=true" up -dir ./deployment/database

migrate.db.down:
	goose mysql "user:user@/bochat?parseTime=true" down -dir ./deployment/database

chglog:
	git-chglog $(shell git describe --tags `git rev-list --tags --max-count=1`) | cat - CHANGELOG.md > temp && mv temp CHANGELOG.md && \
	git add CHANGELOG.md && \
	git commit -m "chore : update change log from $(shell git describe --tags `git rev-list --tags --max-count=1`^) to $(shell git describe --tags `git rev-list --tags --max-count=1`)" && \
	git tag -f $(shell git describe --tags `git rev-list --tags --max-count=1`) 