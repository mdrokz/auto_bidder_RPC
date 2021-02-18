echo "running Bid Service"
go run ./bid/main.go &
echo "running Project Service"
go run ./project/main.go &
echo "running Auth Service"
go run ./auth/main.go &

wait