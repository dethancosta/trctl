testCore:
	go run main.go build -f ~/Developer/tr-cli/test/build.csv
	#go run main.go newCurrent -d "NewTask" -u 23:54:00

testUpdate:
	go run main.go update -f ~/Developer/tr-cli/test/update.csv
