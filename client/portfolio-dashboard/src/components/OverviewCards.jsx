import React from "react";

const OverviewCards = ({ summary, holdingsCount }) => {
  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
      <div className="bg-white shadow p-4 rounded">
        <h3 className="text-sm text-gray-500">Total Portfolio Value</h3>
        <p className="text-2xl font-semibold text-blue-700">
          ₹{summary.totalValue.toLocaleString()}
        </p>
      </div>

      <div className="bg-white shadow p-4 rounded">
        <h3 className="text-sm text-gray-500">Total Gain / Loss</h3>
        <p
          className={`text-2xl font-semibold ${
            summary.totalGainLoss >= 0 ? "text-green-600" : "text-red-600"
          }`}
        >
          ₹{summary.totalGainLoss.toLocaleString()} (
          {summary.totalGainLossPercent.toFixed(2)}%)
        </p>
      </div>

      <div className="bg-white shadow p-4 rounded">
        <h3 className="text-sm text-gray-500">Performance %</h3>
        <p className="text-2xl font-semibold text-purple-600">
          {summary.totalGainLossPercent.toFixed(2)}%
        </p>
      </div>

      <div className="bg-white shadow p-4 rounded">
        <h3 className="text-sm text-gray-500">No. of Holdings</h3>
        <p className="text-2xl font-semibold">{holdingsCount}</p>
      </div>
    </div>
  );
};

export default OverviewCards;
