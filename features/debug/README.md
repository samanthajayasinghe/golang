# Debugging Go Code 

## Debug extensions 
* GDB ( GNU Project Debugger ) 
* [Delve](https://github.com/go-delve/delve) 


## How to install Delve
```
$ git clone https://github.com/go-delve/delve
$ cd delve
$ go install github.com/go-delve/delve/cmd/dlv
```

## How to debug go code with Delve
```
dlv debug main.go
## Add break point 
break ./main.go:14
contunue
print {varname}
```

## How to debug with vscode 
Create a launch Json file with config 
```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Main debug",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/features/debug/main.go",
            "env": {
                "GO111MODULE": "on"
            }
        }
    ]
}
```
More Info 


https://github.com/golang/vscode-go/wiki/debugging

## Remote Debugging 

In launch Json 
```
{
    "name": "Launch Remote",
    "type": "go",
    "request": "attach"
    "mode": "remote",
    "remotePath": "{remote path}",
    "port": 2345,
    "host": "{host}",
    "cwd": "${workspaceFolder}",
    "trace": "verbose"
}
```
