package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/eclipse/paho.mqtt.golang"
)

var mqttClient mqtt.Client

func main() {
    opts := mqtt.NewClientOptions().AddBroker("tcp://broker.hivemq.com:1883")
    mqttClient = mqtt.NewClient(opts)
    if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
        log.Fatal(token.Error())
    }

    http.HandleFunc("/forward", handleForward)
    http.HandleFunc("/backward", handleBackward)
    http.HandleFunc("/left", handleLeft)
    http.HandleFunc("/right", handleRight)

    fmt.Println("Serveur démarré sur le port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleForward(w http.ResponseWriter, r *http.Request) {
    sendCommand("F")
    fmt.Fprintf(w, "moving forward")
}

func handleBackward(w http.ResponseWriter, r *http.Request) {
    sendCommand("B")
    fmt.Fprintf(w, "moving backward")
}

func handleLeft(w http.ResponseWriter, r *http.Request) {
    sendCommand("L")
    fmt.Fprintf(w, "turning left")
}

func handleRight(w http.ResponseWriter, r *http.Request) {
    sendCommand("R")
    fmt.Fprintf(w, "turning right")
}

func sendCommand(command string) {
    token := mqttClient.Publish("rc_car/command", 0, false, command)
    token.Wait()
}
