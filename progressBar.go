package main

import (
	"fmt"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
	"time"
)

// progressBarUnix displays a progress bar in the terminal by printing a period
// every 500 milliseconds until it receives a value on the given channel. The
// function runs in a separate goroutine, so it can be executed concurrently
// with other parts of a Go program. The progress bar can be stopped by sending
// a value on the channel.
//
// Parameters:
// progressBarChannel - the channel to listen for stop signal on.
//
// Example:
//
//	progressBarChannel := make(chan bool)
//	go progressBarUnix(progressBarChannel)
//
//	// Do some work
//	time.Sleep(60 * time.Second)
//
//	// Send a value to the channel to stop the progress bar
//	progressBarChannel <- true
func progressBarUnix(progressBarChannel chan bool) {
	for {
		select {
		case <-progressBarChannel:
			// Stop the progress bar when the channel receives a value
			fmt.Printf("\n")
			return
			//goland:noinspection GrazieInspection
		default:

			// Display a progress bar with 60 steps, each step taking 1 second.
			//goland:noinspection GrazieInspection,GrazieInspection
			for i := 0; i <= 60; i++ {

				// For each step,
				//   1. Print a solid block character █ (Unicode \u2588).
				//   2. Then print the remaining-seconds countdown number (zero-padded
				//      to two digits with %02d).
				//   3. Then move the cursor back two spaces using the ASCII control sequence
				//      \u001B[2D which moves the cursor two characters to the left to overwrite
				//      the number that was printed in the previous step.
				// We stay on one line the whole time.
				fmt.Printf("\u2588%02d\u001B[2D", i)

				time.Sleep(1 * time.Second)
			}
		}
	}
}

func progressBarWindows(progressBarChannel chan bool) {
	for {
		select {
		case <-progressBarChannel:
			// Stop the progress bar when the channel receives a value
			fmt.Printf("\n")
			return
		default:

			// Display a progress bar with 60 steps, each step taking 1 second.
			for i := 0; i <= 60; i++ {

				// For each step,
				//   1. Print a solid block character #.
				//   2. Then print the remaining-seconds countdown number (zero-padded
				//      to two digits with %02d).
				//   3. Then move the cursor back two spaces using the ASCII control sequence
				//      \r which moves the cursor back to the beginning of the line to overwrite
				//      the number that was printed in the previous step.
				// We stay on one line the whole time.
				fmt.Printf("#%02d\r", i)

				time.Sleep(1 * time.Second)
			}
		}
	}
}

// createProgressBar This function creates a progress bar with a name and an
// iteration counter, as well as a percentage complete indicator.
func createProgressBar(progressBarContainer *mpb.Progress, totalIterations int) *mpb.Bar {
	progressBar := progressBarContainer.AddBar(int64(totalIterations),
		mpb.PrependDecorators(
			decor.Name("Progress: "),
			decor.CountersNoUnit("%d/%d", decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.Percentage(decor.WCSyncSpace),
		),
	)
	return progressBar
}
