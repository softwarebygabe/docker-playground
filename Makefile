help:
	@echo 'Usage:'
	@echo ''
	@echo '   make up'
	@echo '   make rm'
	@echo ''

up:
	@docker-compose \
		-f compose/pingpong.docker-compose.yml \
		up \
		--build

rm:
	@docker-compose \
			-f compose/pingpong.docker-compose.yml \
			rm \
			-f
