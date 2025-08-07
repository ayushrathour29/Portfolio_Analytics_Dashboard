import React from "react";

const HoldingsTables = ({ holdings }) => {
  return (
    <div className="overflow-auto bg-white shadow rounded-lg">
      <table className="min-w-full text-sm text-left">
        <thead className="bg-gray-100 text-xs uppercase">
          <tr>
            <th className="px-4 py-2">Symbol</th>
            <th className="px-4 py-2">Name</th>
            <th className="px-4 py-2">Quantity</th>
            <th className="px-4 py-2">Avg Price</th>
            <th className="px-4 py-2">Current Price</th>
            <th className="px-4 py-2">Value</th>
            <th className="px-4 py-2">Gain/Loss</th>
            <th className="px-4 py-2">Gain %</th>
          </tr>
        </thead>
        <tbody>
          {holdings.map((h) => (
            <tr key={h.symbol} className="border-t hover:bg-gray-50">
              <td className="px-4 py-2">{h.symbol}</td>
              <td className="px-4 py-2">{h.name}</td>
              <td className="px-4 py-2">{h.quantity}</td>
              <td className="px-4 py-2">₹{h.avgPrice}</td>
              <td className="px-4 py-2">₹{h.currentPrice}</td>
              <td className="px-4 py-2">₹{h.value.toFixed(2)}</td>
              <td
                className={`px-4 py-2 ${
                  h.gainLoss >= 0 ? "text-green-600" : "text-red-600"
                }`}
              >
                ₹{h.gainLoss.toFixed(2)}
              </td>
              <td
                className={`px-4 py-2 ${
                  h.gainLossPercent >= 0 ? "text-green-600" : "text-red-600"
                }`}
              >
                {h.gainLossPercent.toFixed(2)}%
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default HoldingsTables;
