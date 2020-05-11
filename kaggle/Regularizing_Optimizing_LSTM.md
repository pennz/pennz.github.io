LSTM
Long Short Term Memory

act as a small unit to understand the current state of input words, LSTM stacks up and
undertands the full sentence. 
You just need large data to train your model.

# weight-dropped LSTM
applies recurrent regularization through a DropConnect mask on the hidden-to-hidden
recurrent weights

other modifications like dropout in hidden state vechtor h_t-1 or in updating to the memory state c_t, these modifications to standard LSTM prevent treating it as a black box, might make it slow.
To address this problem, use of DropConnect on the recurrent hidden to hidden weight matrices is proposed.
# randomized-length backpropagation through time(BPTT)

# embedding dropout
# activation regularization(AR)
# temporal activation regularization(TAR)

# averaged SGD (ASGD)
NT-ASGD (none-monotomically triggered variant of ASGD), opt better for LSTM

all others are for regulation.



