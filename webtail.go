package main

import (
  "code.google.com/p/go.net/websocket"
  "flag"
  "log"
  "net/http"
  "text/template"
  "bufio"
  "os"
  "os/exec"
)

var line string
var addr = flag.String("addr", ":8080", "http service address")
var homeTempl = template.Must(template.ParseFiles("template/index.html"))

func homeHandler(c http.ResponseWriter, req *http.Request) {
  homeTempl.Execute(c, req.Host)
}

func main() {
  flag.Parse()
  go h.run()

  go func() {
    reader := bufio.NewReader(os.Stdin);

    for {
      line,_ := reader.ReadString('\n');
      h.broadcast <- line
    }
  }()

  http.HandleFunc("/", homeHandler)
  http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
  http.Handle("/ws", websocket.Handler(wsHandler))

  go func() {
    log.Println("test")
    cmd := exec.Command("open", "http://localhost:8080/")
    cmd.Run()
  }()

  log.Println("http://localhost:8080/ listen start...")
  if err := http.ListenAndServe(*addr, nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
