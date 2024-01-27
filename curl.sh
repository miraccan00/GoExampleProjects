token=$(curl -s -H "Content-Type: application/json" -X POST -d '{"username": "x", "password": "y"}' \
https://hub.docker.com/v2/users/login/ | jq -r '.token')

curl -s -H "Authorization: Bearer $token" https://hub.docker.com/v2/repositories/x/?page_size=100 | jq -r '.results|.[]|.name' > docker_images.txt
