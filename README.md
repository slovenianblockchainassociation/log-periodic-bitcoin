# Log periodic analysis of the bitcoin bubble (2015-2018)

## Introduction

This repo hosts tools for log periodic parameter analysis as described in [Why Stock Markets Crash: Critical Events in Complex Financial Systems](https://www.amazon.com/Why-Stock-Markets-Crash-Financial/dp/0691175950) by [Didier Sornette](https://en.wikipedia.org/wiki/Didier_Sornette).

## Motivation

There is much talk in the general public about the bitcoin bubble due to huge price appreciation in 2017. 
Let's try to quantify this bubble and maybe predict it's end. 
Quick google search shows that a lot of work around financial bubble prediction was done by Didier Sornette. 

## Theory

Didier Sornette came up with a mathematical model that describes price appreciation in a bubble and predicts when is the highest probability that the bubble will pop.
In essence there are noise traders and informed traders trading with each other on the market. 
Usually hearding (everyone trading in the same direction) is very weak, but when strong, big sell-offs or up trends happen.
In a bubble informed traders stay in the market as long as their reward (price appreciation) is higher then the crash hazard rate.
This means that price `p(t)` depends on crash hazard rate. Crash hazard rate is a function of time that also depends on the price. 
There is a higher probability of a price decline (crash), if price went up without any change in the fundamentals.
When reward is lower then crash hazard rate, informed traders exit the market. It is very likely that this will trigger a hearding effect with the noise traders.

## Model

![alt text](https://latex.codecogs.com/gif.latex?\inline&space;\log(p(t))&space;=&space;A&space;&plus;&space;B&space;(t_c&space;-&space;t)^{\beta}&space;(1&space;&plus;&space;C&space;\cos(\omega&space;\log(t_c-t)&plus;\phi)))

## Dependencies

* [golang](https://golang.org/)
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

12.3.2018

Check periodicRandomSearch.csv for periodic search results.

![alt text](https://github.com/slovenianblockchainassociation/log-periodic-bitcoin/blob/master/results/btc_bubble_logperiodic.png)
