echo "########### Creating profile ###########"

aws configure set aws_access_key_id default_access_key --profile=localstack
aws configure set aws_secret_access_key default_secret_key --profile=localstack
aws configure set region us-east-1 --profile=localstack

echo "########### Listing profile ###########"
aws configure list --profile=localstack

echo "########### Creating SNS ###########"
awslocal sns create-topic --name image-recognition-notification

echo "########### Creating S3 ###########"
awslocal s3api create-bucket --bucket image-repository
