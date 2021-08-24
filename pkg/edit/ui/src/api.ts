import axios from "axios";

/**
 * primary address of the backend api
 */
export const addr =
  process.env.NODE_ENV === "development"
    ? "http://localhost:3333/api/v1"
    : "/api/v1";

/**
 * DashboardMeta represents metadata information about a dashboard
 */
export type DashboardMeta = {
  title: string;
  uid: string;
  description?: string;

  file: string;
  name: string;
};

/**
 * getDashboards fetches the list of dashboards the current Polly package has
 * @returns DashboardMeta[]
 */
export async function getDashboards(): Promise<DashboardMeta[]> {
  try {
    const result = await axios.get(`${addr}/dashboards`);
    return result.data;
  } catch (error) {
    throw error;
  }
}
