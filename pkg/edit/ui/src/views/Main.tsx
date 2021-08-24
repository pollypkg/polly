import React from "react";
import { ActiveDashboardList } from "../components/DashboardList";

const Main = () => (
  <div>
    <h1>Polly ODE</h1>

    <h3>Grafana Dashboards</h3>
    <ActiveDashboardList />
  </div>
);

export default Main;
