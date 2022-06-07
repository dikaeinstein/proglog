package log

import (
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndex(t *testing.T) {
	f, err := ioutil.TempFile(os.TempDir(), "index_test")
	require.NoError(t, err)
	defer os.Remove(f.Name())
	c := Config{}
	c.Segment.MaxIndexBytes = 1024
	idx, err := newIndex(f, c)
	require.NoError(t, err)
	_, _, err = idx.Read(-1)
	require.Error(t, err)
	require.Equal(t, f.Name(), idx.Name())
	entries := []struct {
		Off uint32
		Pos uint64
	}{
		{Off: 0, Pos: 0},
		{Off: 1, Pos: 10},
	}
	for _, want := range entries {
		err = idx.Write(want.Off, want.Pos)
		require.NoError(t, err)
		_, pos, err := idx.Read(int64(want.Off))
		require.NoError(t, err)
		require.Equal(t, want.Pos, pos)
	}
	// index and scanner should error when reading past existing entries
	_, _, err = idx.Read(int64(len(entries)))
	require.Equal(t, io.EOF, err)
	_ = idx.Close()
	// index should build its state from the existing file
	f, _ = os.OpenFile(f.Name(), os.O_RDWR, 0600)
	idx, err = newIndex(f, c)
	require.NoError(t, err)
	off, pos, err := idx.Read(-1)
	require.NoError(t, err)
	require.Equal(t, uint32(1), off)
	require.Equal(t, entries[1].Pos, pos)
}

func Test_newIndex(t *testing.T) {
	type args struct {
		f *os.File
		c Config
	}
	tests := []struct {
		name    string
		args    args
		want    *index
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newIndex(tt.args.f, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("newIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_index_Close(t *testing.T) {
	tests := []struct {
		name    string
		i       *index
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.i.Close(); (err != nil) != tt.wantErr {
				t.Errorf("index.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_index_Read(t *testing.T) {
	type args struct {
		in int64
	}
	tests := []struct {
		name    string
		i       *index
		args    args
		wantOut uint32
		wantPos uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, gotPos, err := tt.i.Read(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("index.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut != tt.wantOut {
				t.Errorf("index.Read() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
			if gotPos != tt.wantPos {
				t.Errorf("index.Read() gotPos = %v, want %v", gotPos, tt.wantPos)
			}
		})
	}
}

func Test_index_Write(t *testing.T) {
	type args struct {
		off uint32
		pos uint64
	}
	tests := []struct {
		name    string
		i       *index
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.i.Write(tt.args.off, tt.args.pos); (err != nil) != tt.wantErr {
				t.Errorf("index.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_index_Name(t *testing.T) {
	tests := []struct {
		name string
		i    *index
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Name(); got != tt.want {
				t.Errorf("index.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
