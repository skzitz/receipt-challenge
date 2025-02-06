##
# Project Title
#
# @file
# @version 0.1

GOCMD		=	go
#GOFLAGS	=	-v -work -x
BINDIR		=	${PWD}/bin

.PHONY: all
all: receipt-server # receipt-client

${BINDIR}/receipt-server ${BINDIR}/server: receipt-server
	cp receipt-server ${BINDIR} && cd ${BINDIR} && ln -sf receipt-server server

${BINDIR}/receipt-client ${BINDIR}/client: receipt-client
	cp receipt-client ${BINDIR} && cd ${BINDIR} && ln -sf receipt-client client

.PHONY: install
install: ${BINDIR}/receipt-server # ${BINDIR}/receipt-client

# Go-specific targets
.PHONY: go-test
go-test:
	${GOCMD} test

.PHONY: go-vert
go-vet:
	${GOCMD} vet


# Let go manage the dependency for receipt-server
.PHONY: receipt-server
receipt-server:
	${GOCMD} build ${GOFLAGS} -o $@

#receipt-client:
#	${GOCMD} build -C client ${GOFLAGS} -o $@

erf:
	@echo ${BINDIR}

clean:
	rm -f receipt-server receipt-client
	rm -f ${BINDIR}/receipt-server ${BINDIR}/server
	rm -f ${BINDIR}/receipt-client ${BINDIR}/receipt-client
	rm -f *~ client/*~ __debug*

# end
