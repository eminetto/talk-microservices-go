# -*- mode: Python -*-
local_resource(
    'auth-compile',
    cmd='cd auth; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/auth main.go',
    deps=['auth/main.go', 'auth/security', 'auth/user', 'pkg'],
)

docker_build(
    'auth-image',
    './auth',
    dockerfile='auth/Dockerfile',
    # only=[
    #     'auth',
    # ]
)

local_resource(
    'feedbacks-compile',
    cmd='cd feedbacks; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/feedbacks main.go',
    deps=['feedbacks/main.go', 'feedbacks/feedback', 'pkg'],
)

docker_build(
    'feedbacks-image',
    './feedbacks',
    dockerfile='feedbacks/Dockerfile',
    # only=[
    #     'auth',
    # ]
)

local_resource(
    'votes-compile',
    cmd='cd votes; CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/votes main.go',
    deps=['votes/main.go', 'votes/vote', 'pkg'],
)

docker_build(
    'votes-image',
    './votes',
    dockerfile='votes/Dockerfile',
    # only=[
    #     'auth',
    # ]
)

docker_compose('./docker-compose.yml')

# local_resource('auth', cmd='cd auth; go build -o bin/auth main.go',
#                serve_cmd='auth/bin/auth', deps=['auth/main.go', 'auth/security', 'auth/user', 'pkg'])

# local_resource('feedbacks', cmd='cd feedbacks; go build -o bin/feedbacks main.go',
#                serve_cmd='feedbacks/bin/feedbacks', deps=['feedbacks/main.go', 'feedbacks/feedback', 'pkg'])


# local_resource('votes', cmd='cd votes; go build -o bin/votes main.go',
#                serve_cmd='votes/bin/votes', deps=['votes/main.go', 'votes/vote', 'pkg'])
