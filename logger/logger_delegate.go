package logger

import (
	"context"
	"fmt"
	"time"
)

func NewLoggerDelegate(logger Logger) *Delegate {
	return &Delegate{
		Logger: logger,
	}
}

// Delegate 用于将Logger接口的实现委托给另一个Logger实例
// 主要用于动态修改Logger的实现
type Delegate struct {
	Logger Logger
}

func (l *Delegate) Err(err error) Log {
	return l.Logger.Err(err)
}

func (l *Delegate) Ctx(ctx context.Context) Log {
	return l.Logger.Ctx(ctx)
}

func (l *Delegate) Binary(key string, val []byte) Log {
	return l.Logger.Binary(key, val)
}

func (l *Delegate) ByteString(key string, value []byte) Log {
	return l.Logger.ByteString(key, value)
}

func (l *Delegate) Bool(key string, value bool) Log {
	return l.Logger.Bool(key, value)
}

func (l *Delegate) Boolp(key string, value *bool) Log {
	return l.Logger.Boolp(key, value)
}

func (l *Delegate) Bools(key string, value []bool) Log {
	return l.Logger.Bools(key, value)
}

func (l *Delegate) Complex128(key string, value complex128) Log {
	return l.Logger.Complex128(key, value)
}

func (l *Delegate) Complex128p(key string, value *complex128) Log {
	return l.Logger.Complex128p(key, value)
}

func (l *Delegate) Complex128s(key string, value []complex128) Log {
	return l.Logger.Complex128s(key, value)
}

func (l *Delegate) Complex64(key string, value complex64) Log {
	return l.Logger.Complex64(key, value)
}

func (l *Delegate) Complex64p(key string, value *complex64) Log {
	return l.Logger.Complex64p(key, value)
}

func (l *Delegate) Complex64s(key string, value []complex64) Log {
	return l.Logger.Complex64s(key, value)
}

func (l *Delegate) Duration(key string, value time.Duration) Log {
	return l.Logger.Duration(key, value)
}

func (l *Delegate) Durationp(key string, value *time.Duration) Log {
	return l.Logger.Durationp(key, value)
}

func (l *Delegate) Durations(key string, value []time.Duration) Log {
	return l.Logger.Durations(key, value)
}

func (l *Delegate) Float64(key string, value float64) Log {
	return l.Logger.Float64(key, value)
}

func (l *Delegate) Float64p(key string, value *float64) Log {
	return l.Logger.Float64p(key, value)
}

func (l *Delegate) Float64s(key string, value []float64) Log {
	return l.Logger.Float64s(key, value)
}

func (l *Delegate) Float32(key string, value float32) Log {
	return l.Logger.Float32(key, value)
}

func (l *Delegate) Float32p(key string, value *float32) Log {
	return l.Logger.Float32p(key, value)
}

func (l *Delegate) Float32s(key string, value []float32) Log {
	return l.Logger.Float32s(key, value)
}

func (l *Delegate) Int64(key string, value int64) Log {
	return l.Logger.Int64(key, value)
}

func (l *Delegate) Int64p(key string, value *int64) Log {
	return l.Logger.Int64p(key, value)
}

func (l *Delegate) Int64s(key string, value []int64) Log {
	return l.Logger.Int64s(key, value)
}

func (l *Delegate) Int32(key string, value int32) Log {
	return l.Logger.Int32(key, value)
}

func (l *Delegate) Int32p(key string, value *int32) Log {
	return l.Logger.Int32p(key, value)
}

func (l *Delegate) Int32s(key string, value []int32) Log {
	return l.Logger.Int32s(key, value)
}

func (l *Delegate) Int16(key string, value int16) Log {
	return l.Logger.Int16(key, value)
}

func (l *Delegate) Int16p(key string, value *int16) Log {
	return l.Logger.Int16p(key, value)
}

func (l *Delegate) Int16s(key string, value []int16) Log {
	return l.Logger.Int16s(key, value)
}

func (l *Delegate) Int8(key string, value int8) Log {
	return l.Logger.Int8(key, value)
}

func (l *Delegate) Int8p(key string, value *int8) Log {
	return l.Logger.Int8p(key, value)
}

func (l *Delegate) Int8s(key string, value []int8) Log {
	return l.Logger.Int8s(key, value)
}

func (l *Delegate) Int(key string, value int) Log {
	return l.Logger.Int(key, value)
}

func (l *Delegate) Intp(key string, value *int) Log {
	return l.Logger.Intp(key, value)
}

func (l *Delegate) Ints(key string, value []int) Log {
	return l.Logger.Ints(key, value)
}

func (l *Delegate) String(key string, value string) Log {
	return l.Logger.String(key, value)
}

func (l *Delegate) Stringp(key string, value *string) Log {
	return l.Logger.Stringp(key, value)
}

func (l *Delegate) Strings(key string, value []string) Log {
	return l.Logger.Strings(key, value)
}

func (l *Delegate) Time(key string, value time.Time) Log {
	return l.Logger.Time(key, value)
}

func (l *Delegate) Timep(key string, value *time.Time) Log {
	return l.Logger.Timep(key, value)
}

func (l *Delegate) Times(key string, value []time.Time) Log {
	return l.Logger.Times(key, value)
}

func (l *Delegate) Uint64(key string, value uint64) Log {
	return l.Logger.Uint64(key, value)
}

func (l *Delegate) Uint64p(key string, value *uint64) Log {
	return l.Logger.Uint64p(key, value)
}

func (l *Delegate) Uint64s(key string, value []uint64) Log {
	return l.Logger.Uint64s(key, value)
}

func (l *Delegate) Uint32(key string, value uint32) Log {
	return l.Logger.Uint32(key, value)
}

func (l *Delegate) Uint32p(key string, value *uint32) Log {
	return l.Logger.Uint32p(key, value)
}

func (l *Delegate) Uint32s(key string, value []uint32) Log {
	return l.Logger.Uint32s(key, value)
}

func (l *Delegate) Uint16(key string, value uint16) Log {
	return l.Logger.Uint16(key, value)
}

func (l *Delegate) Uint16p(key string, value *uint16) Log {
	return l.Logger.Uint16p(key, value)
}

func (l *Delegate) Uint16s(key string, value []uint16) Log {
	return l.Logger.Uint16s(key, value)
}

func (l *Delegate) Uint8(key string, value uint8) Log {
	return l.Logger.Uint8(key, value)
}

func (l *Delegate) Uint8p(key string, value *uint8) Log {
	return l.Logger.Uint8p(key, value)
}

func (l *Delegate) Uint8s(key string, value []uint8) Log {
	return l.Logger.Uint8s(key, value)
}

func (l *Delegate) Uint(key string, value uint) Log {
	return l.Logger.Uint(key, value)
}

func (l *Delegate) Uintp(key string, value *uint) Log {
	return l.Logger.Uintp(key, value)
}

func (l *Delegate) Uints(key string, value []uint) Log {
	return l.Logger.Uints(key, value)
}

func (l *Delegate) Uintptr(key string, value uintptr) Log {
	return l.Logger.Uintptr(key, value)
}

func (l *Delegate) Uintptrp(key string, value *uintptr) Log {
	return l.Logger.Uintptrp(key, value)
}

func (l *Delegate) Uintptrs(key string, value []uintptr) Log {
	return l.Logger.Uintptrs(key, value)
}

func (l *Delegate) Reflect(key string, value interface{}) Log {
	return l.Logger.Reflect(key, value)
}

func (l *Delegate) Namespace(key string) Log {
	return l.Logger.Namespace(key)
}

func (l *Delegate) Stringer(key string, value fmt.Stringer) Log {
	return l.Logger.Stringer(key, value)
}

func (l *Delegate) Any(key string, value interface{}) Log {
	return l.Logger.Any(key, value)
}

func (l *Delegate) Skip() Log {
	return l.Logger.Skip()
}

func (l *Delegate) Debug(msg string) {
	l.Logger.Debug(msg)
}

func (l *Delegate) Debugf(msg string, args ...any) {
	l.Logger.Debugf(msg, args...)
}

func (l *Delegate) Info(msg string) {
	l.Logger.Info(msg)
}

func (l *Delegate) Infof(msg string, args ...any) {
	l.Logger.Infof(msg, args...)
}

func (l *Delegate) Warn(msg string) {
	l.Logger.Warn(msg)
}

func (l *Delegate) Warnf(msg string, args ...any) {
	l.Logger.Warnf(msg, args...)
}

func (l *Delegate) Error(msg string) {
	l.Logger.Error(msg)
}

func (l *Delegate) Errorf(msg string, args ...any) {
	l.Logger.Errorf(msg, args...)
}

func (l *Delegate) DPanic(msg string) {
	l.Logger.DPanic(msg)
}

func (l *Delegate) DPanicf(msg string, args ...any) {
	l.Logger.DPanicf(msg, args...)
}

func (l *Delegate) Panic(msg string) {
	l.Logger.Panic(msg)
}

func (l *Delegate) Panicf(msg string, args ...any) {
	l.Logger.Panicf(msg, args...)
}

func (l *Delegate) Fatal(msg string) {
	l.Logger.Fatal(msg)
}

func (l *Delegate) Fatalf(msg string, args ...any) {
	l.Logger.Fatalf(msg, args...)
}
