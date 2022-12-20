package CommandLine

import(
    "os"
    "trab-perso1-2022b-Strke12i/PortScanner"
    "fmt"
    "strconv"
)

func ListPortsCMD(){
    if len(os.Args) < 4 || len(os.Args) > 5 {
        fmt.Println("Argumentos Invalidos!")
        fmt.Println("Formato: ./scan protocol ip range-portas-1 range-portas-2(opcional)")
        return
    }

    var openPorts []int
    protocol := os.Args[1]
    ip := os.Args[2]
    p1, _ := strconv.Atoi(os.Args[3])

    if len(os.Args) == 4 {
        openPorts = PortScanner.OpenPorts(protocol, ip, 0, p1)
    }else{
        p2, _ := strconv.Atoi(os.Args[4])
        openPorts = PortScanner.OpenPorts(protocol, ip, p1, p2)
    }

    if len(openPorts) == 0 {
        fmt.Println("Nenhuma Porta Aberta em",ip)
        return
    }

    fmt.Printf("Host %s \n", ip) 
    for _, i := range openPorts {
        fmt.Printf("Porta %d aberta\n",i);
    }

}
