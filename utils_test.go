package main

import (
	"reflect"
	"testing"
)

func Test_mapp(t *testing.T) {
	type args struct {
		target []int
		f      func(int) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "intから変換",
			args: args{
				target: []int{1, 2, 3},
				f: func(i int) string {
					switch i {
					case 1:
						return "one"
					case 2:
						return "two"
					default:
						return "unknown"
					}
				},
			},
			want: []string{"one", "two", "unknown"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapp(tt.args.target, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filter(t *testing.T) {
	type args struct {
		target []int
		f      func(int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "filter even numbers",
			args: args{
				target: []int{1, 2, 3, 4, 5, 6, 7},
				f: func(i int) bool {
					return i%2 == 0
				},
			},
			want: []int{2, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filter(tt.args.target, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_counts(t *testing.T) {
	type args struct {
		target []int
		f      func(int) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "偶数を数える",
			args: args{
				target: []int{1, 2, 3, 4, 5, 6, 7, 8},
				f: func(i int) bool {
					return i%2 == 0
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := count(tt.args.target, tt.args.f); got != tt.want {
				t.Errorf("count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_span(t *testing.T) {
	type args struct {
		target []int
		f      func(int) bool
	}
	tests := []struct {
		name           string
		args           args
		wantTrueSlice  []int
		wantFalseSlice []int
	}{
		{
			name: "偶数と奇数に分ける",
			args: args{
				target: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				f: func(i int) bool {
					return i%2 == 0
				},
			},
			wantTrueSlice:  []int{2, 4, 6, 8},
			wantFalseSlice: []int{1, 3, 5, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTrueSlice, gotFalseSlice := span(tt.args.target, tt.args.f)
			if !reflect.DeepEqual(gotTrueSlice, tt.wantTrueSlice) {
				t.Errorf("span() gotTrueSlice = %v, want %v", gotTrueSlice, tt.wantTrueSlice)
			}
			if !reflect.DeepEqual(gotFalseSlice, tt.wantFalseSlice) {
				t.Errorf("span() gotFalseSlice = %v, want %v", gotFalseSlice, tt.wantFalseSlice)
			}
		})
	}
}

func Test_flatMap(t *testing.T) {
	type args struct {
		target []int
		f      func(int) []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "正常系",
			args: args{
				target: []int{1, 2, 3, 4, 5, 6, 7, 8},
				f: func(i int) []int {
					return []int{i * 2}
				},
			},
			want: []int{2, 4, 6, 8, 10, 12, 14, 16},
		},
		{
			name: "正常系：複数項目を返す",
			args: args{
				target: []int{1, 2, 3},
				f: func(i int) []int {
					return []int{i, i}
				},
			},
			want: []int{1, 1, 2, 2, 3, 3},
		},
		{
			name: "空配列を渡したケース",
			args: args{
				target: []int{1, 2, 3, 4, 5},
				f: func(i int) []int {
					println("call!")
					return []int{}
				},
			},
			want: []int{},
		},
		{
			name: "nilを渡したケース",
			args: args{
				target: []int{1, 2, 3, 4, 5},
				f: func(i int) []int {
					println("call!!")
					return nil
				},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flatMap(tt.args.target, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flatMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_foldLeft(t *testing.T) {
	fn := func(z string, n int) string {
		switch n {
		case 1:
			return z + "1"
		case 2:
			return z + "2"
		default:
			return z
		}
	}

	type args struct {
		target []int
		z      string
		f      func(string, int) string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "空配列のときは初期値が返るだけ",
			args: args{
				target: []int{},
				z:      "hello",
				f:      fn,
			},
			want: "hello",
		},
		{
			name: "要素1の配列のとき",
			args: args{
				target: []int{1},
				z:      "hello",
				f:      fn,
			},
			want: "hello1",
		},
		{
			name: "要素複数の配列のとき",
			args: args{
				target: []int{1, 2, 3, 4},
				z:      "hello",
				f:      fn,
			},
			want: "hello12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := foldLeft(tt.args.target, tt.args.z, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("foldLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sum(t *testing.T) {
	type args struct {
		target []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "合計値が返るか",
			args: args{
				target: []int{1, 2, 3, 4, 5},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.target); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sum1(t *testing.T) {
	type args struct {
		target []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "合計値が返ること",
			args: args{
				target: []float64{1.1, 2.1, 3.1, 4.1, 5.5},
			},
			want: 15.9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.target); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		target []int
		elem   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "含まれるケース",
			args: args{
				target: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				elem:   8,
			},
			want: true,
		},
		{
			name: "含まれないケース",
			args: args{
				target: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				elem:   10,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.target, tt.args.elem); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
