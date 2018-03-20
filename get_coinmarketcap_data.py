import urllib2
import json

r = urllib2.urlopen('https://graphs2.coinmarketcap.com/global/marketcap-total/1367174820000/1521470820000/')
data = r.read()

parsed_data = [{'date': i[0]/1000., 'close': i[1]} for i in json.loads(data)['market_cap_by_available_supply']]

f = open('data_marketcap.json', 'w')
f.write(json.dumps(parsed_data))
f.close()
