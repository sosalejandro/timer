package timer

import (
	"sync"
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	type args struct {
		duration time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Timer doesn't exceeds 6 hours",
			args:    args{duration: 5*time.Hour + 59*time.Minute},
			wantErr: false,
		},
		{
			name:    "Timer exceeds 6 hours",
			args:    args{duration: 6*time.Hour + 1*time.Minute},
			wantErr: true,
		},
		{
			name:    "Timer equals 6 hours",
			args:    args{duration: 6 * time.Hour},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewTimer(tt.args.duration)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTimer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTimer_Blocked(t1 *testing.T) {
	type args struct {
		duration time.Duration
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Timer is blocked",
			args: args{duration: 1 * time.Nanosecond},
			want: true,
		},
		{
			name: "Timer is not blocked",
			args: args{duration: 1 * time.Hour},
			want: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t, err := NewTimer(tt.args.duration)

			if err != nil {
				t1.Errorf("NewTimer() error = %v", err)
			}
			// delay enough to await for NewTimer to run
			time.Sleep(1 * time.Millisecond)
			if got := t.Blocked(); got != tt.want {
				t1.Errorf("Blocked() = %v, want %v", got, tt.want)
			}
		})
	}
}

//goland:noinspection ALL,GoVetCopyLock,GoVetCopyLock
func TestTimer_Reset(t1 *testing.T) {
	type fields struct {
		mu       sync.Mutex
		blocked  bool
		resetCh  chan struct{}
		duration time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Timer is blocked after reset and duration is 1 nanosecond",
			fields: fields{
				mu:       sync.Mutex{},
				blocked:  false,
				resetCh:  make(chan struct{}),
				duration: 1 * time.Nanosecond,
			},
			want: true,
		},
		{
			name: "Timer is not blocked after reset and duration is 1 hour",
			fields: fields{
				mu:       sync.Mutex{},
				blocked:  true,
				resetCh:  make(chan struct{}),
				duration: 1 * time.Hour,
			},
			want: false,
		},
	}
	//goland:noinspection GoVetCopyLock
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Timer{
				mu:       tt.fields.mu,
				blocked:  tt.fields.blocked,
				resetCh:  tt.fields.resetCh,
				duration: tt.fields.duration,
			}
			go t.run()
			//time.Sleep(100 * time.Millisecond)
			t.Reset()
			//time.Sleep(100 * time.Millisecond)

			// validate timer blocked after reset
			if got := t.Blocked(); got != tt.want {
				t1.Errorf("After the Reset() call, Blocked() = %v, want %v", got, tt.want)
			}
		})
	}
}

//goland:noinspection GoVetCopyLock,GoVetCopyLock
func TestTimer_Stop(t1 *testing.T) {
	type fields struct {
		mu       sync.Mutex
		blocked  bool
		resetCh  chan struct{}
		stopCh   chan struct{}
		duration time.Duration
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Timer is stopped on a non blocked channel",
			fields: fields{
				mu:       sync.Mutex{},
				blocked:  false,
				resetCh:  make(chan struct{}),
				stopCh:   make(chan struct{}),
				duration: 1 * time.Hour,
			},
		},
		{
			name: "Timer is stopped on a blocked channel",
			fields: fields{
				mu:       sync.Mutex{},
				blocked:  true,
				resetCh:  make(chan struct{}),
				stopCh:   make(chan struct{}),
				duration: 1 * time.Hour,
			},
		},
	}
	//goland:noinspection GoVetCopyLock
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Timer{
				mu:       tt.fields.mu,
				blocked:  tt.fields.blocked,
				resetCh:  tt.fields.resetCh,
				stopCh:   tt.fields.stopCh,
				duration: tt.fields.duration,
			}
			t.Stop()
			// delay 1 millisecond to await for Stop() to run
			time.Sleep(100 * time.Millisecond)

			if _, ok := <-t.resetCh; ok {
				t1.Errorf("Reset channel is not closed")
			}
			if _, ok := <-t.stopCh; ok {
				t1.Errorf("Stop channel is not closed")
			}
		})
	}
}
