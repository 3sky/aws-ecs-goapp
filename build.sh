# install AWS SDK
pip install --user awscli
export PATH=$PATH:$HOME/.local/bin

# install ecs-deploy
add-apt-repository ppa:eugenesan/ppa
apt-get update
apt-get install jq -y
curl https://raw.githubusercontent.com/silinternational/ecs-deploy/master/ecs-deploy | \
  sudo tee -a /usr/bin/ecs-deploy
sudo chmod +x /usr/bin/ecs-deploy

# login AWS ECR
eval $(aws ecr get-login --region eu-west-1)


# build the docker image and push to an image repository
docker build -t goapp .
docker tag goapp:latest $IMAGE_REPO_URL:latest