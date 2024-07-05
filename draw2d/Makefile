build:
	tinygo build -o mclock.uf2 -target pyportal 
	ls -l mclock.uf2
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
b:
	sudo mount /dev/sdb /media/rminnich/ && sudo cp mclock.uf2 /media/rminnich/CURRENT.UF2  && sudo umount /media/rminnich 
