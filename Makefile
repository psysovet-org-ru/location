up:
	docker-compose -f ./docker/docker-compose.yaml up -d

down:
	docker-compose -f ./docker/docker-compose.yaml down --remove-orphans

bld:
	@cd ./src &&  go build  -o ../build/location && cd ..

