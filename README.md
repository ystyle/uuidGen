### 使用说明
```
YSTYLE@Y-STYLE MINGW64 /d/Code/Go/uuid
$ uuid -h
Usage of uuid :
  -n int
        生成的uuid数量 ,默认值为 1 . (default 1)

YSTYLE@Y-STYLE MINGW64 /d/Code/Go/uuid
$ uuid
bff66ec6-0773-4caf-9b67-9dd1a0763eea
bff66ec607734caf9b679dd1a0763eea

YSTYLE@Y-STYLE MINGW64 /d/Code/Go/uuid
$ uuid -n 4
80f98f06123c4b2c94c18c723179af2c
7b0ed1f942b7494c8046a5b09328b5f8
8141388209cd466a8210e364f13425d3
0d1c2837a2324cd19d66ab42ab63c328
```

### 编译方式
1. 第一种
  ```shell
  export GOPATH=`pwd`
  go install uuidgen
  ./bin/uuidgen -h
  ```

2. 第二种
  ```shell
  export GOPATH=`pwd`
  go build -ldflags "-s -w" uuidgen
  ./uuidgen -h
  ```
>交叉编译 `export GOOS=linux`

>推荐用第二种方式, 生成的文件更小
