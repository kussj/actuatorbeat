DATE:=$(shell date +%y%m%d%H%M%S)
BUILDID?=$(DATE)

.PHONY: all
all: actuatorbeat/deb actuatorbeat/rpm actuatorbeat/darwin actuatorbeat/win actuatorbeat/bin build/upload/build_id.txt latest

ES_BEATS=../../vendor/github.com/elastic/beats
include $(ES_BEATS)/dev-tools/packer/scripts/Makefile


.PHONY: actuatorbeat
actuatorbeat: build/upload
	# cross compile on ubuntu
	cd build && xgo -image=tudorg/beats-builder \
		-before-build=../$(ES_BEATS)/dev-tools/packer/xgo-scripts/before_build.sh \
		-pkg $@ \
		-source ../../.. \
		github.com/kussj/$@
	# linux builds on debian 6
	cd build && xgo -image=tudorg/beats-builder-deb6 \
		-before-build=../$(ES_BEATS)/dev-tools/packer/xgo-scripts/before_build.sh \
		-pkg $@ \
		-source ../../.. \
		github.com/kussj/$@

