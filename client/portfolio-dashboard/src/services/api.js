import axios from "axios";

const BASE_URL = "https://portfolio-analytics-dashboard-2.onrender.com";

export const getHoldings = () => axios.get(`${BASE_URL}/holdings`);
export const getAllocation = () => axios.get(`${BASE_URL}/allocation`);
export const getPerformance = () => axios.get(`${BASE_URL}/performance`);
export const getSummary = () => axios.get(`${BASE_URL}/summary`);
