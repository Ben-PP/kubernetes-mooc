# log-writer
This is small log writer app to write timestamps with random string into a file.

## Usage
Specify the log directory where the timestamps are written with `LOG_DIR` environment variable. This defaults to `/app/logs` in the container image and if `LOG_DIR` is empty or not provided, it defaults to `./logs` in the app.

To run only this app, use image `benpp/log-writer:1.10`
