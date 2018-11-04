package main

import (
    . "gs-go-email/config"
    . "gs-go-email/mailer"
)

var config = Config{}

func init() {
    config.Read()
}

func main() {
    subject := "Get latest Tech News directly to your inbox"
    receiver:= "ronald.rivera@gmail.com"
    r := NewRequest([]string{receiver}, subject)
    r.Send("templates/template.html", map[string]string{"username": "Conor"})
}
