import json
import datetime
import matplotlib.pyplot as plt
import math
import sys
import numpy as np

labelFormatBasic = 'A={} B={} tc={} beta={}'
labelFormatFull = 'A={} B={} tc={} beta={} C={} omega={} phi={}'

def limitDataSetByMaxDate(minDate, maxDate, data):
	if minDate >= maxDate:
		return []
	if minDate > UnixToDecimal(data[-1]['date']):
		return []
	if maxDate < UnixToDecimal(data[0]['date']):
		return []
	start = -1
	for i in range(len(data)):
		if start < 0 and UnixToDecimal(data[i]['date']) > minDate:
			start = i
		if UnixToDecimal(data[i]['date']) > maxDate:
			return data[start:i]
	return data[start:]

def UnixToDecimal(timestamp):
	dt = datetime.datetime.fromtimestamp(timestamp)
	yearDays = datetime.datetime(dt.year, 12, 31, 0, 0, 0).timetuple().tm_yday
	return dt.year + (float(dt.timetuple().tm_yday) + dt.hour / 24.) / yearDays - 2000

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
	
	if len(sys.argv) not in [4, 8, 11]:
		print "Not enough arguments"
		sys.exit(1)

	filename = sys.argv[1]
	minDate = float(sys.argv[2])
	maxDate = float(sys.argv[3])

	if len(sys.argv) == 8:
		A = float(sys.argv[4])
		B = float(sys.argv[5])
		tc = float(sys.argv[6])
		beta = float(sys.argv[7])

	if len(sys.argv) == 11:
		A = float(sys.argv[4])
		B = float(sys.argv[5])
		tc = float(sys.argv[6])
		beta = float(sys.argv[7])
		C = float(sys.argv[8])
		omega = float(sys.argv[9])
		phi = float(sys.argv[10])

	with open(filename, 'r') as g:
		data = json.loads(g.read())

	data = limitDataSetByMaxDate(minDate, maxDate, data)

	x = [UnixToDecimal(i['date']) for i in data]
	y = [math.log(float(i['close'])) for i in data]
	# y = [float(i['close']) for i in data]

	plt.plot(x, y, label='BTC/USDT price')

	# plt.semilogy()
	# plt.title('BTC/USDT - Poloniex, 19.2.2015-13.12.2017')
	plt.xlabel('time [years]')
	plt.ylabel('log(price) [USDT]')

	if len(sys.argv) == 8:
		y_fit = [f(i, A, B, tc, beta, 0, 0, 0) for i in x]
		print J(data, A, B, tc, beta, 0, 0, 0)

		labelText = labelFormatBasic.format(A, B, Tc, Beta)
		plt.plot(x, y_fit, label=labelText)

	if len(sys.argv) == 11:
		y_fit = [f(i, A, B, tc, beta, C, omega, phi) for i in x]
		print J(data, A, B, tc, beta, C, omega, phi)

		labelText = labelFormatFull.format(A, B, tc, beta, C, omega, phi)
		plt.plot(x, y_fit, label=labelText)

	plt.legend()
	plt.show()
