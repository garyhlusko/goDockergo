package main

import (
	"net"
	"fmt"
	"strconv"
	"os"
	"os/exec"
	"flag"
)


func printError(err error){
	if err != nil {
		fmt.Printf("error: ",err,"\n")
	}
}

func launchGo(appPort int,sqlPort int,webPort int){
	dbPwd := flag.String("dbPwd","toor","Database Password")
	dbName := flag.String("dbName","foo","Database Name")
	dbUser := flag.String("dbUser","foo","Database User")
	network_name := flag.String("network_name","foo","NetworkName")
	flag.Parse()
	fmt.Println(*dbPwd,*dbName,*dbUser)

	dir, _ := os.Getwd()
	
	fmt.Println("Starting docker....")   
	dbPortString := "DB_PORT="+strconv.Itoa(sqlPort)+"\n"
	webPortString := "WEB_PORT="+strconv.Itoa(webPort)+"\n"
	pathString := "PATH="+dir+":/usr/bin:$PATH"+"\n"
	appPortString := "APP_PORT="+strconv.Itoa(appPort)+"\n"
	dbNameString := "DB_NAME="+*dbName+"\n"
	dbUserString := "DB_USER="+*dbUser+"\n"
	dbPwdString := "DB_PASSWORD="+*dbPwd+"\n"
	networkNameString := "network_name="+*network_name+"\n"
	
	f,err := os.Create(".env")
	if err != nil {
		fmt.Println(err)
	}
	
	f.WriteString(appPortString)
	f.WriteString(dbPortString)
	f.WriteString(webPortString)
	f.WriteString(pathString)
	f.WriteString(dbNameString)
	f.WriteString(dbUserString)
	f.WriteString(dbPwdString)
	f.WriteString(networkNameString)
	f.Close()
	
	networkBuildNameString := *network_name+"docker_network"

	exec.Command("docker","network","create",networkBuildNameString)
	cmd := exec.Command("sudo","docker-compose","up","--build")
	cmd.Env= append(os.Environ(),appPortString,dbPortString,
				webPortString,pathString)
	out,err := cmd.CombinedOutput()

	printError(err)
	fmt.Printf("%s\n",out)
}

func checkPort(port int) (int) {
	fmt.Println("Trying Port: "+strconv.Itoa(port))
	ln, err := net.Listen("tcp",":"+strconv.Itoa(port))
	if err != nil{
		fmt.Println("Port: ",port," is close")
		return 0
	} else {	
		fmt.Println("Port:",port," is open")
		ln.Close()
		return port
	}
}

func getPort(startPort int) (int){
	var port int
	for port == 0 {
		port = checkPort(startPort)
		if port == 0 {
			startPort = startPort + 1
		}
	}

	return port 
}

func main(){
	SqlPort := getPort(5432)
	fmt.Println("SQL Port is:", SqlPort)

	WebPort := getPort(81)
	fmt.Println("Sql Port is:", WebPort)

	AppPort:= getPort(8000)
	fmt.Println("App port is:", AppPort)
	
	launchGo(AppPort,SqlPort,WebPort)
	fmt.Println("That'll do pig. That'll do.")

}
