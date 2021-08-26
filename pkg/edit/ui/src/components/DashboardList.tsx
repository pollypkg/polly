import React, { useEffect, useState } from "react";
import { dashboards, Error, isError } from "../api";

import {
  DashboardMeta,
  ListDashboardResponse,
  EditDashboardRequest,
} from "../proto/dashboard_pb";

type ItemProps = {
  data: DashboardMeta;
};

/**
 * Item shows a single Dashboard with an Edit button
 */
const Item = ({ data }: ItemProps) => {
  const edit = async () => {
    const req = new EditDashboardRequest();
    req.setName(data.getName());

    try {
      const res = await dashboards.edit(req, {});
      window.open(res.getEditurl());
    } catch (error) {
      if (isError(error)) {
        window.alert(`Error: ${error.message}`);
      }
    }
  };

  return (
    <li>
      {data.getTitle()} ({data.getFile()}) <button onClick={edit}>Edit</button>
    </li>
  );
};

export type DashboardListProps = {
  data: DashboardMeta[];
};

/**
 * DashboardList shows a list of Dashboards with respective Edit links
 */
export const DashboardList = ({ data }: DashboardListProps) => (
  <ul>
    {data.map((d) => (
      <Item key={d.getUid()} data={d} />
    ))}
  </ul>
);

/**
 * ActiveDashboardList acts like DashboardList, but fetches its own data from the api
 */
export const ActiveDashboardList = () => {
  let emptyMeta: DashboardMeta[] = [];
  let emptyError: Error = {};

  const [meta, setMeta] = useState(emptyMeta);
  const [fetching, setFetching] = useState(false);
  const [error, setError] = useState(emptyError);

  useEffect(() => {
    const load = async () => {
      setFetching(true);
      try {
        const res = await dashboards.list(new ListDashboardResponse(), {});
        setMeta(res.getDashboardsList());
        setError({});
      } catch (err) {
        setError(err);
      }
      setFetching(false);
    }
    load();
  }, []);

  if (fetching) {
    return <div>Loading Dashboards ..</div>;
  }

  if (isError(error)) {
    return <div>{`Error: ${error.message}`}</div>;
  }

  return <DashboardList data={meta} />;
};
