package codec

import "io"

type Header struct {
	ServiceMethod string // 服务名和方法名, 通常与 Go 语言中的结构体和方法相映射
	Seq           uint64 // 某个请求的 ID, 用来区分不同的请求。
	Error         string // 服务端如果如果发生错误, 将错误信息置于 Error 中
}

// Codec 抽象出对消息体进行编解码的接口, 便于实现不同的 Codec 实例
type Codec interface {
	io.Closer
	ReadHeader(*Header) error
	ReadBody(interface{}) error
	Write(*Header, interface{}) error
}

type NewCodecFunc func(closer io.ReadWriteCloser) Codec

type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // not implemented
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	// 和工厂模式类似, 与工厂模式不同的是, 返回的是构造函数, 而非实例
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
