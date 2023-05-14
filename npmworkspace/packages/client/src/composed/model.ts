export interface Info {
  params: Param[];
}

export interface Param {
  path: string;
  key: string;
  handler_config: string;
}
