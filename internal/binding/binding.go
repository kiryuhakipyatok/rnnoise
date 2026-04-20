package binding

// #cgo CFLAGS: -I${SRCDIR}
// #cgo LDFLAGS: ${SRCDIR}/librnnoise.a -lm
// #include "rnnoise.h"
import "C"
import (
	"errors"
	"runtime"
	"unsafe"
)

const frameSize = 480


type State struct {
	ptr *C.DenoiseState
}


func New() *State {
	cState := C.rnnoise_create((*C.RNNModel)(nil))
	s := &State{ptr: cState}

	runtime.SetFinalizer(s, (*State).Close)
	return s
}

func (s *State) Close() error {
	if s.ptr != nil {
		C.rnnoise_destroy(s.ptr)
		s.ptr = nil
	}
	return nil
}


func (s *State) ProcessFrame(out, in []float32) (float32, error) {
	if len(in) != frameSize || len(out) != frameSize {
		return 0, errors.New("rnnoise: in and out slices must have exactly 480 samples")
	}

	vad := C.rnnoise_process_frame(
		s.ptr,
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&in[0])),
	)

	runtime.KeepAlive(s)

	return float32(vad), nil
}