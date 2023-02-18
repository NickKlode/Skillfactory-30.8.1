package storage

import (
	"reflect"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func TestStorage_Tasks(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		taskID   int
		authorID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.Tasks(tt.args.taskID, tt.args.authorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Tasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.Tasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_NewTask(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		t Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.NewTask(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.NewTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_GetListTasksByAuthor(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		authorId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.GetListTasksByAuthor(tt.args.authorId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetListTasksByAuthor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.GetListTasksByAuthor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_GetListTasksByAssignedId(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		assigned int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			got, err := s.GetListTasksByAssignedId(tt.args.assigned)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetListTasksByAssignedId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.GetListTasksByAssignedId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_UpdateTask(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		title   string
		content string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			if err := s.UpdateTask(tt.args.title, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Storage.UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_DeleteTask(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				db: tt.fields.db,
			}
			if err := s.DeleteTask(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Storage.DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
