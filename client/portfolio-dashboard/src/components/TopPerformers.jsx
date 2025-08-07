import React from "react";

const TopPerformers = ({ summary }) => {
  return (
    <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
      <div className="bg-white p-4 shadow rounded">
        <h4 className="text-gray-600 text-sm mb-1">Top Performer</h4>
        <p className="text-lg font-semibold">{summary.topPerformer.name}</p>
        <p className="text-green-600">{summary.topPerformer.gainLossPercent.toFixed(2)}%</p>
      </div>

      <div className="bg-white p-4 shadow rounded">
        <h4 className="text-gray-600 text-sm mb-1">Worst Performer</h4>
        <p className="text-lg font-semibold">{summary.worstPerformer.name}</p>
        <p className="text-red-600">{summary.worstPerformer.gainLossPercent.toFixed(2)}%</p>
      </div>

      <div className="bg-white p-4 shadow rounded">
        <h4 className="text-gray-600 text-sm mb-1">Portfolio Insights</h4>
        <p>Diversification Score: {summary.diversificationScore}/10</p>
        <p>Risk Level: {summary.riskLevel}</p>
      </div>
    </div>
  );
};

export default TopPerformers;
