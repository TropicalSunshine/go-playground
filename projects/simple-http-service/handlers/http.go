package handlers

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Handler http0.9 connection and response with text html
func HandleHttp9(c net.Conn) {
	defer c.Close()

	r := bufio.NewReader(c)
	line, err := r.ReadString('\n')
	if err != nil {
		return // close on any read error per 0.9 simplicity[web:1][web:4]
	}

	// Parse: "GET /path"
	fields := strings.Fields(strings.TrimSpace(line))
	if len(fields) < 2 || strings.ToUpper(fields[0]) != "GET" { // only GET in 0.9[web:1][web:2][web:5][web:8]
		return
	}
	path := fields[1]
	if path == "/" {
		path = "/index.html"
	}

	// For demo just ignore path and always return same body.
	// HTTP/0.9 has no status line or headers, so write body only.[web:1][web:2][web:5][web:8]
	body := fmt.Sprintf("<html>\nYou requested: %s\n</html>\n", path)
	_, _ = c.Write([]byte(body))

}
