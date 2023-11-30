testCore:
	go run main.go build -f ~/Developer/tr-cli/test/build.csv
	go run main.go newCurrent -d "NewTask" -u 23:59:00
