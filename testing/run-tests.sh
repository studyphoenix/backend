#!/bin/bash

# TODO: Start backend
# TODO: Proper CLI interface
# TODO: Run different tests (i.e. either e2e or per-component)

if [[ $1 == "html" ]]; then
	docker run --rm -v $(pwd):/bruno usebruno/cli run --env-var baseUrl="http://$(ip route get 1.1.1.1 | head -1 | cut -d ' ' -f7):8080" --reporter-html report.html --verbose
else
	docker run --rm -v $(pwd):/bruno usebruno/cli run --env-var baseUrl="http://$(ip route get 1.1.1.1 | head -1 | cut -d ' ' -f7):8080" --verbose
fi
