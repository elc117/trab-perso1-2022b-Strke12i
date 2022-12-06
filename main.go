package main

import (
	"fmt"
	"trab-perso1-2022b-Strke12i/PortScanner"
//    "trab-perso1-2022b-Strke12i/DNSLookup"
 //   "github.com/gofiber/fiber/v2"
 //   "strconv"
)

func main() {
    open := PortScanner.OpenPorts("tcp", "scanme.nmap.org", 0, 80)
    fmt.Println(open)

    /*
    app := fiber.New()

    app.Get("scan/:protocol/:hostname/:r1/*", func (c *fiber.Ctx) error{
        if c.Params("hostname") == "" || c.Params("protocol") == "" || c.Params("r1") == "" {
            return c.Status(404).JSON(&fiber.Map{
                "error": "Erro nos parametros",
            })
        }

        var r1 int
        var r2 int

        if c.Params("*") == "" {
            r1 = 0
            r2, _ = strconv.Atoi(c.Params("r1"))
        } else {
            r1, _ = strconv.Atoi(c.Params("r1"))
            r2, _ = strconv.Atoi(c.Params("*"))

            if r1 > 65432 || r1 < 0 || r2 > 65432 || r2 < 0 || (r2 - r1) < 0 {
                return c.Status(401).JSON(&fiber.Map{
                    "error": "Erro no range das portas",
                })
            }
        }
        
        //PortScanner.OpenPorts(protocol,hostname,r1,r2)
        return c.SendString("Test" + strconv.Itoa(r1 + r2))
    })

    app.Listen(":8000")
    */

}
