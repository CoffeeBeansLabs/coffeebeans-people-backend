# coffeebeans-people-backend

#Installation steps

Install golang
mkdir -p $HOME/go/{bin,src}
Set following in .bash_profile export GOPATH=$HOME/go export PATH=$PATH:$GOPATH/bin
Clone coffeebeans-people-backend
Install mongo (or use the docker-compose file) 
In project root run "go mod tidy"
In project root run "go mod vendor"
In project root run "go run main/*.go -file=dev_local.json"
