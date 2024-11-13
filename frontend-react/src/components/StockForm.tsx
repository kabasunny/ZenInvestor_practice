// components/StockForm.tsx
import React, { ChangeEvent, FormEvent } from 'react';

interface StockFormProps {
  ticker: string;
  period: string;
  indicators: any[];
  includeVolume: boolean;
  setTicker: (ticker: string) => void;
  setPeriod: (period: string) => void;
  setIndicators: (indicators: any[]) => void;
  setIncludeVolume: (includeVolume: boolean) => void;
  onSubmit: (e: FormEvent) => void;
}

const StockForm: React.FC<StockFormProps> = ({
  ticker,
  period,
  indicators,
  includeVolume,
  setTicker,
  setPeriod,
  setIndicators,
  setIncludeVolume,
  onSubmit,
}) => {

  const handleIndicatorTypeChange = (index: number, e: ChangeEvent<HTMLSelectElement>) => {
    const newIndicators = [...indicators];
    newIndicators[index] = { ...newIndicators[index], type: e.target.value };
    setIndicators(newIndicators);
  };

  const handleWindowSizeChange = (index: number, e: ChangeEvent<HTMLInputElement>) => {
    const newIndicators = [...indicators];
    newIndicators[index] = {
      ...newIndicators[index],
      params: { ...newIndicators[index].params, window_size: e.target.value },
    };
    setIndicators(newIndicators);
  };

  const handleVolumeChange = (e: ChangeEvent<HTMLInputElement>) => {
    setIncludeVolume(e.target.checked);
  };

  const addIndicator = () => {
    const windowSizes = ["20", "50", "100"];

    if (indicators.length < 3) {
      setIndicators([
        ...indicators,
        { type: 'SMA', params: { window_size: windowSizes[indicators.length] } }
      ]);
    }
  };

  const removeIndicator = (index: number) => {
    const newIndicators = [...indicators];
    newIndicators.splice(index, 1);
    setIndicators(newIndicators);
  };

  return (
    <form onSubmit={onSubmit} className="space-y-4">
      <div className="flex space-x-4">
        {/* 指標選択部分 */}
        <div className="flex space-x-4">
          {indicators.map((indicator, index) => (
            <div key={index} className="border p-2 rounded-md">
              <div>
                <label htmlFor={`indicator-type-${index}`} className="block text-sm font-medium text-gray-700">指標の種類:</label>
                <select
                  id={`indicator-type-${index}`}
                  value={indicator.type}
                  onChange={(e) => handleIndicatorTypeChange(index, e)}
                  className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-teal-500 mb-2"
                >
                  <option value="SMA">SMA : Simple Moving Average</option>
                  <option value="EMA">EMA : Exponential Moving Average</option>
                  <option value="WMA">WMA : Weighted Moving Average</option>
                  <option value="RSI">RSI : Relative Strength Index</option>
                  <option value="BB">BB : Bollinger Bands</option>
                  <option value="MACD">MACD : MA Convergence Divergence</option>
                  <option value="SO">SO : Stochastic Oscillator</option>
                  <option value="ADX">ADX : Average Directional Index</option>
                  <option value="FR">FR : Fibonacci Retracement</option>
                </select>
              </div>
              {indicator.type === 'SMA' && (
                <div>
                  <label htmlFor={`window_size-${index}`} className="block text-sm font-medium text-gray-700">平均幅:</label>
                  <input
                    type="text"
                    id={`window_size-${index}`}
                    value={indicator.params.window_size}
                    onChange={(e) => handleWindowSizeChange(index, e)}
                    className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-teal-500"
                  />
                </div>
              )}
              <button
                type="button"
                onClick={() => removeIndicator(index)}
                className="mt-2 bg-red-500 text-white px-3 py-1 rounded-md hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-300"
              >
                削除
              </button>
            </div>
          ))}
        </div>

        {/* 銘柄コードと期間選択部分 */}
        <div className="flex flex-col space-y-4">
          <label htmlFor="ticker" className="block text-sm font-medium text-gray-700">銘柄コード:</label>
          <div className="flex items-center space-x-2">
            <input
              type="text"
              id="ticker"
              value={ticker}
              onChange={(e) => setTicker(e.target.value)}
              placeholder="例: AAPL"
              className="w-1/3 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-teal-500"
            />
            <button
              type="submit"
              className="ml-2 bg-green-500 text-white px-4 py-2 rounded-md hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-300"
            >
              更新
            </button>
          </div>

          <div>
            <label htmlFor="period" className="block text-sm font-medium text-gray-700">期間:</label>
            <select
              id="period"
              value={period}
              onChange={(e) => setPeriod(e.target.value)}
              className="w-2/3 px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-teal-500"
            >
              <option value="1d">1日</option>
              <option value="5d">5日</option>
              <option value="1mo">1ヶ月</option>
              <option value="3mo">3ヶ月</option>
              <option value="6mo">6ヶ月</option>
              <option value="1y">1年</option>
              <option value="2y">2年</option>
              <option value="5y">5年</option>
              <option value="10y">10年</option>
              <option value="ytd">年初～現在</option>
              <option value="max">最大期間</option>
            </select>
          </div>

          {/* 出来高チェックボックス */}
          <div className="flex items-center space-x-2 mt-4">
            <input
              type="checkbox"
              id="includeVolume"
              checked={includeVolume}
              onChange={handleVolumeChange}
              className="w-6 h-6 text-teal-500 focus:ring-teal-500 border-gray-300 rounded"
            />
            <label htmlFor="includeVolume" className="text-sm font-medium text-gray-700">出来高</label>
          </div>
        </div>
      </div>

      {/* 指標を追加ボタンをフォームの下に移動 */}
      <div className="flex justify-start mt-4">
        <button
          type="button"
          onClick={addIndicator}
          className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-300"
        >
          指標を追加
        </button>
        {indicators.length >= 3 && (
          <p className="text-red-500 mt-1 ml-4">指標は最大3つまで追加できます。</p>
        )}
      </div>
    </form>
  );
};

export default StockForm;
