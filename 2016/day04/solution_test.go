package day04

import "testing"

var in = `aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]`

func TestPart1(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{in, "1514"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := Part1(tt.in); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		room room
		want string
	}{
		{parseRoom("qzmt-zixmtkozy-ivhz-343[something]"), "very encrypted name"},
	}

	for _, tt := range tests {
		t.Run(tt.room.checksum, func(t *testing.T) {
			if got := tt.room.decrypt(); got != tt.want {
				t.Errorf("got: %q; want: %q", got, tt.want)
			}
		})
	}
}
