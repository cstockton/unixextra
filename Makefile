git_repo := https://go.googlesource.com/sys
img_repo := local/posixgen
img_file := gen/.image

archs := $(filter-out sparc64 amd64p32, \
	$(shell go tool dist list|awk -F/ '{print $$2}' | sort | uniq))
ztype_src := $(addprefix gen/sys/unix/ztypes_linux_, $(addsuffix .go, $(archs)))
ztype_dst := $(notdir $(ztype_src))

.PHONY: all
all: $(ztype_dst)

$(ztype_dst): $(ztype_src)
	# make sure all sizes of SiginfoT are 128 bytes before copying
	test -z "$$(grep -r SizeofSiginfo $(ztype_src) | grep -v 0x80)"
	cp $(?) .
	sed -i 's/package unix/package unixextra/g' $(ztype_dst)

$(ztype_src): gen/linux/types.go | $(img_file)
	cp gen/linux/types.go gen/sys/unix/linux/types.go
	docker run -it -v $(CURDIR)/gen/sys/unix:/build $(img_repo):linux

$(img_file): | gen/sys
	docker build -t "$(img_repo):linux" gen/sys/unix/linux
	docker image ls -q "$(img_repo)" > $(@)

gen/sys:
	git -C gen clone $(git_repo)
	cp gen/linux/mkall.go gen/sys/unix/linux/mkall.go

.PHONY: clean
clean:
	rm -f $(ztype_dst)

.PHONY: clean-all
clean-all:
	test ! -f $(img_file) || { \
		docker rmi -f $$(cat $(img_file)) || true; \
		rm $(img_file); \
	};
	rm -rf gen/sys
