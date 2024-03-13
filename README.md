# togglesink

A simple tool to toggle pulseaudio sinks.

This is a simple wrapper around pactl to toggle between pulesaudio sinks.

It is intended to be used with a keyboard shortcut, to quickly switch between audio outputs, e.g between speakers and
headphones.

At the moment, the tool simply switches between all the sinks available.

## Installation

Make sure you have the following installed:

- Go 1.22 or later
- PulseAudio

Then run the following command to install the tool:

```bash
go install github.com/ifross89/togglesink
```

## Usage

```bash
togglesink
```
