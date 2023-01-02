# Golang-Tutorials
## Install Go
## Write some code


Get started with Hello, World.
  
* Open a command prompt and cd to your home directory.
  
* Create a hello directory for your first Go source code.
  
* Enable dependency tracking for your code.  
<br>
  When your code imports packages contained in other modules, you manage those dependencies through your code's own module. That module is defined by a go.mod file that tracks the modules that provide those packages. That go.mod file stays with your code, including in your source code repository.  
<br><br>
  To enable dependency tracking for your code by creating a go.mod file, run the go mod init command, giving it the name of the module your code will be in. The name is the module's module path.  
<br><br>
  In actual development, the module path will typically be the repository location where your source code will be kept. For example, the module path might be github.com/mymodule. If you plan to publish your module for others to use, the module path must be a location from which Go tools can download your module. For more about naming a module with a module path, see Managing dependencies.  
<br><br>
  For the purposes of this tutorial, just use example/hello.  
  <br><br>  

  ```shell
    $ go mod init example/hello
    go: creating new go.mod: module example/hello
  ```
  
* In your text editor, create a file hello.go in which to write your code.
  
* Write the code into your hello.go file and save the file.
<br>
  In this code, you:  
  
  * Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).
  * Import the popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
  * Implement a main function to print a message to the console. A main function executes by default when you run the main package.
  
* Run your code to see the greeting.
  ```shell
    $ go run .
    Hello, World!
  ```

  The go run command is one of many go commands you'll use to get things done with Go. Use the following command to get a list of the others:
  ```shell
    $ go help
  ```
