package logger

import (
	"context"
	"fmt"
	"time"
)

type Level int

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	DPanicLevel
	PanicLevel
	FatalLevel
)

type Log interface {
	// Msg 日志消息，最终执行
	Msg(msg string)
	Err(err error) Log
	Ctx(ctx context.Context) Log

	Binary(key string, val []byte) Log
	ByteString(key string, value []byte) Log

	Bool(key string, value bool) Log
	Boolp(key string, value *bool) Log
	Bools(key string, value []bool) Log

	Complex128(key string, value complex128) Log
	Complex128p(key string, value *complex128) Log
	Complex128s(key string, value []complex128) Log

	Complex64(key string, value complex64) Log
	Complex64p(key string, value *complex64) Log
	Complex64s(key string, value []complex64) Log

	Duration(key string, value time.Duration) Log
	Durationp(key string, value *time.Duration) Log
	Durations(key string, value []time.Duration) Log

	Float64(key string, value float64) Log
	Float64p(key string, value *float64) Log
	Float64s(key string, value []float64) Log

	Float32(key string, value float32) Log
	Float32p(key string, value *float32) Log
	Float32s(key string, value []float32) Log

	Int64(key string, value int64) Log
	Int64p(key string, value *int64) Log
	Int64s(key string, value []int64) Log

	Int32(key string, value int32) Log
	Int32p(key string, value *int32) Log
	Int32s(key string, value []int32) Log

	Int16(key string, value int16) Log
	Int16p(key string, value *int16) Log
	Int16s(key string, value []int16) Log

	Int8(key string, value int8) Log
	Int8p(key string, value *int8) Log
	Int8s(key string, value []int8) Log

	Int(key string, value int) Log
	Intp(key string, value *int) Log
	Ints(key string, value []int) Log

	String(key string, value string) Log
	Stringp(key string, value *string) Log
	Strings(key string, value []string) Log

	Time(key string, value time.Time) Log
	Timep(key string, value *time.Time) Log
	Times(key string, value []time.Time) Log

	Uint64(key string, value uint64) Log
	Uint64p(key string, value *uint64) Log
	Uint64s(key string, value []uint64) Log

	Uint32(key string, value uint32) Log
	Uint32p(key string, value *uint32) Log
	Uint32s(key string, value []uint32) Log

	Uint16(key string, value uint16) Log
	Uint16p(key string, value *uint16) Log
	Uint16s(key string, value []uint16) Log

	Uint8(key string, value uint8) Log
	Uint8p(key string, value *uint8) Log
	Uint8s(key string, value []uint8) Log

	Uint(key string, value uint) Log
	Uintp(key string, value *uint) Log
	Uints(key string, value []uint) Log

	Uintptr(key string, value uintptr) Log
	Uintptrp(key string, value *uintptr) Log
	Uintptrs(key string, value []uintptr) Log

	Reflect(key string, value interface{}) Log
	Namespace(key string) Log
	Stringer(key string, value fmt.Stringer) Log
	Any(key string, value interface{}) Log
	Skip() Log
}

// Logger 日志API
type Logger interface {
	Debug() Log
	Info() Log
	Warn() Log
	Error() Log
	DPanic() Log
	Panic() Log
	Fatal() Log
}
