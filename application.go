package main

import "git.amabanana.com/plancks-cloud/pc-go-brutus/controller/startup"

func main() {

	//Start the web server
	// TODO: start the web server

	//This figures out whether it's normal start or a install start
	startup.Init()
	
}
