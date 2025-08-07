import React from "react";
import { PieChart, Pie, Cell, Tooltip, Legend, ResponsiveContainer } from "recharts";

const COLORS = [
  "#8884d8", "#82ca9d", "#ffc658", "#ff7f7f", "#7fc7ff",
  "#a4de6c", "#d0ed57", "#f78ca2", "#8dd1e1", "#d88884"
];

const pieData = (obj) =>
  Object.entries(obj).map(([key, value]) => ({
    name: key,
    value: value.value,
    percentage: value.percentage,
  }));

const AllocationCharts = ({ allocation }) => {
  const sectorData = pieData(allocation.bySector);
  const capData = pieData(allocation.byMarketCap);

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 gap-10">
      {/* Sector Allocation */}
      <div className="bg-white p-4 shadow rounded">
        <h3 className="text-lg font-semibold mb-2 text-center">Sector Allocation</h3>
        <ResponsiveContainer width="100%" height={350}>
          <PieChart>
            <Pie
              data={sectorData}
              dataKey="value"
              cx="50%"
              cy="50%"
              outerRadius={110}
              label={({ name }) => name}
              labelLine={false}
            >
              {sectorData.map((_, index) => (
                <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
              ))}
            </Pie>
            <Tooltip />
            <Legend
              verticalAlign="bottom"
              height={70}
              layout="horizontal"
              wrapperStyle={{ fontSize: "0.8rem" }}
            />
          </PieChart>
        </ResponsiveContainer>
      </div>

      {/* Market Cap Allocation */}
      <div className="bg-white p-4 shadow rounded">
        <h3 className="text-lg font-semibold mb-2 text-center">Market Cap Allocation</h3>
        <ResponsiveContainer width="100%" height={350}>
          <PieChart>
            <Pie
              data={capData}
              dataKey="value"
              cx="50%"
              cy="50%"
              outerRadius={110}
              label={({ name }) => name}
              labelLine={false}
            >
              {capData.map((_, index) => (
                <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
              ))}
            </Pie>
            <Tooltip />
            <Legend
              verticalAlign="bottom"
              height={70}
              layout="horizontal"
              wrapperStyle={{ fontSize: "0.8rem" }}
            />
          </PieChart>
        </ResponsiveContainer>
      </div>
    </div>
  );
};

export default AllocationCharts;
