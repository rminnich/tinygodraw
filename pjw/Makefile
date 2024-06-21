build:
	tinygo build -target pyportal 
flash:
	tinygo flash -target pyportal 
mon:
	tinygo flash -monitor -target pyportal 
imports:
	goimports -w *.go
	go mod tidy
host:
	go build .
