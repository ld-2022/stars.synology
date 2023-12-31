# Copyright (c) 2000-2020 Synology Inc. All rights reserved.

## You can use CC CFLAGS LD LDFLAGS CXX CXXFLAGS AR RANLIB READELF STRIP after include env.mak
include /env.mak

EXEC= xingxingInstall
SRC= src/main.go

all: $(EXEC)

$(EXEC): $(SRC)
	go build -o $@ $(SRC)

install: $(EXEC)
	mkdir -p $(DESTDIR)/bin
	install $< $(DESTDIR)/bin
	cp -r src/resources $(DESTDIR)/bin

clean:
	rm -rf $(EXEC)