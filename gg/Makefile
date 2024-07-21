build:
	tinygo build -o mclock.uf2 -target pyportal 
	ls -l mclock.uf2

mount:
	-sudo mount -o uid=1000,gid=1000 /dev/sdb /media/PORTALBOOT
	-sudo mount -o uid=1000,gid=1000 /dev/sda /media/PORTALBOOT
	ls -l /media/PORTALBOOT/CURRENT.UF2

flash: mount
	tinygo flash -target pyportal 
	sudo umount -f /media/PORTALBOOT

mon: mount
	tinygo flash -port /dev/ttyACM0 -monitor -target pyportal 
	sudo umount -f /media/PORTALBOOT

imports:
	goimports -w *.go
	go mod tidy
host:
	go build .
f:
	sudo mount /dev/sda /media/rminnich/ && sudo cp mclock.uf2 /media/rminnich/CURRENT.UF2  && sudo umount /media/rminnich 
b:
	sudo mount /dev/sdb -o uid=1000,gid=1000 /media/rminnich/ && sudo cp mclock.uf2 /media/rminnich/CURRENT.UF2  && sudo umount /media/rminnich 
