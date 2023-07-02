SUBDIRS := aliens galaxies gateway stars

.PHONY: all $(SUBDIRS)

all: $(SUBDIRS)

$(SUBDIRS):
	$(MAKE) -C $@ docker-build