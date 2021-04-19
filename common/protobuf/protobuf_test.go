package protobuf

//
//import (
//	pd "boot/common/socketProtoBuff"
//	"fmt"
//	"github.com/golang/protobuf/proto"
//	"net"
//	"strconv"
//	"testing"
//	"time"
//)
//
//func TestProtocolBuffer(t *testing.T) {
//	// MessageEnvelope是models.pb.go的结构体
//	oldData := &pd.MessageEnvelope{
//		TargetId: 1,
//		ID:       "1",
//		Type:     "2",
//		Payload:  []byte("ITDragon protobuf"),
//	}
//
//	data, err := proto.Marshal(oldData)
//	if err != nil {
//		fmt.Println("marshal error: ", err.Error())
//	}
//	fmt.Println("marshal data : ", data)
//
//	newData := &pd.MessageEnvelope{}
//	err = proto.Unmarshal(data, newData)
//	if err != nil {
//		fmt.Println("unmarshal err:", err)
//	}
//	fmt.Println("unmarshal data : ", newData)
//
//}
//
////1.TCP Server端
//func TestTcpServer(t *testing.T) {
//	// 为突出重点，忽略err错误判断
//	addr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:9000")
//	listener, _ := net.ListenTCP("tcp4", addr)
//	for {
//		conn, _ := listener.AcceptTCP()
//		go func() {
//			for {
//				buf := make([]byte, 512)
//				_, _ = conn.Read(buf)
//				newData := &pd.MessageEnvelope{}
//				_ = proto.Unmarshal(buf, newData)
//				fmt.Println("server receive : ", newData)
//			}
//		}()
//	}
//}
//
////2.TCP Client端
//func TestTcpClient(t *testing.T) {
//	// 为突出重点，忽略err错误判断
//	connection, _ := net.Dial("tcp", "127.0.0.1:9000")
//	var targetID int32 = 1
//	for {
//		oldData := &pd.MessageEnvelope{
//			TargetId: targetID,
//			ID:       strconv.Itoa(int(targetID)),
//			Type:     "2",
//			Payload:  []byte(fmt.Sprintf("ITDragon protoBuf-%d", targetID)),
//		}
//		data, _ := proto.Marshal(oldData)
//		_, _ = connection.Write(data)
//		fmt.Println("client send : ", data)
//		time.Sleep(2 * time.Second)
//		targetID++
//	}
//}
