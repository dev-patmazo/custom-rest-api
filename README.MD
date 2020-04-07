# Custom Rest

A custom REST API framework mixed with personal favorite libraries.

## Getting Started

I personally created this framework mixed in with my preferred libraries, with complete configurations for security, routing, databases and much more.

### Prerequisites

To use this custom REST API the following must be installed

```
1. Golang upto most recent version
2. Preferred database (MongoDB, MySQL, Postgres)
3. Golang YAML
4. Golang TOML
5. Goji
```

### Installing

Here are the steps to make this up and running.
Install or pull all required dependencies.

```
1. go get goji.io
2. go get github.com/pelletier/go-toml
3. go get gopkg.in/yaml.v2
4. go get -u github.com/cosmtrek/air
5. go get github.com/sirupsen/logrus
```

### Deployment

```
1. Set runmode.sh as executable file (sudo chmod -Rf 775 runmode.sh)
2. Execute runmode.sh with preferred environment (./runmode.sh dev)
    a. dev
    b. qa
    c. stg
    d. prod
```

## Author

* **Patrick Mark Mazo**