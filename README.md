# Igloo

Igloo can help to build a toolbox image for yourself easily.

## Usage

You can define the packages you would like to install in the `packages.yaml` file.
Then use the following commands to build and push the image.

- make build  : Build the Image
- make push   : Push the Image to the registry
- make run    : Test the Image
- make all    : Build and Push the Image

The whole flow could be described in short as follows:

1. Load `packages.yaml`
2. Generate the `Dockerfile`
3. Build the image