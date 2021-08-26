/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as dashboard_pb from './dashboard_pb';


export class DashboardServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoList = new grpcWeb.AbstractClientBase.MethodInfo(
    dashboard_pb.ListDashboardResponse,
    (request: dashboard_pb.ListDashboardsRequest) => {
      return request.serializeBinary();
    },
    dashboard_pb.ListDashboardResponse.deserializeBinary
  );

  list(
    request: dashboard_pb.ListDashboardsRequest,
    metadata: grpcWeb.Metadata | null): Promise<dashboard_pb.ListDashboardResponse>;

  list(
    request: dashboard_pb.ListDashboardsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: dashboard_pb.ListDashboardResponse) => void): grpcWeb.ClientReadableStream<dashboard_pb.ListDashboardResponse>;

  list(
    request: dashboard_pb.ListDashboardsRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: dashboard_pb.ListDashboardResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/DashboardService/List',
        request,
        metadata || {},
        this.methodInfoList,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/DashboardService/List',
    request,
    metadata || {},
    this.methodInfoList);
  }

  methodInfoEdit = new grpcWeb.AbstractClientBase.MethodInfo(
    dashboard_pb.EditDashboardResponse,
    (request: dashboard_pb.EditDashboardRequest) => {
      return request.serializeBinary();
    },
    dashboard_pb.EditDashboardResponse.deserializeBinary
  );

  edit(
    request: dashboard_pb.EditDashboardRequest,
    metadata: grpcWeb.Metadata | null): Promise<dashboard_pb.EditDashboardResponse>;

  edit(
    request: dashboard_pb.EditDashboardRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: dashboard_pb.EditDashboardResponse) => void): grpcWeb.ClientReadableStream<dashboard_pb.EditDashboardResponse>;

  edit(
    request: dashboard_pb.EditDashboardRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: dashboard_pb.EditDashboardResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/DashboardService/Edit',
        request,
        metadata || {},
        this.methodInfoEdit,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/DashboardService/Edit',
    request,
    metadata || {},
    this.methodInfoEdit);
  }

}

