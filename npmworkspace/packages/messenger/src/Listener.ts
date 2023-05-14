
export class Listener { 
    hooks : ((event: MessageEvent)=> void)[] = []

    setListener() {
        window.addEventListener("message", window.pathServeMessageListener.callHooks, false);
    }

    callHooks(event: MessageEvent) {
        if (!window.pathServeMessageListener.hooks) {
            return;
        }
        for (let i = 0; i < window.pathServeMessageListener.hooks.length; i++) {
            window.pathServeMessageListener.hooks[i](event)
        }
    }

    removeListener() {
        window.removeEventListener("message", window.pathServeMessageListener.callHooks, false);
    }
}