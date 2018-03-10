# Log periodic analysis of the bitcoin bubble (2015-2018)

## Introduction

This repo hosts tools for log periodic parameter analysis as described in [Why Stock Markets Crash: Critical Events in Complex Financial Systems](https://www.amazon.com/Why-Stock-Markets-Crash-Financial/dp/0691175950) by [Didier Sornette](https://en.wikipedia.org/wiki/Didier_Sornette).

## Motivation

There is much talk about the bitcoin bubble due to huge price appreciation in 2017. It reasonable to try to quantify what a bubble means and possibly predicts when it will pop. Quick seach reveals a lot of work around financial bubble prediction was done by Didier Sornette. He took the approach of applying crytical events analysis from physics...

## Model

\log(p(t)) = A + B(t_c - t)^{\beta} (1 + C(\cos(\omega\ln(t_c - t) + \phi)))

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