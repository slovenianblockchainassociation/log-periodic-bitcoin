import urllib2

r = urllib2.urlopen('https://poloniex.com/public?command=returnChartData&currencyPair=USDT_BTC&start=1405699200&end=9999999999&period=86400')
data = r.read()

f = open('data.json', 'w')
f.write(data)
f.close()
