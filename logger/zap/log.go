package zap

import (
	"context"
	"fmt"
	"github.com/aomi-go/common/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type Log struct {
	fields   []zap.Field
	provider *zap.Logger
}

func (l Log) msg(lvl zapcore.Level, msg string) {
	l.provider.Log(lvl, msg, l.fields...)
}

func (l Log) Debug(msg string) {
	l.msg(zapcore.DebugLevel, msg)
}

func (l Log) Debugf(msg string, args ...any) {
	l.msg(zapcore.DebugLevel, fmt.Sprintf(msg, args...))
}

func (l Log) Info(msg string) {
	l.msg(zapcore.InfoLevel, msg)
}
func (l Log) Infof(msg string, args ...any) {
	l.msg(zapcore.InfoLevel, fmt.Sprintf(msg, args...))
}

func (l Log) Warn(msg string) {
	l.msg(zapcore.WarnLevel, msg)
}

func (l Log) Warnf(msg string, args ...any) {
	l.msg(zapcore.WarnLevel, fmt.Sprintf(msg, args...))
}

func (l Log) Error(msg string) {
	l.msg(zapcore.ErrorLevel, msg)
}

func (l Log) Errorf(msg string, args ...any) {
	l.msg(zapcore.ErrorLevel, fmt.Sprintf(msg, args...))
}
func (l Log) DPanic(msg string) {
	l.msg(zapcore.DPanicLevel, msg)
}
func (l Log) DPanicf(msg string, args ...any) {
	l.msg(zapcore.DPanicLevel, fmt.Sprintf(msg, args...))
}

func (l Log) Panic(msg string) {
	l.msg(zapcore.PanicLevel, msg)
}

func (l Log) Panicf(msg string, args ...any) {
	l.msg(zapcore.PanicLevel, fmt.Sprintf(msg, args...))
}

func (l Log) Fatal(msg string) {
	l.msg(zapcore.FatalLevel, msg)
}

func (l Log) Fatalf(msg string, args ...any) {
	l.msg(zapcore.FatalLevel, fmt.Sprintf(msg, args...))
}

func (l Log) Err(err error) logger.Log {
	l.fields = append(l.fields, zap.Error(err))
	return l
}

func (l Log) Ctx(ctx context.Context) logger.Log {
	traceId := ctx.Value(logger.TraceIdCtxKey)
	if nil != traceId {
		if i, ok := traceId.(string); ok {
			l.fields = append(l.fields, zap.String("traceId", i))
		}
	}
	return l
}

func (l Log) Binary(key string, val []byte) logger.Log {
	l.fields = append(l.fields, zap.Binary(key, val))
	return l
}

func (l Log) ByteString(key string, value []byte) logger.Log {
	l.fields = append(l.fields, zap.ByteString(key, value))
	return l
}

func (l Log) Bool(key string, value bool) logger.Log {
	l.fields = append(l.fields, zap.Bool(key, value))
	return l
}

func (l Log) Boolp(key string, value *bool) logger.Log {
	l.fields = append(l.fields, zap.Boolp(key, value))
	return l
}

func (l Log) Bools(key string, value []bool) logger.Log {
	l.fields = append(l.fields, zap.Bools(key, value))
	return l
}

func (l Log) Complex128(key string, value complex128) logger.Log {
	l.fields = append(l.fields, zap.Complex128(key, value))
	return l
}

func (l Log) Complex128p(key string, value *complex128) logger.Log {
	l.fields = append(l.fields, zap.Complex128p(key, value))
	return l
}

func (l Log) Complex128s(key string, value []complex128) logger.Log {
	l.fields = append(l.fields, zap.Complex128s(key, value))
	return l
}

func (l Log) Complex64(key string, value complex64) logger.Log {
	l.fields = append(l.fields, zap.Complex64(key, value))
	return l
}

func (l Log) Complex64p(key string, value *complex64) logger.Log {
	l.fields = append(l.fields, zap.Complex64p(key, value))
	return l
}

func (l Log) Complex64s(key string, value []complex64) logger.Log {
	l.fields = append(l.fields, zap.Complex64s(key, value))
	return l
}

func (l Log) Duration(key string, value time.Duration) logger.Log {
	l.fields = append(l.fields, zap.Duration(key, value))
	return l
}

func (l Log) Durationp(key string, value *time.Duration) logger.Log {
	l.fields = append(l.fields, zap.Durationp(key, value))
	return l
}

func (l Log) Durations(key string, value []time.Duration) logger.Log {
	l.fields = append(l.fields, zap.Durations(key, value))
	return l
}

func (l Log) Float64(key string, value float64) logger.Log {
	l.fields = append(l.fields, zap.Float64(key, value))
	return l
}

func (l Log) Float64p(key string, value *float64) logger.Log {
	l.fields = append(l.fields, zap.Float64p(key, value))
	return l
}

func (l Log) Float64s(key string, value []float64) logger.Log {
	l.fields = append(l.fields, zap.Float64s(key, value))
	return l
}

func (l Log) Float32(key string, value float32) logger.Log {
	l.fields = append(l.fields, zap.Float32(key, value))
	return l
}

func (l Log) Float32p(key string, value *float32) logger.Log {
	l.fields = append(l.fields, zap.Float32p(key, value))
	return l
}

func (l Log) Float32s(key string, value []float32) logger.Log {
	l.fields = append(l.fields, zap.Float32s(key, value))
	return l
}

func (l Log) Int64(key string, value int64) logger.Log {
	l.fields = append(l.fields, zap.Int64(key, value))
	return l
}

func (l Log) Int64p(key string, value *int64) logger.Log {
	l.fields = append(l.fields, zap.Int64p(key, value))
	return l
}

func (l Log) Int64s(key string, value []int64) logger.Log {
	l.fields = append(l.fields, zap.Int64s(key, value))
	return l
}

func (l Log) Int32(key string, value int32) logger.Log {
	l.fields = append(l.fields, zap.Int32(key, value))
	return l
}

func (l Log) Int32p(key string, value *int32) logger.Log {
	l.fields = append(l.fields, zap.Int32p(key, value))
	return l
}

func (l Log) Int32s(key string, value []int32) logger.Log {
	l.fields = append(l.fields, zap.Int32s(key, value))
	return l
}

func (l Log) Int16(key string, value int16) logger.Log {
	l.fields = append(l.fields, zap.Int16(key, value))
	return l
}

func (l Log) Int16p(key string, value *int16) logger.Log {
	l.fields = append(l.fields, zap.Int16p(key, value))
	return l
}

func (l Log) Int16s(key string, value []int16) logger.Log {
	l.fields = append(l.fields, zap.Int16s(key, value))
	return l
}

func (l Log) Int8(key string, value int8) logger.Log {
	l.fields = append(l.fields, zap.Int8(key, value))
	return l
}

func (l Log) Int8p(key string, value *int8) logger.Log {
	l.fields = append(l.fields, zap.Int8p(key, value))
	return l
}

func (l Log) Int8s(key string, value []int8) logger.Log {
	l.fields = append(l.fields, zap.Int8s(key, value))
	return l
}

func (l Log) Int(key string, value int) logger.Log {
	l.fields = append(l.fields, zap.Int(key, value))
	return l
}

func (l Log) Intp(key string, value *int) logger.Log {
	l.fields = append(l.fields, zap.Intp(key, value))
	return l
}

func (l Log) Ints(key string, value []int) logger.Log {
	l.fields = append(l.fields, zap.Ints(key, value))
	return l
}

func (l Log) String(key string, value string) logger.Log {
	l.fields = append(l.fields, zap.String(key, value))
	return l
}

func (l Log) Stringp(key string, value *string) logger.Log {
	l.fields = append(l.fields, zap.Stringp(key, value))
	return l
}

func (l Log) Strings(key string, value []string) logger.Log {
	l.fields = append(l.fields, zap.Strings(key, value))
	return l
}

func (l Log) Time(key string, value time.Time) logger.Log {
	l.fields = append(l.fields, zap.Time(key, value))
	return l
}

func (l Log) Timep(key string, value *time.Time) logger.Log {
	l.fields = append(l.fields, zap.Timep(key, value))
	return l
}

func (l Log) Times(key string, value []time.Time) logger.Log {
	l.fields = append(l.fields, zap.Times(key, value))
	return l
}

func (l Log) Uint64(key string, value uint64) logger.Log {
	l.fields = append(l.fields, zap.Uint64(key, value))
	return l
}

func (l Log) Uint64p(key string, value *uint64) logger.Log {
	l.fields = append(l.fields, zap.Uint64p(key, value))
	return l
}

func (l Log) Uint64s(key string, value []uint64) logger.Log {
	l.fields = append(l.fields, zap.Uint64s(key, value))
	return l
}

func (l Log) Uint32(key string, value uint32) logger.Log {
	l.fields = append(l.fields, zap.Uint32(key, value))
	return l
}

func (l Log) Uint32p(key string, value *uint32) logger.Log {
	l.fields = append(l.fields, zap.Uint32p(key, value))
	return l
}

func (l Log) Uint32s(key string, value []uint32) logger.Log {
	l.fields = append(l.fields, zap.Uint32s(key, value))
	return l
}

func (l Log) Uint16(key string, value uint16) logger.Log {
	l.fields = append(l.fields, zap.Uint16(key, value))
	return l
}

func (l Log) Uint16p(key string, value *uint16) logger.Log {
	l.fields = append(l.fields, zap.Uint16p(key, value))
	return l
}

func (l Log) Uint16s(key string, value []uint16) logger.Log {
	l.fields = append(l.fields, zap.Uint16s(key, value))
	return l
}

func (l Log) Uint8(key string, value uint8) logger.Log {
	l.fields = append(l.fields, zap.Uint8(key, value))
	return l
}

func (l Log) Uint8p(key string, value *uint8) logger.Log {
	l.fields = append(l.fields, zap.Uint8p(key, value))
	return l
}

func (l Log) Uint8s(key string, value []uint8) logger.Log {
	l.fields = append(l.fields, zap.Uint8s(key, value))
	return l
}

func (l Log) Uint(key string, value uint) logger.Log {
	l.fields = append(l.fields, zap.Uint(key, value))
	return l
}

func (l Log) Uintp(key string, value *uint) logger.Log {
	l.fields = append(l.fields, zap.Uintp(key, value))
	return l
}

func (l Log) Uints(key string, value []uint) logger.Log {
	l.fields = append(l.fields, zap.Uints(key, value))
	return l
}

func (l Log) Uintptr(key string, value uintptr) logger.Log {
	l.fields = append(l.fields, zap.Uintptr(key, value))
	return l
}

func (l Log) Uintptrp(key string, value *uintptr) logger.Log {
	l.fields = append(l.fields, zap.Uintptrp(key, value))
	return l
}

func (l Log) Uintptrs(key string, value []uintptr) logger.Log {
	l.fields = append(l.fields, zap.Uintptrs(key, value))
	return l
}

func (l Log) Reflect(key string, value interface{}) logger.Log {
	l.fields = append(l.fields, zap.Reflect(key, value))
	return l
}

func (l Log) Namespace(key string) logger.Log {
	l.fields = append(l.fields, zap.Namespace(key))
	return l
}

func (l Log) Stringer(key string, value fmt.Stringer) logger.Log {
	l.fields = append(l.fields, zap.Stringer(key, value))
	return l
}

func (l Log) Any(key string, value interface{}) logger.Log {
	l.fields = append(l.fields, zap.Any(key, value))
	return l
}

func (l Log) Skip() logger.Log {
	l.fields = append(l.fields, zap.Skip())
	return l
}
