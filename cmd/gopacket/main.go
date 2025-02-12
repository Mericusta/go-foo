package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"slices"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// 参考 https://cloud.tencent.com/developer/article/2334505

func main() {
	// 罗列所有网络设备
	devices, err := pcap.FindAllDevs()
	if err != nil {
		panic(err)
	}

	fmt.Println("devices:")
	for _, device := range devices {
		if device.Name == "en0" {
			fmt.Println("name:", device.Name)
			fmt.Println("flags:", device.Flags)
			fmt.Println("desc:", device.Description)
			fmt.Println("address:")
			for _, address := range device.Addresses {
				fmt.Println("- IP address:", address.IP)
				fmt.Println("- Subnet mask:", address.Netmask)
			}
			fmt.Println()
		}
	}

	// 打开某一个网络设备，比如 en0
	var (
		openDeviceName               = "en0" // 目标网络设备名称
		snapshotLen    int32         = 1024  // 单次抓包长度
		promiscuous    bool          = false // 是否将网口设置为混杂模式，即是否接收目的地址不为本机的包
		timeout        time.Duration = -1    // 抓到包时的返回等待时间，负数不等待
		handle         *pcap.Handle          // 设备句柄
	)

	// 打开某一网络设备，unix 需要修改权限，建议修改临时权限重启后失效（sudo chmod o+r /dev/bpf*）
	handle, err = pcap.OpenLive(openDeviceName, snapshotLen, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		printPacketInfo(packet)
	}
}

func isLayerT[T gopacket.LayerType](t T, packet gopacket.Packet) bool {
	return packet.Layer(gopacket.LayerType(t)) != nil
}

type FilterMode int

const (
	FILTER_MODE_WHITE_LIST = iota + 1
	FILTER_MODE_BLOCK_LIST = iota + 1
)

type packetFilter struct {
	ipv4SrcIP []net.IP
	ipv4DstIP []net.IP

	tcpSrcPort []layers.TCPPort
	tcpDstPort []layers.TCPPort
}

func (f *packetFilter) Filter(packet gopacket.Packet) bool {
	ipv4Layer := packet.Layer(layers.LayerTypeIPv4)
	if ipv4Layer == nil {
		return true
	}
	ipv4Packet, ok := ipv4Layer.(*layers.IPv4)
	if !ok {
		return true
	}
	if len(f.ipv4SrcIP) > 0 && !slices.ContainsFunc(f.ipv4SrcIP, func(filterSrcIP net.IP) bool { return filterSrcIP.Equal(ipv4Packet.SrcIP) }) {
		return true
	}
	if len(f.ipv4DstIP) > 0 && !slices.ContainsFunc(f.ipv4DstIP, func(filterDstIP net.IP) bool { return filterDstIP.Equal(ipv4Packet.DstIP) }) {
		return true
	}

	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer == nil {
		return true
	}
	tcpPacket, ok := tcpLayer.(*layers.TCP)
	if !ok {
		return true
	}
	if len(f.tcpSrcPort) > 0 && !slices.ContainsFunc(f.tcpSrcPort, func(filterSrcPort layers.TCPPort) bool { return filterSrcPort == tcpPacket.SrcPort }) {
		return true
	}
	if len(f.tcpDstPort) > 0 && !slices.ContainsFunc(f.tcpDstPort, func(filterDstPort layers.TCPPort) bool { return filterDstPort == tcpPacket.DstPort }) {
		return true
	}

	return false
}

var globalFilter *packetFilter = &packetFilter{
	ipv4SrcIP:  []net.IP{net.IPv4(18, 167, 163, 74)},
	tcpSrcPort: []layers.TCPPort{443},
}

func printPacketInfo(packet gopacket.Packet) {
	if globalFilter.Filter(packet) {
		return
	}

	if isLayerT(layers.LayerTypeTCP, packet) {
		ipv4Packet := packet.Layer(layers.LayerTypeIPv4).(*layers.IPv4)
		tcpPacket := packet.Layer(layers.LayerTypeTCP).(*layers.TCP)
		if tcpPacket.PSH {
			fmt.Println(time.Now().Format("20060102 15:04:05"))
			fmt.Printf("tcp packet %v:%v -> %v:%v \n", ipv4Packet.SrcIP, tcpPacket.SrcPort, ipv4Packet.DstIP, tcpPacket.DstPort)
			fmt.Printf("PSH %v ACK %v\n", tcpPacket.PSH, tcpPacket.ACK)

			// 拿到的是 TLS 加密之后的数据
			fmt.Printf("tcpPacket.BaseLayer.Payload %v\n", tcpPacket.BaseLayer.Payload)

			fin, opcode, mask, payloadLength, payload, err := parseWebSocketFrame(tcpPacket.BaseLayer.Payload)
			if err != nil {
				fmt.Printf("parseWebSocketFrame occurs error: %v\n", err)
			} else {
				fmt.Printf("FIN: %v, Opcode: %02x, Mask: %v, Payload Length: %d, Payload: %d hex: |%v|\n", fin, opcode, mask, payloadLength, payload, hex.EncodeToString(payload))

			}

			// wsPayload, err := parseWebSocketFrame(tcpPacket.BaseLayer.Payload)
			// if err != nil {
			// 	fmt.Printf("parseWebSocketFrame occurs error: %v\n", err)
			// } else {
			// 	fmt.Printf("tcp payload: %v\n", tcpPacket.BaseLayer.Payload)
			// 	fmt.Printf("ws payload: %v\n", wsPayload)
			// 	fmt.Printf("ws hex string: |%v|\n", hex.EncodeToString(wsPayload))
			// }
			// bytes := make([]byte, 0, len(tcpPacket.BaseLayer.Payload))
			// bytes = append(bytes, tcpPacket.BaseLayer.Payload[0:]...)
			// fmt.Printf("bytes: hex string |%v|\n", hex.EncodeToString(bytes))

			// if l := len(tcpPacket.BaseLayer.LayerPayload()); l > 0 {
			// 	for offset := 0; offset != l; offset++ {
			// 		bytes := make([]byte, 0, len(tcpPacket.BaseLayer.Payload))
			// 		bytes = append(bytes, tcpPacket.BaseLayer.Payload[offset:]...)
			// 		hexString := hex.EncodeToString(bytes)
			// 		if len(hexString) == 0 {
			// 			continue
			// 		}
			// 		fmt.Printf("bytes: hex string |%v|\n", hexString)
			// 		err := protobuffoo.UnmarshalUnknownStruct(bytes, false)
			// 		if err == nil {
			// 		}
			// 	}
			// }
			fmt.Println()
		}
	}

	// // 判断数据包是否为以太网数据包，可解析出源mac地址、目的mac地址、以太网类型（如ip类型）等
	// ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	// if ethernetLayer != nil {
	// 	fmt.Println("Ethernet layer detected.")
	// 	ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)
	// 	fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
	// 	fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
	// 	// Ethernet type is typically IPv4 but could be ARP or other
	// 	fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
	// 	fmt.Println()
	// }

	// // Let's see if the packet is IP (even though the ether type told us)
	// // 判断数据包是否为IP数据包，可解析出源ip、目的ip、协议号等
	// ipLayer := packet.Layer(layers.LayerTypeIPv4)
	// if ipLayer != nil {
	// 	fmt.Println("IPv4 layer detected.")
	// 	ip, _ := ipLayer.(*layers.IPv4)
	// 	// IP layer variables:
	// 	// Version (Either 4 or 6)
	// 	// IHL (IP Header Length in 32-bit words)
	// 	// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
	// 	// Checksum, SrcIP, DstIP
	// 	fmt.Printf("From %s to %s, Protocol: %v\n", ip.SrcIP, ip.DstIP, ip.Protocol)
	// 	fmt.Println()
	// }

	// // Let's see if the packet is TCP
	// // 判断数据包是否为TCP数据包，可解析源端口、目的端口、seq序列号、tcp标志位等
	// tcpLayer := packet.Layer(layers.LayerTypeTCP)
	// if tcpLayer != nil {
	// 	fmt.Println("TCP layer detected.")
	// 	tcp, _ := tcpLayer.(*layers.TCP)
	// 	// TCP layer variables:
	// 	// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
	// 	// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS
	// 	fmt.Printf("From port %d to %d, Sequence number: %v\n", tcp.SrcPort, tcp.DstPort, tcp.Seq)
	// 	fmt.Println()
	// }

	// // // Iterate over all layers, printing out each layer type
	// // fmt.Println("All packet layers:")
	// // for _, layer := range packet.Layers() {
	// // 	fmt.Println("- ", layer.LayerType())
	// // }

	// ///.......................................................
	// // Check for errors
	// // 判断layer是否存在错误
	// if err := packet.ErrorLayer(); err != nil {
	// 	fmt.Println("Error decoding some part of the packet:", err)
	// }
}

// // parseWebSocketFrame 解析 WebSocket 帧
// func parseWebSocketFrame(payload []byte) ([]byte, error) {
// 	if len(payload) < 2 {
// 		return nil, fmt.Errorf("payload is too short")
// 	}

// 	// FIN 和 Opcode
// 	fin := payload[0] >> 7
// 	opcode := payload[0] & 0x0F

// 	// Mask 和 Payload Length
// 	mask := payload[1] >> 7
// 	payloadLength := payload[1] & 0x7F

// 	// 打印帧信息
// 	fmt.Printf("FIN: %d, Opcode: %d, Mask: %d, Payload Length: %d\n", fin, opcode, mask, payloadLength)

// 	// 如果 Mask 位为 1，跳过 4 字节的 Masking Key
// 	var maskingKey []byte
// 	if mask == 1 {
// 		if len(payload) < 6 {
// 			return nil, fmt.Errorf("payload is too short")
// 		}
// 		maskingKey = payload[2:6]
// 		payload = payload[6:]
// 	} else {
// 		payload = payload[2:]
// 	}

// 	// 打印数据内容
// 	fmt.Printf("Masking Key: %x, Payload Data: %x\n", maskingKey, payload)

// 	return payload, nil
// }

// parseWebSocketFrame 解析WebSocket数据帧
func parseWebSocketFrame(frame []byte) (fin bool, opcode byte, mask bool, payloadLength int, payload []byte, err error) {
	// 检查帧长度是否足够
	if len(frame) < 2 {
		return false, 0, false, 0, nil, fmt.Errorf("frame too short")
	}

	// 解析第一个字节
	firstByte := frame[0]
	fin = (firstByte & 0x80) != 0
	opcode = firstByte & 0x0f

	// 解析第二个字节
	secondByte := frame[1]
	mask = (secondByte & 0x80) != 0
	payloadLength = int(secondByte & 0x7f)

	// 根据Payload Length确定实际长度
	var maskKey []byte
	switch payloadLength {
	case 126:
		if len(frame) < 4 {
			return false, 0, false, 0, nil, fmt.Errorf("frame too short for extended payload length")
		}
		payloadLength = int(binary.BigEndian.Uint16(frame[2:4]))
		maskKey = frame[4:8]
	case 127:
		if len(frame) < 10 {
			return false, 0, false, 0, nil, fmt.Errorf("frame too short for extended payload length")
		}
		payloadLength = int(binary.BigEndian.Uint64(frame[2:10]))
		maskKey = frame[10:14]
	default:
		maskKey = frame[2:6]
	}

	// 检查帧长度是否足够
	if len(frame) < payloadLength+maskKeyOffset(payloadLength) {
		return false, 0, false, 0, nil, fmt.Errorf("frame too short for payload")
	}

	// 提取Payload Data
	payload = frame[maskKeyOffset(payloadLength) : maskKeyOffset(payloadLength)+payloadLength]

	// 如果有掩码，解码Payload Data
	if mask {
		for i := 0; i < payloadLength; i++ {
			payload[i] ^= maskKey[i%4]
		}
	}

	return fin, opcode, mask, payloadLength, payload, nil
}

// maskKeyOffset 返回掩码键的偏移量
func maskKeyOffset(payloadLength int) int {
	switch payloadLength {
	case 126:
		return 4
	case 127:
		return 10
	default:
		return 2
	}
}

// func main() {
// 	// 示例数据
// 	frame := []byte{0x81, 0x85, 0x37, 0xfa, 0x21, 0x3d, 0x7f, 0x9f, 0x4d, 0x51, 0x58}

// 	fin, opcode, mask, payloadLength, payload, err := parseWebSocketFrame(frame)
// 	if err != nil {
// 		fmt.Printf("Error parsing frame: %v\n", err)
// 		return
// 	}

// 	fmt.Printf("FIN: %v, Opcode: %02x, Mask: %v, Payload Length: %d, Payload: %s\n", fin, opcode, mask, payloadLength, payload)
// }
