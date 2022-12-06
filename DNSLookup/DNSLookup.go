package DNSLookup

import ( 
    "net"
)

func LookIp(addr string) ([]string, error){
    hosts, err := net.LookupAddr(addr)
    if err != nil {
        return nil, err
    }

    return hosts, nil
}

func LookHostname(hostname string) ([]string, error){
    ip, err := net.LookupHost(hostname)
    if err != nil {
        return nil, err
    }

    return ip, nil
}





