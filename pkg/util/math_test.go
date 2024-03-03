package util

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestFormatSize(t *testing.T) {
	for _, tc := range []struct {
		name string
		size int64
		res  string
	}{
		{
			name: "0",
			size: 0,
			res:  "0.00 B/s",
		},
		{
			name: "512",
			size: 512,
			res:  "512.00 B/s",
		},
		{
			name: "1024",
			size: 1024,
			res:  "1.00 KB/s",
		},
		{
			name: "1048576",
			size: 1048576,
			res:  "1.00 MB/s",
		},
		{
			name: "1073741824",
			size: 1073741824,
			res:  "1.00 GB/s",
		},
		{
			name: "1023",
			size: 1023,
			res:  "1023.00 B/s",
		},
		{
			name: "1025",
			size: 1025,
			res:  "1.00 KB/s",
		},
		{
			name: "1,099,511,627,776", // 1TB
			size: 1099511627776,
			res:  "1024.00 GB/s",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			s := FormatSize(tc.size)

			assert.Equal(t, s, tc.res)
		})
	}
}
