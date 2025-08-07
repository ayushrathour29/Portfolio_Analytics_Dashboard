import React from "react";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from "recharts";

const PerformanceCharts = ({ performance }) => {
  return (
    <div className="bg-white p-4 shadow rounded">
      <h3 className="text-lg font-semibold mb-4">Performance Comparison</h3>
      <ResponsiveContainer width="100%" height={300}>
        <LineChart data={performance.timeline}>
          <XAxis dataKey="date" />
          <YAxis />
          <Tooltip />
          <Legend />
          <Line type="monotone" dataKey="portfolio" stroke="#8884d8" />
          <Line type="monotone" dataKey="nifty50" stroke="#82ca9d" />
          <Line type="monotone" dataKey="gold" stroke="#ffc658" />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
};

export default PerformanceCharts;
