package channelfoo

import (
	"testing"
)

func BenchmarkGoChannelBlock(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			GoChannelBlock()
		}
	}
}

func BenchmarkGoroutineExitThenCloseChannel(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			GoroutineExitThenCloseChannel()
		}
	}
}

func BenchmarkGoroutineExitThenCloseChannelSimpleCase(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			GoroutineExitThenCloseChannelSimpleCase()
		}
	}
}

func BenchmarkGoroutineOutputOrder(b *testing.B) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		// TODO: Add test cases.
		{
			"test case",
			0,
			0,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			GoroutineOutputOrder()
		}
	}
}

func BenchmarkGoroutineOutputOrder2(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			GoroutineOutputOrder2()
		}
	}
}

func BenchmarkListenerBlockedChannel(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"test case",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			ListenerBlockedChannel()
		}
	}
}
