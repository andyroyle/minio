

Build
=====

Run make

    make

Then, download
[generate_cert.go](https://docs.min.io/docs/how-to-secure-access-to-minio-server-with-tls#using-go),
place it `$HOME/.minio`. Then run this:

    $ cd ~/.minio
    $ go run generate_cert.go -ca --host "YOUR-IP-ADDRESS"
    $ mv cert.pem public.crt
    $ mv key.pem private.key

Change the directory where `minio` located, then run minio like this:

    $ mkdir ~/minio-data
    $ minio server ~/minio-data

Download and install AWS cli, and configure it using API-key and Secret key displayed in the output of `minio`, and configure AWS cli and test it after reading [AWS CLI with MinIO](https://docs.min.io/docs/aws-cli-with-minio).

Then, run following command-line to test simple EBS API:

    aws --debug --no-verify-ssl --endpoint-url https://localhost:9000 ec2 describe-volumes
