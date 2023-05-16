import { Listener } from "./Listener";
import { Message as M } from "./Message";
import { cloneDeep as _cloneDeep } from "lodash";

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

export interface Message extends M {}

export let useString = (
  uniqueName: string,
  value: string,
  newValueHook: (newValue: string) => void
) => {
  let msg = {
    fromPathServe: true,
    scenario: "setControlValue",
    name: uniqueName,
    Data: {
      type: "string",
      data: value,
    },
    Form: {
      element: "input",
      type: "text",
    },
  } as Message;
  let messenger = new Messenger(msg);
  sendMessageAndSetListener(msg)

  window.pathServeMessageListener.addHook(
    uniqueName,
    (event: MessageEvent<Message>) => {
      if (
        !event.data.fromPathServe ||
        event.data.name != messenger.state.name
      ) {
        return;
      }
      newValueHook(event.data.Data!.data as string);
    }
  );

  const setValue = (value: string) => {
    messenger.setState({
      Data: { type: "string", data: value },
      scenario: "setControlValue",
    });
    newValueHook(value);
  };

  messenger.subscribe((message: Message) => {
    sendMessage(message);
  });
  return [setValue];
};

export let useNumber = (
  uniqueName: string,
  value: number,
  newValueHook: (newValue: number) => void
) => {
  let msg = {
    fromPathServe: true,
    scenario: "setControlValue",
    name: uniqueName,
    Data: {
      type: "number",
      data: value,
    },
    Form: {
      element: "input",
      type: "number",
    },
  } as Message;
  let messenger = new Messenger(msg);
  sendMessageAndSetListener(msg)

  window.pathServeMessageListener.addHook(
    uniqueName,
    (event: MessageEvent<Message>) => {
      if (
        !event.data.fromPathServe ||
        event.data.name != messenger.state.name
      ) {
        return;
      }
      newValueHook(event.data.Data!.data as number);
    }
  );

  const setValue = (value: number) => {
    messenger.setState({
      Data: { type: "number", data: value },
      scenario: "setControlValue",
    });
    newValueHook(value);
  };

  messenger.subscribe((message: Message) => {
    sendMessage(message);
  });
  return [setValue];
};

export let useObject = (
  uniqueName: string,
  value: object,
  newValueHook: (newValue: object) => void
) => {
  let msg = {
    fromPathServe: true,
    scenario: "setControlValue",
    name: uniqueName,
    Data: {
      type: "object",
      data: value,
    },
    Form: {
      element: "input",
      type: "textarea",
    },
  } as Message;
  let messenger = new Messenger(msg);

  sendMessageAndSetListener(msg)

  window.pathServeMessageListener.addHook(
    uniqueName,
    (event: MessageEvent<Message>) => {
      if (
        !event.data.fromPathServe ||
        event.data.name != messenger.state.name
      ) {
        return;
      }
      newValueHook(event.data.Data!.data as object);
    }
  );

  const setValue = (value: object) => {
    const clonedValue = _cloneDeep(value);
    messenger.setState({
      Data: { type: "object", data: clonedValue },
      scenario: "setControlValue",
    });
    newValueHook(clonedValue);
  };

  messenger.subscribe((message: Message) => {
    sendMessage(message);
  });
  return [setValue];
};

export let useButton = (uniqueName: string, clickedHook: () => void) => {
  let msg = {
    fromPathServe: true,
    scenario: "setControlValue",
    name: uniqueName,
    Data: null,
    Form: {
      element: "button",
    },
  } as Message;
  let messenger = new Messenger(msg);

  sendMessageAndSetListener(msg)
  
  window.pathServeMessageListener.addHook(
    uniqueName,
    (event: MessageEvent<Message>) => {
      if (
        !event.data.fromPathServe ||
        event.data.name != messenger.state.name
      ) {
        return;
      }
      clickedHook();
    }
  );

  messenger.subscribe((message: Message) => {
    sendMessage(message);
  });
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

const sendMessageAndSetListener = (message: Message) => {
  setListenerVariable()
  sendMessage(message)
}

const setListenerVariable = () => {
  if (!window.pathServeMessageListener) {
    let msg = {
      fromPathServe: true,
      scenario: "clearControl",
      name: "",
    } as Message;
    sendMessage(msg)
    window.pathServeMessageListener = new Listener();
  }
};

const sendMessage = (message: Message) => {
  try {
    window.parent.postMessage(message, "*");
  } catch (e) {
    console.error(`can't send message to parent `, message, e);
  }
};
