Creating a greetings module:

create a directory greetings and inside it create a file greetings.go.

Initialize the greeing package using this command (go mod init example.com/greetings)

Now to use this package we will create a directory hello

create a hello.go file and inside it import the package as "example.com/greetings"

Now when we are working on prod, we will deploy the packages separately but since we are currently using locally then we will use this command:
go mod edit -replace example.com/greetings=../greetings

Now we can use the function of the package