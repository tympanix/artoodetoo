# Artoodetoo
Awesome self-programmable task runner for automating every day tasks

## Development
### Prerequisites
Development requires that you have set up golang, npm and nodejs on your machine

First, get the source code with the go get utility
```shell
go get github.com/Tympanix/artoodetoo
```

Then install the development dependencies
```shell
cd web && npm install
```

Now install the `@angular/cli` utility
```
npm install -g @angular/cli@latests
```

### Running
Start the server with

(Windows)
```shell
go build && artoodetoo
```

(Unix/Linux)
```shell
go build && ./artoodetoo
```

Now start the web application development server
```shell
cd web && npm start
```
