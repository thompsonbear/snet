package main

import (
    "fmt"
    "flag"
    "regexp"
    "math"
    "math/bits"
    "strings" 
    "strconv"
    "text/tabwriter"
    "os"
)

type Net4 struct {
    addr [4]byte
    mask [4]byte
}    

func (n Net4) printNetList(){
    if n.maskValid(){
        w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
        fmt.Fprintln(w, "NETWORK\tUSEABLE\tBROADCAST\tMASK\tCIDR\t")    
        fmt.Fprintf(w, "%v\t%v\t%v\t%v\t/%v\t\n", n.networkStr(), n.useableStr(), n.broadcastStr(), n.maskStr(), n.maskBits())
        w.Flush()
    } else {
        fmt.Println("Error: Subnet Mask is not valid.")
    }
}

func (n Net4) network() [4]byte{
    var na [4]byte
    for i := 0; i < 4; i++{
        if n.mask[i] == 255{
            na[i] = n.addr[i]
        } else if n.mask[i] == 0{
            na[i] = 0
        } else{
            step := 0 - n.mask[i] 
            for (na[i] + step) < n.addr[i]{
                na[i] += step
            }
        }
    }
    return na
}

func (n Net4) networkStr() string{
    na := n.network()
    return fmt.Sprintf("%v.%v.%v.%v", na[0], na[1], na[2], na[3])
}

func (n Net4) useableStr() string{
    na := n.network() 
    bc := n.broadcast()
    var u [4]string

    if na[3] + 1 < bc[3] - 1{
        for i := 0; i < 4; i++{
            if na[i] == bc[i]{
                u[i] = fmt.Sprint(na[i])
            } else if i == 3 {
                u[i] = fmt.Sprintf("%v-%v", na[i]+1, bc[i]-1)
            } else {
                u[i] = fmt.Sprintf("%v-%v", na[i], bc[i])
            }
        }
        return fmt.Sprintf("%v.%v.%v.%v", u[0], u[1], u[2], u[3])
    } else {
        return "N/A"
    }
}

func (n Net4) broadcast() [4]byte{
    var bc [4]byte
    for i := 0; i < 4; i++{
        if n.mask[i] == 255{
            bc[i] = n.addr[i]
        } else if n.mask[i] == 0{
            bc[i] = 255
        } else {
            step := 0 - n.mask[i] 
            for bc[i] < n.addr[i]{
                bc[i] += step
            }
            bc[i]--
        }
    }
    return bc
}

func (n Net4) broadcastStr() string{
    bc := n.broadcast()
    return fmt.Sprintf("%v.%v.%v.%v", bc[0], bc[1], bc[2], bc[3])
}

func (n Net4) maskStr() string{
    return fmt.Sprintf("%v.%v.%v.%v", n.mask[0], n.mask[1], n.mask[2], n.mask[3])
}

func (n Net4) maskBits() byte{
    var b byte
    for i := 0; i < 4; i++{
        b += byte(bits.OnesCount8(n.mask[i]))  
    }
    return b
}

func (n Net4) maskValid() bool{
    str := fmt.Sprintf("%08b%08b%08b%08b", n.mask[0], n.mask[1], n.mask[2], n.mask[3])
    for i := 1; i < 32; i++{
        if str[i] > str[i-1]{ return false }
    } 
    return true
}

// Constructors

func newNet4(a string, m string) Net4{
    var mask [4]byte
    var addr [4]byte    

    addrs := strings.Split(a, ".")
    masks := strings.Split(m, ".")

    for i := 0; i < 4; i++{
        x, _ := strconv.Atoi(addrs[i])
        y, _ := strconv.Atoi(masks[i])
        addr[i] = byte(x)
        mask[i] = byte(y)
    } 
    return Net4{addr, mask}
}

func newNet4CIDR(c string) Net4{
    var mask [4]byte
    var addr [4]byte    

    s := strings.Split(c,"/")

    ip := s[0]
    mb, _ := strconv.Atoi(s[1])

    ips := strings.Split(ip, ".")

    for i := 0; i < 4; i++{
        x, _ := strconv.Atoi(ips[i])
        addr[i] = byte(x)
        if mb >= 8{
            mask[i] = 255   
            mb -= 8
        } else {
            bits := 8
            for mb >= 0{
                mask[i] += byte(math.Pow(2, float64(bits)))
                mb--
                bits--
            }
        }
    } 
    return Net4{addr, mask}
}

// Tests if an input string is in the Address + CIDR format 0-255.0-255.0-255.0-255/0-32
func isCIDRFormat(c string) bool{
    regCIDRMatch := fmt.Sprintf("^%[1]v.%[1]v.%[1]v.%[1]v/%[2]v$", "(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)" , "([0-9]|[1-2][0-9]|3[0-2])")
    matched, _ := regexp.MatchString(regCIDRMatch, c) 
    return matched
}

// Tests if an address or mask string is in the ipv4 format 0-255.0-255.0-255.0-255
func isIPv4Format(a string) bool{
    regIPMatch := fmt.Sprintf("^%[1]v.%[1]v.%[1]v.%[1]v$", "(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)")
    matched, _ := regexp.MatchString(regIPMatch, a) 
    return matched
}

func main(){

    colorWarn := "\033[33m" //ANSI Yellow
    colorError := "\033[31m" //ANSI Red
    colorReset := "\033[0m" //Resets ANSI color

    cidrString := flag.String("c", "", "Specify a Network in Address w/ CIDR Format: x.x.x.x/x")
    netAddrString := flag.String("a", "", "Specify an IP address on the network (Use in conjunction with -m) - Format: x.x.x.x")
    snetMaskString := flag.String("m", "", "Specify the Subnet Mask for the network (Use in conjunction with -a) - Format: x.x.x.x")

    flag.Parse()

    if *cidrString != ""{
        if *netAddrString != "" || *snetMaskString != ""{
            fmt.Println(colorWarn + "Warning: Addresses (-a) and Masks (-m) are ignored when you specify a CIDR (-c)" + colorReset)
        }
        if isCIDRFormat(*cidrString){
            net := newNet4CIDR(*cidrString)
            net.printNetList()
        } else{
            fmt.Println(colorError + "Error: Address with CIDR (-c) is not in the format x.x.x.x/x or is out of range." + colorReset)
        }
    } else if *netAddrString != "" && *snetMaskString != ""{
        if isIPv4Format(*netAddrString) && isIPv4Format(*snetMaskString){
            net := newNet4(*netAddrString, *snetMaskString)
            net.printNetList()
        } else {
            fmt.Println(colorError + "Error: Network Address or Subnet Mask is not in the format x.x.x.x or is out of range" + colorReset)
        }
    } else if (*netAddrString != "") != (*snetMaskString != ""){
        fmt.Println(colorError + "Error: Subnet Mask (-m) must be specified with a Network Address (-a)" + colorReset)
    } else{
        fmt.Println(colorError + "Error: The snet command requires flag arguments." + colorReset)
        flag.PrintDefaults()
    }
}


