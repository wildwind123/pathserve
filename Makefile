build_pathserve:
	docker build -t pathserve . && docker run -v $(CURDIR)/npmworkspace/packages/bin:/build -t pathserve cp -R /app/bin /build

build_client:
	cd npmworkspace/packages/client && pnpm install && pnpm build

publish_packages:
	pnpm publish --recursive --filter "@pathserve/messenger" --filter "@pathserve/client" --filter="@pathserve/bin"

change-version:
	cd $(CURDIR)/npmworkspace/packages/bin && pnpm version ${version} && cd $(CURDIR)/npmworkspace/packages/client && pnpm version ${version} && cd $(CURDIR)/npmworkspace/packages/messenger && pnpm version ${version} 
	