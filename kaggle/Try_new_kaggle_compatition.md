Try a new competition

[Ion Channel Wikipedia](https://en.wikipedia.org/wiki/Ion_channel)

reading the paper:

model: hybrid recurrent CNN(RCNN)
data: an analogue synthetic ion channel record generator system for training and validation
    two kinetic schemes, M1 low channel open probability, M2, high open probability. 
    There are real data and synthesized data. 
    

And we have the code for tensorflow, we will change it to fastai.

How to handle the generalization problem?

In the paper, there is golden dataset produced by five ion channel experts

# model

input(to 1D convolution layer): 3D data with ion channel current recordings(raw), time steps(n=1), and features (n=1)

# TODO
review LSTM, the figure
advantage, why it is brought up

CONV1D

maybe try n-1, the output will be more reasonable.
And we found some people use u-net, and input is 4000\*1, output 4000\*1.
In the paper, it is 1\*1->model->1\*1 

We will try to replicate it in fastai.
