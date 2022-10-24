help:
	@echo "    server"
	@echo "        Run server"
	@echo "    migrate"
	@echo "        Run migrations"
	@echo "    test"
	@echo "        Run tests"
	@echo "    swagger"
	@echo "        Generate Swagger documentation"


.PHONY: go-env
go-env:
	go env
