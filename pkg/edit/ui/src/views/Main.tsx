import React from "react";
import { ActiveDashboardList } from "../components/DashboardList";

const Main = () => (
  <div style={{ width: "100%", display: "flex", justifyContent: "center" }}>
    <div style={{ maxWidth: "1000px", width: "100vw" }}>
      <h1>Polly ODE</h1>

      <h2>Grafana</h2>
      <h3>Dashboards</h3>
      <ActiveDashboardList />
    </div>
  </div>
);

export default Main;
