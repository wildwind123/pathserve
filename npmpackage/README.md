# Proxypath

Alternative of storybook. Is alpha, experimental project. Maybe someone will create project like this.

This app required editable dns server, you should add to dns like this record 

```
*.component.com
```

## config
Add in root dir config, vite vue3

```
# pathserve.config.yaml
server:
  host: :8085
handler_configs:
  book:
    handler: index
    work_dir: node_modules/pathserve/client
  ps:
    handler: vite
    work_dir: './'
    watch_dir: './src'
    params:
      params:
      html_template: node_modules/pathserve/templates/vue3-template.html
      script_template: node_modules/pathserve/templates/vue3-template.ts
      auto_gen_dir: node_modules/pathserve/autogen
      dir_public: public
      host: http://localhost:5023
host_params:
  index: node_modules/pathserve/client/index.html
```
on package.json

```
  "scripts": {
     ...
    "pathserve": "pathserve-amd64-linux"
  },
```


Tested on linux

## Examples

```
https://github.com/wildwind123/pathserve
```