package servive

import (
	"fmt"
	"io"
	"net"
)

type Proxy struct {
	laddr, raddr *net.TCPAddr
	lconn, rconn io.ReadWriteCloser
	closeSig     chan bool
	connClosed   bool
}

func NewProxy(lconn *net.TCPConn, laddr, raddr *net.TCPAddr) *Proxy {
	return &Proxy{
		lconn:      lconn,
		laddr:      laddr,
		raddr:      raddr,
		closeSig:   make(chan bool),
		connClosed: false,
	}
}

func (p *Proxy) Run() {
	defer p.lconn.Close()
	var err error
	p.rconn, err = net.DialTCP("tcp", nil, p.raddr)
	if err != nil {
		fmt.Println("Remote connection err: ", err.Error())
		return
	}
	defer p.rconn.Close()
	go p.pipe(p.lconn, p.rconn)
	go p.pipe(p.rconn, p.lconn)
	// waite close
	<-p.closeSig
}

func (p *Proxy) closeConn(msg string, err error) {
	if p.connClosed {
		return
	}
	if err != io.EOF {
		fmt.Println(msg, err.Error())
	}
	p.closeSig <- true
	p.connClosed = true
}

func (p *Proxy) pipe(src io.ReadWriter, dst io.ReadWriter) {
	//64k
	buff := make([]byte, 0xffff)
	for {
		n, err := src.Read(buff)
		if err != nil {
			p.closeConn("Read conn err: ", err)
			return
		}
		b := buff[:n]
		//write out result
		n, err = dst.Write(b)
		if err != nil {
			p.closeConn("Write conn err: ", err)
			return
		}
	}
}
