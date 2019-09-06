package bitmap

import "testing"

var (
	bitmapFortest *bitmap
)

func init() {
	bitmapFortest = NewBitmap(100000)
}

func Test_bitmap_Set(t *testing.T) {

	type args struct {
		index uint
	}
	tests := []struct {
		name    string
		fields  *bitmap
		args    args
		wantErr bool
	}{
		{
			"1",
			bitmapFortest,
			args{index: 1},
			false,
		},
		{
			"1000",
			bitmapFortest,
			args{index: 1000},
			false,
		},
		{
			"1000000",
			bitmapFortest,
			args{index: 1000000},
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fields.Set(tt.args.index); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_bitmap_Check(t *testing.T) {

	type args struct {
		index uint
	}
	tests := []struct {
		name    string
		fields  *bitmap
		args    args
		want    bool
		wantErr bool
	}{
		{
			"1",
			bitmapFortest,
			args{index: 1},
			true,
			false,
		},
		{
			"1000",
			bitmapFortest,
			args{index: 1000},
			true,
			false,
		},
		{
			"1000000",
			bitmapFortest,
			args{index: 1000000},
			false,
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.fields.Check(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Check() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bitmap_Unset(t *testing.T) {

	type args struct {
		index uint
	}
	tests := []struct {
		name    string
		fields  *bitmap
		args    args
		wantErr bool
	}{
		{
			"1",
			bitmapFortest,
			args{index: 1},
			false,
		},
		{
			"1000",
			bitmapFortest,
			args{index: 1000},
			false,
		},
		{
			"1000000",
			bitmapFortest,
			args{index: 1000000},
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := tt.fields.Unset(tt.args.index); (err != nil) != tt.wantErr {
				t.Errorf("Unset() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_bitmap_Check2(t *testing.T) {

	type args struct {
		index uint
	}
	tests := []struct {
		name    string
		fields  *bitmap
		args    args
		want    bool
		wantErr bool
	}{
		{
			"1",
			bitmapFortest,
			args{index: 1},
			false,
			false,
		},
		{
			"1000",
			bitmapFortest,
			args{index: 1000},
			false,
			false,
		},
		{
			"1000000",
			bitmapFortest,
			args{index: 1000000},
			false,
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.fields.Check(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Check() got = %v, want %v", got, tt.want)
			}
		})
	}
}
