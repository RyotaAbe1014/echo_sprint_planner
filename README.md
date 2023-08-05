## project start
```shell
$ docker-compose build
$ docker-compose up -d
```

## project stop
```shell
$ docker-compose down
```

## project restart
```shell
$ docker-compose restart
```

## project logs
```shell
$ docker-compose logs -f --tail=100 web
```

## project bash
```shell
$ docker-compose exec web bash
```

## migrations
```shell
$ GO_ENV=dev go run infra/db/migration/migrate.go
```

## web server
```shell
$ GO_ENV=dev go run main.go
```
