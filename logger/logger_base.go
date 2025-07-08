package logger

import (
	"context"
	"fmt"
	"time"
)

// BaseLogger 基础日志实现,主要实现链式调用功能
type BaseLogger struct {
	// NewLog 创建新的log对象，用于一次日志打印
	NewLog func() Log
}

func (l *BaseLogger) Err(err error) Log {
	return l.NewLog().Err(err)
}

func (l *BaseLogger) Ctx(ctx context.Context) Log {
	return l.NewLog().Ctx(ctx)
}

func (l *BaseLogger) Binary(key string, val []byte) Log {
	return l.NewLog().Binary(key, val)
}

func (l *BaseLogger) ByteString(key string, value []byte) Log {
	return l.NewLog().ByteString(key, value)
}

func (l *BaseLogger) Bool(key string, value bool) Log {
	return l.NewLog().Bool(key, value)
}

func (l *BaseLogger) Boolp(key string, value *bool) Log {
	return l.NewLog().Boolp(key, value)
}

func (l *BaseLogger) Bools(key string, value []bool) Log {
	return l.NewLog().Bools(key, value)
}

func (l *BaseLogger) Complex128(key string, value complex128) Log {
	return l.NewLog().Complex128(key, value)
}

func (l *BaseLogger) Complex128p(key string, value *complex128) Log {
	return l.NewLog().Complex128p(key, value)
}

func (l *BaseLogger) Complex128s(key string, value []complex128) Log {
	return l.NewLog().Complex128s(key, value)
}

func (l *BaseLogger) Complex64(key string, value complex64) Log {
	return l.NewLog().Complex64(key, value)
}

func (l *BaseLogger) Complex64p(key string, value *complex64) Log {
	return l.NewLog().Complex64p(key, value)
}

func (l *BaseLogger) Complex64s(key string, value []complex64) Log {
	return l.NewLog().Complex64s(key, value)
}

func (l *BaseLogger) Duration(key string, value time.Duration) Log {
	return l.NewLog().Duration(key, value)
}

func (l *BaseLogger) Durationp(key string, value *time.Duration) Log {
	return l.NewLog().Durationp(key, value)
}

func (l *BaseLogger) Durations(key string, value []time.Duration) Log {
	return l.NewLog().Durations(key, value)
}

func (l *BaseLogger) Float64(key string, value float64) Log {
	return l.NewLog().Float64(key, value)
}

func (l *BaseLogger) Float64p(key string, value *float64) Log {
	return l.NewLog().Float64p(key, value)
}

func (l *BaseLogger) Float64s(key string, value []float64) Log {
	return l.NewLog().Float64s(key, value)
}

func (l *BaseLogger) Float32(key string, value float32) Log {
	return l.NewLog().Float32(key, value)
}

func (l *BaseLogger) Float32p(key string, value *float32) Log {
	return l.NewLog().Float32p(key, value)
}

func (l *BaseLogger) Float32s(key string, value []float32) Log {
	return l.NewLog().Float32s(key, value)
}

func (l *BaseLogger) Int64(key string, value int64) Log {
	return l.NewLog().Int64(key, value)
}

func (l *BaseLogger) Int64p(key string, value *int64) Log {
	return l.NewLog().Int64p(key, value)
}

func (l *BaseLogger) Int64s(key string, value []int64) Log {
	return l.NewLog().Int64s(key, value)
}

func (l *BaseLogger) Int32(key string, value int32) Log {
	return l.NewLog().Int32(key, value)
}

func (l *BaseLogger) Int32p(key string, value *int32) Log {
	return l.NewLog().Int32p(key, value)
}

func (l *BaseLogger) Int32s(key string, value []int32) Log {
	return l.NewLog().Int32s(key, value)
}

func (l *BaseLogger) Int16(key string, value int16) Log {
	return l.NewLog().Int16(key, value)
}

func (l *BaseLogger) Int16p(key string, value *int16) Log {
	return l.NewLog().Int16p(key, value)
}

func (l *BaseLogger) Int16s(key string, value []int16) Log {
	return l.NewLog().Int16s(key, value)
}

func (l *BaseLogger) Int8(key string, value int8) Log {
	return l.NewLog().Int8(key, value)
}

func (l *BaseLogger) Int8p(key string, value *int8) Log {
	return l.NewLog().Int8p(key, value)
}

func (l *BaseLogger) Int8s(key string, value []int8) Log {
	return l.NewLog().Int8s(key, value)
}

func (l *BaseLogger) Int(key string, value int) Log {
	return l.NewLog().Int(key, value)
}

func (l *BaseLogger) Intp(key string, value *int) Log {
	return l.NewLog().Intp(key, value)
}

func (l *BaseLogger) Ints(key string, value []int) Log {
	return l.NewLog().Ints(key, value)
}

func (l *BaseLogger) String(key string, value string) Log {
	return l.NewLog().String(key, value)
}

func (l *BaseLogger) Stringp(key string, value *string) Log {
	return l.NewLog().Stringp(key, value)
}

func (l *BaseLogger) Strings(key string, value []string) Log {
	return l.NewLog().Strings(key, value)
}

func (l *BaseLogger) Time(key string, value time.Time) Log {
	return l.NewLog().Time(key, value)
}

func (l *BaseLogger) Timep(key string, value *time.Time) Log {
	return l.NewLog().Timep(key, value)
}

func (l *BaseLogger) Times(key string, value []time.Time) Log {
	return l.NewLog().Times(key, value)
}

func (l *BaseLogger) Uint64(key string, value uint64) Log {
	return l.NewLog().Uint64(key, value)
}

func (l *BaseLogger) Uint64p(key string, value *uint64) Log {
	return l.NewLog().Uint64p(key, value)
}

func (l *BaseLogger) Uint64s(key string, value []uint64) Log {
	return l.NewLog().Uint64s(key, value)
}

func (l *BaseLogger) Uint32(key string, value uint32) Log {
	return l.NewLog().Uint32(key, value)
}

func (l *BaseLogger) Uint32p(key string, value *uint32) Log {
	return l.NewLog().Uint32p(key, value)
}

func (l *BaseLogger) Uint32s(key string, value []uint32) Log {
	return l.NewLog().Uint32s(key, value)
}

func (l *BaseLogger) Uint16(key string, value uint16) Log {
	return l.NewLog().Uint16(key, value)
}

func (l *BaseLogger) Uint16p(key string, value *uint16) Log {
	return l.NewLog().Uint16p(key, value)
}

func (l *BaseLogger) Uint16s(key string, value []uint16) Log {
	return l.NewLog().Uint16s(key, value)
}

func (l *BaseLogger) Uint8(key string, value uint8) Log {
	return l.NewLog().Uint8(key, value)
}

func (l *BaseLogger) Uint8p(key string, value *uint8) Log {
	return l.NewLog().Uint8p(key, value)
}

func (l *BaseLogger) Uint8s(key string, value []uint8) Log {
	return l.NewLog().Uint8s(key, value)
}

func (l *BaseLogger) Uint(key string, value uint) Log {
	return l.NewLog().Uint(key, value)
}

func (l *BaseLogger) Uintp(key string, value *uint) Log {
	return l.NewLog().Uintp(key, value)
}

func (l *BaseLogger) Uints(key string, value []uint) Log {
	return l.NewLog().Uints(key, value)
}

func (l *BaseLogger) Uintptr(key string, value uintptr) Log {
	return l.NewLog().Uintptr(key, value)
}

func (l *BaseLogger) Uintptrp(key string, value *uintptr) Log {
	return l.NewLog().Uintptrp(key, value)
}

func (l *BaseLogger) Uintptrs(key string, value []uintptr) Log {
	return l.NewLog().Uintptrs(key, value)
}

func (l *BaseLogger) Reflect(key string, value interface{}) Log {
	return l.NewLog().Reflect(key, value)
}

func (l *BaseLogger) Namespace(key string) Log {
	return l.NewLog().Namespace(key)
}

func (l *BaseLogger) Stringer(key string, value fmt.Stringer) Log {
	return l.NewLog().Stringer(key, value)
}

func (l *BaseLogger) Any(key string, value interface{}) Log {
	return l.NewLog().Any(key, value)
}

func (l *BaseLogger) Skip() Log {
	return l.NewLog().Skip()
}

func (l *BaseLogger) Debug(msg string) {
	l.NewLog().Debug(msg)
}

func (l *BaseLogger) Debugf(msg string, args ...any) {
	l.NewLog().Debugf(msg, args...)
}

func (l *BaseLogger) Info(msg string) {
	l.NewLog().Info(msg)
}

func (l *BaseLogger) Infof(msg string, args ...any) {
	l.NewLog().Infof(msg, args...)
}

func (l *BaseLogger) Warn(msg string) {
	l.NewLog().Warn(msg)
}

func (l *BaseLogger) Warnf(msg string, args ...any) {
	l.NewLog().Warnf(msg, args...)
}

func (l *BaseLogger) Error(msg string) {
	l.NewLog().Error(msg)
}

func (l *BaseLogger) Errorf(msg string, args ...any) {
	l.NewLog().Errorf(msg, args...)
}

func (l *BaseLogger) DPanic(msg string) {
	l.NewLog().DPanic(msg)
}

func (l *BaseLogger) DPanicf(msg string, args ...any) {
	l.NewLog().DPanicf(msg, args...)
}

func (l *BaseLogger) Panic(msg string) {
	l.NewLog().Panic(msg)
}

func (l *BaseLogger) Panicf(msg string, args ...any) {
	l.NewLog().Panicf(msg, args...)
}

func (l *BaseLogger) Fatal(msg string) {
	l.NewLog().Fatal(msg)
}

func (l *BaseLogger) Fatalf(msg string, args ...any) {
	l.NewLog().Fatalf(msg, args...)
}
