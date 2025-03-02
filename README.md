# bechallenge
Surfe BE coding Challenge

## startup instructions
To get started clone this repo and cd into the directory you have cloned it to

run `go build .` to build the service 

then `.\bechallenge.exe -port=8080` to run it on windows, on port 8080. If port is not supplied the default is 3000

if using VSCode a launch configuration is set up in the .vscode directory. Just open the main.go file and hit F5

## endpoints

the endpoints defined are 
* /users/{id}
* /users/{id}/actions/count
* /actions/{type}/nextactions
* /referralindexes

And can be accessed by prefixing with localhost:PORT

for example if running on port 8080
* localhost:8080/users/1
* localhost:8080/users/1/actions/count
* localhost:8080/actions/WELCOME/nextactions
* localhost:8080/referralindexes


