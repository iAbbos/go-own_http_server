package parser

import (
	"bufio"
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/internal/entity"
	"io"
	"log"
	"strings"
)

type Parser struct {
	reader *bufio.Reader
}

func NewParser(rd io.Reader) *Parser {
	return &Parser{reader: bufio.NewReader(rd)}
}

func (p *Parser) Parse() (entity.Request, error) {

	for {
		b, err := p.reader.ReadByte()
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				log.Fatal(err)
			}
		}
		fmt.Printf("byte: %v | string: %v\n", b, string(b))
	}

	method, target, version, err := p.readRequestLine(p.reader)
	if err != nil {
		fmt.Println("Error: ", err)
		return entity.Request{}, err
	}

	fmt.Println("Method: ", method)
	fmt.Println("Target: ", target)
	fmt.Println("Version: ", version)

	headers, err := p.readHeaders(p.reader)
	if err != nil {
		return entity.Request{}, err
	}

	fmt.Println("Headers: ", headers)

	return entity.Request{
		Method:  method,
		Target:  target,
		Version: version,
		Headers: headers,
		Body:    p.reader,
	}, nil
}

func (p *Parser) readRequestLine(reader *bufio.Reader) (method string, target string, httpVersion string, err error) {
	line, err := p.readLine(reader)
	if err != nil {
		return
	}
	requestLine := strings.Split(line, " ")
	if len(requestLine) != 3 {
		err = fmt.Errorf("invalid request line")
		return
	}
	method = strings.ToUpper(requestLine[0])
	target = requestLine[1]
	httpVersion = requestLine[2]
	return
}

func (p *Parser) readHeaders(reader *bufio.Reader) (headers map[string]string, err error) {
	headers = make(map[string]string)
	var line = ""
	for {
		line, err = p.readLine(reader)
		if err != nil {
			return
		}
		if line == "" {
			break
		}
		header := strings.SplitN(line, ":", 2)
		if len(header) != 2 {
			err = fmt.Errorf("invalid header")
			return
		}
		header[0] = strings.TrimSpace(header[0])
		header[1] = strings.TrimSpace(header[1])
		headers[header[0]] = header[1]
	}
	return
}

func (p *Parser) readLine(reader *bufio.Reader) (line string, err error) {
	var tmp []byte
	var isPrefix bool
	for {
		tmp, isPrefix, err = reader.ReadLine()
		if err != nil {
			return
		}
		line += string(tmp)
		if !isPrefix {
			break
		}
	}
	return
}
