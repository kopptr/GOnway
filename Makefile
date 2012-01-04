6.out: board.6 main.6
	6l main.6
board.6: board.go
	6g board.go
main.6: main.go board.go
	6g main.go
clean:
	rm -f *.6 6.out
rebuild: clean 6.out
