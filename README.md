# numpy2go

This is a simple example of how to write a go libary (shared object) with exported functions that can accept numpy arrays as passed by ctypes. It requires go 1.5.

To try it run:
```bash
go build -buildmode=c-shared -o numpy2go.so
python numpy2go.py
```

Which should output:
```
Python says [ 0.  1.  2.]
Go says [0 1 2]
```

Code is based on these two examples:
https://scipy-lectures.github.io/advanced/interfacing_with_c/interfacing_with_c.html#id5
https://github.com/golang/go/wiki/cgo#turning-c-arrays-into-go-slices




