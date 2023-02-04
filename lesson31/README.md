## HOWTO on clean Ubuntu 18.04

```
    1  apt update
    2  apt upgrade
    3  apt install git
    4  wget  https://go.dev/dl/go1.19.linux-amd64.tar.gz
    5  sudo tar -xvf go1.19.linux-amd64.tar.gz
    6  mv go /usr/local
    7  export GOROOT=/usr/local/go
    8  mkdir $HOME/go
    9  export GOPATH=$HOME/go
   10  export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
   11  git clone https://github.com/toffguy77/skillboxGo.git
   12  apt install make
   13  apt install docker.io
   14  apt install docker-compose
   15  apt install gcc
   16  apt install telnet
   17  cd skillboxGo/
   18  git checkout lesson31
   19  cd lesson31
   20  make
```

## Makefile

```
all     clean, test & up
up      build cli & docker compose up
cli     vendor & build cli 
vendor  vendoring (+tidy)
down    docker compose down
mongo   docker compose up -d mongodb
test    up mongo & run tests
clean   docker compose down, prune containers, prune images, rm bin files, erase database
```

## friends-cli usage

```
create, c        create new person
get, g           get person's info
get-all, ga      get all person records
delete, d, del   delete person
update, u        update existing person
make-friend, mf  create friendship relationship for a person
friends, f       get friends list for a person
```
