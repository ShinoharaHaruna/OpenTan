package utils

import "testing"

func TestFileExist(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "TestFileExist", args: args{"file.go"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileExist(tt.args.path); got != tt.want {
				t.Errorf("FileExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilesExist(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "TestFilesExist", args: args{[]string{"file.go", "file_test.go"}}, want: true},
		{name: "NotAllFilesExist", args: args{[]string{"file.go", "file_test.go", "NonExists.go"}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilesExist(tt.args.paths...); got != tt.want {
				t.Errorf("FilesExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
