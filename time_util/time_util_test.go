package time_util

import (
	"reflect"
	"testing"
	"time"
)

func TestWeekFirst(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "test0",
			args: args{
				t: time.Date(2024, 07, 21, 0, 0, 0, 0, time.Local),
			},
			want: time.Date(2024, 07, 21, 0, 0, 0, 0, time.Local),
		},
		{
			name: "test1",
			args: args{
				t: time.Date(2024, 07, 22, 0, 0, 0, 0, time.Local),
			},
			want: time.Date(2024, 07, 21, 0, 0, 0, 0, time.Local),
		},
		{
			name: "test2",
			args: args{
				t: time.Date(2024, 07, 23, 0, 0, 0, 0, time.Local),
			},
			want: time.Date(2024, 07, 21, 0, 0, 0, 0, time.Local),
		},
		{
			name: "test3",
			args: args{
				t: time.Date(2024, 07, 24, 0, 0, 0, 0, time.Local),
			},
			want: time.Date(2024, 07, 21, 0, 0, 0, 0, time.Local),
		},
		{
			name: "test4",
			args: args{
				t: time.Date(2024, 07, 25, 0, 0, 0, 0, time.Local),
			},
			want: time.Date(2024, 07, 21, 0, 0, 0, 0, time.Local),
		},
		{
			name: "test5",
			args: args{
				t: time.Date(2024, 07, 26, 0, 0, 0, 0, time.Local),
			},
			want: time.Date(2024, 07, 21, 0, 0, 0, 0, time.Local),
		},
		{
			name: "test6",
			args: args{
				t: time.Date(2024, 07, 27, 0, 0, 0, 0, time.Local),
			},
			want: time.Date(2024, 07, 21, 0, 0, 0, 0, time.Local),
		},
		{
			name: "test7",
			args: args{
				t: time.Date(2024, 07, 28, 0, 0, 0, 0, time.Local),
			},
			want: time.Date(2024, 07, 28, 0, 0, 0, 0, time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WeekFirst(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WeekFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeekLast(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "test0",
			args: args{
				t: time.Date(2024, 07, 21, 0, 0, 0, 0, time.Local),
			},
			want: time.Date(2024, 07, 27, 0, 0, 0, 0, time.Local),
		},
		{
			name: "test1",
			args: args{
				t: time.Date(2024, 07, 22, 0, 0, 0, 0, time.Local),
			},
			want: time.Date(2024, 07, 27, 0, 0, 0, 0, time.Local),
		},
		{
			name: "test2",
			args: args{
				t: time.Date(2024, 07, 23, 0, 0, 0, 0, time.Local),
			},
			want: time.Date(2024, 07, 27, 0, 0, 0, 0, time.Local),
		},
		{
			name: "test3",
			args: args{
				t: time.Date(2024, 07, 24, 0, 0, 0, 0, time.Local),
			},
			want: time.Date(2024, 07, 27, 0, 0, 0, 0, time.Local),
		},
		{
			name: "test4",
			args: args{
				t: time.Date(2024, 07, 25, 0, 0, 0, 0, time.Local),
			},
			want: time.Date(2024, 07, 27, 0, 0, 0, 0, time.Local),
		},
		{
			name: "test5",
			args: args{
				t: time.Date(2024, 07, 26, 0, 0, 0, 0, time.Local),
			},
			want: time.Date(2024, 07, 27, 0, 0, 0, 0, time.Local),
		},
		{
			name: "test6",
			args: args{
				t: time.Date(2024, 07, 27, 0, 0, 0, 0, time.Local),
			},
			want: time.Date(2024, 07, 27, 0, 0, 0, 0, time.Local),
		},
		{
			name: "test7",
			args: args{
				t: time.Date(2024, 07, 27, 0, 0, 0, 0, time.Local),
			},
			want: time.Date(2024, 07, 27, 0, 0, 0, 0, time.Local),
		},
		{
			name: "test8",
			args: args{
				t: time.Date(2024, 07, 28, 0, 0, 0, 0, time.Local),
			},
			want: time.Date(2024, 07, 04, 0, 0, 0, 0, time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WeekLast(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WeekLast() = %v, want %v", got, tt.want)
			}
		})
	}
}
