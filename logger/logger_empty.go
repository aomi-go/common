package logger

import (
	"context"
	"fmt"
	"time"
)

type emptyLogger struct{}

func (l emptyLogger) Err(err error) Log {
	return l
}

func (l emptyLogger) Ctx(ctx context.Context) Log {
	return l
}

func (l emptyLogger) Binary(key string, val []byte) Log {
	return l
}

func (l emptyLogger) ByteString(key string, value []byte) Log {
	return l
}

func (l emptyLogger) Bool(key string, value bool) Log {
	return l
}

func (l emptyLogger) Boolp(key string, value *bool) Log {
	return l
}

func (l emptyLogger) Bools(key string, value []bool) Log {
	return l
}

func (l emptyLogger) Complex128(key string, value complex128) Log {
	return l
}

func (l emptyLogger) Complex128p(key string, value *complex128) Log {
	return l
}

func (l emptyLogger) Complex128s(key string, value []complex128) Log {
	return l
}

func (l emptyLogger) Complex64(key string, value complex64) Log {
	return l
}

func (l emptyLogger) Complex64p(key string, value *complex64) Log {
	return l
}

func (l emptyLogger) Complex64s(key string, value []complex64) Log {
	return l
}

func (l emptyLogger) Duration(key string, value time.Duration) Log {
	return l
}

func (l emptyLogger) Durationp(key string, value *time.Duration) Log {
	return l
}

func (l emptyLogger) Durations(key string, value []time.Duration) Log {
	return l
}

func (l emptyLogger) Float64(key string, value float64) Log {
	return l
}

func (l emptyLogger) Float64p(key string, value *float64) Log {
	return l
}

func (l emptyLogger) Float64s(key string, value []float64) Log {
	return l
}

func (l emptyLogger) Float32(key string, value float32) Log {
	return l
}

func (l emptyLogger) Float32p(key string, value *float32) Log {
	return l
}

func (l emptyLogger) Float32s(key string, value []float32) Log {
	return l
}

func (l emptyLogger) Int64(key string, value int64) Log {
	return l
}

func (l emptyLogger) Int64p(key string, value *int64) Log {
	return l
}

func (l emptyLogger) Int64s(key string, value []int64) Log {
	return l
}

func (l emptyLogger) Int32(key string, value int32) Log {
	return l
}

func (l emptyLogger) Int32p(key string, value *int32) Log {
	return l
}

func (l emptyLogger) Int32s(key string, value []int32) Log {
	return l
}

func (l emptyLogger) Int16(key string, value int16) Log {
	return l
}

func (l emptyLogger) Int16p(key string, value *int16) Log {
	return l
}

func (l emptyLogger) Int16s(key string, value []int16) Log {
	return l
}

func (l emptyLogger) Int8(key string, value int8) Log {
	return l
}

func (l emptyLogger) Int8p(key string, value *int8) Log {
	return l
}

func (l emptyLogger) Int8s(key string, value []int8) Log {
	return l
}

func (l emptyLogger) Int(key string, value int) Log {
	return l
}

func (l emptyLogger) Intp(key string, value *int) Log {
	return l
}

func (l emptyLogger) Ints(key string, value []int) Log {
	return l
}

func (l emptyLogger) String(key string, value string) Log {
	return l
}

func (l emptyLogger) Stringp(key string, value *string) Log {
	return l
}

func (l emptyLogger) Strings(key string, value []string) Log {

	return l
}

func (l emptyLogger) Time(key string, value time.Time) Log {

	return l
}

func (l emptyLogger) Timep(key string, value *time.Time) Log {

	return l
}

func (l emptyLogger) Times(key string, value []time.Time) Log {

	return l
}

func (l emptyLogger) Uint64(key string, value uint64) Log {

	return l
}

func (l emptyLogger) Uint64p(key string, value *uint64) Log {

	return l
}

func (l emptyLogger) Uint64s(key string, value []uint64) Log {

	return l
}

func (l emptyLogger) Uint32(key string, value uint32) Log {

	return l
}

func (l emptyLogger) Uint32p(key string, value *uint32) Log {

	return l
}

func (l emptyLogger) Uint32s(key string, value []uint32) Log {

	return l
}

func (l emptyLogger) Uint16(key string, value uint16) Log {

	return l
}

func (l emptyLogger) Uint16p(key string, value *uint16) Log {

	return l
}

func (l emptyLogger) Uint16s(key string, value []uint16) Log {

	return l
}

func (l emptyLogger) Uint8(key string, value uint8) Log {

	return l
}

func (l emptyLogger) Uint8p(key string, value *uint8) Log {

	return l
}

func (l emptyLogger) Uint8s(key string, value []uint8) Log {

	return l
}

func (l emptyLogger) Uint(key string, value uint) Log {

	return l
}

func (l emptyLogger) Uintp(key string, value *uint) Log {

	return l
}

func (l emptyLogger) Uints(key string, value []uint) Log {

	return l
}

func (l emptyLogger) Uintptr(key string, value uintptr) Log {

	return l
}

func (l emptyLogger) Uintptrp(key string, value *uintptr) Log {

	return l
}

func (l emptyLogger) Uintptrs(key string, value []uintptr) Log {

	return l
}

func (l emptyLogger) Reflect(key string, value interface{}) Log {

	return l
}

func (l emptyLogger) Namespace(key string) Log {

	return l
}

func (l emptyLogger) Stringer(key string, value fmt.Stringer) Log {

	return l
}

func (l emptyLogger) Any(key string, value interface{}) Log {

	return l
}

func (l emptyLogger) Skip() Log {

	return l
}

func (l emptyLogger) Debug(msg string) {
}

func (l emptyLogger) Debugf(msg string, args ...any) {
}

func (l emptyLogger) Info(msg string) {
}

func (l emptyLogger) Infof(msg string, args ...any) {
}

func (l emptyLogger) Warn(msg string) {
}

func (l emptyLogger) Warnf(msg string, args ...any) {
}

func (l emptyLogger) Error(msg string) {
}

func (l emptyLogger) Errorf(msg string, args ...any) {
}

func (l emptyLogger) DPanic(msg string) {
}

func (l emptyLogger) DPanicf(msg string, args ...any) {
}

func (l emptyLogger) Panic(msg string) {
}

func (l emptyLogger) Panicf(msg string, args ...any) {
}

func (l emptyLogger) Fatal(msg string) {
}

func (l emptyLogger) Fatalf(msg string, args ...any) {
}
