# Simple Cross-Site Scripting Lab

## App

Application vulnerable to XSS.

## Attacker

Attacker has only one endpoint '/', which listens for incoming
connections and logs all the query parameters to `stdout`.

## Setup

### With Docker

* Starting the vulnerable application

```bash
docker run -it \
    -e PORT=8080 \
    -e JWT_SECRET="$(head -c 50 /dev/urandom | base64)" \
    -p 8080:8080 \
    murtazau/lab-xss app
```

* Starting the attacker

```bash
docker run -it \
    -e ATTACKER_PORT=5000 \
    -p 5000:5000 \
    murtazau/lab-xss attack
```

### Without Docker

```bash
git clone github.com/murtaza-u/lab-xss
cd lab-xss
go build -o lab-xss cmd/lab-xss/main.go
```

* Starting the vulnerable application

```bash
export JWT_SECRET="$(head -c 50 /dev/urandom | base64)"
export PORT=8080
./lab-xss app
```

* Starting the attacker

```bash
export ATTACKER_PORT=5000
./lab-xss attack
```
