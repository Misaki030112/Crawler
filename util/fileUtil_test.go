package util

import "testing"

func TestFileDownloadFromUrl(t *testing.T) {
	type args struct {
		url          string
		wantFileName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "downLoad Test",
			args: args{
				url:          "http://rs.sfacg.com/web/novel/images/NovelCover/Big/2022/08/bb06f93e-9dd1-4372-b9cc-ab8e562a9271.jpg",
				wantFileName: "novel.jpg",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileDownloadFromUrl(tt.args.url, tt.args.wantFileName); got != tt.want {
				t.Errorf("FileDownloadFromUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
