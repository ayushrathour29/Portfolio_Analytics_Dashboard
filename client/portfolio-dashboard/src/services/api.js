import axios from "axios";

const BASE_URL = "http://localhost:8080/api/portfolio";

export const getHoldings = () => axios.get(`${BASE_URL}/holdings`);
export const getAllocation = () => axios.get(`${BASE_URL}/allocation`);
export const getPerformance = () => axios.get(`${BASE_URL}/performance`);
export const getSummary = () => axios.get(`${BASE_URL}/summary`);
