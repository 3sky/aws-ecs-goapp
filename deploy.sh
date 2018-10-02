# install AWS SDK
export PATH=$PATH:$HOME/.local/bin

# update an AWS ECS service with the new image
ecs-deploy -c $CLUSTER_NAME -n $SERVICE_NAME -i $IMAGE_REPO_URL:latest