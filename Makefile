BUILDDIR  ?= $(CURDIR)/build
LDFLAGS   +=-s -w
VERSION   = 1.0.1
PREFIX	  = /usr/bin
MANPREFIX = /usr/local/share/man

${BUILDDIR}/meowfetch:
	mkdir -p ${BUILDDIR}
	GOOS=linux GOARCH=${GOARCH} go build \
	     -ldflags="${LDFLAGS}" \
	     -o ${BUILDDIR}/meowfetch main.go

build: ${BUILDDIR}/meowfetch

install: build
	cp -f ${BUILDDIR}/meowfetch ${DESTDIR}${PREFIX}/
	chmod 755 ${DESTDIR}${PREFIX}/meowfetch
	mkdir -p ${DESTDIR}${MANPREFIX}/man1
	sed "s/{VERSION}/${VERSION}/g" < meowfetch.1 > ${DESTDIR}${MANPREFIX}/man1/meowfetch.1
	chmod 644 ${DESTDIR}${MANPREFIX}/man1/meowfetch.1

uninstall:
	rm -rf ${PREFIX}/meowfetch
	rm -rf ${MANPREFIX}/man1/meowfetch

.PHONY: run
run:
	go run main.go

.PHONY: version
opts:
	@echo "BUILDDIR:${BUILDDIR}"
	@echo "LDFLAGS:${LDFLAGS}"
	@echo "VERSION:${VERSION}"

.PHONY: clean
clean:
	rm -rf ${BUILDDIR}/*
	go clean -v ./...
	@echo "Done."
