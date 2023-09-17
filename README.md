# LB
Basic Load Balancer in GoLang 
server folder contains basic golang http server, to create servers run below command with arguments as server name and port
go run main.go server-1 :5001
create N number of servers inside the server folder
now, run modify the list of servers in main.go file in MY-LB folder according to your created servers earlier from line-11
newServer("server-1", "http://127.0.0.1:5001"),

save changes and run below command

go run main.go healthCheck.go server.go
