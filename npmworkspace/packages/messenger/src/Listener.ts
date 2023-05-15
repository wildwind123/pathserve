import { Message } from "./Message";

export class Listener { 
    hooks : ((event: MessageEvent)=> void)[] = []
    scenario = ""
    setListener(scenario : Message['scenario'] = 'setComponentValue') {
        this.scenario = scenario
        window.addEventListener("message", window.pathServeMessageListener.callHooks.bind(this), false);
    }

    callHooks(event: MessageEvent<Message>) {
        if (!window.pathServeMessageListener.hooks) {
            return;
        }
        if (!event.data.fromPathServe) {
            return
        }
        if (this.scenario != "" && this.scenario != event.data.scenario) {
            return
        }

        for (let i = 0; i < window.pathServeMessageListener.hooks.length; i++) {
            window.pathServeMessageListener.hooks[i](event)
        }
    }

    removeListener() {
        window.removeEventListener("message", window.pathServeMessageListener.callHooks.bind(this), false);
    }
}