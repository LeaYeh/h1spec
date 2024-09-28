package spec

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/LeaYeh/h1spec/config"
)

const (
	DefaultTimeout = 10
	DefaultMaxHeaderLength = 128
	DefaultMaxBodyLength = 1024
)

// Conn represent a basic connection.
type Conn struct {
	net.Conn
	Timeout  		time.Duration
	MaxHeaderLength int
	MaxBodyLength 	int
	Closed   		bool
	server 			bool
}

// Dial connects to the server based on configuration.
func Dial(c *config.Config) (*Conn, error) {
	var baseConn net.Conn
	var err error

	baseConn, err = net.DialTimeout("tcp", c.Addr(), c.Timeout)
	if err != nil {
		return nil, err
	}

	conn := &Conn{
		Conn: baseConn,
		Timeout: c.Timeout,
		MaxHeaderLength: c.MaxHeaderLen,
		MaxBodyLength: c.MaxBodyLen,
		server: false,
	}

	return conn, nil
}

func Accept(c *config.Config, baseConn net.Conn) (*Conn, error) {
	conn := &Conn{
		Conn: baseConn,
		Timeout: c.Timeout,
		MaxHeaderLength: c.MaxHeaderLen,
		MaxBodyLength: c.MaxBodyLen,
		server: true,
	}

	return conn, nil
}

type Request struct {
	RequestLine string
	Headers  	[]string
	Body     	[]byte
}

func (conn *Conn) ReadRequest() (*Request, error) {
	headers := make([]string, 0, 256)

	// Read request line
	requestLine, err := conn.ReadLine()
	if err != nil {
		return nil, err
	}

	// Read headers
	for {
		line, err := conn.ReadLine()
		if err != nil {
			return nil, err
		}

		// Break the loop if an empty line is encountered (end of headers)
		if line == "" {
            break
        }

		headers = append(headers, line)
	}

	// Read body
	body, err := io.ReadAll(conn)
	if err != nil && err != io.EOF {
		return nil, err
	}

	request := &Request{
		RequestLine: requestLine,
		Headers: headers,
		Body: body,
	}
	return request, nil
}

func (conn *Conn) Close() error {
	return conn.Conn.Close()
}

// Send sends a byte sequense. This function is used to send a raw
// data in tests.
func (conn *Conn) Send(payload []byte) error {
	_, err := conn.Write(payload)
	return err
}

func (conn *Conn) WriteHeaders(headers []HeaderField) error {
	for _, header := range headers {
		_, err := conn.Write([]byte(fmt.Sprintf("%s: %s\r\n", header.Name, header.Value)))
		if err != nil {
			return err
		}
	}
	_, err := conn.Write([]byte("\r\n"))
	if err != nil {
		return err
	}

	return nil
}


func (conn *Conn) WriteBytes(b []byte) error {
	_, err := conn.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func (conn *Conn) ReadByte() (byte, error) {
	var b [1]byte
	_, err := conn.Read(b[:])
	if err != nil {
		return 0, err
	}

	return b[0], nil
}

func (conn *Conn) ReadLine() (string, error) {
	var buffer bytes.Buffer
	for {
		b, err := conn.ReadByte()
		if err != nil {
			return "", err
		}

		if b == '\n' {
			break
		}

		buffer.WriteByte(b)
	}

	return buffer.String(), nil
}

func (conn *Conn) readBytes(size int) ([]byte, error) {
	var remain = size
	buffer := make([]byte, 0, size)

	for remain > 0 {
		tmp := make([]byte, remain)
		n, err := conn.Read(tmp)
		if err != nil {
			return nil, err
		}

		buffer = append(buffer, tmp[:n]...)
		remain = remain - n
	}
	return buffer, nil
}
