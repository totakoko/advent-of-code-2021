package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	input := strings.Split(string(content), "\n")
	input = input[:len(input)-1]

	fmt.Println("# Part 1", part1(input[0]))
	fmt.Println("# Part 2", part2(input[0]))
}

func part1(input string) int {
	binaryString := convertToBinaryString(input)
	packet, _ := NewPacketDecoder(binaryString).decodeNextPacket()
	return sumPacketsVersion(packet)
}

type PacketDecoder struct {
	binaryData string
	readIndex  int
}

type Packet struct {
	Version    int
	TypeID     int
	SubPackets []Packet
}

func NewPacketDecoder(binaryString string) *PacketDecoder {
	return &PacketDecoder{
		binaryData: binaryString,
		readIndex:  0,
	}
}

func (decoder *PacketDecoder) readNextBits(nbBits int) string {
	bits := decoder.binaryData[decoder.readIndex : decoder.readIndex+nbBits]
	decoder.readIndex += nbBits
	return bits
}

func (decoder *PacketDecoder) readNextNumber(nbBits int) int {
	bits := decoder.readNextBits(nbBits)
	return convertBinaryToNumber(bits)
}

func (decoder *PacketDecoder) getBytesRead() int {
	return decoder.readIndex
}

func (decoder *PacketDecoder) decodePacketsNBits(length int) []Packet {
	totalBytesRead := 0
	packets := []Packet{}
	for totalBytesRead < length {
		packet, bytesRead := decoder.decodeNextPacket()
		packets = append(packets, packet)
		totalBytesRead += bytesRead
	}
	return packets
}

func (decoder *PacketDecoder) decodePacketsN(count int) []Packet {
	packets := []Packet{}
	for len(packets) < count {
		packet, _ := decoder.decodeNextPacket()
		packets = append(packets, packet)
	}
	return packets
}

func (decoder *PacketDecoder) decodeNextPacket() (Packet, int) {
	initialReadIndex := decoder.readIndex
	version := decoder.readNextNumber(3)
	typeID := decoder.readNextNumber(3)

	packet := Packet{
		Version: version,
		TypeID:  typeID,
	}
	switch typeID {
	case 4: // literal value
		var bits string
		for decoder.readNextNumber(1) == 1 {
			bits += decoder.readNextBits(4)
		}
		bits += decoder.readNextBits(4)

	default: // operator
		lengthTypeID := decoder.readNextNumber(1)
		if lengthTypeID == 0 {
			subPacketsTotalLength := decoder.readNextNumber(15)
			packet.SubPackets = decoder.decodePacketsNBits(subPacketsTotalLength)
		} else {
			subPacketsCount := decoder.readNextNumber(11)
			packet.SubPackets = decoder.decodePacketsN(subPacketsCount)
		}
	}
	return packet, decoder.getBytesRead() - initialReadIndex
}

func convertToBinaryString(hexString string) string {
	sb := strings.Builder{}
	for _, c := range hexString {
		number, err := strconv.ParseInt(string(c), 16, 64)
		if err != nil {
			panic(err)
		}
		sb.WriteString(fmt.Sprintf("%04b", number))
	}
	return sb.String()
}

func part2(input string) int {
	return 0
}

func convertBinaryToNumber(binary string) int {
	number, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(number)
}

func sumPacketsVersion(packet Packet) int {
	sum := packet.Version
	for _, subPacket := range packet.SubPackets {
		sum += sumPacketsVersion(subPacket)
	}
	return sum
}
