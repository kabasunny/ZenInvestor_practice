import React, { useState } from 'react';
import axios from 'axios';
import '../assets/styles/LossCutSimulator.css'; // カスタムCSSファイルをインポート

const LossCutSimulator: React.FC = () => {
    const [ticker, setTicker] = useState('AAPL');
    const [simulationDate, setSimulationDate] = useState('2023-12-21');
    const [stopLossPercentage, setStopLossPercentage] = useState(2.0);
    const [trailingStopTrigger, setTrailingStopTrigger] = useState(5.0);
    const [trailingStopUpdate, setTrailingStopUpdate] = useState(2.0);
    const [chartData, setChartData] = useState('');
    const [profitLoss, setProfitLoss] = useState(0);

    const handleSubmit = async () => {
        try {
            const response = await axios.post('http://localhost:8086/losscutSimulation', {
                ticker,
                simulationDate,
                stopLossPercentage,
                trailingStopTrigger,
                trailingStopUpdate,
            });
            setChartData(response.data.chart_data.chart_data);
            setProfitLoss(response.data.profitLoss);
        } catch (error) {
            console.error('Error fetching simulation data:', error);
        }
    };

    return (
        <div className="bg-teal-50 p-6 rounded-lg shadow-lg">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                    <div className="mb-4">
                        <label className="block mb-2">ティッカー:</label>
                        <input
                            type="text"
                            value={ticker}
                            onChange={(e) => setTicker(e.target.value)}
                            className="w-full p-2 border rounded"
                        />
                    </div>
                    <div className="mb-4">
                        <label className="block mb-2">シミュレーション日付:</label>
                        <input
                            type="date"
                            value={simulationDate}
                            onChange={(e) => setSimulationDate(e.target.value)}
                            className="w-full p-2 border rounded"
                        />
                        <p
                            className={`mt-4 text-2xl ${profitLoss < 0 ? 'text-red-500' : 'text-teal-500'}`}
                        >
                            損益率: {profitLoss}%
                        </p>
                    </div>
                </div>
                <div>
                    <div className="mb-4">
                        <label className="block mb-2">ストップロスパーセンテージ: {stopLossPercentage}%</label>
                        <input
                            type="range"
                            min="1"
                            max="10"
                            value={stopLossPercentage}
                            onChange={(e) => setStopLossPercentage(parseFloat(e.target.value))}
                            className="w-full"
                        />
                        <div className="flex justify-between text-xs">
                            <span>1%</span>
                            <span>5.5%</span>
                            <span>10%</span>
                        </div>
                    </div>
                    <div className="mb-4">
                        <label className="block mb-2">トレーリングストップトリガー: {trailingStopTrigger}%</label>
                        <input
                            type="range"
                            min="5"
                            max="20"
                            value={trailingStopTrigger}
                            onChange={(e) => setTrailingStopTrigger(parseFloat(e.target.value))}
                            className="w-full"
                        />
                        <div className="flex justify-between text-xs">
                            <span>5%</span>
                            <span>12.5%</span>
                            <span>20%</span>
                        </div>
                    </div>
                    <div className="mb-4">
                        <label className="block mb-2">トレーリングストップ更新: {trailingStopUpdate}%</label>
                        <input
                            type="range"
                            min="1"
                            max="10"
                            value={trailingStopUpdate}
                            onChange={(e) => setTrailingStopUpdate(parseFloat(e.target.value))}
                            className="w-full"
                        />
                        <div className="flex justify-between text-xs">
                            <span>1%</span>
                            <span>5.5%</span>
                            <span>10%</span>
                        </div>
                    </div>
                </div>
            </div>
            <button onClick={handleSubmit} className="bg-teal-500 text-white p-2 rounded">
                シミュレーション実行
            </button>
            {chartData && (
                <div className="mt-4">
                    <h3 className="text-lg font-semibold">シミュレーション結果</h3>
                    <img src={`data:image/png;base64,${chartData}`} alt="チャート" />
                </div>
            )}
        </div>
    );
};

export default LossCutSimulator;
