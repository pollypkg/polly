import {DashboardServiceClient} from "./proto/DashboardServiceClientPb"
import {Error as GrpcError} from "grpc-web"

/**
 * primary address of the backend grpc-web api
 */
export const addr =
  process.env.NODE_ENV === "development"
    ? "http://localhost:3333/api"
    : "/api";

export const dashboards = new DashboardServiceClient(addr)

export type Error = GrpcError | {}
export const isError = (T: GrpcError | {}): T is GrpcError => {
  if ((T as GrpcError).code) {
    return true;
  }

  return false;
};