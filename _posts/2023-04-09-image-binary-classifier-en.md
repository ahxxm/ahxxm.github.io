---
title: "Train a Personal Image Scoring Model"
author: ahxxm
layout: post
permalink: /179.moew/
categories:
  - IT
  - ML
---

I have many images, and I have also deleted many images, I am likely to continue classifying images for preservation or deletion in the future. Can I utilize modern technology to simplify this task, such as ordering by the likelihood to delete before I actually review them?

Yes to an extent now, and it should become better as my preferences become clearer with more training data.

*Warning: moderate amount of Python codes inside.*

<!--more-->

## Training

The training process is boringly traditional: collect source data and labels, transform into vectors and labels, remove data leakage, experiment ~~play~~ with different model structures and hyperparameters, and verify through feedback loop.

### Dataset

Since the model needs to express my preferences, all labels are assigned by me, resulting in a small sample size:
- 4000+ positive: memes, infographics, professional photos(mostly landscape and some animals), a few photos taken with a phone, screenshots
- ~1000 photos taken by my DSLR, mostly landscape, some animals <!--try not to imply professionalism-->
- 20000+ negative, intentionally kept since I wanted this experiment: photos(some professional ones, most taken with a phone), screenshots

Meme and screenshots are too diverse to describe.

### Model Justification

There aren't too many options, CLIP limits to 224x224, which was the major blocker of this experiment. I chose `uform-vl-english` because it allows images of arbitrary shapes, and I do not mind if it generates 224x224 thumbnails [behind the scene](https://github.com/unum-cloud/uform/blob/main/src/uform.py#L362).

Another important factor is their excellent article: [Beating OpenAI CLIP with 100x less data and compute](https://www.unum.cloud/blog/2023-02-20-efficient-multimodality).

### Transform

The model encodes image and text into a joint embedding of 768 float numbers.

```python
# PIL.Image.MAX_IMAGE_PIXELS=None
# PIL.Image.DecompressionBombError: Image size (252349376 pixels) exceeds limit of 178956970 pixels, could be decompression bomb DOS attack.
uf_model = uform.get_model("unum-cloud/uform-vl-english")
DIMENSION = 768
def get_image_embedding(path: pathlib.PosixPath) -> typing.List[float]:
    img = Image.open(path)
    image_data = uf_model.preprocess_image(img)
    text_data = uf_model.preprocess_text(path.name) # filenames aren't necessarily useful, encode anyway
    memb = uf_model.encode_multimodal(image=image_data, text=text_data).detach().numpy()
    assert memb.shape == (1, DIMENSION)
    return memb[0]

def list_dir_recur(dir_path: str):
    path = pathlib.Path(dir_path)
    for p in path.iterdir():
        if p.is_dir():
            yield from list_dir_recur(p)
        else:
            yield p

def generate_embeddings_for_dir(dir_path: str):
    data = [(p.name, sha256(p), get_image_embedding(p))
            for p in list_dir_recur(dir_path) if p.suffix.lower() in {".png", ".jpg", ".jpeg"}]
    df = pd.DataFrame(data)
    df.columns = ["filename", "sig", "uform_emb"]
    return df
```

To prevent data leakage, I checked file signatures for an exact match, I should also check embeddings similarities.

Inferencing 8000 images with CPU took about 30 minutes, 20000 images with GPU took slightly shorter, the GPU utilization rate wasn't high at all. (Quote from [here](https://timdettmers.com/2023/01/30/which-gpu-for-deep-learning/): "So fast, in fact, that they(tensor cores) are idle most of the time as they are waiting for memory to arrive from global memory.")

```python
# simple label
df_accepted = generate_embeddings_for_dir(accepted_path)
df_rejected = generate_embeddings_for_dir(rejected_path)
df_accepted["val"] = 1.0
df_rejected["val"] = 0.0

# To prevent duplicated work, save embeddings to parquet files
# df_accepted.to_parquet("accepted-{path-description-here}.pq")

df = pd.concat([df_accepted, df_rejected]).drop_duplicates(["sig", "val"])
df = df[["uform_emb", "val"]]

def split_train_test(df: pd.DataFrame, test_size: float = 0.1):
    sdf = df.sample(frac=1).reset_index(drop=True)
    train_size = int(len(sdf) * (1 - test_size))
    return sdf.iloc[:train_size], sdf.iloc[train_size:]

df_train, df_test = split_train_test(df)
```

### Hyperparameters

Except for 768 and 1, the model structure is purely arbitrary.

```python
class ImageBinaryClassifier(nn.Module):
    def __init__(self):
        super().__init__()
        self.layers = nn.Sequential(
            nn.Linear(768, 2048),
            nn.ReLU(),
            nn.Linear(2048, 2048),
            nn.ReLU(),
            nn.Linear(2048, 2048),
            nn.ReLU(),
            nn.Linear(2048, 256),
            nn.ReLU(),
            nn.Linear(256, 1),
        )
    def forward(self, x):
        return self.layers(x)
```

There isn't Sigmoid because the loss function `BCEWithLogitsLoss` -- claims to be numerically stable -- covers it, at the cost of needing to sigmoid predicted values.

```python
# HACK: to train with a graphics card, bad practice
# torch.set_default_tensor_type('torch.cuda.FloatTensor')
model = ImageBinaryClassifier()
# sum([x.reshape(-1).shape[0] for x in model.parameters()]) # 10492417 params
learning_rate = 0.0003
num_epochs = 15
batch_size = 1
train_loader = DataLoader(train_dataset, batch_size=batch_size, shuffle=True, generator=torch.Generator(device='cuda') if cuda else None)

criterion = nn.BCEWithLogitsLoss()
optimizer = optim.SGD(model.parameters(), lr=learning_rate, momentum=0.0)

model.train()
for epoch in range(num_epochs):
    train_loss = 0.0
    for i, (inputs, labels) in enumerate(train_loader):
        optimizer.zero_grad()
        outputs = model(inputs)
        loss = criterion(outputs, labels)
        loss.backward()
        optimizer.step()
        train_loss += loss.item()
    print(f"Epoch {epoch}, total loss {train_loss}")

# save the weights
MODEL_PATH = "model.pt"
torch.save(model.state_dict(), MODEL_PATH)
```

The training finished in 10 minutes using Tesla T4, its output looks like:
```
Epoch 0, total loss 9449.750139231794
Epoch 1, total loss 3176.7831058780284
Epoch 2, total loss 1782.7670029629808
Epoch 3, total loss 1066.8573923760205
Epoch 4, total loss 795.308564953793
Epoch 5, total loss 650.6638346871068
Epoch 6, total loss 562.0393543641796
Epoch 7, total loss 489.4068465524059
Epoch 8, total loss 433.97104152526447
Epoch 9, total loss 393.40820040911666
Epoch 10, total loss 360.4532831643036
Epoch 11, total loss 322.57149341895547
Epoch 12, total loss 307.90023628154034
Epoch 13, total loss 269.49978952276723
Epoch 14, total loss 254.2954651907422
```

This specific combination of `learning_rate` and `num_epochs` allowed the total loss to converge, with diminishing returns(and an increasing risk of overfitting?). I've no idea on the optimizer choice, as with `momentum=0.9` the loss looks lower, and `Adam` works fine too.

During training, the GPU utilization rate reported by `nvidia-smi` is around 78%.

### Verify

After each training, first evaluate on the test dataset,
```python
test_loader = DataLoader(test_dataset, batch_size=1, shuffle=False, generator=torch.Generator(device='cuda') if cuda else None)
model.eval()
predictions = []
test_labels = []
with torch.no_grad():
    correct = 0
    total = 0
    for i, (inputs, labels) in enumerate(test_loader):
        outputs = torch.sigmoid(model(inputs)) # required by BCEWithLogitsLoss
        predictions.append(outputs[0][0]) # batch size 1 to avoid flattening
        test_labels.append(labels[0][0])
        total += labels.size(0)
        correct += (torch.round(outputs) == torch.round(labels)).sum().item()
    print('Accuracy of the model on the validation set: {:.2f} %'.format(100 * correct / total))

# round() is a bad criteria, >95% accuracy makes me doubt where went wrong
# examine individual predictions difference
sorted([(i, j) for i,j in zip(predictions, test_labels) if abs(i-j) > 0.4])
```

If the performance is too bad or too good, investigate where went wrong, else continue to evaluate new images with the model and eyes.

### Evaluation

Optionally, save weights by `torch.save(model.state_dict(), "model.pt")` and transfer back to evaluate local images,
```bash
# remote, about 40MB
sudo tailscale file cp model.pt laptop-ae88:
# laptop-ae88 to receive
sudo tailscale file get .
```

Load to infer with CPU:
```python
saved_model = ImageBinaryClassifier()
saved_model.load_state_dict(torch.load(MODEL_PATH, map_location=torch.device('cpu')))
saved_model = torch.compile(saved_model)
saved_model.eval()

IMG_PATH = "/path-to-some-new-images"
order_emb = generate_embeddings_for_dir(IMG_PATH)
order_emb["val"] = order_emb["uform_emb"].map(lambda x: float(torch.sigmoid(saved_model(torch.tensor(x)))))
order_emb.sort_values(by="val", ascending=True).head(10)
```

After checking individual predictions, I found some wrong labels, with a little surprise! For example, it suggested with a very high confidence of `6.15e-07` that an Italian infograph should be deleted instead, which is true, I don't remember or understand it now.

I also run this against a new batch of very random images, of the top 10 most "unwanted" images, there are:
- 8 to delete at 0.00008~0.10: true negative
- 1 to keep, 0.04: very chill landscape photography, 1200mm(estimated) on a Japanese shrine with full moon as the background
- 1 meme can't decide on: I tend to believe I've better ones to express the same feeling

By the way, [When your binary classification model outputs 0.5](https://twitter.com/mldcmu/status/1048995493848776705) is a high qualify meme regardless of the image quality.

## Thoughts

The good:
- The model can capture some of my preferences on a variety of subjects: meme, screenshot, photography
- The model can even correct some wrong labels(sometimes)
- To sort by preferences for the image viewer: sort by the prediction values and "touch" files one by one

The bad:
- More labels needed: low performing image types, write descriptive filename
- Trust but verify: The model can also miss my preferences, the full moon picture example above

The ugly, engineering efforts to make it easier to train use:
- Already some hacks for CPU and GPU
- It's stateful, unlike LLM summarizer for podcast episodes(not ready to open-source) and [browser tab](https://github.com/merrickluo/summer-web-clipper/), so it can be hard to manage existing embeddings and new images
- How to fix for wrong labels?
- How to run easily against new images in a lightweight manner? The dependency `uform` contains about 1.7G transitive dependencies, getting embeddings requires it
