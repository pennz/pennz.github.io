my Question:
Overfit to current problem? Not good for private test case.

Mannually center and crop the images, it won't help much, even worse...
What I did wrong? Overfit?


# Lesson
## Validation Strategy

combine the whole 2015 and 2019 data as train set, solely relied on public LB for validation.

## Preprocessing
Just plain resizing... (image qualities are perfect for deep neural networks)


## Models and input sizes
simple average of the following eight models...
inceptions and resnets usually blend well

the input size was mainly determined by observations in the 2015 competition
that larger *input size* brought better performance

384 / 512 , 384 enough for public LB, for private set, use 512

## Loss, aug, pool
only nn.SmoothL1Loss() ....
*stick to this single loss just to simplify the emsembling process*

augmentation: contrast, brightness, hue, saturation, blur and sharpen, rotate,
scale, sheer, shift, mirror.

For the last pooling layer, the *generalized mean pooling*, which is better
than original average pooling.

## traing and testing

Two stage of traing:

first just routinely trained the eight models and validated each of them.

2x for each type of model to get more stable resulte, each of the two run with
different seeds.

tried to reduce the degree of freedom of hyperparemeters to alleviate overfitting, e.g., to determine the stopping epoch number, use step size of five.
(really good result...)

second
added more external data

mitigate the labeling bias: smoothed the labels by averaging the provided
labels with the predicted labels from stage 1 models.

and another data mapping thing to train more data.


# TODO

check the paper about generalized mean pooling
compare the loss usage, and analysis
learn about pseudo labels


# another one, the strategies
## Pseudo lables

