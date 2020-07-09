# CAH - API

## Develop

### Help

```bash
make
```

### Configuration

Application can be configure throught env variables:
```bash
CAH_DB_PASSWORD=testingOnLocal1234 make start
```

### DB

I use `/docker-entrypoint-initdb.d` in order to init the DB with schema and cards data.

```
docker-compose -f scripts/docker-compose.yml -p cah-pg up
```

## Contribution

You have to enable [Go modules](https://github.com/golang/go/wiki/Modules) for compiling this project.
