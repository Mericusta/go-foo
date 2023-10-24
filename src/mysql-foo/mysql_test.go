package mysqlfoo

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenMySQLDatabase(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect(tt.args.URI); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenMySQLDatabase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prayResult_GetItem(t *testing.T) {
	tests := []struct {
		name string
		p    *prayResult
		want map[int32]int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.GetItem(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prayResult.GetItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBatchInsertPrayRecordData(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BatchInsertPrayRecordData()
		})
	}
}

func Test_insertOrUpdateFoo(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				url: MYSQL_DATABASE_FOO,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertOrUpdateFoo(tt.args.url)
		})
	}
}

func Test_concurrentlyInsertOrUpdateFoo(t *testing.T) {
	type args struct {
		url            string
		goroutineCount int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test case 1",
			args{
				url:            MYSQL_DATABASE_FOO,
				goroutineCount: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			concurrentlyInsertOrUpdateFoo(tt.args.url, tt.args.goroutineCount)
		})
	}
}
