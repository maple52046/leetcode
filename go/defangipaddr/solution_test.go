package defangipaddr

import "testing"

func TestDefangIPaddr(t *testing.T) {
	tests := []struct {
		name    string
		address string
		want    string
	}{
		{
			name:    "example 1",
			address: "1.1.1.1",
			want:    "1[.]1[.]1[.]1",
		},
		{
			name:    "example 2",
			address: "255.100.50.0",
			want:    "255[.]100[.]50[.]0",
		},
		{
			name:    "example 3",
			address: "192.168.1.1",
			want:    "192[.]168[.]1[.]1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := defangIPaddr(tt.address); got != tt.want {
				t.Errorf("defangIPaddr(%v) = %v, want %v",
					tt.address, got, tt.want)
			}
		})
	}
}
