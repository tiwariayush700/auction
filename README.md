# Auction


####ABOUT
This is a golang library which allows you to bid against items.
It follows :
 
``1:  Admins can create items and auctions. Auctions will have start and end time``

``2:  Any user can take part in any auction provided he bids price is more than the base(start price).``

``3:   You can only bid between the auction timings set by ADMIN``




#### Run Locally
```bash
1) Install golang
2) mkdir -p $HOME/go/{bin,src} 
3) Set following in .bash_profile 
	export GOPATH=$HOME/go
	export PATH=$PATH:$GOPATH/bin
4) Clone auction
5) docker-compose up
6) go run /cmd/main.go -file=local.json
```

#### Import Locally
```
go get -u github.com/tiwariayush700/auction
```

### APIS CAN BE TESTED FROM THE
`/apiTestLocal directory`
 