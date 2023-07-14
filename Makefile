# Copyright (c) 2000-2020 Synology Inc. All rights reserved.

## You can use CC CFLAGS LD LDFLAGS CXX CXXFLAGS AR RANLIB READELF STRIP after include env.mak
include /env.mak

EXEC= examplePkg
SRC= src

all: $(EXEC)

$(EXEC): $(SRC)
	go env -w GO111MODULE=on
    go env -w GOPROXY=https://goproxy.cn,direct
	cd $(SRC) && go build -o $@ . && cd -

install: $(EXEC)
	mkdir -p $(DESTDIR)/bin
	install $< $(DESTDIR)/bin

clean:
	rm -rf $(EXEC)
