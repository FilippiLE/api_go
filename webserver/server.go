package main

import(
"fmt"
//"os"
"runtime"
"github.com/gin-gonic/gin"
)

var(
Router *gin.Engine
)
	
func main() {
	startApp()
}

func startApp(){
configureRuntime()
mapURLsToControllers()
Router = gin.Default()
Router.RedirectFixedPath = false 
Router.RedirectTrailingSlash = false
Router.Run(":8080")
}

func configureRuntime(){
	numCPU := runtime.NumCPU()
	fmt.Println(numCPU)
	runtime.GOMAXPROCS(numCPU)
}




