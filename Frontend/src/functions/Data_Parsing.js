export function OhlcDataParsing(data){
    const candlestickData = data.ohlc.map(item => ({
        x: new Date(item.date),
        o: item.open,
        h: item.high,
        l: item.low,
        c: item.close,
      }));
    return candlestickData
}

export function IndicatorsDataParsing(data){
  const ohlc = data.ohlc;
  const indicators = data.indicators;

  const parsedIndicators = {};

  for (const key in indicators) {
    parsedIndicators[key] = indicators[key].map((val, idx) => ({
      x: new Date(ohlc[idx].date),
      y: val,
    }));
  }

  return parsedIndicators;
}

export function OhlcIndicatorsDataParsing(data) {
  if (!data || !Array.isArray(data.quotes)) {
    return { candlestickData: [], indicators: {} };
  }

  const candlestickData = data.quotes.map(item => ({
    x: new Date(item.date).getTime(),
    o: item.open,
    h: item.high,
    l: item.low,
    c: item.close,
  }));

  const indicatorsRaw = data.indicators || {};
  const indicators = {};

  for (const key in indicatorsRaw) {
    if (Array.isArray(indicatorsRaw[key])) {
      indicators[key] = indicatorsRaw[key].map((val, idx) => ({
        x: new Date(data.quotes[idx].date),
        y: val,
      }));
    }
  }

  return { candlestickData, indicators };
}

