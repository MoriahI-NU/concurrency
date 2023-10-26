# Concurrency

## The Goal
This experiment compares two similar programs whose aim is to train two predictive models. One program has each model running one after the other ("sequential.exe") and another program was written to be identical EXCEPT that it introduces concurrency where the models are trained essentially at the same time.

## Data and Models
The data used comes from the Boston Housing Study ("boston.csv"). It includes multiple variables that could affect the median price of homes.

Because the main goal of this experiment is to show the effects of concurrency, I did not think too much about which models I used or worry about changing parameters to increase accuracy. Accuracy is not the goal here, concurrency is.
The first model I used was a Random Tree Model. I chose this one because it can handle non-linear data and multiple types of data. "boston.csv" contains both float values and string values so I thought that random trees would be appropriate.
The second model I used was a Linear Regression Model. I chose this model because it is simple and has clear results. It will show correlations between variables and works well when predicting numerical estimations.

## Experiment Files and Outcomes
main.exe = Go application that runs both models WITH concurrency
sequential.exe = Go application that runs both models WITHOUT concurrency

Each file will run each model one hundred times and print out the results for each run. (That's 200 total runs and 200 printed lines). The last line printed will show the total runtime from when the application started running.

On my system:
main.exe (WITH concurrency) completed all runs in 2.29 seconds
sequential.exe (WITHOUT concurrency) completed all runs in 2.91 seconds

Implementing concurrency shaved 0.62 seconds off the total runtime.

## Management and Future Studies
If hypothetical management were interested as to the benefits and detriments of using concurrency with Go, I would say that it definitely is worth looking into. Concurrency saves time, as illustrated here. Granted, my experiment saved a fraction of a second but my experiment was also quite simple. The code itself was short, the tasks easily achievable, and the data was not incredible large. If you are dealing with a project that is much more complex, takes a lot longer to complete, and time is an absolute priority, then concurrency could definitely help. However, for something like this project - a simple prediction that takes less than three seconds, it's not necessary. Additionally, implementing concurrency requires more code which could come with its own learning curve and make things less easily readable.
Some things that could help expand on this experiment would be using a much larger dataset (this one was only ~500 rows long) and increasing the task complexity of the code. In addition to CPU time, memory usage should also be measured. Memory usage could be visualized to see how it is affected throughout the duration of the runtime for each experiment.