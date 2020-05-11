# Jigsaw Biased Comments

Learn from others:

## 2nd place solution

### multiple models...

### LOSS

different _weights_ ... for different groups, toxic loss and identity loss

### Predict

split the score to 11 classes, 0.0, 0.1 ~ 1.0 and the prediction is toxic\_logits \* cls\_vals zipwise product

### then blend....

Power 3.5 weighted sum ...

## 3rd place solution

### blend

optuna to determine blending weight

