# ref: https://www.digitalocean.com/community/tutorials/openssl-essentials-working-with-ssl-certificates-private-keys-and-csrs
# 1. create private key and self signed certificate
dhcp-10-191-18-135:cert ronglucao$ openssl req        -newkey rsa:2048 -nodes -keyout localhost.key        -x509 -days 365 -out localhost.crt -subj "/C=CN/ST=Beijing/L=Beijing/O=Learn Golang/CN=localhost"
Generating a 2048 bit RSA private key
...................................................+++
...+++
writing new private key to 'localhost.key'
-----

# 2. check certificate entries
dhcp-10-191-18-135:cert ronglucao$ openssl x509 -text -noout -in localhost.crt
Certificate:
    Data:
        Version: 1 (0x0)
        Serial Number: 17590621234533135220 (0xf41e70d04eca7374)
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: C=CN, ST=Beijing, L=Beijing, O=Learn Golang, CN=localhost
        Validity
            Not Before: Apr 25 13:46:54 2019 GMT
            Not After : Apr 24 13:46:54 2020 GMT
        Subject: C=CN, ST=Beijing, L=Beijing, O=Learn Golang, CN=localhost
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:df:2e:25:7d:a0:ea:7a:03:c5:50:fe:eb:ff:2e:
                    86:84:9f:cf:bb:4e:29:ba:e4:f9:55:56:df:28:7e:
                    f0:e9:dc:8b:2a:82:9b:d3:e6:4a:8b:20:b9:87:c4:
                    d2:d1:0c:77:0c:cc:92:a8:40:1a:94:fc:84:87:e5:
                    8b:eb:a4:2a:00:fe:bf:3c:ec:98:dc:ac:37:c1:0d:
                    10:f8:30:62:3a:87:c6:9c:e2:88:0f:4c:17:b3:d3:
                    40:bb:34:00:b4:79:a0:ce:46:ae:5a:88:53:ab:c7:
                    e9:eb:49:d2:66:5b:83:72:e5:71:f5:01:49:48:80:
                    5a:91:d0:10:a5:5a:68:42:0c:ee:17:0a:43:c2:26:
                    09:28:75:a8:d3:43:f4:16:87:3f:6e:c5:00:d5:71:
                    71:e6:b8:e1:74:c8:4b:a4:2e:60:c2:f9:99:99:05:
                    6e:39:e4:a4:df:de:93:5e:4a:c8:37:0e:86:fb:77:
                    d5:27:3c:fd:97:d6:23:7b:14:79:63:f3:7c:97:4f:
                    b9:88:79:92:3b:e2:ca:fb:8f:07:e6:82:e2:7e:86:
                    a0:52:1e:0e:a7:85:bb:38:cb:15:41:23:79:12:6e:
                    33:19:3e:80:e6:6d:3d:a5:63:9c:4c:33:46:c3:e8:
                    9a:e0:80:58:30:32:68:9d:e9:bc:4c:41:1d:d7:9a:
                    10:43
                Exponent: 65537 (0x10001)
    Signature Algorithm: sha256WithRSAEncryption
         5a:b6:6e:71:74:e6:af:9d:f5:8a:48:a7:6f:9a:3c:fe:73:f4:
         93:5f:f7:1f:da:76:06:21:2e:fb:f2:af:96:86:7f:36:b7:04:
         2c:37:10:66:92:6c:34:9d:eb:b4:e6:0e:7f:f9:92:4d:32:db:
         e2:d4:6a:55:de:8b:9d:21:44:4d:8a:1f:4a:61:0f:b2:40:4d:
         26:f5:9a:59:88:e2:f4:e4:d9:27:90:24:17:ad:ea:24:08:e8:
         24:da:36:e5:24:14:81:90:60:9c:02:89:05:56:4e:fd:2b:87:
         91:ec:64:36:4b:02:bc:f8:16:ab:73:f6:11:57:89:36:0d:92:
         ee:3c:f4:fd:e9:54:ba:01:89:49:7d:75:d8:ef:9d:35:69:a6:
         2a:45:6a:b5:0b:c4:ec:a3:6e:b6:35:32:39:20:b1:58:62:b1:
         69:8f:d2:a7:a1:7d:0e:ca:ed:b9:fb:8c:28:f3:3b:82:18:f4:
         53:c2:d1:c6:a0:c9:7b:61:0a:32:53:a0:6f:73:c1:b0:00:0c:
         db:54:84:53:00:9e:c5:79:73:69:ed:a7:ef:24:00:f5:85:26:
         a3:9b:a5:24:6d:89:f3:73:a3:6b:23:a6:83:c1:5f:69:c6:4a:
         16:b5:b8:25:56:18:3f:7f:19:14:31:01:81:13:30:3c:ac:e4:
         e7:7b:79:a6

# 3. check private key
dhcp-10-191-18-135:cert ronglucao$ openssl rsa -check -in localhost.key
RSA key ok
writing RSA key
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA3y4lfaDqegPFUP7r/y6GhJ/Pu04puuT5VVbfKH7w6dyLKoKb
0+ZKiyC5h8TS0Qx3DMySqEAalPyEh+WL66QqAP6/POyY3Kw3wQ0Q+DBiOofGnOKI
D0wXs9NAuzQAtHmgzkauWohTq8fp60nSZluDcuVx9QFJSIBakdAQpVpoQgzuFwpD
wiYJKHWo00P0Foc/bsUA1XFx5rjhdMhLpC5gwvmZmQVuOeSk396TXkrINw6G+3fV
Jzz9l9YjexR5Y/N8l0+5iHmSO+LK+48H5oLifoagUh4Op4W7OMsVQSN5Em4zGT6A
5m09pWOcTDNGw+ia4IBYMDJonem8TEEd15oQQwIDAQABAoIBAQCcre/h/2s9V5JZ
p22Ui1wilXMnotsG+yLTt5YY/u0wsQT+Rg0RuRXSEJpPZ209F4wobyHyfnZq69A6
+3q8Zlaatj0Zj1xNZ5YIsJyeMJF1V8amcx5j5t6o/wDtq7dm/BBZYCOdKSHccg+V
gBD2jP5+Vfgnscx52qo6vaBCBXM2Ej+QcdUGs+hygdqGxVhzTMo7hlccNYHw6+aL
I0/vjNLyzGg0Y0ei2pTH0DiZ2P39MtZvcCKvUHYmXQp+fF5D2zu0snNtiaXE/6uA
yAu/uCOjS/edeQ90BdWF4VE/JmHd4tlBawMfc+/nr7NEblfh/Gf7Oa50oDETa00w
HAUQUfrRAoGBAP1h5QuYQR2f6n0cCx1nOhuK3hD0drnpu74sghy6ku2cVci9I0Vi
s8b7rF9MNFBXAKSnbf7GHxu9qFpkC14f2g+p0jZGxz9bihLrL0FxIXj1iGlH2Z6r
+b1AoTD+XrNNVADSE378MK/GatRyOk9OqsGCQ9mFnl4u0TN+8vGMehXrAoGBAOF8
YMDBJPCpQ70AtUu6ashs40RImi4j3zBoaz7ZVV9VZ4i4nBcfqCnpcW0hYsUjPIqW
A8K9tWV9SN9azUbdgr19uQakeHdZ51SUlJCSPJkVcHxqsgutdr1X97r2FI/iqodz
w1P7uWlHHie56gfgOiB6z1suojsfTG7mDxu0dCEJAoGBAJ7NZfHYYKcifKIgm064
TZDJfdf+fxKRzNqppno+7KsC2jjPYXWxIJ9LSIMJjZf0jzCixqtwnDqUqRjNrto0
+EPs2RSvU10AEA8/WwSW1LWsnOvu3hM2EXVtNhkws0WI52cEQrfJcvIXVwkC9kyS
Ly7MZ53CwrigevjdTYHZ2wI1AoGAfyIYjzWXNVXKSotWrN7rjBvQu9RPkgpJscp+
BU6WkwzdNjoYT/VaZF8pw8UaIlqurNjOQAOkfhd4ee6BZR5Js7tqEjOzdF+tTYQE
i8rX/dwXsx6ZSnpAQ5uQospSZ5n+WibMD9MTBCfD+dQZ/tzPCTcVKK0PbVRFTRXL
mOxSubkCgYEAs5fUCnaIvfK2knamzQ8oWHCZd++Zs9zt/9BDDOcL/dpo+3BdoU+m
Ij6cwBUykl99j4xo+MyZLdy8s7mRe2vZiJPZ+JcAvdIEM44ItFqT6Df9024Epase
bVgIN+pt5EqLuLrJD9OkaKipEVn7yJdXM4XMP5wOfZeSZ0hj+n0iLKs=
-----END RSA PRIVATE KEY-----

# 4. check if private key and certificate are related
dhcp-10-191-18-135:cert ronglucao$ openssl rsa -noout -modulus -in localhost.crt | openssl md5
unable to load Private Key
4600551020:error:09FFF06C:PEM routines:CRYPTO_internal:no start line:/BuildRoot/Library/Caches/com.apple.xbs/Sources/libressl/libressl-22.200.4/libressl-2.6/crypto/pem/pem_lib.c:683:Expecting: ANY PRIVATE KEY
d41d8cd98f00b204e9800998ecf8427e
dhcp-10-191-18-135:cert ronglucao$ openssl rsa -noout -modulus -in localhost.csr | openssl md5
unable to load Private Key
4632618604:error:09FFF06C:PEM routines:CRYPTO_internal:no start line:/BuildRoot/Library/Caches/com.apple.xbs/Sources/libressl/libressl-22.200.4/libressl-2.6/crypto/pem/pem_lib.c:683:Expecting: ANY PRIVATE KEY
d41d8cd98f00b204e9800998ecf8427e
dhcp-10-191-18-135:cert ronglucao$ ls -rlth
total 56
-rw-r--r--@ 1 ronglucao  staff   2.0K Apr 25 16:31 create-certs-recordings.md
-rw-r--r--  1 ronglucao  staff   1.7K Apr 25 21:46 localhost.key
-rw-r--r--  1 ronglucao  staff   1.2K Apr 25 21:46 localhost.crt
