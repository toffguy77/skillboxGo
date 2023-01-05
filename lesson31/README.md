
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