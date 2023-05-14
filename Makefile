build_pathserve:
	docker build -t pathserve . && docker run -v $(CURDIR)/npmpackage:/build -t pathserve cp -R /app/bin /build
build_client:
	rm -r /home/ganbatte/Desktop/project/proxypath2/npmworkspace/packages/client 2> /dev/null || true && cd client && npm run build && cp -R dist/* ../npmpackage/client/