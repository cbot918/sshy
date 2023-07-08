# SSHY
a convenient ssh client for experiment purpose
# .env
for example
- URL="127.0.0.1:22"
- USER="userr"
- PASSWORD="abcde12345"

# Run
```
go run .
```
or 
```
go build . && ./sshy
```
# SUPPORT COMMAND
只能用一些簡單指令
```
ls
mkdir
touch 
```
etc
# UNSUPPORT COMMAND
```
cd
clear
vim
```

# Todo
- a local server keep the ssh connection for cli operate
