# file-upload

Prototyping image uploading in golang. This was mainly used to prototype image uploading and offline functionality (including storing channels and posts locally while offline).

### Setup

Follow the setup described for [Imagick](https://github.com/gographics/imagick). Imagick is needed to convert the files uploaded by the user into jpegs.

After installing all the other required dependencies, running the prototype should be as simple as entering `go run *.go` within the `server/` directory. After that, the upload site will be accessible at `localhost:9090`.