# Log periodic analysis of the bitcoin bubble (2015-2018)

## Introduction

This repo hosts tools for log periodic parameter analysis as described in [Why Stock Markets Crash: Critical Events in Complex Financial Systems](https://www.amazon.com/Why-Stock-Markets-Crash-Financial/dp/0691175950) by [Didier Sornette](https://en.wikipedia.org/wiki/Didier_Sornette).

## Motivation

There is much talk about the bitcoin bubble due to huge price appreciation in 2017. It reasonable to try to quantify what a bubble means and possibly predicts when it will pop. Quick seach reveals a lot of work around financial bubble prediction was done by Didier Sornette. He took the approach of applying crytical events analysis from physics...

## Model

![alt text](https://www.codecogs.com/eqnedit.php?latex=\inline&space;\log(p(t))&space;=&space;A&space;&plus;&space;B&space;(t_c&space;-&space;t)^{\beta}&space;(1&space;&plus;&space;C&space;\cos(\omega&space;\log(t_c-t)&plus;\phi))" target="_blank"><img src="https://latex.codecogs.com/gif.latex?\inline&space;\log(p(t))&space;=&space;A&space;&plus;&space;B&space;(t_c&space;-&space;t)^{\beta}&space;(1&space;&plus;&space;C&space;\cos(\omega&space;\log(t_c-t)&plus;\phi))

## Dependencies

* [matplotlib](https://matplotlib.org/)

## Usage

Run
```bash
$ python get_data.py
``` 

to get the latest data from poloniex. Note that current version needs daily data.

Run
```bash
go build
```
to build parameter search engine.

Run
```bash
./log-periodic-bitcoin --help
```
to check all available flags.

Run
```bash
./log-periodic-bitcoin 
```
to run a search for basic parameters. See help to run with more parallel processes, other search modes, etc.

Results will be saved in a .csv file which has the following header
```
cost;A B tc beta C omega phi 
```

Run
```bash
python plot.py A B tc beta
```
or
```
python plot.py A B tc beta C omega phi
```
to plot results. You need the same data file as for the analysis. (There is a maxDate limit in plot.py)

## Results

10.3.2018

Analysis was done with bitcoin data from poloniex (19.2.2015 - 13.12.2017).
There are several different parameter combinations that yield similar results (in terms of optimisation function mean squared error). 
basicRandomSearch.csv lists best results from recent run.
We are interested in `tc` (critical time), which predicts the end of the bubble.
Average critical time in basicRandomSearch.csv is `tc = 18.13 -> 16.2.2018`.
More analysis needs to be done around the parameter space where mentioned solutions were found.
Finding a log periodic pattern would further strengthen any predictions.

![alt text](https://github.com/slovenianblockchainassociation/log-periodic-bitcoin/blob/master/results/btc_bubble.png)
