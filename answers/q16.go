package answers

import (
	"fmt"
	"math"
)

func Day16() []int {
	data := ReadInputAsStr(16)
	hexa := ""
	lookup := map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}
	for _, dataRune := range data[0] {
		hexa += lookup[dataRune]
	}

	packet, hexa := ProcessString(hexa)
	return []int{
		packet.VersionSum(),
		packet.CalculateValue(),
	}
}

type Packet struct {
	Version int
	TypeID  int
	Value   int
	Packets []Packet
}

func (p Packet) Print() {
	fmt.Printf("Version: %d, TypeID %d, Value %d, Packets Length %d \n", p.Version, p.TypeID, p.Value, len(p.Packets))
	for _, packet := range p.Packets {
		packet.Print()
	}
}

func (p Packet) VersionSum() int {
	value := p.Version
	for _, packet := range p.Packets {
		value += packet.VersionSum()
	}
	return value
}

func (p Packet) CalculateValue() int {

	if p.TypeID == 0 {
		// Sum
		sum := 0
		for _, subp := range p.Packets {
			sum += subp.CalculateValue()
		}
		return sum
	} else if p.TypeID == 1 {
		// Product
		prod := 1
		for _, subp := range p.Packets {
			prod = prod * subp.CalculateValue()
		}
		return prod
	} else if p.TypeID == 2 {
		// Minimum
		minVal := 9999999
		for _, subp := range p.Packets {
			if subp.CalculateValue() < minVal {
				minVal = subp.CalculateValue()
			}
		}
		return minVal
	} else if p.TypeID == 3 {
		maxVal := -9999999
		for _, subp := range p.Packets {
			if subp.CalculateValue() > maxVal {
				maxVal = subp.CalculateValue()
			}
		}
		return maxVal
	} else if p.TypeID == 4 {
		return p.Value
	} else if p.TypeID == 5 {
		if p.Packets[0].CalculateValue() > p.Packets[1].CalculateValue() {
			return 1
		}
		return 0
	} else if p.TypeID == 6 {
		if p.Packets[0].CalculateValue() < p.Packets[1].CalculateValue() {
			return 1
		}
		return 0
	} else if p.TypeID == 7 {
		if p.Packets[0].CalculateValue() == p.Packets[1].CalculateValue() {
			return 1
		}
		return 0
	}
	panic("Shouldn't get here")
}

func BitsToNum(data string) int {
	// Takes input like 000101, and converts to number -> 5
	num := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '1' {
			num += int(math.Pow(2, float64(len(data)-i-1)))
		}
	}
	return num
}
func ReadLiteralValue(data string) (string, string) {
	// Returns the hexadecimal string, and the popped data string
	if data[0] == '1' {
		solution := data[1:5]
		value, data := ReadLiteralValue(data[5:])
		solution += value
		return solution, data
	}
	return data[1:5], data[5:]
}

func ProcessString(hexa string) (Packet, string) {
	// Returns a packet, and all the unprocessed string
	version, typeid, hexa := BitsToNum(hexa[0:3]), BitsToNum(hexa[3:6]), hexa[6:]
	packet := Packet{
		Version: version,
		TypeID:  typeid,
	}

	if typeid == 4 {
		value, hexa := ReadLiteralValue(hexa)
		packet.Value = BitsToNum(value)
		return packet, hexa
	} else {
		lengthTypeId, hexa := hexa[0], hexa[1:]

		if lengthTypeId == '0' {
			numberOfBits, hexa := BitsToNum(hexa[0:15]), hexa[15:]

			var subpacket Packet
			subStr, hexa := hexa[0:numberOfBits], hexa[numberOfBits:]
			for len(subStr) > 0 {
				subpacket, subStr = ProcessString(subStr)
				packet.Packets = append(packet.Packets, subpacket)
			}
			return packet, hexa
		} else {
			numberOfSubPackets, hexa := BitsToNum(hexa[:11]), hexa[11:]
			packetNum := 0
			var subpacket Packet
			for packetNum < numberOfSubPackets {
				subpacket, hexa = ProcessString(hexa)
				packet.Packets = append(packet.Packets, subpacket)
				packetNum++
			}
			return packet, hexa
		}
	}
}
