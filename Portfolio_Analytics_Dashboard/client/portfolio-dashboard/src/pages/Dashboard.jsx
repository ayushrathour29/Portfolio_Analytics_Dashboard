import React, { useEffect, useState } from "react";
import {
  getHoldings,
  getAllocation,
  getPerformance,
  getSummary,
} from "../services/api";

import OverviewCards from "../components/OverviewCards";
import HoldingsTable from "../components/HoldingsTables";
import AllocationCharts from "../components/AllocationCharts";
import PerformanceChart from "../components/PerformanceCharts";
import TopPerformers from "../components/TopPerformers";

const Dashboard = () => {
  const [holdings, setHoldings] = useState([]);
  const [allocation, setAllocation] = useState(null);
  const [performance, setPerformance] = useState(null);
  const [summary, setSummary] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchAll = async () => {
      try {
        const [h, a, p, s] = await Promise.all([
          getHoldings(),
          getAllocation(),
          getPerformance(),
          getSummary(),
        ]);
        setHoldings(h.data);
        setAllocation(a.data);
        setPerformance(p.data);
        setSummary(s.data);
        setLoading(false);
      } catch (err) {
        console.error("Failed to fetch API:", err);
      }
    };
    fetchAll();
  }, []);

  if (loading) return <p className="text-center text-xl mt-10">Loading...</p>;

  return (
    <div className="p-6 space-y-10">
      <OverviewCards summary={summary} holdingsCount={holdings.length} />
      <AllocationCharts allocation={allocation} />
      <HoldingsTable holdings={holdings} />
      <PerformanceChart performance={performance} />
      <TopPerformers summary={summary} />
    </div>
  );
};

export default Dashboard;
