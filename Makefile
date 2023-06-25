build-mockgen:
	@echo "Building mockgen started"
	docker build -f mockgen.Dockerfile --tag timer-mockgen .
	@echo "Building mockgen done"

mockgen: build-mockgen
	@echo "Generating mocks"
	docker run --rm --volume "$$(pwd):/home/mockgen/src" timer-mockgen
	@echo "Generating mocks done"