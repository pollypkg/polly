import * as jspb from 'google-protobuf'



export class DashboardMeta extends jspb.Message {
  getFile(): string;
  setFile(value: string): DashboardMeta;

  getName(): string;
  setName(value: string): DashboardMeta;

  getTitle(): string;
  setTitle(value: string): DashboardMeta;

  getUid(): string;
  setUid(value: string): DashboardMeta;

  getDesc(): string;
  setDesc(value: string): DashboardMeta;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DashboardMeta.AsObject;
  static toObject(includeInstance: boolean, msg: DashboardMeta): DashboardMeta.AsObject;
  static serializeBinaryToWriter(message: DashboardMeta, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DashboardMeta;
  static deserializeBinaryFromReader(message: DashboardMeta, reader: jspb.BinaryReader): DashboardMeta;
}

export namespace DashboardMeta {
  export type AsObject = {
    file: string,
    name: string,
    title: string,
    uid: string,
    desc: string,
  }
}

export class ListDashboardsRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListDashboardsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListDashboardsRequest): ListDashboardsRequest.AsObject;
  static serializeBinaryToWriter(message: ListDashboardsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListDashboardsRequest;
  static deserializeBinaryFromReader(message: ListDashboardsRequest, reader: jspb.BinaryReader): ListDashboardsRequest;
}

export namespace ListDashboardsRequest {
  export type AsObject = {
  }
}

export class ListDashboardResponse extends jspb.Message {
  getDashboardsList(): Array<DashboardMeta>;
  setDashboardsList(value: Array<DashboardMeta>): ListDashboardResponse;
  clearDashboardsList(): ListDashboardResponse;
  addDashboards(value?: DashboardMeta, index?: number): DashboardMeta;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListDashboardResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListDashboardResponse): ListDashboardResponse.AsObject;
  static serializeBinaryToWriter(message: ListDashboardResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListDashboardResponse;
  static deserializeBinaryFromReader(message: ListDashboardResponse, reader: jspb.BinaryReader): ListDashboardResponse;
}

export namespace ListDashboardResponse {
  export type AsObject = {
    dashboardsList: Array<DashboardMeta.AsObject>,
  }
}

export class EditDashboardRequest extends jspb.Message {
  getName(): string;
  setName(value: string): EditDashboardRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EditDashboardRequest.AsObject;
  static toObject(includeInstance: boolean, msg: EditDashboardRequest): EditDashboardRequest.AsObject;
  static serializeBinaryToWriter(message: EditDashboardRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EditDashboardRequest;
  static deserializeBinaryFromReader(message: EditDashboardRequest, reader: jspb.BinaryReader): EditDashboardRequest;
}

export namespace EditDashboardRequest {
  export type AsObject = {
    name: string,
  }
}

export class EditDashboardResponse extends jspb.Message {
  getEditurl(): string;
  setEditurl(value: string): EditDashboardResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EditDashboardResponse.AsObject;
  static toObject(includeInstance: boolean, msg: EditDashboardResponse): EditDashboardResponse.AsObject;
  static serializeBinaryToWriter(message: EditDashboardResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EditDashboardResponse;
  static deserializeBinaryFromReader(message: EditDashboardResponse, reader: jspb.BinaryReader): EditDashboardResponse;
}

export namespace EditDashboardResponse {
  export type AsObject = {
    editurl: string,
  }
}

