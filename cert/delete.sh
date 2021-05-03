# disable CA
# aws acm-pca update-certificate-authority \
# --certificate-authority-arn $certificateAuthorityArn \
# --status DISABLED

# Delete CA
aws acm-pca delete-certificate-authority \
--certificate-authority-arn arn:aws:acm-pca:us-east-1:259459167566:certificate-authority/a44cb2a3-9f72-424e-b971-6a6fcc47415c


# List CA
 aws acm-pca list-certificate-authorities | jq .CertificateAuthorities