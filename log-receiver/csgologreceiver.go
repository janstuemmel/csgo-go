package csgologreceiver

import (
	"bytes"
	"errors"
	"net"
	"regexp"
)

type (
	Receiver interface {
		Close() error
		Read() (Response, error)
	}

	receiver struct {
		addr net.UDPAddr
		conn *net.UDPConn
	}

	// Response holds the log message data
	Response struct {
		Secret string
		Token  string
		Line   string
		Addr   *net.UDPAddr
	}
)

// Close closes the udp connection
func (r receiver) Close() error {
	return r.conn.Close()
}

// Read reads bytes from the udp connection
// and returns a Response struct or error
func (r receiver) Read() (Response, error) {

	p := make([]byte, 1024)
	_, remoteaddr, err := r.conn.ReadFromUDP(p)

	// remove zero bytes
	p = bytes.Trim(p, "\x00")

	if err != nil {
		return Response{}, err
	}

	secret, token, line, err := ParseHeader(p)

	return Response{Secret: secret, Token: token, Line: line, Addr: remoteaddr}, nil
}

// New initializes a udp connection to ip:port
func New(ip string, port int) (Receiver, error) {

	addr := net.UDPAddr{Port: port, IP: net.ParseIP(ip)}

	conn, err := net.ListenUDP("udp", &addr)

	if err != nil {
		return nil, err
	}

	return receiver{addr: addr, conn: conn}, nil
}

/*
	RegexNoSecretNoToken regular expression for parsing the
	log header if there's no secret nor a token set:

	Example: ����RL ...
*/
var RegexNoSecretNoToken = regexp.MustCompile(`.{4}R(L .*)`)

/*
	RegexSecretToken regular expression for parsing the
	log header if there's a secret and a token set:

	Example: ����S2fooTB8032A3B450FB7A6 ...
*/
var RegexSecretToken = regexp.MustCompile(`.{4}S(\w+)T(\w+) (L .*)`)

/*
	RegexSecretToken regular expression for parsing the
	log header if there's a secret but no token set:

	Example: ����S2fooL ...
*/
var RegexSecretNoToken = regexp.MustCompile(`.{4}S(\w+)(L .*)`)

/*
	RegexSecretToken regular expression for parsing the
	log header if there's no secret but a token set:

	Example: ����RTB8032A3B450FB7A6 L
*/
var RegexNoSecretToken = regexp.MustCompile(`.{4}RT(\w+) (L .*)`)

// ErrorNoMatch throw when log message header not valid
var ErrorNoMatch = errors.New("no match")

/*
	ParseHeader parses the header of a log
	message and returns: secret, token, line, error
*/
func ParseHeader(line []byte) (string, string, string, error) {

	// go doesn't support negative lookahead expression,
	// so parse 4 possiblities for logline header
	if result := RegexNoSecretNoToken.FindStringSubmatch(string(line)); result != nil {
		return "", "", result[1], nil
	}

	if result := RegexSecretToken.FindStringSubmatch(string(line)); result != nil {
		return result[1], result[2], result[3], nil
	}

	if result := RegexSecretNoToken.FindStringSubmatch(string(line)); result != nil {
		return result[1], "", result[2], nil
	}

	if result := RegexNoSecretToken.FindStringSubmatch(string(line)); result != nil {
		return "", result[1], result[2], nil
	}

	return "", "", "", ErrorNoMatch
}
