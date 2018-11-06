# s3md5

Calculate MD5 hash for files uploaded (with Multipart Upload API) to Amazon S3 
compatible services.

### Why?

When using the multipart upload API, Amazon (and other services like Digital Ocean
Spaces) use a modified MD5 calculation as the ETag of the file. This tool calculate
that MD5 hash so you can validate the checksum of your downloads or uploads.

### How to use?

```bash
# Usage of s3md5:
#   -file string
#         file name to check
#   -size int
#        chunk size in MB (default 15)

$ s3cmd -file my_big_file -size 15
```

### How it works?

The algorithm is basically an MD5 of all MD5s of all generated file parts.


### Thank you

This is inspired by https://github.com/antespi/s3md5 which does the same thing 
(probably better) in bash.

It does run a bit faster then bash, but the point of this is not to be the
fastests possible implementation (the goal was to be as simple as possible).


```bash
# test_file has 901605836 bytes (around 880 MB)

# using BASH s3md5 (2.8 seconds)
$ time s3md5 15 test_file 
231c7a0ea512ea6ce07e771bd287cf1c-58

real	0m2.881s
user	0m2.268s
sys	0m1.077s

# using Go s3md5 (1.3 second)
$ time s3md5 -file test_file -size 15
231c7a0ea512ea6ce07e771bd287cf1c-58

real	0m1.332s
user	0m1.167s
sys	0m0.165s
```
