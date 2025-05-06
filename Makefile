APP_NAME=server
GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING= "root:root1234@tcp(127.0.0.1:33306)/shopdevgo"
GOOSE_MIGRATION_DIR ?= sql/schema


# docker_build:
# 	docker-compose up -d --build
# 	docker-compose ps
# docker_stop:
# 	docker-compose down	
# docker_up:
# 	docker-compose up -d	

# Mục tiêu để khởi động tất cả các container
docker_start_all:
	@docker start $(shell docker ps -a -q)

# Mục tiêu để dừng tất cả các container (tuỳ chọn thêm)
docker_stop_all:
	@docker stop $(shell docker ps -a -q)

# tạo các bảng dưới db
up_by_one:
	goose -dir=$(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up-by-one
# create new a migration
create_migration:
	goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql
run:
	go run ./cmd/$(APP_NAME)/

upse:
	goose -dir=$(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up
downse:
	goose -dir=$(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) down
resetse:
	goose -dir=$(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) reset
sqlgen:
	sqlc generate			
swag:
	swag init -g ./cmd/server/main.go -o ./cmd/swag/docs

.PHONY: docker_start_all docker_stop_all run upse downse resetse docker_build docker_stop docker_up sqlgen swag docker_start_container docker_stop_container
