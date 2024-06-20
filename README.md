# picsum.photos CLI

## Description

This is a CLI tool that allows you to download images from [picsum.photos](https://picsum.photos/).

## Installation

```bash
git clone https://github.com/afroborg/picsum-cli.git
cd picsum-cli
make install
```

This will install the CLI tool in `/usr/local/bin`.

## Usage

Help: `picsum-cli -h`

### Dimensions

You can specify the width and height of the image you want to download using the `-w` and `-h` flags.

```bash
picsum-cli -w 200 -h 200
```

### Output

You can specify the output directory using the `-d` flag, and the output file name using the `-o` flag.

```bash
picsum-cli -d /path/to/output -o image # saves the image as /path/to/output/image.jpg
```

### Format

The default format is `jpg`, but you can download a `webp` image using the `-webp` flag.

```bash
picsum-cli -webp
```

### Blur and grayscale

Blur and grayscale can be applied to the image using the `-b` and `-g` flags.

```bash
picsum-cli -b 5 -g
```

> Blur is a number between 0 and 10.

### Seed and Id

You can specify a seed using the `-s` flag. This will generate a random image based on the seed.

You can also specify an image id using the `-id` flag. This will download the image with the specified id. Complete list of images can be found [here](https://picsum.photos/images).

```bash
picsum-cli -s 123
```

```bash
picsum-cli -id 100
```

## License

MIT
