build_publish_all: build_all publish_packages

build_all: build_messenger build_client build_pathserve

build_pathserve:
	docker build -t pathserve . && docker run -v $(CURDIR)/npmworkspace/packages/bin:/build -t pathserve cp -R /app/bin /build
build_client:
	cd npmworkspace/packages/client && pnpm install && pnpm build
build_messenger:
	cd npmworkspace/packages/messenger && pnpm install && pnpm build

publish_packages:
	pnpm publish --recursive --filter "@pathserve/messenger" --filter "@pathserve/client" --filter="pathserve" --no-git-checks

change-update-version: change-version update-version

change-version:
	cd $(CURDIR)/npmworkspace/packages/bin && \
	pnpm version ${version} && \
	cd $(CURDIR)/npmworkspace/packages/client && \
	pnpm version ${version} && \
	cd $(CURDIR)/npmworkspace/packages/messenger && \
	pnpm version ${version} 
update-version:
	cd $(CURDIR)/npmworkspace/packages/client && \
	pnpm update @pathserve/messenger