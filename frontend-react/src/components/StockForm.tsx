import React, { useState, ChangeEvent, FormEvent } from 'react';

interface StockFormProps {
  ticker: string;
  period: string;
  indicators: any[];
  setTicker: (ticker: string) => void;
  setPeriod: (period: string) => void;
  setIndicators: (indicators: any[]) => void;
  onSubmit: (e: FormEvent) => void;
}


const StockForm: React.FC<StockFormProps> = ({
  ticker,
  period,
  indicators,
  setTicker,
  setPeriod,
  setIndicators,
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

  const addIndicator = () => {
    setIndicators([...indicators, { type: 'SMA', params: { window_size: '20' } }]);
  };

  const removeIndicator = (index: number) => {
    const newIndicators = [...indicators];
    newIndicators.splice(index, 1);
    setIndicators(newIndicators);
  };


  return (
    <form onSubmit={onSubmit} className="space-y-4">
      <input
        type="text"
        value={ticker}
        onChange={(e) => setTicker(e.target.value)}
        placeholder="銘柄コード (例: AAPL)"
        className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-teal-500"
      />
      <select
        value={period}
        onChange={(e) => setPeriod(e.target.value)}
        className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-teal-500"
      >
        <option value="1d">1日</option>
        <option value="5d">5日</option>
        <option value="1mo">1ヶ月</option>
        <option value="1y">1年</option>
        {/* 他の期間を追加 */}
      </select>

      {indicators.map((indicator, index) => (
        <div key={index} className="border p-2 rounded-md">
        <select
          value={indicator.type}
          onChange={(e) => handleIndicatorTypeChange(index, e)}
          className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-teal-500 mb-2"
        >
          <option value="SMA">SMA</option>
          <option value="EMA">EMA</option>
          {/* 他の指標を追加 */}
        </select>
        {indicator.type === 'SMA' && ( // SMAの場合のみwindow_sizeを表示
            <div>
          <label htmlFor={`window_size-${index}`}>Window Size:</label>
          <input
            type="text"
            id={`window_size-${index}`}
            value={indicator.params.window_size}
            onChange={(e) => handleWindowSizeChange(index, e)}
            className="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-1 focus:ring-teal-500"
          />
                </div>
        )}
          <button type="button" onClick={() => removeIndicator(index)} className="text-red-500">
            削除
          </button>

        </div>
      ))}


      <div>
          <button type="button" onClick={addIndicator}>指標を追加</button>
      </div>




      <button
        type="submit"
        className="bg-teal-500 text-white px-4 py-2 rounded-md hover:bg-teal-600 focus:outline-none focus:ring-2 focus:ring-teal-300"
      >
        データ取得
      </button>
    </form>
  );
};

export default StockForm;