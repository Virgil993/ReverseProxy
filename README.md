# ReverseProxy
A reverse proxy application written in Golang 

To run the project , you must first have installed Golang, MySQL workbench and preferably Visual Studio Code with the Golang extension installed. Follow the steps bellow to run the application: 

1. Download the ReverseProxy folder into your src folder in the directory where GO is installed.
2. Download the ReverseProxyDump folder .
3. Create a new schema called `proxy_requests` in MySQL workbench 
4. Click on settings, data import then choose the folder ReverseProxyDump and click start import
5. Lastly at line 54 in the main.go file there is a function called sql.Open with 2 parameters, the second parameter must match the following structure "username:password@protocol(addres)/dbname" where you need to fill in the username and password for your database, as well as the address and database name, in this case being `proxy_requests`
6. Now, from a terminal go to the ReverseProxy directory and if you run `go run main.go` you should have the server up and running.

My application simulates a reverse proxy server which takes requests from the user and then sends those requests to an API that contains data for testing. After recieving the response from the API , my application modifies that JSON response and sends it back to the user. My project also stores in a database all requests and responses that are sent through the proxy server.
