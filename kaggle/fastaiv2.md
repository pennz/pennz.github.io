# fastai V2

layered structure

two main design goals:

* to be approachable and rapidly productive
* while also being deeply hackable and configurable

through: layered architecture High-level API: ready-to-use with sensible defaults and customizable models lower-level: composable building blocks

and the picture says alot.

### High-level

four main application areas: vision, text, tabular and time-series analysis and collaborative filtering.

_Learner_ class, compose/glue architecture, optimizer, data together, and automatically loss function, all can come with default values ready for use \(with best practice\) e.g., DataLoader class, load validation and training data with split, and shuffle training data.

transfer learning tailored functions are built in: transfer learning optimised batch-normalization \[3\] training, layer freezing, and discriminative learning rates.

### mid-level

provides core deep learning and data-processing methods

### low-level

a library of optimized primitives and functional and object-oriented foundations, which allows the mid-level to be developed and customised. Built on PyTorch, Numpy, PIL, pandas and so on. low-level is hackable too.

## Applications

### Vision

example of fine-tune an ImageNet

code walkthrough: 1. import fastai library 2. get the data from internet and untar it 3. create dataloader from the untared data \(will setup batch size, filters for the file names to be loaded, augmentations, and batch trms, normalize thing\) 4. create the learner \( data, network , loss\) 5. let the learner learn from the training data

above, batch trms are applied to a mini-batch \(might be in GPU, faster\): a great hack

besideds, we can easily look at the data with one line of code:

```python
dls.show_batch()
```

...

### Text

almost the same code data -&gt; TextDataLoader -&gt; learner -&gt; fit the learner the same behaviour for different applications. and many manuvours for NLP is added to fastai, like some token processing techs

### Tabular

code to create and train a model for tabular data looks familiar, there is just information specific to tabular data requires when building the DataLoader object.

### Collaborative filtering

a colaborative filtering model in fastai can be simply seen as a tabular model with high cardinality categorical variables.

### Deployment

Learner.export will serialize the model \(PyTorch model\)

## High-level API design considerations

all the fastai applications share some basic components: e.g. visualisation API \(show\_batch, show\_resultsa\), learn.lr\_find\(\) data block API\(changed to Functional API, you don't need to consider the order in fluent api\) to do the usual thing for loading data to better computing.

### Incrementally adapting PyTorch code

### Consistency across domains

## Mid-level APIs

beautiful layered structure

### Learner

requires a PyTorch model, and optimizer, a loss function ans a DataLoader object also handles transfer learning functionality \(along with Optimizer, different update rate\)

### two-way callbacks

2-way callback system allows gradients, data, losses, control flow, and anything else to be read

and changed at any point during training.

callback in the training framework, every critical points there is a callback

#### Case study: generative

### Generic optimizer

## Low-level APIs

middle layer is based on this set of abstraction. 1. Pipeline of transforms 2. Type-dispatch based on the needs of data processing pipelines 3. Attach semantics to tensor objects, besides, the semantics are maintained throughout a pipepline. 4. GPU-optimized computer vision operations 5. Convenience functionality

### PyTorch foundations

pytorch is easier to extend. So fastai move from tensorflow to pytorch.

