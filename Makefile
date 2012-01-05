6.out: conway.6 main.6
	6l main.6
conway.6: conway.go
	6g conway.go
cell.6: cell.go
	6g cell.go
main.6: main.go conway.go
	6g main.go
clean:
	rm -f *.6 6.out
rebuild: clean 6.out
