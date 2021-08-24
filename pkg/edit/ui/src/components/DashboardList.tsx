import React, { useEffect, useState } from "react";
import { DashboardMeta, getDashboards, addr } from "../api";

type ItemProps = {
  data: DashboardMeta;
};

/**
 * Item shows a single Dashboard with an Edit link
 */
const Item = ({ data }: ItemProps) => (
  <li>
    {data.title} ({data.file}){" "}
    <a href={`${addr}/edit?dashboard=${data.name}`}>Edit</a>
  </li>
);

export type DashboardListProps = {
  data: DashboardMeta[];
};

/**
 * DashboardList shows a list of Dashboards with respective Edit links
 */
export const DashboardList = ({ data }: DashboardListProps) => (
  <ul>
    {data.map((d) => (
      <Item key={d.uid} data={d} />
    ))}
  </ul>
);

/**
 * ActiveDashboardList acts like DashboardList, but fetches its own data from the api
 */
export const ActiveDashboardList = () => {
  let empty: DashboardMeta[] = [];
  const [dashboards, setDashboards] = useState(empty);
  const [fetching, setFetching] = useState(false);

  useEffect(() => {
    const load = async () => {
      setFetching(true);
      try {
        const data = await getDashboards();
        setDashboards(data);
        setFetching(false);
      } catch (error) {
        console.log(error);
      }
    };
    load();
  }, []);

  if (fetching) {
    return <div>Loading Dashboards ..</div>;
  }

  return <DashboardList data={dashboards} />;
};
