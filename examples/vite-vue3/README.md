1. on dns server add record like 
```
*.component.com
```

1. on package.json change on yourself os, but current time tested on linux x64

```
"pathserve": "pathserve-amd64-linux"
or
"pathserve": "pathserve-amd64-mac"
or
"pathserve": "pathserve-amd64-windows"
```

2. run command
```
npm run dev
```

3. run command on different terminal
```
npm run pathserve
```

4. open url

```
http://book.index.component.com:8085
```