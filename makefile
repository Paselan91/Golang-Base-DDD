attach-app:
	docker-compose exec myapp bash
attach-db:
	docker-compose exec myapp-db bash
fmt:
	docker-compose exec myapp go fmt ./...
