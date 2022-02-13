package worldservices

import (
	deep "github.com/patrikeh/go-deep"
	"github.com/patrikeh/go-deep/training"
)

var (
	// DEFINE A BASIC TRAINING NETWORK TO MOVE CREATURES
	BasicTraining = training.Examples{
		{[]float64{0.0, 0.1231, 0.84, 1.0}, []float64{0.0, 0.0, 0.0, 1.0}},
		{[]float64{0.2, 0.1231, 0.9, 0.4}, []float64{0.0, 0.0, 1.0, 0.0}},
		{[]float64{0.2, 0.74, 0.2, 0.4}, []float64{0.0, 1.0, 0.0, 0.0}},
		{[]float64{0.32, 0.2, 0.1, 0.0}, []float64{1.0, 0.0, 0.0, 0.0}},
	}

	LeftOnlyTraining = training.Examples{
		{[]float64{0.0, 0.1231, 0.84, 1.0}, []float64{0.0, 1.0, 0.0, 0.0}},
		{[]float64{0.2, 0.1231, 0.9, 0.4}, []float64{0.0, 1.0, 0.0, 0.0}},
		{[]float64{0.2, 0.74, 0.2, 0.4}, []float64{0.0, 1.0, 0.0, 0.0}},
		{[]float64{0.32, 0.2, 0.1, 0.0}, []float64{0.0, 1.0, 0.0, 0.0}},
	}

	BasicTrainingWOscilation = training.Examples{
		{[]float64{0.0, 0.1231, 0.84, 1.0, 0.1}, []float64{0.0, 0.0, 0.0, 1.0}},
		{[]float64{0.2, 0.1231, 0.9, 0.4, 0.1}, []float64{0.0, 0.0, 1.0, 0.0}},
		{[]float64{0.2, 0.74, 0.2, 0.4, 0.1}, []float64{0.0, 1.0, 0.0, 0.0}},
		{[]float64{0.32, 0.2, 0.1, 0.0, 0.1}, []float64{1.0, 0.0, 0.0, 0.0}},
		{[]float64{0.0, 0.1231, 0.84, 1.0, 0.9}, []float64{0.7, 0.7, 0.7, 1.0}},
		{[]float64{0.2, 0.1231, 0.9, 0.4, 0.9}, []float64{0.8, 0.7, 1.0, 0.8}},
		{[]float64{0.2, 0.74, 0.2, 0.4, 0.9}, []float64{0.77, 1.0, 0.78, 0.7}},
		{[]float64{0.32, 0.2, 0.1, 0.0, 0.9}, []float64{1.0, 0.78, 0.77, 0.66}},
	}
)

func createNetwork(inputs int, layout []int) *deep.Neural {
	return deep.NewNeural(&deep.Config{
		/* Input dimensionality */
		Inputs: inputs,
		/* Two hidden layers consisting of two neurons each, and a single output */
		Layout: layout,
		/* Activation functions: Sigmoid, Tanh, ReLU, Linear */
		Activation: deep.ActivationSigmoid,
		/* Determines output layer activation & loss function:
		ModeRegression: linear outputs with MSE loss
		ModeMultiClass: softmax output with Cross Entropy loss
		ModeMultiLabel: sigmoid output with Cross Entropy loss
		ModeBinary: sigmoid output with binary CE loss */
		Mode: deep.ModeBinary,
		/* Weight initializers: {deep.NewNormal(μ, σ), deep.NewUniform(μ, σ)} */
		Weight: deep.NewNormal(1.0, 0.0),
		/* Apply bias */
		Bias: true,
	})
}

func trainNetwork(n *deep.Neural, data training.Examples) {
	// params: learning rate, momentum, alpha decay, nesterov
	optimizer := training.NewSGD(0.05, 0.1, 1e-6, true)
	// params: optimizer, verbosity (print stats at every 50th iteration)
	trainer := training.NewTrainer(optimizer, 50)

	training, heldout := data.Split(0.5)
	trainer.Train(n, training, heldout, 1000) // training, validation, iterations
}
