testCore:
	go run main.go build -f ~/Developer/tr-cli/test/build.csv
	#go run main.go newCurrent -d "NewTask" -e 23:54:00

testUpdate:
	go run main.go update -f ~/Developer/tr-cli/test/update.csv

testInit:
	#rm ~/.timeruler/config.json
	go run main.go init -u "User" -p "PAssword" -s "localhost:69" -b "buildybuild.csv"
