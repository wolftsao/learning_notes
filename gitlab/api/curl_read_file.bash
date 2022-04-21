curl -s -H "Authorization: Bearer <token>" https://gitlab.paradise-soft.com.tw/api/v4/projects/:id/repository/files/:filepath/raw?ref=:branch

# Replace below:
# :id -> project id
# :filepath -> URL encoded filepath
# :branch -> branch name

# /raw in the end of url means raw content
# without /raw in the end, the output will be JSON with base64 encoded file content
