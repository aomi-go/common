package zap

import (
	"context"
	"fmt"
	"github.com/aomi-go/common/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type ZapLog struct {
	fields   []zap.Field
	provider *zap.Logger
}

func (l ZapLog) msg(lvl zapcore.Level, msg string) {
	l.provider.Log(lvl, msg, l.fields...)
}

func (l ZapLog) Debug(msg string) {
	l.msg(zapcore.DebugLevel, msg)
}
func (l ZapLog) Info(msg string) {
	l.msg(zapcore.InfoLevel, msg)
}
func (l ZapLog) Warn(msg string) {
	l.msg(zapcore.WarnLevel, msg)
}
func (l ZapLog) Error(msg string) {
	l.msg(zapcore.ErrorLevel, msg)
}
func (l ZapLog) DPanic(msg string) {
	l.msg(zapcore.DPanicLevel, msg)
}
func (l ZapLog) Panic(msg string) {
	l.msg(zapcore.PanicLevel, msg)
}
func (l ZapLog) Fatal(msg string) {
	l.msg(zapcore.FatalLevel, msg)
}

func (l ZapLog) Err(err error) logger.Log {
	l.fields = append(l.fields, zap.Error(err))
	return l
}

func (l ZapLog) Ctx(ctx context.Context) logger.Log {
	traceId := ctx.Value(logger.TraceIdCtxKey)
	if nil != traceId {
		if i, ok := traceId.(string); ok {
			l.fields = append(l.fields, zap.String("traceId", i))
		}
	}
	return l
}

func (l ZapLog) Binary(key string, val []byte) logger.Log {
	l.fields = append(l.fields, zap.Binary(key, val))
	return l
}

func (l ZapLog) ByteString(key string, value []byte) logger.Log {
	l.fields = append(l.fields, zap.ByteString(key, value))
	return l
}

func (l ZapLog) Bool(key string, value bool) logger.Log {
	l.fields = append(l.fields, zap.Bool(key, value))
	return l
}

func (l ZapLog) Boolp(key string, value *bool) logger.Log {
	l.fields = append(l.fields, zap.Boolp(key, value))
	return l
}

func (l ZapLog) Bools(key string, value []bool) logger.Log {
	l.fields = append(l.fields, zap.Bools(key, value))
	return l
}

func (l ZapLog) Complex128(key string, value complex128) logger.Log {
	l.fields = append(l.fields, zap.Complex128(key, value))
	return l
}

func (l ZapLog) Complex128p(key string, value *complex128) logger.Log {
	l.fields = append(l.fields, zap.Complex128p(key, value))
	return l
}

func (l ZapLog) Complex128s(key string, value []complex128) logger.Log {
	l.fields = append(l.fields, zap.Complex128s(key, value))
	return l
}

func (l ZapLog) Complex64(key string, value complex64) logger.Log {
	l.fields = append(l.fields, zap.Complex64(key, value))
	return l
}

func (l ZapLog) Complex64p(key string, value *complex64) logger.Log {
	l.fields = append(l.fields, zap.Complex64p(key, value))
	return l
}

func (l ZapLog) Complex64s(key string, value []complex64) logger.Log {
	l.fields = append(l.fields, zap.Complex64s(key, value))
	return l
}

func (l ZapLog) Duration(key string, value time.Duration) logger.Log {
	l.fields = append(l.fields, zap.Duration(key, value))
	return l
}

func (l ZapLog) Durationp(key string, value *time.Duration) logger.Log {
	l.fields = append(l.fields, zap.Durationp(key, value))
	return l
}

func (l ZapLog) Durations(key string, value []time.Duration) logger.Log {
	l.fields = append(l.fields, zap.Durations(key, value))
	return l
}

func (l ZapLog) Float64(key string, value float64) logger.Log {
	l.fields = append(l.fields, zap.Float64(key, value))
	return l
}

func (l ZapLog) Float64p(key string, value *float64) logger.Log {
	l.fields = append(l.fields, zap.Float64p(key, value))
	return l
}

func (l ZapLog) Float64s(key string, value []float64) logger.Log {
	l.fields = append(l.fields, zap.Float64s(key, value))
	return l
}

func (l ZapLog) Float32(key string, value float32) logger.Log {
	l.fields = append(l.fields, zap.Float32(key, value))
	return l
}

func (l ZapLog) Float32p(key string, value *float32) logger.Log {
	l.fields = append(l.fields, zap.Float32p(key, value))
	return l
}

func (l ZapLog) Float32s(key string, value []float32) logger.Log {
	l.fields = append(l.fields, zap.Float32s(key, value))
	return l
}

func (l ZapLog) Int64(key string, value int64) logger.Log {
	l.fields = append(l.fields, zap.Int64(key, value))
	return l
}

func (l ZapLog) Int64p(key string, value *int64) logger.Log {
	l.fields = append(l.fields, zap.Int64p(key, value))
	return l
}

func (l ZapLog) Int64s(key string, value []int64) logger.Log {
	l.fields = append(l.fields, zap.Int64s(key, value))
	return l
}

func (l ZapLog) Int32(key string, value int32) logger.Log {
	l.fields = append(l.fields, zap.Int32(key, value))
	return l
}

func (l ZapLog) Int32p(key string, value *int32) logger.Log {
	l.fields = append(l.fields, zap.Int32p(key, value))
	return l
}

func (l ZapLog) Int32s(key string, value []int32) logger.Log {
	l.fields = append(l.fields, zap.Int32s(key, value))
	return l
}

func (l ZapLog) Int16(key string, value int16) logger.Log {
	l.fields = append(l.fields, zap.Int16(key, value))
	return l
}

func (l ZapLog) Int16p(key string, value *int16) logger.Log {
	l.fields = append(l.fields, zap.Int16p(key, value))
	return l
}

func (l ZapLog) Int16s(key string, value []int16) logger.Log {
	l.fields = append(l.fields, zap.Int16s(key, value))
	return l
}

func (l ZapLog) Int8(key string, value int8) logger.Log {
	l.fields = append(l.fields, zap.Int8(key, value))
	return l
}

func (l ZapLog) Int8p(key string, value *int8) logger.Log {
	l.fields = append(l.fields, zap.Int8p(key, value))
	return l
}

func (l ZapLog) Int8s(key string, value []int8) logger.Log {
	l.fields = append(l.fields, zap.Int8s(key, value))
	return l
}

func (l ZapLog) Int(key string, value int) logger.Log {
	l.fields = append(l.fields, zap.Int(key, value))
	return l
}

func (l ZapLog) Intp(key string, value *int) logger.Log {
	l.fields = append(l.fields, zap.Intp(key, value))
	return l
}

func (l ZapLog) Ints(key string, value []int) logger.Log {
	l.fields = append(l.fields, zap.Ints(key, value))
	return l
}

func (l ZapLog) String(key string, value string) logger.Log {
	l.fields = append(l.fields, zap.String(key, value))
	return l
}

func (l ZapLog) Stringp(key string, value *string) logger.Log {
	l.fields = append(l.fields, zap.Stringp(key, value))
	return l
}

func (l ZapLog) Strings(key string, value []string) logger.Log {
	l.fields = append(l.fields, zap.Strings(key, value))
	return l
}

func (l ZapLog) Time(key string, value time.Time) logger.Log {
	l.fields = append(l.fields, zap.Time(key, value))
	return l
}

func (l ZapLog) Timep(key string, value *time.Time) logger.Log {
	l.fields = append(l.fields, zap.Timep(key, value))
	return l
}

func (l ZapLog) Times(key string, value []time.Time) logger.Log {
	l.fields = append(l.fields, zap.Times(key, value))
	return l
}

func (l ZapLog) Uint64(key string, value uint64) logger.Log {
	l.fields = append(l.fields, zap.Uint64(key, value))
	return l
}

func (l ZapLog) Uint64p(key string, value *uint64) logger.Log {
	l.fields = append(l.fields, zap.Uint64p(key, value))
	return l
}

func (l ZapLog) Uint64s(key string, value []uint64) logger.Log {
	l.fields = append(l.fields, zap.Uint64s(key, value))
	return l
}

func (l ZapLog) Uint32(key string, value uint32) logger.Log {
	l.fields = append(l.fields, zap.Uint32(key, value))
	return l
}

func (l ZapLog) Uint32p(key string, value *uint32) logger.Log {
	l.fields = append(l.fields, zap.Uint32p(key, value))
	return l
}

func (l ZapLog) Uint32s(key string, value []uint32) logger.Log {
	l.fields = append(l.fields, zap.Uint32s(key, value))
	return l
}

func (l ZapLog) Uint16(key string, value uint16) logger.Log {
	l.fields = append(l.fields, zap.Uint16(key, value))
	return l
}

func (l ZapLog) Uint16p(key string, value *uint16) logger.Log {
	l.fields = append(l.fields, zap.Uint16p(key, value))
	return l
}

func (l ZapLog) Uint16s(key string, value []uint16) logger.Log {
	l.fields = append(l.fields, zap.Uint16s(key, value))
	return l
}

func (l ZapLog) Uint8(key string, value uint8) logger.Log {
	l.fields = append(l.fields, zap.Uint8(key, value))
	return l
}

func (l ZapLog) Uint8p(key string, value *uint8) logger.Log {
	l.fields = append(l.fields, zap.Uint8p(key, value))
	return l
}

func (l ZapLog) Uint8s(key string, value []uint8) logger.Log {
	l.fields = append(l.fields, zap.Uint8s(key, value))
	return l
}

func (l ZapLog) Uint(key string, value uint) logger.Log {
	l.fields = append(l.fields, zap.Uint(key, value))
	return l
}

func (l ZapLog) Uintp(key string, value *uint) logger.Log {
	l.fields = append(l.fields, zap.Uintp(key, value))
	return l
}

func (l ZapLog) Uints(key string, value []uint) logger.Log {
	l.fields = append(l.fields, zap.Uints(key, value))
	return l
}

func (l ZapLog) Uintptr(key string, value uintptr) logger.Log {
	l.fields = append(l.fields, zap.Uintptr(key, value))
	return l
}

func (l ZapLog) Uintptrp(key string, value *uintptr) logger.Log {
	l.fields = append(l.fields, zap.Uintptrp(key, value))
	return l
}

func (l ZapLog) Uintptrs(key string, value []uintptr) logger.Log {
	l.fields = append(l.fields, zap.Uintptrs(key, value))
	return l
}

func (l ZapLog) Reflect(key string, value interface{}) logger.Log {
	l.fields = append(l.fields, zap.Reflect(key, value))
	return l
}

func (l ZapLog) Namespace(key string) logger.Log {
	l.fields = append(l.fields, zap.Namespace(key))
	return l
}

func (l ZapLog) Stringer(key string, value fmt.Stringer) logger.Log {
	l.fields = append(l.fields, zap.Stringer(key, value))
	return l
}

func (l ZapLog) Any(key string, value interface{}) logger.Log {
	l.fields = append(l.fields, zap.Any(key, value))
	return l
}

func (l ZapLog) Skip() logger.Log {
	l.fields = append(l.fields, zap.Skip())
	return l
}
