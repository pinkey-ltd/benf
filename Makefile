SUBDIRS := services

.PHONY: all clean

all:
	$(MAKE) -C $@

clean:
	for dir in $(SUBDIRS); do \
        $(MAKE) -C $$dir clean; \
    done