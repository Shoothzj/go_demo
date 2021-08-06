### command
```bash
 cat cn.json|grep name|awk -F ':' '{print $2}'|tr '\n' ' '
```