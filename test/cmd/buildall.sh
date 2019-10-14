GOOS=windows GOARCH=amd64 go build -o ./eventloghook/eventloghook.exe ./eventloghook/main.go
go build -o ./fluentdformatter/fluentdformatter ./fluentdformatter/main.go
go build -o ./lfshook/lfshook ./lfshook/main.go
go build -o ./lfshook-simple/lfshook-simple ./lfshook-simple/main.go
go build -o ./nestedformatter/nestedformatter ./nestedformatter/main.go
go build -o ./textformatter/textformatter ./textformatter/main.go
