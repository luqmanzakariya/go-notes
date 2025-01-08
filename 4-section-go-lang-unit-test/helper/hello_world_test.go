package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// func TestHelloWorld(t *testing.T){
// 	result := HelloWorld("Eko")

// 	if result != "Hello Eko" {
// 		// unit test failed
// 		panic("Result is not Hello Eko")
// 	}
// }

// func TestHelloWorldKhannedy(t *testing.T){
// 	result := HelloWorld("Khannedy")

// 	if result != "Hello Khannedy" {
// 		// unit test failed
// 		panic("Result is not Hello Khannedy")
// 	}
// }

// # Test t.Fail()
// func TestHelloWorldEko(t *testing.T){
// 	result := HelloWorld("Eko")

// 	if result != "Hello Eko asfdf" {
// 		// error
// 		t.Fail()
// 	}

// 	fmt.Println("TestHelloWorldEko Done") // Kondisi ini tetap dieksekusi walaupun terjadi error
// }

// # Test t.FailNow()
// func TestHelloWorldKhannedy(t *testing.T){
// 	result := HelloWorld("Khannedy")

// 	if result != "Hello Khannedy sfasdf" {
// 		// unit test failed
// 		t.FailNow()
// 	}

// 	fmt.Println("TestHelloWorldKhannedy Done") // Fungsi ini tidak akan dijalankan
// }

// Test t.Error()
func TestHelloWorldEko(t *testing.T) {
	result := HelloWorld("Eko")

	if result != "Hello Eko sdasdf" {
		// error
		t.Error("Result must be 'Hello Eko'")
	}

	fmt.Println("TestHelloWorldEko Done") // Kondisi ini tetap dieksekusi walaupun terjadi error
}

// # Test t.Fatal()
func TestHelloWorldKhannedy(t *testing.T) {
	result := HelloWorld("Khannedy")

	if result != "Hello Khannedy sfasdf" {
		// unit test failed
		t.Fatal("Result must be 'Hello Khannedy'")
	}

	fmt.Println("TestHelloWorldKhannedy Done") // Fungsi ini tidak akan dijalankan
}

// # Test Assertion
func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko sadfasdf", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Assert Done")
}

// # Test Require
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko sadfasdf", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Assert Done")
}

// # Skip Test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test tidak bisa jalan di Mac")
	}

	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result)
}

// # Before and After Test
func TestMain(m *testing.M) {
	// Before Unit Test
	fmt.Println("SEBELUM UNIT TEST")

	m.Run() // eksekusi semua unit test

	// After Unit Test
	fmt.Println("SETELAH UNIT TEST")
}

// # Membuat Sub Test
func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		require.Equal(t, "Hello Eko", result)
	})

	t.Run("Kurniawan", func(t *testing.T) {
		result := HelloWorld("Kurniawan")
		require.Equal(t, "Hello Kurniawan", result)
	})
}

// # Table Test
func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Eko",
			request:  "Eko",
			expected: "Hello Eko",
		},
		{
			name:     "Kurniawan",
			request:  "Kurniawan",
			expected: "Hello Kurniawan",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

// # Benchmark Test 1
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}

// # Benchmark Test 1
func BenchmarkHelloWorldKurniawan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kurniawan")
	}
}

// # Sub Benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Eko", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Eko")
		}
	})
	b.Run("Kurniawan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Kurniawan")
		}
	})
}

// # Benchmark Table
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Eko",
			request: "Eko",
		}, {
			name:    "Kurniawan",
			request: "Kurniawan",
		},
		{
			name:    "EkoKurniawanKhannedy",
			request: "Eko Kurniawan Khannedy",
		},
	}

	for _,benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i:= 0; i<b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
