package goBench

import (
	"bytes"
	"testing"
)

var sampleString = `<html>
	<head>
		<title>sample title</title>
	</head>
	<body>
		<h1>sample head</h1>
		<p>sample paragraph</p>
	</body>
</html>`

var sampleBytes = []byte(sampleString)

func Benchmark(b *testing.B) {
	b.Run("Write", func(b *testing.B) {
		b.Run("Bytes", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				buf := bytes.NewBuffer(nil)
				_, _ = buf.Write(sampleBytes)
				_ = buf.Bytes()
			}
		})
		b.Run("String", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				buf := bytes.NewBuffer(nil)
				_, _ = buf.Write(sampleBytes)
				_ = buf.String()
			}
		})
	})

	b.Run("WriteString", func(b *testing.B) {
		b.Run("Bytes", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				buf := bytes.NewBuffer(nil)
				_, _ = buf.WriteString(sampleString)
				_ = buf.Bytes()
			}
		})
		b.Run("String", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				buf := bytes.NewBuffer(nil)
				_, _ = buf.WriteString(sampleString)
				_ = buf.String()
			}
		})
	})

	b.Run("WriteByte", func(b *testing.B) {
		b.Run("Bytes", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				buf := bytes.NewBuffer(nil)
				for j := 0; j < len(sampleBytes); j++ {
					_ = buf.WriteByte(sampleBytes[j])
				}
				_ = buf.Bytes()
			}
		})
		b.Run("String", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				buf := bytes.NewBuffer(nil)
				for j := 0; j < len(sampleBytes); j++ {
					_ = buf.WriteByte(sampleBytes[j])
				}
				_ = buf.String()
			}
		})
	})
}
