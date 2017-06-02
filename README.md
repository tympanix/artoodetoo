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
Build the server application using:
```shell
go build
```

Follow the instruction in [Quick Installation](#quick-installation) to deploy the server

Now start the web application development server
```shell
cd web && npm start
```

## Quick Installation
First, create an application secret. Type the following command:
```shell
artoodetoo gensecret
```

Now, create a user for the application. Execute and follow the prompt instructions:
```shell
artoodetoo adduser
```

Now start the application itself by executing:
```
artoodetoo
```
