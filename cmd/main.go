package rnnoise

import (
	"github.com/kiryuhakipyatok/rnnoise/internal/binding"
)

type RNNoise struct {
	state *binding.State
}

func NewRNNoise() *RNNoise {
	state := binding.New()

	rnn := RNNoise{
		state: state,
	}

	return &rnn
}

func (rnn *RNNoise) Denoise(clear, noisy []float32) (float32, error) {
	vad, err := rnn.state.ProcessFrame(clear, noisy)
	if err != nil {
		return 0, err
	}
	return vad, nil
}

func (rnn *RNNoise) Close() error {
	return rnn.state.Close()
}

func main() {}
