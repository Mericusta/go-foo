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
			if got := OpenMySQLDatabase(tt.args.URI); !reflect.DeepEqual(got, tt.want) {
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
