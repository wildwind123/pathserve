import { computed, reactive } from "vue";
import axios from "axios";
import { Info, Param } from "@/composed/model";

const storageKeySelectedKey = "pathproxy-client_selected-key"

export const getApiHost = () => {
  if (import.meta.env.VITE_API_HOST) {
    return import.meta.env.VITE_API_HOST;
  }
  const regex = /(http|https):\/\/([^.]*).([^.]*).(.*)$/gm;
  const m = regex.exec(document.location.origin)
  if (!m) {
    console.error(`cant parse host`, document.location.origin, m)
    return ""
  }
  if (!m[4]) {
    console.error(`cant parse host`, document.location.origin, m)
    return ""
  }
  return m[1] + '://api.' +m[4] ;
};

const info = reactive<{
  loading: boolean;
  params: Info["params"];
  selectedKye: string;
}>({
  loading: false,
  params: [],
  selectedKye: localStorage.getItem(storageKeySelectedKey) ?? ''
});

const selectedKey = computed({
  get() {
    return info.selectedKye;
  },
   set(newValue: string)  {
    localStorage.setItem(storageKeySelectedKey, newValue)
    info.selectedKye = newValue
   }
})

const selectedParam = computed(() => {
  if (!info.params) {
    return null
  }
  for (let i = 0; i < info.params.length; i++) {
    if (info.params[i].key == selectedKey.value) {
      return info.params[i];
    }
  }
  return null;
});

export const useApp = () => {
  const request = () => {
    return new Promise((resolve, reject) => {
      axios<Info>({
        url: getApiHost() + "/info",
      })
        .then((r) => {
          info.params = r.data.params;
          resolve(r);
        })
        .catch((e) => {
          console.error(e);
          reject(e);
        })
        .finally(() => (info.loading = false));
    });
  };

  return {
    info,
    request,
    selectedParam,
    selectedKey,
  };
};

export const paramUrl = (p: Param) => {
  return getApiHost().replace("api", p.handler_config + "." + p.key);
};
