package checkmass

import (
	//"errors"
	"fmt"
	"net"
	"os"
	//"path/filepath"
	//"regexp"
	//"strconv"
	//"strings"
	//"github.com/Icinga/icingadb/configobject"
	//"github.com/Icinga/icingadb/connection"
	//"github.com/Icinga/icingadb/utils"
	"github.com/jessevdk/go-flags"
	//"github.com/mackerelio/checkers"
)

func isIPv6(host string) bool {
	addr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		os.Exit(1)
	}
	if ip4 := addr.IP.To4(); len(ip4) != net.IPv4len {
		return true
	}
	return false
}

var opts struct {
	Host     string `long:"host" short:"H" description:"check target IP Address"`
	Count    int    `long:"count" short:"n" default:"1" description:"sending (and receiving) count ping packets"`
	WaitTime int    `long:"wait-time" short:"w" default:"1000" description:"wait time, Max RTT(ms)"`
	Output   int  `long:"output" short:"o" description:"Type of output"`
}

type example interface{

}
//
//const (
//	OK Status = iota
//	WARNING
//	CRITICAL
//	UNKNOWN
//)

const (
	PERFORMANCPERFEDATA = 0
	OUTPUT = 1
	LONGOUTPUT = 2
)

func run(args []string) (string, string){                     //*checkers.Checker
	_, err := flags.ParseArgs(&opts, args)
	if err != nil {
		os.Exit(1)
	}

	var msg string
	chkSt := "OK"     //checkers.OK

	switch opts.Output {
	case PERFORMANCEDATA:
		msg = fmt.Sprintf("The performance data is dummy performance")
	case OUTPUT:
		msg = fmt.Sprintf("The output is dummy output")
	case LONGOUTPUT:
		msg = fmt.Sprintf("The long output is dummy output")
	default:
		msg = fmt.Sprintf("The default output is dummy output")
	}
	return chkSt, msg  //checkers.NewChecker(chkSt, msg)
}

// Do the plugin
func Do() {
	chkSt, msg := run(os.Args[1:])
	//ckr.Name = "Mass-Performance-Data"
	//ckr.Exit()
	fmt.Printf(chkSt, msg)
}
