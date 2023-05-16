import { Message } from "./Message";

export class Listener { 
    hooks : Map<string, (event: MessageEvent) => void> | undefined 
    scenario = ""
    listenerIsSet = false
    setListener(scenario : Message['scenario'] = 'setComponentValue') {
        if (this.listenerIsSet) {
            return
        }
        this.scenario = scenario
        window.addEventListener("message", this.callHooks.bind(this), false);
        this.listenerIsSet = true
    }

    setHooksIfNotDefined() {
        if (!this.hooks) {
            this.hooks = new Map([])
        }
    }

    callHooks(event: MessageEvent<Message>) {
        if (!this.hooks) {
            return;
        }
        if (!event.data.fromPathServe) {
            return
        }
        if (this.scenario != "" && this.scenario != event.data.scenario) {
            return
        }

        for (const x of this.hooks.keys()) {
            const hook = this.hooks.get(x)
            if (!hook) {
                console.error('cant find hook with name', event.data.name)
                return
            }
            hook(event)
        }
    }

    addHook(name: string, hook : (event: MessageEvent) => void) {
        this.setHooksIfNotDefined()
        if (!this.hooks) {
            console.error('cant add hook, this.hooks is not defined')
            return
        }
        this.hooks.set(name, hook)
    }

    removeListener() {
        window.removeEventListener("message", window.pathServeMessageListener.callHooks.bind(this), false);
    }
}