To run a MySQL container on port 3307, detached mod, accessible through terminal, with local volume.

```
podman run --detach --tty --name orderdb -p 3307:3306 -v ~/mysql_data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD='verysecretpass' -e MYSQL_DATABASE=order mysql:latest
```

Login into the container and enter mysql:
```
podman exec -it orderdb /bin/bash
mysql -uroot -pverysecretpass
```

Delete orderdb container
```
podman rm -f -v orderdb
```

In this case, the data source URL is:

```
root:verysecretpass@tcp(127.0.0.1:3307)/order
```
To run the payment service application, inside of payment folder, use the following:

```sh
DB_DRIVER=mysql \
DATA_SOURCE_URL=root:verysecretpass@tcp(127.0.0.1:3307)/payment \
APPLICATION_PORT=3001 \
ENV=development \
go run cmd/main.go
```

To run the Order service application, you can use the following:
```sh
DB_DRIVER=mysql \
DATA_SOURCE_URL=root:verysecretpass@tcp(127.0.0.1:3307)/order \
APPLICATION_PORT=3000 \
ENV=development \
PAYMENT_SERVICE_URL=localhost:3001 \
go run cmd/main.go
```

Login into the orderdb container
```bash
podman exec -it orderdb /bin/bash
```

Show logs from orderdb
```bash
podman logs orderdb --tail 80
```

Run queries from host:

```bash
docker exec orderdb mysql -uroot -pverysecretpass -e 'SELECT 1;'
```

Check TCP services for port 3306 in your machine:
```bash
lsof -nP -iTCP:3306 -sTCP:LISTEN
```

## Calling a gRPC endpoint

```sh
grpcurl -d '{"user_id": 123, "order_items": [{"product_code": "prod", "quantity":4, "unit_price":12}]}' -plaintext localhost:3000 Order/Create
```
