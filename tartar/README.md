#tartar

A simple tool that can create a tar archive with:

```
tartar -c -f filename.tar ./path1 /path/2 path/n
```

Or, extract archived files with
```
tartar -x -f filename.tar path/to/extract/files
```

To trigger gz compression simply add `.gz` to the file name:

```
tartar -c -f filename.tar.gz ./path1 /path/2 path/n
```