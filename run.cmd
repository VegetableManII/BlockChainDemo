del /q .\bcdb\*.*
del /q blockchain.exe
go build -o blockchain.exe .\*.go
.\blockchain.exe