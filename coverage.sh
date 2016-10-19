t="/tmp/go-cover.$$.tmp"
go test ./nes -coverprofile=$t $@ && go tool cover -html=$t && unlink $t
