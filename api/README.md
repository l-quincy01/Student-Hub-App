## Setup Environment

### 1. Configure dependecies
Install the `wire` package tool:
```bash
go install wire@latest
```

***Note**
For macOS please add the following to your .zshrc
```
export GOPATH="$HOME/go"
PATH="$GOPATH/bin:$PATH"
```

To configure the `api` dependecies run the following make command:
```bash
make wire
```

### 2. Run Mongodb Container

