.PHONY: deps run

deps:
	go mod download

run:
	go run main.go config.go

demo:
	@echo "Response for repositories with languages=go&limit=4:"
	@curl -s "http://localhost:5000/repositories?languages=go&limit=4" > response_go.json && cat response_go.json
	@echo "\nResponse for repositories with languages=go,javascript&limit=4:"
	@curl -s "http://localhost:5000/repositories?languages=go,javascript&limit=4" > response_go_javascript.json && cat response_go_javascript.json
	@echo "\nResponse for repositories with languages=go&min_stars=5000&limit=4:"
	@curl -s "http://localhost:5000/repositories?languages=go&min_stars=5000&limit=4" > response_go_min_stars_5000.json && cat response_go_min_stars_5000.json
