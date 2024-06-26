package strconv2

import (
	"github.com/go-kid/strings2"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestParseStringSlice(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				val: "[hello,world,foo,bar]",
			},
			want:    []string{"hello", "world", "foo", "bar"},
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				val: "[]",
			},
			want:    []string{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseStringSlice(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseAnySlice(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name    string
		args    args
		want    []any
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				val: "[abc,123,true,false,1.111]",
			},
			want:    []any{"abc", float64(123), true, false, 1.111},
			wantErr: false,
		},
		{
			name: "1",
			args: args{
				val: `["abc",123,true,false,1.111]`,
			},
			want:    []any{"abc", float64(123), true, false, 1.111},
			wantErr: false,
		},
		{
			name: "1",
			args: args{
				val: `[["abc"],[123,true],[false,1.111]]`,
			},
			want:    []any{[]any{"abc"}, []any{float64(123), true}, []any{false, 1.111}},
			wantErr: false,
		},
		{
			name: "1",
			args: args{
				val: `[{"a":"b"},{"c":123}]`,
			},
			want: []any{
				map[string]any{
					"a": "b",
				},
				map[string]any{
					"c": float64(123),
				},
			},
			wantErr: false,
		},
		{
			name: "1",
			args: args{
				val: `[map[a:b],map[c:-123]]`,
			},
			want: []any{
				map[string]any{
					"a": "b",
				},
				map[string]any{
					"c": float64(-123),
				},
			},
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				val: "[]",
			},
			want:    []any{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAnySlice(tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAnySlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParseSlice(t *testing.T) {
	t.Run("ParseString", func(t *testing.T) {
		parse, err := ParseSlice[string]("[hello,world,foo,bar]")
		assert.NoError(t, err)
		assert.Equal(t, []string{"hello", "world", "foo", "bar"}, parse)
	})
	t.Run("ParseInt", func(t *testing.T) {
		i, err := ParseSlice[int]("[123,456]")
		assert.NoError(t, err)
		assert.Equal(t, []int{123, 456}, i)
		i8, err := ParseSlice[int8]("[1,2,3]")
		assert.NoError(t, err)
		assert.Equal(t, []int8{1, 2, 3}, i8)
		ui, err := ParseSlice[uint]("[123,456]")
		assert.NoError(t, err)
		assert.Equal(t, []uint{123, 456}, ui)
		ui8, err := ParseSlice[uint8]("[1,2,3]")
		assert.NoError(t, err)
		assert.Equal(t, []uint8{1, 2, 3}, ui8)
	})
	t.Run("ParseFloat", func(t *testing.T) {
		f64, err := ParseSlice[float64]("[123.456,789.01]")
		assert.NoError(t, err)
		assert.Equal(t, []float64{123.456, 789.01}, f64)
		f32, err := ParseSlice[float32]("[123.456,789.01]")
		assert.NoError(t, err)
		assert.Equal(t, []float32{123.456, 789.01}, f32)
	})
	t.Run("ParseBool", func(t *testing.T) {
		parse, err := ParseSlice[bool]("[true,false,true]")
		assert.NoError(t, err)
		assert.Equal(t, []bool{true, false, true}, parse)
	})
}

func Test_splitSlicePart(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				val: `{"x-request-id":"123"},{"x-cross-origin":["*"]},{"x-allowed-method":["POST","GET"]}`,
			},
			want: []string{`{"x-request-id":"123"}`, `{"x-cross-origin":["*"]}`, `{"x-allowed-method":["POST","GET"]}`},
		},
		{
			name: "2",
			args: args{
				val: "1,2,3",
			},
			want: []string{"1", "2", "3"},
		},
		{
			name: "2",
			args: args{
				val: "1,,2,3",
			},
			want: []string{"1", "", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, strings2.SplitWithConfig(tt.args.val, sliceSplitConfig), "splitSlicePart(%v)", tt.args.val)
		})
	}
}
