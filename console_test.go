package gonsole

import (
	"testing"
)

func Test_color_Foreground(t *testing.T) {
	tests := []struct {
		name string
		c    color
		want string
	}{
		{name: "0", c: color(0), want: "\x1b[38;5;0m"},
		{name: "10", c: color(10), want: "\x1b[38;5;10m"},
		{name: "100", c: color(100), want: "\x1b[38;5;100m"},
		{name: "255", c: color(255), want: "\x1b[38;5;255m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Foreground(); got != tt.want {
				t.Errorf("color.Foreground() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_color_Background(t *testing.T) {
	tests := []struct {
		name string
		c    color
		want string
	}{
		{name: "0", c: color(0), want: "\x1b[48;5;0m"},
		{name: "10", c: color(10), want: "\x1b[48;5;10m"},
		{name: "100", c: color(100), want: "\x1b[48;5;100m"},
		{name: "255", c: color(255), want: "\x1b[48;5;255m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Background(); got != tt.want {
				t.Errorf("color.Background() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_color_Underline(t *testing.T) {
	tests := []struct {
		name string
		c    color
		want string
	}{
		{name: "0", c: color(0), want: "\x1b[58;5;0m"},
		{name: "10", c: color(10), want: "\x1b[58;5;10m"},
		{name: "100", c: color(100), want: "\x1b[58;5;100m"},
		{name: "255", c: color(255), want: "\x1b[58;5;255m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Underline(); got != tt.want {
				t.Errorf("color.Underline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForeground(t *testing.T) {
	type args struct {
		c color
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0", args: args{c: color(0)}, want: "\x1b[38;5;0m"},
		{name: "10", args: args{c: color(10)}, want: "\x1b[38;5;10m"},
		{name: "100", args: args{c: color(100)}, want: "\x1b[38;5;100m"},
		{name: "255", args: args{c: color(255)}, want: "\x1b[38;5;255m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Foreground(tt.args.c); got != tt.want {
				t.Errorf("Foreground() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBackground(t *testing.T) {
	type args struct {
		c color
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0", args: args{c: color(0)}, want: "\x1b[48;5;0m"},
		{name: "10", args: args{c: color(10)}, want: "\x1b[48;5;10m"},
		{name: "100", args: args{c: color(100)}, want: "\x1b[48;5;100m"},
		{name: "255", args: args{c: color(255)}, want: "\x1b[48;5;255m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Background(tt.args.c); got != tt.want {
				t.Errorf("Background() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnderline(t *testing.T) {
	type args struct {
		c color
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0", args: args{c: color(0)}, want: "\x1b[58;5;0m"},
		{name: "10", args: args{c: color(10)}, want: "\x1b[58;5;10m"},
		{name: "100", args: args{c: color(100)}, want: "\x1b[58;5;100m"},
		{name: "255", args: args{c: color(255)}, want: "\x1b[58;5;255m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Underline(tt.args.c); got != tt.want {
				t.Errorf("Underline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRGBForeground(t *testing.T) {
	type args struct {
		r int
		g int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "minus color", args: args{r: -1, g: -1, b: -1}, want: "", wantErr: true},
		{name: "0", args: args{r: 0, g: 0, b: 0}, want: "\x1b[38;2;0;0;0m", wantErr: false},
		{name: "255", args: args{r: 255, g: 255, b: 255}, want: "\x1b[38;2;255;255;255m", wantErr: false},
		{name: "over", args: args{r: 256, g: 256, b: 256}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RGBForeground(tt.args.r, tt.args.g, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("RGBForeground() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RGBForeground() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRGBBackground(t *testing.T) {
	type args struct {
		r int
		g int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "minus color", args: args{r: -1, g: -1, b: -1}, want: "", wantErr: true},
		{name: "0", args: args{r: 0, g: 0, b: 0}, want: "\x1b[48;2;0;0;0m", wantErr: false},
		{name: "255", args: args{r: 255, g: 255, b: 255}, want: "\x1b[48;2;255;255;255m", wantErr: false},
		{name: "over", args: args{r: 256, g: 256, b: 256}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RGBBackground(tt.args.r, tt.args.g, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("RGBBackground() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RGBBackground() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRGBUnderline(t *testing.T) {
	type args struct {
		r int
		g int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "minus color", args: args{r: -1, g: -1, b: -1}, want: "", wantErr: true},
		{name: "0", args: args{r: 0, g: 0, b: 0}, want: "\x1b[58;2;0;0;0m", wantErr: false},
		{name: "255", args: args{r: 255, g: 255, b: 255}, want: "\x1b[58;2;255;255;255m", wantErr: false},
		{name: "over", args: args{r: 256, g: 256, b: 256}, want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RGBUnderline(tt.args.r, tt.args.g, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("RGBUnderline() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RGBUnderline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDemo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "general"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Demo(); len(got) < 1 {
				t.Errorf("Demo() = %v, want not empty", got)
			}
		})
	}
}
