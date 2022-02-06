package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/Terry-Mao/goim/api/protocol"
	"github.com/Terry-Mao/goim/pkg/bufio"
)

func main() {
	msg := "Hello, 世界"
	p := &protocol.Proto{Ver: 1, Op: 2, Seq: 3, Body: []byte(msg)}
	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)
	_ = p.WriteTCP(wr)
	_ = wr.Flush()
	fmt.Printf("Data: %#v\n", buf.String())

	//rr := bufio.NewReader(&buf)
	//p2 := &protocol.Proto{}
	//_ = p2.ReadTCP(rr)
	//fmt.Printf("Body: %#v\n", string(p2.Body))

	decode(buf.Bytes())
}

func decode(buf []byte) {
	if len(buf) <= 16 {
		return
	}

	packLen := binary.BigEndian.Uint32(buf[:4])
	headerLen := binary.BigEndian.Uint16(buf[4:6])
	ver := binary.BigEndian.Uint16(buf[6:8])
	op := binary.BigEndian.Uint32(buf[8:12])
	seq := binary.BigEndian.Uint32(buf[12:16])
	body := string(buf[16:])

	fmt.Printf("Package Length: %v\n", packLen)
	fmt.Printf("Header Length: %v\n", headerLen)
	fmt.Printf("Protocol Version: %v\n", ver)
	fmt.Printf("Operation: %v\n", op)
	fmt.Printf("Sequence Id: %v\n", seq)
	fmt.Printf("Body: %#v\n", body)
}
