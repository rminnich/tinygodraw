build:
	tinygo build -o mclock.uf2 -target pyportal 
flash:
	tinygo flash -target pyportal 
mon:
	tinygo flash -monitor -target pyportal 
imports:
	goimports -w *.go
	go mod tidy
host:
	go build .
f:
	sudo mount /dev/sda /media/rminnich/ && sudo cp mclock.uf2 /media/rminnich/CURRENT.UF2  && sudo umount /media/rminnich
