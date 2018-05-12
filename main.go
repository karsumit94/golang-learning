package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"errors"
	//"os/exec"
	"strconv"
	"time"
	"github.com/sparrc/go-ping"
)
func Cors(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=ascii")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")
          w.Write([]byte("Hello, World!"))
}

type Block struct {
    Try     func()
    Catch   func(Exception)
    Finally func()
}
type Exception interface{}
 
func Throw(up Exception) {
    panic(up)
}

func (tcf Block) Do() {
    if tcf.Finally != nil {
 
        defer tcf.Finally()
    }
    if tcf.Catch != nil {
        defer func() {
            if r := recover(); r != nil {
                tcf.Catch(r)
            }
        }()
    }
    tcf.Try()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting service in port 8081")
	li, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handle(conn)
	}
}

// Get preferred outbound ip of this machine
func outboundIP() string {
	ip := "NA"
	Block{
		Try: func() {
			conn, err := net.Dial("udp", "8.8.8.8:80")
			if err != nil {
		        log.Fatal(err)
		    } else {
				defer conn.Close()
				localAddr := conn.LocalAddr().(*net.UDPAddr)
				ip = localAddr.IP.String()
			}
		},
		Catch: func(e Exception) {
			log.Fatal("Exception occured")
		},
	}.Do()
    return ip
}

func externalIPString() string {
	ip, err := externalIP()
	if err != nil {
        log.Fatal(err)
    }
	return ip
}
func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("Not connected to the network")
}


func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		//log.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	// request line
	m := strings.Fields(ln)[0] // method
	u := strings.Fields(ln)[1] // uri
	//log.Println("***METHOD", m)
	//log.Println("***URI", u)

	// multiplexer
	if m == "GET" && u == "/" {
		index(conn)
	}
	if m == "GET" && u == "/ping" {
		ping_web(conn)
	}
	if m == "GET" && u == "/date" {
		date(conn)
	}
}

func index(conn net.Conn) {
	externalIPString := externalIPString()
	body := externalIPString
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "Access-Control-Allow-Origin *")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}



func ping_web(conn net.Conn){
	// send two pings and send the ouput to STDOUT
	// body, err := exec.Command("ping", "-c", "2", "www.broadsoft.com").Output()
	// bodyString := "Error"
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	//     bodyString = string(body[:])
 //    }
	// log.Printf("\n%s\n", bodyString)
 // 	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	// fmt.Fprintf(conn, "Content-Length: %d\r\n", len(bodyString))
	// fmt.Fprint(conn, "Content-Type: text/html\r\n")
	// fmt.Fprint(conn, "\r\n")
	// fmt.Fprint(conn, bodyString)
	bodyString := ""
	pinger, err := ping.NewPinger("www.broadsoft.com")
	if err != nil {
	        panic(err)
	}
	pinger.Count = 5
	pinger.OnRecv = func(pkt *ping.Packet) {
	        fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
	                pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}
	pinger.OnFinish = func(stats *ping.Statistics) {
	        fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
	        fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
	                stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
	        fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
	                stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	        bodyString = "PacketLoss : " + strconv.FormatFloat(stats.PacketLoss, 'f', 2, 64) + "%"
	        fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprintf(conn, "Content-Length: %d\r\n", len(bodyString))
			fmt.Fprint(conn, "Content-Type: text/html\r\n")
			fmt.Fprint(conn, "\r\n")
			fmt.Fprint(conn, bodyString)
	}
	
	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	pinger.Run()
	
}


func date(conn net.Conn){
	current_time := time.Now().Local()
	current_time_string := current_time.Format("Monday, 02-Jan-06 15:04:05 MST")
	log.Printf("\n %s ", current_time_string)
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(current_time_string))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, current_time_string)
}

