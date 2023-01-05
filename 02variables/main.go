package main

import "fmt"

const LoginToken string = "asdf jkl;" // Public

func main() {
  var username string = "ruban"
  fmt.Println(username)
  fmt.Printf("Variable is of type: %T \n", username)

  var isLoggedIn bool = true
  fmt.Println(isLoggedIn)
  fmt.Printf("Variable is of type: %T \n", isLoggedIn)
  
  var smallVal uint8 = 255
  fmt.Println(smallVal)
  fmt.Printf("Variable is of type: %T \n", smallVal)

  var smallFloat float64 = 255.45487698763
  fmt.Println(smallFloat)
  fmt.Printf("Variable is of type: %T \n", smallFloat)

  // default values and some aliases
  var anotherVariable int
  fmt.Println(anotherVariable)
  fmt.Printf("Variable is of type: %T \n", anotherVariable)

  var anotherVariable2 float64
  fmt.Println(anotherVariable2)
  fmt.Printf("Variable is of type: %T \n", anotherVariable2)

  var anotherVariable3 string
  fmt.Println(anotherVariable3)
  fmt.Printf("Variable is of type: %T \n", anotherVariable3)

  // implicit types
  var website = "rubansahoo.com"
  fmt.Println(website)
  fmt.Printf("Variable is of type: %T \n", website)

  // no var style
  // walrus operator (:=) not allowed outside the method
  numberOfUser := 300000.0
  fmt.Println(numberOfUser)
  fmt.Printf("Variable is of type: %T \n", numberOfUser)

  fmt.Println(LoginToken)
  fmt.Printf("Variable is of type: %T \n", LoginToken)

}
