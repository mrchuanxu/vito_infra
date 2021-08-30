package util

import (
	"testing"
	"time"
)

func TestGetNumberByStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{str: "weibo/12344566/hahah.html"}, "12344566"},
		{"test2", args{str: "weibo/12344/1222/hahah.html"}, "12344"},
		{"test3", args{str: "http://weibo/2/12344/1222/hahah.html"}, "2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNumberByStr(tt.args.str); got != tt.want {
				t.Errorf("GetNumberByStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_TimeStr(t *testing.T) {
	loc, _ := time.LoadLocation("Local")
	fieldTime, _ := time.ParseInLocation("20060102", "20211221", loc)
	t.Log(fieldTime)
}
