.SUFFIXES: .go .s

all: main1 main2 main1.s main2.s

main1: main1.go
	go build main1.go

main2: main2.go
	go build main2.go

.go.s:
	go tool compile -S $< > $@

clean:
	rm -f *.s *.o main1 main2
