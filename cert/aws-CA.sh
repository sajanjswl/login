# rm *.key
# rm *.crt
# rm *.csr 

# # 1. Generate CA's private key and self-signed certificate
# openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=FR/ST=Occitanie/L=Toulouse/O=Tech School/OU=Education/CN=*.kubeosc.com/emailAddress=sjnjaiswal@gmail.com"

# # echo "CA's self-signed certificate"
# openssl x509 -in ca-cert.pem -noout -text

# # # 2. Generate web server's private key and certificate signing request (CSR)
# openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=FR/ST=Ile de France/L=Paris/O=PC Book/OU=Computer/CN=*.kubeosc.com/emailAddress=sjnjaiswal@gmail.com"

# # # 3. Use CA's private key to sign web server's CSR and get back the signed certificate
# openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf

# # echo "Server's signed certificate"
# openssl x509 -in server-cert.pem -noout -text

# # 4. Generate client's private key and certificate signing request (CSR)
# openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem -subj "/C=FR/ST=Alsace/L=Strasbourg/O=PC Client/OU=Computer/CN=*.kubeosc.com/emailAddress=sjnjaiswal@gmail.com"

# # 5. Use CA's private key to sign client's CSR and get back the signed certificate
# openssl x509 -req -in client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile client-ext.cnf

# echo "Client's signed certificate"
# openssl x509 -in client-cert.pem -noout -text


 
#  export csr="MIIE3TCCAsUCAQAwgZcxCzAJBgNVBAYTAkZSMQ8wDQYDVQQIDAZBbHNhY2UxEzAR
# BgNVBAcMClN0cmFzYm91cmcxEjAQBgNVBAoMCVBDIENsaWVudDERMA8GA1UECwwI
# Q29tcHV0ZXIxFjAUBgNVBAMMDSoua3ViZW9zYy5jb20xIzAhBgkqhkiG9w0BCQEW
# FHNqbmphaXN3YWxAZ21haWwuY29tMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIIC
# CgKCAgEAn0eIIrbuak95pPhrzIwQRbs7Aj0TSro8+M68UM67WkJYV7qotzBSf/wq
# 8gvOE3mhdeH4vyzOrV083SxGLW2rRc5dCAMQMGeQapxxLxULjgjifwQDaR/QElPg
# u7CsDngHzR7dh0monolggKHi0zYYPSku5MN2Hkefh0S2qHT4LhjdLWQEwkldrXDW
# hhcDSTioMAMjCnukEkIf7jY5uDvdmHYfkyc0Umtx1rGPeRVYtP9n7IFWWO5T/ZO+
# GDe4NDF/XcGpXUV/a02iGZADZAQB/767z/CyCzowZTH+yJjPbCa0i7oqERq3Y6i5
# uecFAOwZnNZ5hEPogw1yzCbSKoo/mE5Rp/CvhwQ+2VTi1Ks355xIi7acte9eOYEq
# sYSP/j20A/OU9TM7lpp5O969DkfB2NJ9BjQqR8ZP+/r53iuKRRZblFSjS+M93hiz
# wqbz8qnYcGueLzdu/RnwlFh5FB51mNBX/+6EYyH3JTLjT3a+8717SWo2u+m5Z+70
# S4x00UfC/BLXJiR/YSzp6At8ekcKm7RqUTfDksw++RIK/6PgwFARrdMvycrxr0ov
# kl4uyIIOEGKzb8n7gkddOGQLWmV0IyUxUwrrihF7BKfWNm2wSdP2W1fYNkV5wQAf
# U/ldGIsPEulaweGmkHqg6TbaVhXdrs/wu2hpFrk0QDl1D0PuPUkCAwEAAaAAMA0G
# CSqGSIb3DQEBCwUAA4ICAQA75HOJG+LABiF9FfcZj7ZyBjf/R8YDFKyPJ4Dd9elA
# eulfVSgj05PQyBAMqFcn2mXCELKpzaMpV9agXwh3GJmToCCGsmI4FPwQrd5XWVkc
# nVvK1S7mcfYvGJkXnYKK+zoCKhs1L/SyZmZtt4KJM87EUE6gEJJ9hBqvwLngXJZw
# UBeio/H/4HkU3T2Dy7RBQ2+bS0ze0hn0/zlUsjgaeNeVdhRf7XWe8mgWTbgoRjgJ
# D670F9GzXIBB/wDM/6KzOXM80xTOkWuKi3hKyZPDtW52irRtNCrpgF4kvRIyP5D1
# O2oz6vkyE1tHLjJp8VlUBJcpePHYkdJkFYK9Av1aVDYGunSOt67es57rxlwczcrs
# 6A4cy+1ylZSAJaDHtytam1RtS5vKSxYa/v+j20A87TiGYh/o2bsc7nKZfeVFVThO
# flZY7pCPvxDM0oG9Sz/kL+WNyxcKwz4iIicBY2PuxQACi/ytFztQ7/GhM8EvuvpZ
# RDqNsEbfO/uxDLBLPo1BWSGWjw85gFqjxtKAWgKlEWWOqcC3o4F+kh/ASpiB1FA5
# BSBr0ZpyjJ69TqN5VJnqMbtiB6DAX1b2Xg7JRxa12SklfsU4vQmAUXvlPjTuhCqM
# 983fz+gXuG8x8AaLMCzvpQYgZ8OGqGT1ikoVeprp5gqgDk7QhAWnKgBSchONhf/y
# ug=="
# aws acm-pca issue-certificate \
# --certificate-authority-arn arn:aws:acm-pca:us-west-2:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012 \
# --csr $(echo $csr | base64) \
# --signing-algorithm "SHA256WITHRSA" \
# --template-arn “arn:aws:acm-pca:::template/EndEntityServerAuthCertificate/V1” \
# --validity Value=365,Type="DAYS"



# # Create a root key.
# openssl genrsa -out ExampleInternalCA-Root.key 4096
# # Create and self sign the Root Certificate.
# openssl req -x509 -new -nodes -key ExampleInternalCA-Root.key \
# -sha256 -days 3650 -out ExampleInternalCA-Root.crt \
# -subj "/CN=ExampleInternalCA-Root"


# Create the CA in 'Pending Certificate' state
#  certificateAuthorityArn=$(aws acm-pca create-certificate-authority \
#  --certificate-authority-configuration KeyAlgorithm=RSA_2048,SigningAlgorithm=SHA256WITHRSA,Subject={CommonName=ExampleInternalCA-TLS} \
# --certificate-authority-type "SUBORDINATE" \
# | jq -r '.CertificateAuthorityArn')

# # Wait for the CSR to become available.
# aws acm-pca wait certificate-authority-csr-created \
# --certificate-authority-arn $certificateAuthorityArn

# # Retrieve the CSR
# aws acm-pca get-certificate-authority-csr \
#  --certificate-authority-arn $certificateAuthorityArn \
# | jq -r ".Csr" >  ExampleInternalCA-TLS.csr


# # Prepare the extensions file for OpenSSL
# echo "basicConstraints=critical,CA:TRUE" > openssl-ca-extensions.ext
# # Generate the Private CA Certificate
# openssl x509 -req -in ExampleInternalCA-TLS.csr \
# -CA ExampleInternalCA-Root.crt -CAkey ExampleInternalCA-Root.key \
# -CAcreateserial -out ExampleInternalCA-TLS.crt -days 3650 -sha256 \
# -extfile openssl-ca-extensions.ext
# # Import the CA certificate into ACM PCA.
aws acm-pca import-certificate-authority-certificate \
--certificate-authority-arn arn:aws:acm-pca:us-east-1:259459167566:certificate-authority/47186ef2-d0d6-4d37-9c1e-856a993043a0 \
--certificate fileb://ExampleInternalCA-TLS.crt \
--certificate-chain fileb://ExampleInternalCA-Root.crt




export exampleinternalcatlscrt="\n-----BEGIN CERTIFICATE-----
MIID1zCCAb+gAwIBAgIJAMk6lh0FT79yMA0GCSqGSIb3DQEBCwUAMCExHzAdBgNV
BAMMFkV4YW1wbGVJbnRlcm5hbENBLVJvb3QwHhcNMjEwNTAyMTkzMzI4WhcNMzEw
NDMwMTkzMzI4WjAgMR4wHAYDVQQDDBVFeGFtcGxlSW50ZXJuYWxDQS1UTFMwggEi
MA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCYK+WeZ1pCvzzoE7YtOxqujNJz
Vzfubb+9koNjm3/Y6jJc6RYXHIxMoD2PvP3fXuHCPHHkdGYb/Vr3vsFQjrx4FyGT
z1Q7NGWeyM6KyJI2uqAnh8hcBwlA/KXcbamxC552939B0XIZawW+9jvK2tQvuq1B
reu8k/aSWde9Tx3RnCYuF62WkItzZMlKKKPjZR+ztsFLF3vF4DP2oAhjlKt54Qrf
J+1cnJzMxhVtKnxz37WD8bNKNhgUOv+Dm9lmsca6KUbrJwY743NTdsOeKU38WlFK
scwkQxaGUFUZTowXbY12c5yjNDQco+g1GuQakyHydBD31KHFR5ACU2Nsw+SfAgMB
AAGjEzARMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggIBAJMqumXp
C2veR+/QZ+nEq6W9p6eJl+jO8vKCVk+odCmXvJNm4U2lhQa2dCO0R2BTpF/i6MSI
nYTawNexq7jok5UrojHL2NBBP1XDnE5QRhhWxcalzdR+VwQVh80Cz8yVKcg9TtBY
2q6+Ey6ic7j3YHroLgXPA11aoAS1EoIkoftCI56RKBR5l8lWyM7aI2jpJrS5VOpg
TQDCza7TAJlJL2NAMMGzQlc8dmMoeRJhkqa8tPlWcZOCGEZHfe5kas2eD2239hcv
/NIa06/N/ciEqw8Pws6UAu/ez71CLuMl3OgfwgrH9NCOlqxVi4qOevo0Q/A++ohJ
zYhdKVuGWkKoq+Y420stK72FZ+Gm3b5BkamGF0mLv0x0ZyMdzos9C2T4sbEia6qv
K4XmF1nbnTxm5YItCXJismxIBGsjWBW8T0IUhV05kY4uVa6MHRXS/n0n//QNQU1Z
woLyg2RgRHMEuahomsdzYVtfUqr8qWfi76kltcNzn3UgZDqKEm8rIrEIUjJ0KdqA
KZjZT9Y8X/SkP7l23XaH5VuCbrWIyFpypNjq8VA//uexs5ME8KJzbDc2+vZWaBvb
G3QCwrI/vfhzjVYqzv1iKDF1ld1BaHBE354leAejNuuv9cWcQXlr3SBvNDmUnORv
jlXP6oaH16wWj+qkhENV6X1PtP+iU79fYSBY
-----END CERTIFICATE-----\n"


export exampleinternalcarootcrt="\n-----BEGIN CERTIFICATE-----
MIIEvjCCAqYCCQCCE7/cuZQ9dDANBgkqhkiG9w0BAQsFADAhMR8wHQYDVQQDDBZF
eGFtcGxlSW50ZXJuYWxDQS1Sb290MB4XDTIxMDUwMjE5MzMyMloXDTMxMDQzMDE5
MzMyMlowITEfMB0GA1UEAwwWRXhhbXBsZUludGVybmFsQ0EtUm9vdDCCAiIwDQYJ
KoZIhvcNAQEBBQADggIPADCCAgoCggIBAJ1zb8z4n8WP0OQ/fRkxr0s+czRiybQT
mIVk5/xPgN6vLsHW+h7h6lD3IJvckWhY6merop2KsZEsGVNQBkl9v8xQZCJCfgY1
WIIQVFsxmSCdU4t651K4yOS4X7elW9CRn2Z+y216FT1jUwZbMmR6Ho5m6Xg8HSlq
XEaOwp9IdknitYOr3R3ahb96kQ6cmZ5up9Yv+4VAvaPZjmqm+SqShDXK0982+R9r
a9TgbW2F6sofVbcXzjpV0nhnxwmhnFg1qwAqz2hICA4ThVUyVuYmOciTSbXfh51k
+QN5EuaCdvqIqSpqB3iOC9/pbb7mPc1VyyuX7OrW/bPhu80tY2QyGJsaE0y7pMlJ
/+IQAhJJ2Kx2ecrKbpZJJv78m8O7ENBgKLEuQ/TPFg3tF7VC49fjSoQKpaMmR/CC
euRLgC+T5Z+EYkiR+tm4vsHebaUmrydeOmLDYY25Y26XQzTTrZIH3iOhMyOftzaO
E0/RfvlDXUmBSrcZ5CqPqj2ixMG0tVZa77NgUSqHjOnAaouVsLcl3KD9dlBh6ef7
rjNsXNQyDBFW97OuC2++qV4tpMX8c6AW7X3wHAGRWgfVZ01hLUMx4oE8j9DHzNNB
J2pGB9gWRTZR2vwa2yNUWLpgpUfUDd0QMHbv1/ueH7yds+MD4wtiG8k5iWS1c0qt
G5Gu/zQkzcELAgMBAAEwDQYJKoZIhvcNAQELBQADggIBADfngbPiWrWjz5wm1sNq
R/YL/h47qnQ+DBjOy/F/zp7999SJdg+SeDLYUT60A3e1x0yf27M5m9nPwgHg83le
Xo4uQJ4ltFJx/+UuTRL4p9DniH5R5EC9XkDE5+RFXuRuppeFm9RS5GT8ZcdfgX8o
zaaTJA4mCtJUM38Q2Dq3rZejBP8Xh2S2NAI/mEQwP65go64zzIUrvw3Jc3XF0F/t
qStn+QxWtMaDRyUUvG8yPvVWP/H4HKuH5Xl62wDMpOoKut5H/Fd4ISHMC2k06xXJ
T/jQw1AxwQ6s8z22kUMmgnEVdysI4TIRkjchfmUWoMTGp0xtr18kBctupc/mWD6Q
16Gy8fqbamlee9mLsmDxCbthmHXinwK6lY0xfJlDC/33Xxsl25orJrzpss5AHbco
VZuuaLwOitVpx1GtpWhNBjaK+3EcL/8pjOa3zHBQ/FtfyG2irzv3czoQa7TlQzPB
g1t0+JLNXphN0ZGO+GqzijO4UAyAFszps1gr6sCdZVduYGBQYQ+vDp+FxtRb49Zj
NgJIa2UZNOP4msyIdjjRRYChbcVohGwU19AoBKLtodP3TM9xD360dqmqbpjGUH0x
nr8JviYb5cM6sZMOEu7EdPkQlqhjLHyHVbNdfn+rN7Yu2wd0hxyJf6HaNHtStltr
v/CaUpi3OZiUmqI5ibPWHecL
-----END CERTIFICATE-----\n"


# aws acm-pca import-certificate-authority-certificate \
# --certificate-authority-arn arn:aws:acm-pca:us-east-1:259459167566:certificate-authority/47186ef2-d0d6-4d37-9c1e-856a993043a0 \
# --certificate $(echo $exampleinternalcatlscrt | base64) \
# --certificate-chain $(echo $exampleinternalcarootcrt | base64)