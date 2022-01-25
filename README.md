# fetchcert: easy fetching of cert chains from a server

Its like `openssl s_client`, except you don't have to lookup how to run it every time.

`fetchcert` will tell you if it thinks the cert chain is valid or not.

You can also specify a server name with the `-sni` flag.


## Usage

Fetch a cert for a domain name:
```
Dail google.com:443
-----BEGIN CERTIFICATE-----
MIINsTCCDJmgAwIBAgIRALThYcbfDJL1CgAAAAErfHwwDQYJKoZIhvcNAQELBQAw
RjELMAkGA1UEBhMCVVMxIjAgBgNVBAoTGUdvb2dsZSBUcnVzdCBTZXJ2aWNlcyBM
TEMxEzARBgNVBAMTCkdUUyBDQSAxQzMwHhcNMjExMjI3MDYwMjExWhcNMjIwMzIx
MDYwMjEwWjAXMRUwEwYDVQQDDAwqLmdvb2dsZS5jb20wWTATBgcqhkjOPQIBBggq
hkjOPQMBBwNCAAQ6GQZ05QpPgtfvyii4Sl8wt8dNams1L+5Xn7yVq7XELSx1gdUZ
hFe3+VmJYJiKUqW3h2BaX7KmVpkzloKnzgf5o4ILkjCCC44wDgYDVR0PAQH/BAQD
AgeAMBMGA1UdJQQMMAoGCCsGAQUFBwMBMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYE
FOcfhJDdj+n43YzPMZtiUkYRJ+bEMB8GA1UdIwQYMBaAFIp0f6+Fze6VzT2c0OJG
FPNxNR0nMGoGCCsGAQUFBwEBBF4wXDAnBggrBgEFBQcwAYYbaHR0cDovL29jc3Au
cGtpLmdvb2cvZ3RzMWMzMDEGCCsGAQUFBzAChiVodHRwOi8vcGtpLmdvb2cvcmVw
by9jZXJ0cy9ndHMxYzMuZGVyMIIJQgYDVR0RBIIJOTCCCTWCDCouZ29vZ2xlLmNv
bYIWKi5hcHBlbmdpbmUuZ29vZ2xlLmNvbYIJKi5iZG4uZGV2ghIqLmNsb3VkLmdv
b2dsZS5jb22CGCouY3Jvd2Rzb3VyY2UuZ29vZ2xlLmNvbYIYKi5kYXRhY29tcHV0
ZS5nb29nbGUuY29tggsqLmdvb2dsZS5jYYILKi5nb29nbGUuY2yCDiouZ29vZ2xl
LmNvLmlugg4qLmdvb2dsZS5jby5qcIIOKi5nb29nbGUuY28udWuCDyouZ29vZ2xl
LmNvbS5hcoIPKi5nb29nbGUuY29tLmF1gg8qLmdvb2dsZS5jb20uYnKCDyouZ29v
Z2xlLmNvbS5jb4IPKi5nb29nbGUuY29tLm14gg8qLmdvb2dsZS5jb20udHKCDyou
Z29vZ2xlLmNvbS52boILKi5nb29nbGUuZGWCCyouZ29vZ2xlLmVzggsqLmdvb2ds
ZS5mcoILKi5nb29nbGUuaHWCCyouZ29vZ2xlLml0ggsqLmdvb2dsZS5ubIILKi5n
b29nbGUucGyCCyouZ29vZ2xlLnB0ghIqLmdvb2dsZWFkYXBpcy5jb22CDyouZ29v
Z2xlYXBpcy5jboIRKi5nb29nbGV2aWRlby5jb22CDCouZ3N0YXRpYy5jboIQKi5n
c3RhdGljLWNuLmNvbYIPZ29vZ2xlY25hcHBzLmNughEqLmdvb2dsZWNuYXBwcy5j
boIRZ29vZ2xlYXBwcy1jbi5jb22CEyouZ29vZ2xlYXBwcy1jbi5jb22CDGdrZWNu
YXBwcy5jboIOKi5na2VjbmFwcHMuY26CEmdvb2dsZWRvd25sb2Fkcy5jboIUKi5n
b29nbGVkb3dubG9hZHMuY26CEHJlY2FwdGNoYS5uZXQuY26CEioucmVjYXB0Y2hh
Lm5ldC5jboILd2lkZXZpbmUuY26CDSoud2lkZXZpbmUuY26CEWFtcHByb2plY3Qu
b3JnLmNughMqLmFtcHByb2plY3Qub3JnLmNughFhbXBwcm9qZWN0Lm5ldC5jboIT
Ki5hbXBwcm9qZWN0Lm5ldC5jboIXZ29vZ2xlLWFuYWx5dGljcy1jbi5jb22CGSou
Z29vZ2xlLWFuYWx5dGljcy1jbi5jb22CF2dvb2dsZWFkc2VydmljZXMtY24uY29t
ghkqLmdvb2dsZWFkc2VydmljZXMtY24uY29tghFnb29nbGV2YWRzLWNuLmNvbYIT
Ki5nb29nbGV2YWRzLWNuLmNvbYIRZ29vZ2xlYXBpcy1jbi5jb22CEyouZ29vZ2xl
YXBpcy1jbi5jb22CFWdvb2dsZW9wdGltaXplLWNuLmNvbYIXKi5nb29nbGVvcHRp
bWl6ZS1jbi5jb22CEmRvdWJsZWNsaWNrLWNuLm5ldIIUKi5kb3VibGVjbGljay1j
bi5uZXSCGCouZmxzLmRvdWJsZWNsaWNrLWNuLm5ldIIWKi5nLmRvdWJsZWNsaWNr
LWNuLm5ldIIOZG91YmxlY2xpY2suY26CECouZG91YmxlY2xpY2suY26CFCouZmxz
LmRvdWJsZWNsaWNrLmNughIqLmcuZG91YmxlY2xpY2suY26CEWRhcnRzZWFyY2gt
Y24ubmV0ghMqLmRhcnRzZWFyY2gtY24ubmV0gh1nb29nbGV0cmF2ZWxhZHNlcnZp
Y2VzLWNuLmNvbYIfKi5nb29nbGV0cmF2ZWxhZHNlcnZpY2VzLWNuLmNvbYIYZ29v
Z2xldGFnc2VydmljZXMtY24uY29tghoqLmdvb2dsZXRhZ3NlcnZpY2VzLWNuLmNv
bYIXZ29vZ2xldGFnbWFuYWdlci1jbi5jb22CGSouZ29vZ2xldGFnbWFuYWdlci1j
bi5jb22CGGdvb2dsZXN5bmRpY2F0aW9uLWNuLmNvbYIaKi5nb29nbGVzeW5kaWNh
dGlvbi1jbi5jb22CJCouc2FmZWZyYW1lLmdvb2dsZXN5bmRpY2F0aW9uLWNuLmNv
bYIWYXBwLW1lYXN1cmVtZW50LWNuLmNvbYIYKi5hcHAtbWVhc3VyZW1lbnQtY24u
Y29tggtndnQxLWNuLmNvbYINKi5ndnQxLWNuLmNvbYILZ3Z0Mi1jbi5jb22CDSou
Z3Z0Mi1jbi5jb22CCzJtZG4tY24ubmV0gg0qLjJtZG4tY24ubmV0ghRnb29nbGVm
bGlnaHRzLWNuLm5ldIIWKi5nb29nbGVmbGlnaHRzLWNuLm5ldIIMYWRtb2ItY24u
Y29tgg4qLmFkbW9iLWNuLmNvbYINKi5nc3RhdGljLmNvbYIUKi5tZXRyaWMuZ3N0
YXRpYy5jb22CCiouZ3Z0MS5jb22CESouZ2NwY2RuLmd2dDEuY29tggoqLmd2dDIu
Y29tgg4qLmdjcC5ndnQyLmNvbYIQKi51cmwuZ29vZ2xlLmNvbYIWKi55b3V0dWJl
LW5vY29va2llLmNvbYILKi55dGltZy5jb22CC2FuZHJvaWQuY29tgg0qLmFuZHJv
aWQuY29tghMqLmZsYXNoLmFuZHJvaWQuY29tggRnLmNuggYqLmcuY26CBGcuY2+C
BiouZy5jb4IGZ29vLmdsggp3d3cuZ29vLmdsghRnb29nbGUtYW5hbHl0aWNzLmNv
bYIWKi5nb29nbGUtYW5hbHl0aWNzLmNvbYIKZ29vZ2xlLmNvbYISZ29vZ2xlY29t
bWVyY2UuY29tghQqLmdvb2dsZWNvbW1lcmNlLmNvbYIIZ2dwaHQuY26CCiouZ2dw
aHQuY26CCnVyY2hpbi5jb22CDCoudXJjaGluLmNvbYIIeW91dHUuYmWCC3lvdXR1
YmUuY29tgg0qLnlvdXR1YmUuY29tghR5b3V0dWJlZWR1Y2F0aW9uLmNvbYIWKi55
b3V0dWJlZWR1Y2F0aW9uLmNvbYIPeW91dHViZWtpZHMuY29tghEqLnlvdXR1YmVr
aWRzLmNvbYIFeXQuYmWCByoueXQuYmWCGmFuZHJvaWQuY2xpZW50cy5nb29nbGUu
Y29tghtkZXZlbG9wZXIuYW5kcm9pZC5nb29nbGUuY26CHGRldmVsb3BlcnMuYW5k
cm9pZC5nb29nbGUuY26CGHNvdXJjZS5hbmRyb2lkLmdvb2dsZS5jbjAhBgNVHSAE
GjAYMAgGBmeBDAECATAMBgorBgEEAdZ5AgUDMDwGA1UdHwQ1MDMwMaAvoC2GK2h0
dHA6Ly9jcmxzLnBraS5nb29nL2d0czFjMy9mVkp4YlYtS3Rtay5jcmwwggEEBgor
BgEEAdZ5AgQCBIH1BIHyAPAAdgBGpVXrdfqRIDC1oolp9PN9ESxBdL79SbiFq/L8
cP5tRwAAAX36smjIAAAEAwBHMEUCIQD++9b++xcJD1hRHAvi5w+dNFsg+Qz3oFXq
DN0GBijKXwIgHeTLxsb+Aych7+gOMcreX3ZpweeGvcBl9aCvVMYeIOkAdgBByMqx
3yJGShDGoToJQodeTjGLGwPr60vHaPCQYpYG9gAAAX36smbwAAAEAwBHMEUCIQCN
gIaWx3F7/WlA/bQ+bnoFrgNwPjte/Tnsk4R0TITzfgIgXpUVPlAFhmDHrCuOooiF
PdquxKE9rSGt75eVcx8HqKYwDQYJKoZIhvcNAQELBQADggEBAL952Ouz/2ynqw90
lE3i9tuYyG+zqOFXa+zxJT0gy6E0Oe5UI7FInL1pYRl1v0vPN4qkrBALamiftQ0A
M/6KDug71fB2NwYb8NO78idjbNpS0EMqjdWvjrAQxgDcqfAOhtTZkLm57PQPhzGE
HZr+mQQCvUybmwrjGDO4VyztdXOks4ioVJfbj3EPMhMwpNnNFVev1gZNJdSwMjOd
xfZcPvss3QWSzm4NEkNGiC2MnvWT8uONMRcVWemnaSrrUXl/Fb1rjjB9tRmQnJph
l88QveHp9An7CXgHol6IK+dj7NjkVfF/Y/N6O2tNyev++D/9rNuIfMEElQeMVWUQ
QQufuxk=
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIFljCCA36gAwIBAgINAgO8U1lrNMcY9QFQZjANBgkqhkiG9w0BAQsFADBHMQsw
CQYDVQQGEwJVUzEiMCAGA1UEChMZR29vZ2xlIFRydXN0IFNlcnZpY2VzIExMQzEU
MBIGA1UEAxMLR1RTIFJvb3QgUjEwHhcNMjAwODEzMDAwMDQyWhcNMjcwOTMwMDAw
MDQyWjBGMQswCQYDVQQGEwJVUzEiMCAGA1UEChMZR29vZ2xlIFRydXN0IFNlcnZp
Y2VzIExMQzETMBEGA1UEAxMKR1RTIENBIDFDMzCCASIwDQYJKoZIhvcNAQEBBQAD
ggEPADCCAQoCggEBAPWI3+dijB43+DdCkH9sh9D7ZYIl/ejLa6T/belaI+KZ9hzp
kgOZE3wJCor6QtZeViSqejOEH9Hpabu5dOxXTGZok3c3VVP+ORBNtzS7XyV3NzsX
lOo85Z3VvMO0Q+sup0fvsEQRY9i0QYXdQTBIkxu/t/bgRQIh4JZCF8/ZK2VWNAcm
BA2o/X3KLu/qSHw3TT8An4Pf73WELnlXXPxXbhqW//yMmqaZviXZf5YsBvcRKgKA
gOtjGDxQSYflispfGStZloEAoPtR28p3CwvJlk/vcEnHXG0g/Zm0tOLKLnf9LdwL
tmsTDIwZKxeWmLnwi/agJ7u2441Rj72ux5uxiZ0CAwEAAaOCAYAwggF8MA4GA1Ud
DwEB/wQEAwIBhjAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwEgYDVR0T
AQH/BAgwBgEB/wIBADAdBgNVHQ4EFgQUinR/r4XN7pXNPZzQ4kYU83E1HScwHwYD
VR0jBBgwFoAU5K8rJnEaK0gnhS9SZizv8IkTcT4waAYIKwYBBQUHAQEEXDBaMCYG
CCsGAQUFBzABhhpodHRwOi8vb2NzcC5wa2kuZ29vZy9ndHNyMTAwBggrBgEFBQcw
AoYkaHR0cDovL3BraS5nb29nL3JlcG8vY2VydHMvZ3RzcjEuZGVyMDQGA1UdHwQt
MCswKaAnoCWGI2h0dHA6Ly9jcmwucGtpLmdvb2cvZ3RzcjEvZ3RzcjEuY3JsMFcG
A1UdIARQME4wOAYKKwYBBAHWeQIFAzAqMCgGCCsGAQUFBwIBFhxodHRwczovL3Br
aS5nb29nL3JlcG9zaXRvcnkvMAgGBmeBDAECATAIBgZngQwBAgIwDQYJKoZIhvcN
AQELBQADggIBAIl9rCBcDDy+mqhXlRu0rvqrpXJxtDaV/d9AEQNMwkYUuxQkq/BQ
cSLbrcRuf8/xam/IgxvYzolfh2yHuKkMo5uhYpSTld9brmYZCwKWnvy15xBpPnrL
RklfRuFBsdeYTWU0AIAaP0+fbH9JAIFTQaSSIYKCGvGjRFsqUBITTcFTNvNCCK9U
+o53UxtkOCcXCb1YyRt8OS1b887U7ZfbFAO/CVMkH8IMBHmYJvJh8VNS/UKMG2Yr
PxWhu//2m+OBmgEGcYk1KCTd4b3rGS3hSMs9WYNRtHTGnXzGsYZbr8w0xNPM1IER
lQCh9BIiAfq0g3GvjLeMcySsN1PCAJA/Ef5c7TaUEDu9Ka7ixzpiO2xj2YC/WXGs
Yye5TBeg2vZzFb8q3o/zpWwygTMD0IZRcZk0upONXbVRWPeyk+gB9lm+cZv9TSjO
z23HFtz30dZGm6fKa+l3D/2gthsjgx0QGtkJAITgRNOidSOzNIb2ILCkXhAd4FJG
AJ2xDx8hcFH1mt0G/FX0Kw4zd8NLQsLxdxP8c4CU6x+7Nz/OAipmsHMdMqUybDKw
juDEI/9bfU1lcKwrmz3O2+BtjjKAvpafkmO8l7tdufThcV4q5O8DIrGKZTqPwJNl
1IXNDw9bg1kWRxYtnCQ6yICmJhSFm/Y3m6xv+cXDBlHz4n/FsRC6UfTd
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIFYjCCBEqgAwIBAgIQd70NbNs2+RrqIQ/E8FjTDTANBgkqhkiG9w0BAQsFADBX
MQswCQYDVQQGEwJCRTEZMBcGA1UEChMQR2xvYmFsU2lnbiBudi1zYTEQMA4GA1UE
CxMHUm9vdCBDQTEbMBkGA1UEAxMSR2xvYmFsU2lnbiBSb290IENBMB4XDTIwMDYx
OTAwMDA0MloXDTI4MDEyODAwMDA0MlowRzELMAkGA1UEBhMCVVMxIjAgBgNVBAoT
GUdvb2dsZSBUcnVzdCBTZXJ2aWNlcyBMTEMxFDASBgNVBAMTC0dUUyBSb290IFIx
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAthECix7joXebO9y/lD63
ladAPKH9gvl9MgaCcfb2jH/76Nu8ai6Xl6OMS/kr9rH5zoQdsfnFl97vufKj6bwS
iV6nqlKr+CMny6SxnGPb15l+8Ape62im9MZaRw1NEDPjTrETo8gYbEvs/AmQ351k
KSUjB6G00j0uYODP0gmHu81I8E3CwnqIiru6z1kZ1q+PsAewnjHxgsHA3y6mbWwZ
DrXYfiYaRQM9sHmklCitD38m5agI/pboPGiUU+6DOogrFZYJsuB6jC511pzrp1Zk
j5ZPaK49l8KEj8C8QMALXL32h7M1bKwYUH+E4EzNktMg6TO8UpmvMrUpsyUqtEj5
cuHKZPfmghCN6J3Cioj6OGaK/GP5Afl4/Xtcd/p2h/rs37EOeZVXtL0m79YB0esW
CruOC7XFxYpVq9Os6pFLKcwZpDIlTirxZUTQAs6qzkm06p98g7BAe+dDq6dso499
iYH6TKX/1Y7DzkvgtdizjkXPdsDtQCv9Uw+wp9U7DbGKogPeMa3Md+pvez7W35Ei
Eua++tgy/BBjFFFy3l3WFpO9KWgz7zpm7AeKJt8T11dleCfeXkkUAKIAf5qoIbap
sZWwpbkNFhHax2xIPEDgfg1azVY80ZcFuctL7TlLnMQ/0lUTbiSw1nH69MG6zO0b
9f6BQdgAmD06yK56mDcYBZUCAwEAAaOCATgwggE0MA4GA1UdDwEB/wQEAwIBhjAP
BgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBTkrysmcRorSCeFL1JmLO/wiRNxPjAf
BgNVHSMEGDAWgBRge2YaRQ2XyolQL30EzTSo//z9SzBgBggrBgEFBQcBAQRUMFIw
JQYIKwYBBQUHMAGGGWh0dHA6Ly9vY3NwLnBraS5nb29nL2dzcjEwKQYIKwYBBQUH
MAKGHWh0dHA6Ly9wa2kuZ29vZy9nc3IxL2dzcjEuY3J0MDIGA1UdHwQrMCkwJ6Al
oCOGIWh0dHA6Ly9jcmwucGtpLmdvb2cvZ3NyMS9nc3IxLmNybDA7BgNVHSAENDAy
MAgGBmeBDAECATAIBgZngQwBAgIwDQYLKwYBBAHWeQIFAwIwDQYLKwYBBAHWeQIF
AwMwDQYJKoZIhvcNAQELBQADggEBADSkHrEoo9C0dhemMXoh6dFSPsjbdBZBiLg9
NR3t5P+T4Vxfq7vqfM/b5A3Ri1fyJm9bvhdGaJQ3b2t6yMAYN/olUazsaL+yyEn9
WprKASOshIArAoyZl+tJaox118fessmXn1hIVw41oeQa1v1vg4Fv74zPl6/AhSrw
9U5pCZEt4Wi4wStz6dTZ/CLANx8LZh1J7QJVj2fhMtfTJr9w4z30Z209fOU0iOMy
+qduBmpvvYuR7hZL6Dupszfnw0Skfths18dG9ZKb59UhvmaSGZRVbNQpsg3BZlvi
d0lIKO2d1xozclOzgjXPYovJJIultzkMu34qQb9Sz/yilrbCgj8=
-----END CERTIFICATE-----
========================================
Serial Number : 240431289184393283661149297772776356988
Subject: CN=*.google.com
Issuer: CN=GTS CA 1C3,O=Google Trust Services LLC,C=US
Not Before: 2021-12-27 06:02:11 +0000 UTC
Not After: 2022-03-21 06:02:10 +0000 UTC
Subject Alt Names: *.google.com;*.appengine.google.com;*.bdn.dev;*.cloud.google.com;*.crowdsource.google.com;*.datacompute.google.com;*.google.ca;*.google.cl;*.google.co.in;*.google.co.jp;*.google.co.uk;*.google.com.ar;*.google.com.au;*.google.com.br;*.google.com.co;*.google.com.mx;*.google.com.tr;*.google.com.vn;*.google.de;*.google.es;*.google.fr;*.google.hu;*.google.it;*.google.nl;*.google.pl;*.google.pt;*.googleadapis.com;*.googleapis.cn;*.googlevideo.com;*.gstatic.cn;*.gstatic-cn.com;googlecnapps.cn;*.googlecnapps.cn;googleapps-cn.com;*.googleapps-cn.com;gkecnapps.cn;*.gkecnapps.cn;googledownloads.cn;*.googledownloads.cn;recaptcha.net.cn;*.recaptcha.net.cn;widevine.cn;*.widevine.cn;ampproject.org.cn;*.ampproject.org.cn;ampproject.net.cn;*.ampproject.net.cn;google-analytics-cn.com;*.google-analytics-cn.com;googleadservices-cn.com;*.googleadservices-cn.com;googlevads-cn.com;*.googlevads-cn.com;googleapis-cn.com;*.googleapis-cn.com;googleoptimize-cn.com;*.googleoptimize-cn.com;doubleclick-cn.net;*.doubleclick-cn.net;*.fls.doubleclick-cn.net;*.g.doubleclick-cn.net;doubleclick.cn;*.doubleclick.cn;*.fls.doubleclick.cn;*.g.doubleclick.cn;dartsearch-cn.net;*.dartsearch-cn.net;googletraveladservices-cn.com;*.googletraveladservices-cn.com;googletagservices-cn.com;*.googletagservices-cn.com;googletagmanager-cn.com;*.googletagmanager-cn.com;googlesyndication-cn.com;*.googlesyndication-cn.com;*.safeframe.googlesyndication-cn.com;app-measurement-cn.com;*.app-measurement-cn.com;gvt1-cn.com;*.gvt1-cn.com;gvt2-cn.com;*.gvt2-cn.com;2mdn-cn.net;*.2mdn-cn.net;googleflights-cn.net;*.googleflights-cn.net;admob-cn.com;*.admob-cn.com;*.gstatic.com;*.metric.gstatic.com;*.gvt1.com;*.gcpcdn.gvt1.com;*.gvt2.com;*.gcp.gvt2.com;*.url.google.com;*.youtube-nocookie.com;*.ytimg.com;android.com;*.android.com;*.flash.android.com;g.cn;*.g.cn;g.co;*.g.co;goo.gl;www.goo.gl;google-analytics.com;*.google-analytics.com;google.com;googlecommerce.com;*.googlecommerce.com;ggpht.cn;*.ggpht.cn;urchin.com;*.urchin.com;youtu.be;youtube.com;*.youtube.com;youtubeeducation.com;*.youtubeeducation.com;youtubekids.com;*.youtubekids.com;yt.be;*.yt.be;android.clients.google.com;developer.android.google.cn;developers.android.google.cn;source.android.google.cn
========================================
Serial Number : 159612451717983579589660725350
Subject: CN=GTS CA 1C3,O=Google Trust Services LLC,C=US
Issuer: CN=GTS Root R1,O=Google Trust Services LLC,C=US
Not Before: 2020-08-13 00:00:42 +0000 UTC
Not After: 2027-09-30 00:00:42 +0000 UTC
Subject Alt Names:
========================================
Serial Number : 159159747900478145820483398898491642637
Subject: CN=GTS Root R1,O=Google Trust Services LLC,C=US
Issuer: CN=GlobalSign Root CA,OU=Root CA,O=GlobalSign nv-sa,C=BE
Not Before: 2020-06-19 00:00:42 +0000 UTC
Not After: 2028-01-28 00:00:42 +0000 UTC
Subject Alt Names:
========================================
```

Fetch a cert from an ip address:
```
$ ./fetchcert 1.1.1.1
Dail 1.1.1.1:443
-----BEGIN CERTIFICATE-----
MIIF+TCCBYCgAwIBAgIQD3WjbTLBawPHyl9fcUoDcDAKBggqhkjOPQQDAzBWMQsw
CQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMTAwLgYDVQQDEydEaWdp
Q2VydCBUTFMgSHlicmlkIEVDQyBTSEEzODQgMjAyMCBDQTEwHhcNMjExMDI1MDAw
MDAwWhcNMjIxMDI1MjM1OTU5WjByMQswCQYDVQQGEwJVUzETMBEGA1UECBMKQ2Fs
aWZvcm5pYTEWMBQGA1UEBxMNU2FuIEZyYW5jaXNjbzEZMBcGA1UEChMQQ2xvdWRm
bGFyZSwgSW5jLjEbMBkGA1UEAxMSY2xvdWRmbGFyZS1kbnMuY29tMFkwEwYHKoZI
zj0CAQYIKoZIzj0DAQcDQgAE+ylE8pg/2L2CVtMsvY4JnzErmCaeIpaNe0v82sV7
eymqjjVsnApIBWyJc+0gDs1GIfDsTbOl6a8bOJnl9NrxhKOCBBIwggQOMB8GA1Ud
IwQYMBaAFAq8CCkXjKU5bXoOzjPHLrPt+8N6MB0GA1UdDgQWBBQZRRsjGPh02iIU
y0Zr4hOzYBWCQDCBpgYDVR0RBIGeMIGbghJjbG91ZGZsYXJlLWRucy5jb22CFCou
Y2xvdWRmbGFyZS1kbnMuY29tgg9vbmUub25lLm9uZS5vbmWHBAEBAQGHBAEAAAGH
BKKfJAGHBKKfLgGHECYGRwBHAAAAAAAAAAAAERGHECYGRwBHAAAAAAAAAAAAEAGH
ECYGRwBHAAAAAAAAAAAAAGSHECYGRwBHAAAAAAAAAAAAZAAwDgYDVR0PAQH/BAQD
AgeAMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjCBmwYDVR0fBIGTMIGQ
MEagRKBChkBodHRwOi8vY3JsMy5kaWdpY2VydC5jb20vRGlnaUNlcnRUTFNIeWJy
aWRFQ0NTSEEzODQyMDIwQ0ExLTEuY3JsMEagRKBChkBodHRwOi8vY3JsNC5kaWdp
Y2VydC5jb20vRGlnaUNlcnRUTFNIeWJyaWRFQ0NTSEEzODQyMDIwQ0ExLTEuY3Js
MD4GA1UdIAQ3MDUwMwYGZ4EMAQICMCkwJwYIKwYBBQUHAgEWG2h0dHA6Ly93d3cu
ZGlnaWNlcnQuY29tL0NQUzCBhQYIKwYBBQUHAQEEeTB3MCQGCCsGAQUFBzABhhho
dHRwOi8vb2NzcC5kaWdpY2VydC5jb20wTwYIKwYBBQUHMAKGQ2h0dHA6Ly9jYWNl
cnRzLmRpZ2ljZXJ0LmNvbS9EaWdpQ2VydFRMU0h5YnJpZEVDQ1NIQTM4NDIwMjBD
QTEtMS5jcnQwDAYDVR0TAQH/BAIwADCCAX4GCisGAQQB1nkCBAIEggFuBIIBagFo
AHcAKXm+8J45OSHwVnOfY6V35b5XfZxgCvj5TV0mXCVdx4QAAAF8uWHB7wAABAMA
SDBGAiEAywu8Te3xzKNYwRup8r2yqBarJOqzTPv+43ay4/u8lZACIQD2CmGDyBOo
WMWrdSORUhPvAAGGDQ+AiVEr29SGK9RRCQB1AFGjsPX9AXmcVm24N3iPDKR6zBsn
y/eeiEKaDf7UiwXlAAABfLlhwfoAAAQDAEYwRAIgf8AhijjUmD/AY3HXlZcqNyFB
g6ZUlYjAAy4bcyqmYfgCIAfiUkbqx0lV3F9nm2KFCqrQ7T1qTQmd0xssEP0IXe70
AHYAQcjKsd8iRkoQxqE6CUKHXk4xixsD6+tLx2jwkGKWBvYAAAF8uWHBuAAABAMA
RzBFAiBk2pXuLzS2KFlvrgsWpBKinb/3AG+pgWSc1k4A8a0QeAIhAI8C1YuZelsI
zYV3zpHiyIlkxKTdlYg0lnP+/UwdHTaAMAoGCCqGSM49BAMDA2cAMGQCMBFSpqTJ
DKHJXG+cnPDLLcY/H0IwpBHKQ8M/qSHx0tRRW1uK5oRoGiDDDHYuDxtKGwIwSaH5
q4tyVvs3fmARob3fVmfy1gq53ktKYhZQwGFwr5q065Zfi4QRB9OVqiRbtx4a
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIEFzCCAv+gAwIBAgIQB/LzXIeod6967+lHmTUlvTANBgkqhkiG9w0BAQwFADBh
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBD
QTAeFw0yMTA0MTQwMDAwMDBaFw0zMTA0MTMyMzU5NTlaMFYxCzAJBgNVBAYTAlVT
MRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxMDAuBgNVBAMTJ0RpZ2lDZXJ0IFRMUyBI
eWJyaWQgRUNDIFNIQTM4NCAyMDIwIENBMTB2MBAGByqGSM49AgEGBSuBBAAiA2IA
BMEbxppbmNmkKaDp1AS12+umsmxVwP/tmMZJLwYnUcu/cMEFesOxnYeJuq20ExfJ
qLSDyLiQ0cx0NTY8g3KwtdD3ImnI8YDEe0CPz2iHJlw5ifFNkU3aiYvkA8ND5b8v
c6OCAYIwggF+MBIGA1UdEwEB/wQIMAYBAf8CAQAwHQYDVR0OBBYEFAq8CCkXjKU5
bXoOzjPHLrPt+8N6MB8GA1UdIwQYMBaAFAPeUDVW0Uy7ZvCj4hsbw5eyPdFVMA4G
A1UdDwEB/wQEAwIBhjAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwdgYI
KwYBBQUHAQEEajBoMCQGCCsGAQUFBzABhhhodHRwOi8vb2NzcC5kaWdpY2VydC5j
b20wQAYIKwYBBQUHMAKGNGh0dHA6Ly9jYWNlcnRzLmRpZ2ljZXJ0LmNvbS9EaWdp
Q2VydEdsb2JhbFJvb3RDQS5jcnQwQgYDVR0fBDswOTA3oDWgM4YxaHR0cDovL2Ny
bDMuZGlnaWNlcnQuY29tL0RpZ2lDZXJ0R2xvYmFsUm9vdENBLmNybDA9BgNVHSAE
NjA0MAsGCWCGSAGG/WwCATAHBgVngQwBATAIBgZngQwBAgEwCAYGZ4EMAQICMAgG
BmeBDAECAzANBgkqhkiG9w0BAQwFAAOCAQEAR1mBf9QbH7Bx9phdGLqYR5iwfnYr
6v8ai6wms0KNMeZK6BnQ79oU59cUkqGS8qcuLa/7Hfb7U7CKP/zYFgrpsC62pQsY
kDUmotr2qLcy/JUjS8ZFucTP5Hzu5sn4kL1y45nDHQsFfGqXbbKrAjbYwrwsAZI/
BKOLdRHHuSm8EdCGupK8JvllyDfNJvaGEwwEqonleLHBTnm8dqMLUeTF0J5q/hos
Vq4GNiejcxwIfZMy0MJEGdqN9A57HSgDKwmKdsp33Id6rHtSJlWncg+d0ohP/rEh
xRqhqjn1VtvChMQ1H3Dau0bwhr9kAMQ+959GG50jBbl9s08PqUU643QwmA==
-----END CERTIFICATE-----
========================================
Serial Number : 20549233353565500740390275235468411760
Subject: CN=cloudflare-dns.com,O=Cloudflare\, Inc.,L=San Francisco,ST=California,C=US
Issuer: CN=DigiCert TLS Hybrid ECC SHA384 2020 CA1,O=DigiCert Inc,C=US
Not Before: 2021-10-25 00:00:00 +0000 UTC
Not After: 2022-10-25 23:59:59 +0000 UTC
Subject Alt Names: cloudflare-dns.com;*.cloudflare-dns.com;one.one.one.one
========================================
Serial Number : 10566067766768619126890179173671052733
Subject: CN=DigiCert TLS Hybrid ECC SHA384 2020 CA1,O=DigiCert Inc,C=US
Issuer: CN=DigiCert Global Root CA,OU=www.digicert.com,O=DigiCert Inc,C=US
Not Before: 2021-04-14 00:00:00 +0000 UTC
Not After: 2031-04-13 23:59:59 +0000 UTC
Subject Alt Names:
========================================
```
