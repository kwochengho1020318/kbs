### Server 要做的準備
* 在gitlab 加入 ssh key 不然要輸密碼
* 更改權限 不然也要輸密碼
```
cd .git/objects

sudo chown -R yourname:yourgroup *
```