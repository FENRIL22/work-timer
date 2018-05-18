all:
	@echo "hoge"

init:
	go get github.com/airking05/termui
	#go get -u github.com/gizak/termui

init-statik:
	go get github.com/rakyll/statik
	statik -src=html/
