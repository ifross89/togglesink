package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"os/signal"
)

type listOutput struct {
	Sinks []struct {
		State string `json:"state"`
		Name  string `json:"name"`
	} `json:"sinks"`
}

func getListOutput(ctx context.Context) (listOutput, error) {
	cmd := exec.CommandContext(ctx, "pactl", "-f", "json", "list")
	b, err := cmd.CombinedOutput()
	if err != nil {
		return listOutput{}, fmt.Errorf("error listing from pactl, out=%s: %v", string(b), err)
	}

	var lo listOutput
	if err := json.Unmarshal(b, &lo); err != nil {
		return listOutput{}, fmt.Errorf("error parsing list output: %v", err)
	}

	return lo, nil
}

func setDefaultSink(ctx context.Context, sink string) error {
	cmd := exec.CommandContext(ctx, "pactl", "set-default-sink", sink)

	b, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error setting default sink to %s, out=%s: %v", sink, string(b), err)
	}

	return nil
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	lo, err := getListOutput(ctx)
	if err != nil {
		slog.Error("error getting pactl list output", "error", err)
		os.Exit(1)
	}

	var nextIndex int
	for i, sink := range lo.Sinks {
		if sink.State == "RUNNING" {
			nextIndex = (i + 1) % len(lo.Sinks)
			break
		}
	}

	if err := setDefaultSink(ctx, lo.Sinks[nextIndex].Name); err != nil {
		slog.Error("error setting default sink", "error", err)
		os.Exit(1)
	}

	slog.Info("set default sink to", "sink", lo.Sinks[nextIndex].Name)
}
