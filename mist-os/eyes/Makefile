OS := $(shell uname)
CAM_ID := 0

XML_FILE := haarcascade_frontalface_default.xml

ifeq ($(OS), Darwin)
    PKG_PATH=/opt/homebrew/lib/pkgconfig
else
    PKG_PATH=/usr/local/lib/pkgconfig:/usr/lib/aarch64-linux-gnu/pkgconfig
endif

run:
	PKG_CONFIG_PATH=$(PKG_PATH) go run main.go $(CAM_ID) $(XML_FILE)