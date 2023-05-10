1. on dns server add record like 
```
*.component.com
```

2. on package.json change on yourself os, but current time tested on linux x64

```
"pathserve": "pathserve-amd64-linux"
or
"pathserve": "pathserve-amd64-mac"
or
"pathserve": "pathserve-amd64-windows"
```

3. run command
```
npm run dev
```

4. run command on different terminal
```
npm run pathserve
```

5. open url

```
http://book.index.component.com:8085
```