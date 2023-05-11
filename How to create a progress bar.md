# How to create a progress bar

Do this inside of the print function where you want a progress bar to print before the table renders. Place this just before the for loop that builds the table to be printed.

```go
// Define the total number of iterations. This will be used by the progress bar
totalIterations := (consoleHeight / 2) - 1

// Create a new progress bar container
progressBarContainer := mpb.New()

// Create a progress bar called progressBar
progressBar := createProgressBar(progressBarContainer, totalIterations)
```

Then, inside the printing loop, just before the end of the loop:

```go
// Increment the progress progressBar
progressBar.Increment()
```

Then, outside the loop, just before rendering the table:

```go
// Wait for the progress progressBar to finish rendering
progressBarContainer.Wait()
```

----------------------------------------------------------

All of the above relies on a function that is already defined in the `progressBar.go` file:

```go
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
```