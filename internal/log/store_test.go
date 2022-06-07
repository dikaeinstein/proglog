package log

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	write = []byte("hello world")
	width = uint64(len(write)) + lenWidth
)

func TestStoreAppendRead(t *testing.T) {
	f, err := ioutil.TempFile("", "store_append_read_test")
	require.NoError(t, err)
	defer os.Remove(f.Name())
	s, err := newStore(f)
	require.NoError(t, err)
	testAppend(t, s)
	testRead(t, s)
	testReadAt(t, s)
	s, err = newStore(f)
	require.NoError(t, err)
	testRead(t, s)
}

func TestStoreClose(t *testing.T) {
	f, err := ioutil.TempFile("", "store_close_test")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	s, err := newStore(f)
	require.NoError(t, err)

	_, _, err = s.Append(write)
	require.NoError(t, err)
	f, beforeSize, err := openFile(f.Name())
	require.NoError(t, err)
	err = s.Close()
	require.NoError(t, err)
	_, afterSize, err := openFile(f.Name())
	require.NoError(t, err)
	require.True(t, afterSize > beforeSize)
}

func testAppend(t *testing.T, s *store) {
	t.Helper()
	for i := uint64(1); i < 4; i++ {
		n, pos, err := s.Append(write)
		require.NoError(t, err)
		require.Equal(t, pos+n, width*i)
	}
}

func testRead(t *testing.T, s *store) {
	t.Helper()
	var pos uint64
	for i := uint64(1); i < 4; i++ {
		read, err := s.Read(pos)
		require.NoError(t, err)
		require.Equal(t, write, read)
		pos += width
	}
}

func testReadAt(t *testing.T, s *store) {
	t.Helper()
	for i, off := uint64(1), int64(0); i < 4; i++ {
		b := make([]byte, lenWidth)
		n, err := s.ReadAt(b, off)
		require.NoError(t, err)
		require.Equal(t, lenWidth, n)
		off += int64(n)
		size := enc.Uint64(b)
		b = make([]byte, size)
		n, err = s.ReadAt(b, off)
		require.NoError(t, err)
		require.Equal(t, write, b)
		require.Equal(t, int(size), n)
		off += int64(n)
	}
}

func openFile(name string) (file *os.File, size int64, err error) {
	f, err := os.OpenFile(
		name,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		return nil, 0, err
	}
	fi, err := f.Stat()
	if err != nil {
		return nil, 0, err
	}
	return f, fi.Size(), nil
}

func Test_newStore(t *testing.T) {
	type args struct {
		f *os.File
	}
	tests := []struct {
		name    string
		args    args
		want    *store
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newStore(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("newStore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_store_Append(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		s       *store
		args    args
		wantN   uint64
		wantPos uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, gotPos, err := tt.s.Append(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("store.Append() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("store.Append() gotN = %v, want %v", gotN, tt.wantN)
			}
			if gotPos != tt.wantPos {
				t.Errorf("store.Append() gotPos = %v, want %v", gotPos, tt.wantPos)
			}
		})
	}
}

func Test_store_Read(t *testing.T) {
	type args struct {
		pos uint64
	}
	tests := []struct {
		name    string
		s       *store
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Read(tt.args.pos)
			if (err != nil) != tt.wantErr {
				t.Errorf("store.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("store.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_store_ReadAt(t *testing.T) {
	type args struct {
		p   []byte
		off int64
	}
	tests := []struct {
		name    string
		s       *store
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ReadAt(tt.args.p, tt.args.off)
			if (err != nil) != tt.wantErr {
				t.Errorf("store.ReadAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("store.ReadAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_store_Close(t *testing.T) {
	tests := []struct {
		name    string
		s       *store
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Close(); (err != nil) != tt.wantErr {
				t.Errorf("store.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_store_Name(t *testing.T) {
	tests := []struct {
		name string
		s    *store
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Name(); got != tt.want {
				t.Errorf("store.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
