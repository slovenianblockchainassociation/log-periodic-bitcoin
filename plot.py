import json
import datetime
import matplotlib.pyplot as plt
import math
import sys
import numpy as np

labelFormatBasic = 'A={} B={} tc={} beta={}'
labelFormatFull = 'A={} B={} tc={} beta={} C={} omega={} phi={}'

def limitDataSetByMaxDate(minDate, maxDate, data):
	start = -1
	for i in range(len(data)):
		if start < 0 and UnixToDecimal(data[i]['date']) > minDate:
			start = i
		if UnixToDecimal(data[i]['date']) > maxDate:
			return data[start:i]
	return data

def UnixToDecimal(timestamp):
	dt = datetime.datetime.fromtimestamp(timestamp)
	yearDays = datetime.datetime(dt.year, 12, 31, 0, 0, 0).timetuple().tm_yday
	return dt.year + float(dt.timetuple().tm_yday) / yearDays - 2000

def f(t, A, B, tc, beta, C, omega, phi):
	return A + B * np.power(tc - t, beta) * (1 + C*np.cos(omega*np.log(tc - t) + phi))

def J(data, A, B, tc, beta, C, omega, phi):
	j = 0
	for i in data:
		j += np.power(f(UnixToDecimal(i['date']), A, B, tc, beta, C, omega, phi) - math.log(float(i['close'])), 2)
	return j / len(data) / 2

# def model(t, A, B, tc, beta, C, omega, phi):
# 	return A + B * np.power(tc - t, beta) * (1 + C * np.cos(omega * np.log(tc - t) + phi))

if __name__ == '__main__':
	
	if len(sys.argv) not in [5, 8]:
		print "Not enough arguments"
		sys.exit(1)

	A = float(sys.argv[1])
	B = float(sys.argv[2])
	tc = float(sys.argv[3])
	beta = float(sys.argv[4])

	if len(sys.argv) == 8:
		C = float(sys.argv[5])
		omega = float(sys.argv[6])
		phi = float(sys.argv[7])

	with open('data.json', 'r') as g:
		data = json.loads(g.read())

	data = limitDataSetByMaxDate(17.2, 17.95, data)

	x = [UnixToDecimal(i['date']) for i in data]
	y = [math.log(float(i['close'])) for i in data]

	if len(sys.argv) == 5:
		y_fit = [f(i, A, B, tc, beta, 0, 0, 0) for i in x]
		print J(data, A, B, tc, beta, 0, 0, 0)

	if len(sys.argv) == 8:
		y_fit = [f(i, A, B, tc, beta, C, omega, phi) for i in x]
		print J(data, A, B, tc, beta, C, omega, phi)

	labelText = labelFormatBasic.format(sys.argv[1], sys.argv[2], sys.argv[3], sys.argv[4])

	if len(sys.argv) == 8:
		labelText = labelFormatFull.format(sys.argv[1], sys.argv[2], sys.argv[3], sys.argv[4], sys.argv[5], sys.argv[6], sys.argv[7])

	plt.plot(x, y, label='BTC/USDT price')
	plt.plot(x, y_fit, label=labelText)

	plt.semilogy()
	plt.title('BTC/USDT - Poloniex, 19.2.2015-13.12.2017')
	plt.xlabel('time [years]')
	plt.ylabel('log(price) [USDT]')
	plt.legend()
	plt.show()

	plt.show()
