import { Listener } from "./Listener";
import { Message as M } from "./Message";

import { v4 as uuidv4 } from "uuid";

declare global {
  interface Window {
    pathServeMessageListener: Listener;
  }
}

class Messenger {
  state: Message;
  subscribers: Set<(state: Message) => void>;

  constructor(initialState: Message) {
    this.state = initialState;
    this.subscribers = new Set();
  }

  setState(newState: any) {
    this.state = { ...this.state, ...newState };
    this.notifySubscribers();
  }

  subscribe(callback: (state: Message) => void) {
    this.subscribers.add(callback);
    return () => {
      this.subscribers.delete(callback);
    };
  }

  notifySubscribers() {
    for (const callback of this.subscribers) {
      callback(this.state);
    }
  }
}

export interface Message extends M{}

export let useString = (value: string , name : string, newValueHook: (newValue: string) => void) => {
  let key = uuidv4();
  let msg = {
    fromPathServe: true,
    scenario: "setParentValue",
    key: key,
    Data: {
      type: "string",
      data: value,
    },
    Form: {
      element: "input",
      type: "text",
      name: name
    },
  } as Message;
  let messenger = new Messenger(msg);
  sendMessage(msg);
  setListenerVariable()
  
  window.pathServeMessageListener!.hooks.push((event: MessageEvent<Message>) => {

    if (!event.data.fromPathServe || event.data.key != messenger.state.key) {
      return;
    }

    newValueHook(event.data.Data.data as string)
  });

  const setValue = (value: string) => {
    messenger.setState({ Data: { type: "string", data: value } });
  };

  messenger.subscribe((message: Message) => {
    sendMessage(message);
  });
  return {
    setValue,
    message: messenger.state,
  };
};

export const setListener = () => {
  if (!window.pathServeMessageListener) {
    console.error(
      "listener is null, call is method after document onload and after call method use*"
    );
    return;
  }
  window.pathServeMessageListener.setListener();
};

const setListenerVariable = () => {
  if (!window.pathServeMessageListener) {
    window.pathServeMessageListener = new Listener();
  }
};

const sendMessage = (message: Message) => {
  window.parent.postMessage(message, "*");
};
