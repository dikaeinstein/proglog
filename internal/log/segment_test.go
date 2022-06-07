package log

import (
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	api "github.com/dikaeinstein/proglog/api/v1"
	"github.com/stretchr/testify/require"
)

func Test_newSegment(t *testing.T) {
	type args struct {
		dir        string
		baseOffset uint64
		c          Config
	}
	tests := []struct {
		name    string
		args    args
		want    *segment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newSegment(tt.args.dir, tt.args.baseOffset, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("newSegment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_segment_Append(t *testing.T) {
	dir, _ := ioutil.TempDir("", "segment-test")
	defer os.RemoveAll(dir)
	want := &api.Record{Value: []byte("hello world")}
	c := Config{}
	c.Segment.MaxStoreBytes = 1024
	c.Segment.MaxIndexBytes = entWidth * 3
	s, err := newSegment(dir, 16, c)
	require.NoError(t, err)
	require.Equal(t, uint64(16), s.nextOffset, s.nextOffset)
	require.False(t, s.IsMaxed())
	for i := uint64(0); i < 3; i++ {
		off, err := s.Append(want)
		require.NoError(t, err)
		require.Equal(t, 16+i, off)
		got, err := s.Read(off)
		require.NoError(t, err)
		require.Equal(t, want.Value, got.Value)
	}
	_, err = s.Append(want)
	require.Equal(t, io.EOF, err)
	// maxed index
	require.True(t, s.IsMaxed())
	c.Segment.MaxStoreBytes = uint64(len(want.Value) * 3)
	c.Segment.MaxIndexBytes = 1024
	s, err = newSegment(dir, 16, c)
	require.NoError(t, err)
	// maxed store
	require.True(t, s.IsMaxed())
	err = s.Remove()
	require.NoError(t, err)
	s, err = newSegment(dir, 16, c)
	require.NoError(t, err)
	require.False(t, s.IsMaxed())
}

func Test_segment_Read(t *testing.T) {
	type args struct {
		off uint64
	}
	tests := []struct {
		name    string
		s       *segment
		args    args
		want    *api.Record
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Read(tt.args.off)
			if (err != nil) != tt.wantErr {
				t.Errorf("segment.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("segment.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_segment_IsMaxed(t *testing.T) {
	tests := []struct {
		name string
		s    *segment
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsMaxed(); got != tt.want {
				t.Errorf("segment.IsMaxed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_segment_Remove(t *testing.T) {
	tests := []struct {
		name    string
		s       *segment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Remove(); (err != nil) != tt.wantErr {
				t.Errorf("segment.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_segment_Close(t *testing.T) {
	tests := []struct {
		name    string
		s       *segment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Close(); (err != nil) != tt.wantErr {
				t.Errorf("segment.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_nearestMultiple(t *testing.T) {
	type args struct {
		j uint64
		k uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nearestMultiple(tt.args.j, tt.args.k); got != tt.want {
				t.Errorf("nearestMultiple() = %v, want %v", got, tt.want)
			}
		})
	}
}
