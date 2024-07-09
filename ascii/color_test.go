package ascii

import (
	"testing"
)

func Test_rgbToANSI(t *testing.T) {
	tests := []struct {
		name    string
		rgb     string
		want    string
		wantErr bool
	}{
		{
			name:    "Valid RGB String - Full Red",
			rgb:     "rgb(255, 0, 0)",
			want:    "38;2;255;0;0",
			wantErr: false,
		},
		{
			name:    "Valid RGB String - Full Green",
			rgb:     "rgb(0, 255, 0)",
			want:    "38;2;0;255;0",
			wantErr: false,
		},
		{
			name:    "Invalid RGB String - Out of Range Values",
			rgb:     "rgb(300, -20, 500)",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := rgbToANSI(tt.rgb)
			if (err != nil) != tt.wantErr {
				t.Errorf("rgbToANSI() error = %v and (err != nil) = %v, wantErr %v", err, (err != nil), tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("rgbToANSI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hexToANSI(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		want    string
		wantErr bool
	}{
		{
			name:    "Valid Hex String - Bright Pink",
			hex:     "#ff69b4",
			want:    "38;2;255;105;180",
			wantErr: false,
		},
		{
			name:    "Valid Hex String - Light Blue",
			hex:     "#add8e6",
			want:    "38;2;173;216;230",
			wantErr: false,
		},
		{
			name:    "Invalid Hex String - Invalid Characters",
			hex:     "#zzzzzz",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hexToANSI(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("hexToANSI() error = %v (err != nil) = %v, wantErr %v", err, (err != nil), tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hexToANSI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseColor(t *testing.T) {
	tests := []struct {
		name    string
		color   string
		want    string
		wantErr bool
	}{
		{
			name:    "Hex Color Yellow",
			color:   "#ffff00",
			want:    "38;2;255;255;0",
			wantErr: false,
		},
		{
			name:    "RGB Color Coral",
			color:   "rgb(240, 128, 128)",
			want:    "38;2;240;128;128",
			wantErr: false,
		},
		{
			name:    "Common Color Name - Blue",
			color:   "blue",
			want:    "34",
			wantErr: false,
		},
		{
			name:    "Invalid Hex Color",
			color:   "#xxx",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseColor(tt.color)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
