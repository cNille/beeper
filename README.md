# Beeper

Beeper is a simple command-line tool written in Go that plays a notification sound to alert you when a process has finished executing. You can use it in conjunction with other terminal commands to add an audible alert at the end of a long-running process.

## Usage

```bash
long_running_process && beeper
```

Or, an example with the sleep command:
```bash
sleep 5 && beeper
```

## Installation

To use Beeper, follow these steps:

Clone the repository or download the source code.

Build the executable using Go:

```bash
go build -o beeper main.go
``` 

Copy the beeper executable to a directory in your PATH so you can run it from anywhere in your terminal.

Optionally, you can customize the notification sound by replacing the notification.mp3 file with your own audio file. Make sure the audio file is in MP3 format.

## Requirements

Go programming language (to build the executable)
Speakers or headphones connected to your computer to hear the notification sound

## License

This project is licensed under the MIT License. See the LICENSE file for details.

##  Acknowledgments

Beeper uses the [gopxl/beep](https://github.com/gopxl/beep) package for audio playback.

