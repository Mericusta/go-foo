package mysqlfoo

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func BenchmarkBatchInsertPrayRecordData(b *testing.B) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for index := 0; index != len(tests); index++ {
			BatchInsertPrayRecordData()
		}
	}
}

func BenchmarkOpenMySQLDatabase(b *testing.B) {
	type args struct {
		URI string
	}
	tests := []struct {
		name string
		args args
		want *sql.DB
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			connect(tt.args.URI)
		}
	}
}

func Benchmark_prayResult_GetItem(b *testing.B) {
	tests := []struct {
		name string
		p    *prayResult
		want map[int32]int64
	}{
		// TODO: Add test cases.
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			tt.p.GetItem()
		}
	}
}
