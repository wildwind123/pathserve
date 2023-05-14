build_pathserve:
	docker build -t pathserve . && docker run -v $(CURDIR)/npmworkspace/packages/bin:/build -t pathserve cp -R /app/bin /build