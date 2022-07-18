attach-app:
	docker-compose exec app-c bash
attach-db:
	docker-compose exec mysql-c bash
fmt:
	docker-compose exec app-c go fmt ./...
