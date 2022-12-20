package main

import (
    "github.com/gofiber/fiber/v2"
    "trab-perso1-2022b-Strke12i/Routes"
    "trab-perso1-2022b-Strke12i/CommandLine"
    "os"
)

func main() {
    
    if len(os.Args) > 1 {
        CommandLine.ListPortsCMD()
        return
    }


    app := fiber.New()

    app.Get("scan/:protocol/:hostname/:r1/*", Routes.Handler );

    app.Listen(":8000")
}
