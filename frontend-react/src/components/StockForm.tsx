import React from "react";

interface StockFormProps {
  ticker: string;
  period: string;
  setTicker: (value: string) => void;
  setPeriod: (value: string) => void;
  onSubmit: (e: React.FormEvent) => void;
}

const StockForm: React.FC<StockFormProps> = ({ ticker, period, setTicker, setPeriod, onSubmit }) => {
  return (
    <form onSubmit={onSubmit} className="mb-4">
      <label className="mr-2">
        銘柄コード:
        <input
          type="text"
          value={ticker}
          onChange={(e) => setTicker(e.target.value)}
          className="ml-2 p-1 border rounded"
        />
      </label>
      <label className="mr-2">
        期間:
        <input
          type="text"
          value={period}
          onChange={(e) => setPeriod(e.target.value)}
          className="ml-2 p-1 border rounded"
        />
      </label>
      <button type="submit" className="ml-2 px-4 py-1 bg-blue-500 text-white rounded">
        Get
      </button>
    </form>
  );
};

export default StockForm;
